package slideWindow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 借助双端队列
// 队列没有必要维护窗口里的所有元素，只需要**维护有可能成为窗口里最大值的元素就可以了**，同时保证队里的元素数值是由大到小的。
// 1. 队列中存储索引而非 num，是为了方便判断元素是否在当前窗口之外
// 窗口从左往右移动，假设当前元素的索引为 i，则窗口最左侧元素的索引为 i-k+1
// 2. 队列中的元素是有序的，递减
// 3. 每个元素入队时，需要从右侧开始逐个将小于该元素的出队
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k || k == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}
	result := make([]int, len(nums)-k+1)
	// 双端队列中保存的是索引
	dq := &deque{}
	for i := range nums {
		// 判断双端队列左侧首元素是否在当前窗口内，如果不在，则将其从队列中移除
		if !dq.empty() && (i-k+1 > dq.lPeek()) {
			dq.lPop()
		}
		// 将队列中小于当前元素的都出队
		for !dq.empty() && nums[dq.rPeek()] < nums[i] {
			dq.rPop()
		}
		// 当前元素的索引入队
		dq.push(i)
		// 队列的左侧首元素代表的就是当前窗口的最大值
		if i >= k-1 {
			result[i-k+1] = nums[dq.lPeek()]
		}
	}
	return result
}

// 双端队列
type deque struct {
	data []int
}

func (q *deque) push(v int) {
	q.data = append(q.data, v)
}

func (q *deque) lPeek() int {
	return q.data[0]
}

func (q *deque) lPop() {
	q.data = q.data[1:]
}

func (q *deque) rPeek() int {
	return q.data[len(q.data)-1]
}

func (q *deque) rPop() {
	q.data = q.data[:len(q.data)-1]
}

func (q *deque) empty() bool {
	return len(q.data) == 0
}

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "1",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "2",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, maxSlidingWindow(test.nums, test.k))
		})
	}
}
