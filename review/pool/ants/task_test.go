package ants

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Task struct {
	index int
	nums  []int
	sum   int
	wg    *sync.WaitGroup
}

func (t *Task) Do() {
	for _, num := range t.nums {
		t.sum += num
	}
	t.wg.Done()
}

// worker pool 定义了对指定数据类型固定的处理逻辑，需要使用 Invoke() 来对指定数据处理
func TestTaskPool(t *testing.T) {
	// 创建 goroutine pool
	pool, _ := ants.NewPoolWithFunc(10, func(data interface{}) {
		task := data.(*Task)
		task.Do()
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("task: %d, sum: %d\n", task.index, task.sum)
	})
	// 释放！
	defer pool.Release()
	// 随机生成 1w 个整数，等分为 100 组，每组生成一个 Task
	nums := make([]int, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}
	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)
	tasks := make([]*Task, 0, DataSize/DataPerTask)
	for i := 0; i < DataSize/DataPerTask; i++ {
		task := &Task{index: i + 1, nums: nums[i*DataPerTask : (i+1)*DataPerTask], wg: &wg}
		tasks = append(tasks, task)
		// 处理
		pool.Invoke(task)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", pool.Running())
	// 对比结果
	var sum int
	for _, task := range tasks {
		sum += task.sum
	}
	var expect int
	for _, num := range nums {
		expect += num
	}
	fmt.Printf("finish all tasks, result is %d expect:%d\n", sum, expect)
}
