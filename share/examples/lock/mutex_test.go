package lock

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
	"sync"
	"testing"
	"time"
)

func TestMutex_Base(t *testing.T) {
	var (
		mu           sync.Mutex // 不需要额外的初始化
		count        uint64
		wg           sync.WaitGroup
		goroutineNum = 10
		batchNum     = 1_000
	)
	wg.Add(goroutineNum)
	for i := 0; i < goroutineNum; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < batchNum; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}

		}()
	}
	wg.Wait()
	assert.Equal(t, uint64(goroutineNum*batchNum), count)
}

// 嵌入 Mutex

// 线程安全的计数器类型
type Counter struct {
	CounterType string

	mu    sync.Mutex // Mutex 嵌入结构体时，一般会把 Mutex 放在要控制的字段上，并使用空格把字段分隔开
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func TestMutex_InStruct(t *testing.T) {
	var (
		counter      Counter
		wg           sync.WaitGroup
		goroutineNum = 10
		batchNum     = 1_000
	)
	wg.Add(goroutineNum)
	for i := 0; i < goroutineNum; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < batchNum; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	assert.Equal(t, uint64(goroutineNum*batchNum), counter.Count())
}

// TryLock

func TestMutex_TryLock(t *testing.T) {
	var mu sync.Mutex
	// 启动一个 goroutine 持有一段时间的锁
	go func() {
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)

	if mu.TryLock() { // 尝试获取锁
		// 进入这里代表获取锁成功
		t.Log("get the lock")
		// do something
		mu.Unlock()
		return
	}
	t.Log("can't get the lock")
}
