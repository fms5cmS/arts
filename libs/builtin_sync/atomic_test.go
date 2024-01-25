package builtin_sync

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestCOW(t *testing.T) {
	type Map map[string]string
	var m atomic.Value
	m.Store(make(Map))
	var mu sync.Mutex // 只有写操作时才使用
	read := func(key string) string {
		m1 := m.Load().(Map)
		return m1[key]
	}
	// Copy On Write
	insert := func(key, val string) {
		mu.Lock()
		defer mu.Unlock()
		m1 := m.Load().(Map)
		m2 := make(Map)
		for k, v := range m1 {
			m2[k] = v
		}
		m2[key] = val
		m.Store(m2)
	}
	_, _ = read, insert
}

// 无符号整数利用 atomic 进行减法
// 利用计算机补码规则，将减法变成加法
func TestAtomicSub(t *testing.T) {
	var x uint32 = 12

	// 减 1
	atomic.AddUint32(&x, ^uint32(0))
	t.Log(x)
	assert.Equal(t, x, uint32(11))

	// 此时 x = 11
	// 减 3
	atomic.AddUint32(&x, ^uint32(3-1))
	t.Log(x)
	assert.Equal(t, x, uint32(8))
}

func TestAtomicValue(t *testing.T) {
	type Config struct {
		NodeName string
		Addr     string
		Count    int32
	}
	loadNewConfig := func() Config {
		return Config{
			NodeName: "BJ",
			Addr:     "10.77.95.27",
			Count:    rand.Int31(),
		}
	}

	var config atomic.Value
	config.Store(loadNewConfig())
	var cond = sync.NewCond(&sync.Mutex{})

	go func() {
		for {
			time.Sleep(time.Duration(5+rand.Int63n(5)) * time.Second)
			config.Store(loadNewConfig())
			cond.Broadcast() // 通知新配置已更新
		}
	}()
	go func() {
		for {
			cond.L.Lock()
			cond.Wait()                 // 等待更新信号
			c := config.Load().(Config) // 读取新的配置
			fmt.Printf("new config: %+v\n", c)
			cond.L.Unlock()
		}
	}()
	select {}
}
