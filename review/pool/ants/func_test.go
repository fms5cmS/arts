package ants

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"sync"
	"testing"
)

type taskFunc func()

func taskFuncWrapper(nums []int, i int, sum *int, wg *sync.WaitGroup) taskFunc {
	return func() {
		for _, num := range nums[i*DataPerTask : (i+1)*DataPerTask] {
			*sum += num
		}
		fmt.Printf("task:%d sum:%d\n", i+1, *sum)
		wg.Done()
	}
}

// worker pool 中放入任意的 worker（即 func 类型）
func TestFuncPool(t *testing.T) {
	pool, _ := ants.NewPool(10)
	defer pool.Release()
	// 模拟数据
	nums := make([]int, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}
	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)
	// 暂存每组数据的和
	partSums := make([]int, DataSize/DataPerTask, DataSize/DataPerTask)
	for i := 0; i < DataSize/DataPerTask; i++ {
		// 提交任务！
		pool.Submit(taskFuncWrapper(nums, i, &partSums[i], &wg))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", pool.Running())
	// 对比结果
	var sum int
	for _, partSum := range partSums {
		sum += partSum
	}
	var expect int
	for _, num := range nums {
		expect += num
	}
	fmt.Printf("finish all tasks, result is %d expect:%d\n", sum, expect)
}
