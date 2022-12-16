package lock

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestChannelMutex(t *testing.T) {
	lock := NewChannelMutex()
	sum := 0
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			sum += 1
		}()
	}
	wg.Wait()
	assert.Equal(t, 1000, sum)
}

func TestChannelMutex_TryLock(t *testing.T) {
	lock := NewChannelMutex()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		time.Sleep(2*time.Second)
		fmt.Println("first goroutine finish!")
	}()
	time.Sleep(time.Second)
	result := lock.TryLock()
	wg.Wait()
	assert.False(t, result)
}

func TestChannelMutex_TryLockWithTimeout(t *testing.T) {
	lock := NewChannelMutex()
	go func() {
		lock.Lock()
		defer lock.Unlock()
		time.Sleep(2*time.Second)
		fmt.Println("first goroutine finish!")
	}()
	time.Sleep(500*time.Millisecond)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		if lock.TryLockWithTimeout(500*time.Millisecond) {
			defer lock.Unlock()
			fmt.Println("second goroutine get lock")
		} else {
			fmt.Println("second goroutine get lock failed")
		}
	}()
	go func() {
		defer wg.Done()
		if lock.TryLockWithTimeout(2*time.Second) {
			defer lock.Unlock()
			fmt.Println("third goroutine get lock")
		} else {
			fmt.Println("third goroutine get lock failed")
		}
	}()
	wg.Wait()
}