package greedy

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"sort"
	"testing"
)

func getMinCoinCountByGreedy(coins []int, k int) int {
	sort.Ints(coins)
	count, sum := 0, 0
	path := make([]int, 0) // 可以不用，仅为了 debug
	var backtracking func(start int)
	backtracking = func(start int) {
		if sum > k {
			return
		}
		for i := start; i >= 0 && sum < k; i-- {
			sum += coins[i]
			count++
			path = append(path, coins[i])
			fmt.Println("sum = ", sum, ", path = ", path)
			backtracking(start)
			if sum != k {
				sum -= coins[i]
				count--
				path = path[:len(path)-1]
			}
		}
	}
	backtracking(len(coins) - 1)
	if sum != k {
		return -1
	}
	return count
}

// 递归+枚举
func getMinCoinCountByRecursion(coins []int, k int) int {
	coinSets := make([][]int, 0)
	// 1. 递归求所有的组合
	// coinSet[i] 记录了当前组合中需要多少枚 coins[i] 面额的硬币
	var getMinCountHelper func(total int, coinSet []int)
	getMinCountHelper = func(total int, coinSet []int) {
		if total == 0 {
			coinSets = append(coinSets, coinSet)
			return
		}
		for i, coin := range coins {
			if coin > total {
				continue
			}
			newSet := make([]int, len(coins))
			copy(newSet, coinSet)
			newSet[i]++
			getMinCountHelper(total-coin, newSet)
		}
	}
	initSet := make([]int, len(coins))
	getMinCountHelper(k, initSet)

	if len(coinSets) == 0 {
		return -1
	}
	// 2. 枚举所有的组合，找到硬币数最少的组合
	minCount := k + 1
	for _, set := range coinSets {
		totalCount := 0
		for _, v := range set {
			totalCount += v
		}
		if totalCount < minCount {
			minCount = totalCount
		}
	}
	return minCount
}

// 递归+备忘录剪枝
func getMinCoinCountByMemorization(coins []int, k int) int {
	memo := make([]int, k+1)
	for i := 1; i < len(memo); i++ {
		memo[i] = -2
	}
	var getMinCoinCount func(total int) int
	getMinCoinCount = func(total int) int {
		if memo[total] != -2 {
			return memo[total]
		}
		minCount := k + 1
		for i := 0; i < len(coins); i++ {
			if coins[i] > total {
				continue
			}
			restCount := getMinCoinCount(total - coins[i])
			if restCount == -1 {
				continue
			}
			totalCount := restCount + 1
			if totalCount < minCount {
				minCount = totalCount
			}
		}
		if minCount == k+1 {
			memo[total] = -1
			return -1
		}
		memo[total] = minCount
		return minCount
	}
	return getMinCoinCount(k)
}

// DP
func getMinCoinCountByDP(coins []int, k int) int {
	dp := make([]int, k+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = k + 1
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for total := 1; total < len(dp); total++ {
		for _, coin := range coins {
			if total-coin < 0 {
				continue
			}
			dp[total] = min(dp[total], dp[total-coin]+1)
		}
	}
	fmt.Println("dp = ",dp)
	if dp[k] == k+1 {
		return -1
	}
	return dp[k]
}

func TestGetMinCoinCount(t *testing.T) {
	tests := []struct {
		name  string
		coins []int
		k     int
		want  int
	}{
		{
			name:  "1",
			coins: []int{5, 3},
			k:     11,
			want:  3, // 5+3+3
		},
		{
			name:  "2",
			coins: []int{5, 3},
			k:     7,
			want:  -1,
		},
		{
			name:  "3",
			coins: []int{1, 2, 5},
			k:     12,
			want:  3, // 5+5+2
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, getMinCoinCountByGreedy(test.coins, test.k))
			assert.Equal(t, test.want, getMinCoinCountByRecursion(test.coins, test.k))
			assert.Equal(t, test.want, getMinCoinCountByMemorization(test.coins, test.k))
			assert.Equal(t, test.want, getMinCoinCountByDP(test.coins, test.k))
		})
	}

}
