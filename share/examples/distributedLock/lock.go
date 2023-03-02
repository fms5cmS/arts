package distributedLock

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

//go:generate mockgen -destination ./lock_mock.go -package distributedLock -source $GOFILE

var (
	// embed 是 Go1.16 加入的包，可以在编译阶段将静态资源文件打包仅编译好的程序中，并提供访问这些文件的能力
	//go:embed script/lua/unlock.lua
	luaUnlock string
	//go:embed script/lua/refresh.lua
	luaRefresh string

	//go:embed script/lua/lock.lua
	luaLock string

	ErrFailedToPreemptLock = errors.New("rlock: 抢锁失败")
	// ErrLockNotHold 一般是出现在你预期你本来持有锁，结果却没有持有锁的地方
	// 比如说当你尝试释放锁的时候，可能得到这个错误
	// 这一般意味着有人绕开了 rlock 的控制，直接操作了 Redis
	ErrLockNotHold = errors.New("rlock: 未持有锁")
)

type Client struct {
	client redis.Cmdable
	// valuer 用于生成值，将来可以考虑暴露出去允许用户自定义
	valuer func() string

	g singleflight.Group
}

func NewClient(client redis.Cmdable) *Client {
	return &Client{
		client: client,
		valuer: func() string {
			return uuid.New().String()
		},
	}
}

func (c *Client) TryLock(ctx context.Context, key string, expiration time.Duration) (*Lock, error) {
	val := uuid.NewString()
	ok, err := c.client.SetNX(ctx, key, val, expiration).Result()
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrFailedToPreemptLock
	}
	return newLock(c.client, key, val, expiration), nil
}

// Lock 是尽可能重试减少加锁失败的可能
// Lock 会在超时或者锁正被人持有的时候进行重试
// 最后返回的 error 使用 errors.Is 判断，可能是：
// - context.DeadlineExceeded: Lock 整体调用超时
// - ErrFailedToPreemptLock: 超过重试次数，但是整个重试过程都没有出现错误
// - DeadlineExceeded 和 ErrFailedToPreemptLock: 超过重试次数，但是最后一次重试超时了
// 你在使用的过程中，应该注意：
// - 如果 errors.Is(err, context.DeadlineExceeded) 那么最终有没有加锁成功，谁也不知道
// - 如果 errors.Is(err, ErrFailedToPreemptLock) 说明肯定没成功，而且超过了重试次数
// - 否则，和 Redis 通信出了问题
func (c *Client) Lock(ctx context.Context, key string, expiration time.Duration, retry RetryStrategy, timeout time.Duration) (*Lock, error) {
	val := c.valuer()
	var timer *time.Timer
	defer func() {
		if timer != nil {
			timer.Stop()
		}
	}()
	for {
		lctx, cancel := context.WithTimeout(ctx, timeout)
		res, err := c.client.Eval(lctx, luaLock, []string{key}, val, expiration.Seconds()).Result()
		cancel()
		if err != nil && !errors.Is(err, context.DeadlineExceeded) {
			// 非超时错误，如 Redis server 宕机、EOF 等
			return nil, err
		}
		if res == "OK" {
			return newLock(c.client, key, val, expiration), nil
		}
		interval, ok := retry.Next()
		if !ok {
			if err != nil {
				err = fmt.Errorf("最后一次重试错误: %w", err)
			} else {
				err = fmt.Errorf("锁被人持有: %w", ErrFailedToPreemptLock)
			}
			return nil, fmt.Errorf("rlock: 重试机会耗尽，%w", err)
		}
		if timer == nil {
			timer = time.NewTimer(interval)
		} else {
			timer.Reset(interval)
		}
		select {
		case <-timer.C:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

// SingleFlightLock 非常高并发，且热点集中的情况下，结合 singleFlight 优化，即 本地所有的 goroutine 先竞争，胜利者再去抢全局的分布式锁
func (c *Client) SingleFlightLock(ctx context.Context, key string, expiration time.Duration, retry RetryStrategy, timeout time.Duration) (*Lock, error) {
	for {
		flag := false // 标记是不是自己拿到了锁
		result := c.g.DoChan(key, func() (interface{}, error) {
			flag = true
			return c.Lock(ctx, key, expiration, retry, timeout)
		})
		select {
		case res := <-result:
			if flag { // 自己拿到了锁
				c.g.Forget(key)
				if res.Err != nil {
					return nil, res.Err
				}
				return res.Val.(*Lock), nil
			}
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

type Lock struct {
	client           redis.Cmdable
	key              string
	value            string
	expiration       time.Duration
	unlock           chan struct{}
	signalUnlockOnce sync.Once
}

func newLock(client redis.Cmdable, key string, value string, expiration time.Duration) *Lock {
	return &Lock{
		client:     client,
		key:        key,
		value:      value,
		expiration: expiration,
		unlock:     make(chan struct{}, 1),
	}
}

func (l *Lock) Unlock(ctx context.Context) error {
	res, err := l.client.Eval(ctx, luaUnlock, []string{l.key}, l.value).Int64()
	defer func() {
		l.unlock <- struct{}{}
	}()
	if errors.Is(err, redis.Nil) {
		return ErrLockNotHold
	}
	if err != nil {
		return err
	}
	if res != 1 {
		return ErrLockNotHold
	}
	return nil
}

// Refresh 手动续约，使用见 lock_test.go ExampleLock_Refresh
// 1. 间隔多久续约一次？
// 2. 如果 Refresh 返回了 error，怎么处理？
// 3. 如果确认续约失败，怎么中断后续的业务？无解，业务代码一旦执行，除非手动检测分布式锁，并手动中断，不然是没办法的！
func (l *Lock) Refresh(ctx context.Context) error {
	res, err := l.client.Eval(ctx, luaRefresh, []string{l.key}, l.value, l.expiration.Seconds()).Int64()
	if err != nil {
		return err
	}
	if res != 1 {
		return ErrLockNotHold
	}
	return nil
}

func (l *Lock) AutoRefresh(interval time.Duration, timeout time.Duration) error {
	ticker := time.NewTicker(interval)
	ch := make(chan struct{}, 1)
	defer func() {
		ticker.Stop()
		close(ch)
	}()
	for {
		select {
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			err := l.Refresh(ctx)
			cancel()
			// 超时，可以继续尝试续约
			if errors.Is(err, context.DeadlineExceeded) {
				// 因为有两个可能的地方要写入数据，而 ch 容量仅为 1，所以，如果写不进去，说明前一次调用超时，且还没被处理，同时计时器也触发了
				select {
				case ch <- struct{}{}:
				default:
				}
				continue
			}
			// 仅处理超时引起的续约失败
			if err != nil {
				return err
			}
		case <-ch:
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			err := l.Refresh(ctx)
			cancel()
			if errors.Is(err, context.DeadlineExceeded) {
				select {
				case ch <- struct{}{}:
				default:
				}
				continue
			}
			if err != nil {
				return err
			}
		case <-l.unlock:
			return nil
		}
	}
}
