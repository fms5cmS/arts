
硬币找零问题

给定 n 中不同面值的硬币 coins，以及一个总金额 k，计算出**最少**需要几枚硬币凑出这个金额 k。每种硬币的个数不限，且如果没有任何一种硬币组合能组成总金额时，返回 -1。

> case1:
> input: coins = [1, 2, 5], k = 12
> output: 3
> explanation: 12 = 2 + 5 + 5

穷举所有的组合。

一般来说穷举从来都不是一个好方法。除非需要的结果就是所有的不同组合， 而不是一个最值。但即便是求所有的不同组合，计算的过程中也仍然会出现重复计算的问题，这种现象称之为**重叠子问题**。

**贪心算法的的思想是局部最优**，每一步计算作出的都是在当前看起来最好的选择，即在一定条件下的最优解。

贪心算法的基本思路：
1. 根据问题来建立数学模型;
2. 把待求解问题划分成若干个子问题，对每个子问题进行求解，得到子问题的局部最优解;
3. 把子问题的局部最优解进行合并，得到最后基于局部最优解的一个解，即原问题的答案。

假设 `coins = [5, 3], k = 11`，根据贪心，先挑面值最大的 5 放入组合中，还剩 6 需要凑；再次贪心，放入 5 ...

```go
package main

import "sort"

func getMinCoinCount(coins []int, k int) int {
    total, count := k, 0
    length := len(coins)
    sort.Ints(coins)
    for i := length - 1; i >= 0; i-- {
        curCount := total / coins[i] // 计算当前面值最多可以放几个
        total -= curCount * coins[i] // 计算余额
        count += curCount
        if total == 0 {
            return count
        }
    }
    return -1 // 未凑出
}
```

贪心算法很容易出现"过于贪心"的问题，比如上例的解法中，最后会发现返回值为 -1，但实际是存在返回结果的（硬币组合为 5+3+3），而上面的代码中先放入 2 个 5，导致仅余 1，所以无法得到正确的结果。可以**利用回溯来尝试当前阶段的"局部次优方案"来解决过于贪心的问题**。

```go
package main

import "sort"

func getMinCoinCount(coins []int, k int) int {
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
```

