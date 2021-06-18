package treeRela

import (
	"fmt"
	"testing"
)

var numTreesMem map[string]int

func numTrees(n int) int {
	numTreesMem = make(map[string]int)
	return countNumTrees(1, n)
}

// 没有剪枝操作的话会超时
// 计算 [low, high] 范围内数字构成的 BST 总数
func countNumTrees(low, high int) int {
	if low > high {
		return 1
	}
	// 剪枝操作
	key := fmt.Sprintf("%d-%d", low, high)
	if v, exists := numTreesMem[key]; exists{
		return v
	}

	ret := 0
	// 计算 [low, high] 范围内每个数字作为根节点时构成的 BST 总数
	for i := low; i <= high; i++ {
		leftNum := countNumTrees(low, i-1)
		rightNum := countNumTrees(i+1, high)
		// 注意：这里是左右子树各自组合的乘积
		ret += leftNum * rightNum
	}
	// 将计算结果保存
	numTreesMem[key] = ret
	return ret
}

func TestNumTrees(t *testing.T) {
	fmt.Println(numTrees(19))
}