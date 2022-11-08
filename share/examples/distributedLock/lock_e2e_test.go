package distributedLock

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

// E2E (End to End) 测试，集成测试

type ClientE2ESuite struct {
	suite.Suite
	rdb redis.Cmdable
}

func (s *ClientE2ESuite) SetupSuite() {
	s.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// 用于确保测试的 Redis 已正常启动并建立连接
	for s.rdb.Ping(context.Background()).Err() != nil {

	}
}

func TestClientE2E(t *testing.T) {
	suite.Run(t, &ClientE2ESuite{})
}

func (s *ClientE2ESuite) TestLock() {
	t := s.T()
	rdb := s.rdb
	testCases := []struct {
		name       string
		key        string
		expiration time.Duration
		retry      RetryStrategy
		timeout    time.Duration
		client     *Client
		wantLock   *Lock
		wantErr    error
		before     func()
		after      func()
	}{
		{
			name:       "locked",
			key:        "locked-key",
			retry:      &FixIntervalRetry{Interval: time.Second, Max: 1},
			timeout:    time.Second,
			expiration: time.Minute,
			client:     NewClient(rdb),
			before:     func() {},
			after: func() {
				res, err := rdb.Del(context.Background(), "locked-key").Result()
				require.NoError(t, err)
				require.Equal(t, int64(1), res)
			},
		},
		{
			// 锁已被其他人持有导致加锁不成功
			name:       "failed",
			key:        "failed-key",
			retry:      &FixIntervalRetry{Interval: time.Second, Max: 3},
			timeout:    time.Second,
			expiration: time.Minute,
			client:     NewClient(rdb),
			before: func() {
				res, err := rdb.Set(context.Background(), "failed-key", "123", time.Minute).Result()
				require.NoError(t, err)
				require.Equal(t, "OK", res)
			},
			after: func() {
				res, err := rdb.Get(context.Background(), "failed-key").Result()
				require.NoError(t, err)
				require.Equal(t, "123", res)
				delRes, err := rdb.Del(context.Background(), "failed-key").Result()
				require.NoError(t, err)
				require.Equal(t, int64(1), delRes)
			},
			wantErr: ErrFailedToPreemptLock,
		},
		{
			// 已经加锁，再加锁就是刷新超时时间
			name:       "already-locked",
			key:        "already-key",
			retry:      &FixIntervalRetry{Interval: time.Second, Max: 3},
			timeout:    time.Second,
			expiration: time.Minute,
			client: func() *Client {
				client := NewClient(rdb)
				client.valuer = func() string {
					return "123"
				}
				return client
			}(),
			before: func() {
				res, err := rdb.Set(context.Background(), "already-key", "123", time.Minute).Result()
				require.NoError(t, err)
				require.Equal(t, "OK", res)
			},
			after: func() {
				res, err := rdb.Get(context.Background(), "already-key").Result()
				require.NoError(t, err)
				require.Equal(t, "123", res)
				delRes, err := rdb.Del(context.Background(), "already-key").Result()
				require.NoError(t, err)
				require.Equal(t, int64(1), delRes)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.before()
			l, err := testCase.client.Lock(context.Background(), testCase.key, testCase.expiration, testCase.retry, testCase.timeout)
			assert.True(t, errors.Is(err, testCase.wantErr))
			if err != nil {
				return
			}
			assert.Equal(t, testCase.key, l.key)
			assert.Equal(t, testCase.expiration, l.expiration)
			assert.NotEmpty(t, l.value)
			testCase.after()
		})
	}
}

func (s *ClientE2ESuite) TestTryLock() {
	t := s.T()
	rdb := s.rdb
	client := NewClient(rdb)
	testCases := []struct {
		name       string
		key        string
		expiration time.Duration
		wantLock   *Lock
		wantErr    error
		before     func()
		after      func()
	}{
		{
			// 加锁成功
			name:       "locked",
			key:        "locked-key",
			expiration: time.Minute,
			before:     func() {},
			after: func() {
				res, err := rdb.Del(context.Background(), "locked-key").Result()
				require.NoError(t, err)
				require.Equal(t, int64(1), res)
			},
			wantLock: &Lock{
				key:        "locked-key",
				expiration: time.Minute,
			},
		},
		{
			// 模拟并发竞争失败
			name:       "failed",
			key:        "failed-key",
			expiration: time.Minute,
			before: func() {
				// 假设已经有人设置了分布式锁
				val, err := rdb.Set(context.Background(), "failed-key", "123", time.Minute).Result()
				require.NoError(t, err)
				require.Equal(t, "OK", val)
			},
			after: func() {
				res, err := rdb.Del(context.Background(), "failed-key").Result()
				require.NoError(t, err)
				require.Equal(t, int64(1), res)
			},
			wantErr: ErrFailedToPreemptLock,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before()
			l, err := client.TryLock(context.Background(), tc.key, tc.expiration)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.key, l.key)
			assert.Equal(t, tc.expiration, l.expiration)
			assert.NotEmpty(t, l.value)
			tc.after()
		})
	}
}

func (s *ClientE2ESuite) TestRefresh() {
	t := s.T()
	rdb := s.rdb
	testCases := []struct {
		name    string
		timeout time.Duration
		lock    *Lock
		wantErr error
		before  func()
		after   func()
	}{
		{
			name: "refresh success",
			lock: &Lock{
				key:        "refresh-key",
				value:      "123",
				expiration: time.Minute,
				unlock:     make(chan struct{}, 1),
				client:     rdb,
			},
			before: func() {
				// 设置一个比较短的过期时间
				res, err := rdb.SetNX(context.Background(), "refresh-key", "123", time.Second*10).Result()
				require.NoError(t, err)
				assert.True(t, res)
			},
			after: func() {
				res, err := rdb.TTL(context.Background(), "refresh-key").Result()
				require.NoError(t, err)
				// 刷新完过期时间后判断现在的过期时间
				assert.True(t, res.Seconds() > 50)
				// 清理数据
				rdb.Del(context.Background(), "refresh-key")
			},
			timeout: time.Minute,
		},
		{
			name: "refresh failed", // 锁被人持有
			lock: &Lock{
				key:        "refresh-key",
				value:      "123",
				expiration: time.Minute,
				unlock:     make(chan struct{}, 1),
				client:     rdb,
			},
			before: func() {
				// 设置一个比较短的过期时间
				res, err := rdb.SetNX(context.Background(), "refresh-key", "456", time.Second*10).Result()
				require.NoError(t, err)
				assert.True(t, res)
			},
			after: func() {
				res, err := rdb.Get(context.Background(), "refresh-key").Result()
				require.NoError(t, err)
				require.Equal(t, "456", res)
				rdb.Del(context.Background(), "refresh-key")
			},
			timeout: time.Minute,
			wantErr: ErrLockNotHold,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.before()
			ctx, cancel := context.WithTimeout(context.Background(), testCase.timeout)
			err := testCase.lock.Refresh(ctx)
			cancel()
			assert.Equal(t, testCase.wantErr, err)
			testCase.after()
		})
	}
}
