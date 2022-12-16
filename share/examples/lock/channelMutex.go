package lock

import "time"

type ChannelMutex struct {
	ch chan struct{}
}

func NewChannelMutex() *ChannelMutex {
	mu := &ChannelMutex{ch: make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

func (m *ChannelMutex) Lock() {
	<-m.ch
}

func (m *ChannelMutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

// TryLock 如果成功，则返回 true；如果失败，则返回 false。
// 无论成败都会立即返回，获取不到锁（锁已被其他线程获取）时不会一直等待！
func (m *ChannelMutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

// TryLockWithTimeout 可定时获取锁，类似于 TryLock
// 这个方法在获取不到锁时会等待一定的时间，在时间期限之内如果还获取不到锁，就返回 false。如果如果一开始拿到锁或者在等待期间内拿到了锁，则返回 true
func (m *ChannelMutex) TryLockWithTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:
	}
	return false
}
