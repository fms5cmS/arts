package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
	if v, exists := numTreesMem[key]; exists {
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

// 动态规划解法
// 画出 n 分别为 1、2、3 时可以组成的二叉搜索树，会发现 n = 3 的二叉搜索树可以由 n = 1、2 时的情况推导得出
// dp[3]，就是 元素1为头结点搜索树的数量 + 元素2为头结点搜索树的数量 + 元素3为头结点搜索树的数量
//
//	元素1为头结点搜索树的数量 = 右子树有2个元素的搜索树数量(dp[2]) * 左子树有0个元素的搜索树数量(dp[0])
//	元素2为头结点搜索树的数量 = 右子树有1个元素的搜索树数量(dp[1]) * 左子树有1个元素的搜索树数量(dp[1])
//	元素3为头结点搜索树的数量 = 右子树有0个元素的搜索树数量(dp[0]) * 左子树有2个元素的搜索树数量(dp[2])
//
// 时间 O(n^2)，空间 O(n)
func numTreesByDP(n int) int {
	// dp[i] 代表 i 个节点组成 BST 的个数
	dp := make([]int, n+1)
	// 3. 初始化，空节点也是一颗二叉搜索树
	dp[0], dp[1] = 1, 1
	// 4. 遍历，节点数为 i 的状态依赖于 i 之前节点数的状态
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// 2. 递推公式
			// j-1 代表了 j 为头节点左子树节点的数量
			// i-j 代表 j 为头节点右子树节点的数量
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func TestNumTrees(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "1",
			n:    3,
			want: 5,
		},
		{
			name: "2",
			n:    1,
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, numTrees(test.n))
			assert.Equal(t, test.want, numTreesByDP(test.n))
		})
	}
}
