package builtin_sync

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

// 线程安全的计数器
type Counter struct {
	mu    sync.Mutex
	count uint64
}

// 对计数值加一
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 获取当前的计数值
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// sleep 1秒，然后计数值加1
// 注意，这里传入的是 WaitGroup 指针
func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}

func TestWaitGroup(t *testing.T) {
	var (
		counter Counter
		wg      sync.WaitGroup
		num     = 10
	)
	wg.Add(num) // 添加等待的 goroutine 总数

	for i := 0; i < num; i++ { // 启动 num 个 goroutine 执行"加1"任务
		go worker(&counter, &wg)
	}

	wg.Wait() // 检查点，等待 goroutine 都完成任务
	assert.Equal(t, uint64(num), counter.Count())
}

func TestWaitGroup_Err(t *testing.T) {
	var wg sync.WaitGroup
	go doSomething(100, &wg)
	go doSomething(110, &wg)
	go doSomething(120, &wg)
	go doSomething(130, &wg)

	wg.Wait()
	t.Log("Done")
}

func doSomething(sleep time.Duration, wg *sync.WaitGroup) {
	duration := sleep * time.Millisecond
	time.Sleep(duration) // 故意 sleep 一段时间
	wg.Add(1)
	println("后台执行, duration:", duration)
	wg.Done()
}
