package arrayRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 由于必然有一个众数，其出现的次数 > len(nums)/2，所以，可以先对其排序，然后去 len(nums)/2 位置的值即可

func majorityElementForce(nums []int) int {
	set := make(map[int]int)
	for _, num := range nums {
		set[num]++
	}
	beforeFreq, result := 0, 0
	for num, freq := range set {
		if freq > beforeFreq {
			result = num
			beforeFreq = freq
		}
	}
	return result
}

// 摩尔投票算法
func majorityElement(nums []int) int {
	major, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count++
			major = nums[i]
		} else if major == nums[i] {
			count++
		} else {
			count--
		}
	}
	return major
}

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{
			nums:   []int{3, 2, 3},
			output: 3,
		},
		{
			nums:   []int{2, 2, 1, 1, 1, 2, 2},
			output: 2,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, majorityElementForce(test.nums))
	}
}
