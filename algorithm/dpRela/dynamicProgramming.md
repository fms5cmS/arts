DP(Dynamic Programming) 

# 题目特点

- 计数
    - 如：多少种方式走到右下角；多少种方法选出 k 个数使得和为 sum
- 求最大最小值
    - 如：左上角走到右下角路径的最大数字和；最长上升子序列长度
- 求存在性
    - 如：取石子游戏，先手是否必胜；能不能选出 k 个数使得和为 sum




解题步骤：

1. 确定dp数组（dp table）以及下标的含义
2. 确定递推公式
3. dp数组如何初始化（有些情况是递推公式决定了dp数组要如何初始化，所以初始化放在递推公式后面）
4. 确定遍历顺序
5. 举例推导dp数组

- 题目：
    - 斐波那契数列：509
    - 交换硬币：518、322
    - 爬楼梯：70、746
    - 找路径：62、73
    - 二叉搜索树：96

# 背包问题

## 01 背包

- 题目：
    - 正常：416、1049、494
    - 变体：474（物品的重量有两个维度）


[01背包](https://mp.weixin.qq.com/s?__biz=MzUxNjY5NTYxNA==&mid=2247486598&idx=1&sn=dd7d0530dd7a5caef7ce70cc3d6eee3f&scene=21#wechat_redirect)

数组：weight 代表物品重量，value 代表物品价值

1. dp 数组（二维）：`dp[i][j]` 代表了从下标为 [0-i] 的物品中任取，将其放入容量为 j 的背包后的**最大价值**
2. 递推公式：`dp[i][j]` 是最大价值，
    1. 如果容量为 j 的背包装不下下标为 i 的物品，则 `dp[i][j] = dp[i-1][j]`
    2. 如果容量为 j 的背包装得下下标为 i 的物品，则最大价值可以从以下两个值中比较得到（假设容量为 5，物品重量为 3）
        1. `dp[i-1][j]` 容量为 j 的背包**剩余的容量装不下**下标为 i 的物品时的最大价值（先装了重量为 4 的多个物品，剩余容量不够再装）
        2. `dp[i-1][j-weight[i]] + value[i]` 容量为 j 的背包将下标为 i 的物品装进去时的最大价值（刚好装下重量为 3 的物品，其价值为 `value[i]`，也即已经装了重量为 2 的物品，其价值为 `dp[i-1][j-weight[i]]`）
3. dp 数组初始化：
    1. 容量为 0 时，无论物品怎么选，背包价值一定为 0，所以 `dp[i][0] = 0`
    2. 由推导公式可以看出，i 是由 i-1 推导出的，所以 `i == 0` 时一定要初始化：如果背包装不下物品(`j < weight[0]`)，则 `dp[0][j] = 0`；如果背包装得下物品(`j >= weight[0]`)，则 `dp[0][j] = value[0]`
    3. 其他位置初始化为 0 即可，在遍历过程中会被覆盖
4. 遍历顺序（以下两种都可以，但存在区别）
    1. 先遍历物品再遍历背包
    2. 先遍历背包再遍历物品

```go
package main

// 已知条件
func zeroOneBag(weight, value []int, bagWeight int) int {
	weightLength := len(weight)
	// dp 数组
	dp := make([][]int, weightLength+1)
	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
	}
	// 初始化
	for j := weight[0]; j <= bagWeight; j++ {
		dp[0][j] = value[0]
	}
	// 先遍历物品再遍历背包
	for i := 1; i < weightLength; i++ {
		for j := 0; j <= bagWeight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[weightLength-1][bagWeight]
}
```

如果先遍历背包再遍历物品：

```go
	// 先遍历背包再遍历物品
	for j := 0; j <= bagWeight; j++ {
		for i := 1; i < weightLength; i++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
```

---

[01背包-滚动数组](https://mp.weixin.qq.com/s?__biz=MzUxNjY5NTYxNA==&mid=2247486624&idx=2&sn=96e8c6344dc25f57462b675b55ccd6e7&scene=21#wechat_redirect)

1. dp 数组（一维），`dp[j]` 代表了容量为 j 的背包，所装的物品最大价值为 `dp[j]`
2. 递推公式：`dp[j] = max(dp[j], dp[j - weight[i]] + value[i])`
3. 初始化：`dp[0] = 0`，由于递推公式中取得是最大价值的数，所以下标非 0 的都初始化为0，如果存在价值为负数的物品，那么非 0 下标的值需要初始化为负无穷
4. 遍历顺序：只能先遍历物品再遍历背包！

```go
// 必须是先遍历物品再遍历背包！
func zeroOneBag(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	for i, w := range weight {
		for j := bagWeight; j >= w; j-- {
			dp[j] = max(dp[j], dp[j - weight[i]] + value[i])
		}
	}
	return dp[bagWeight]
}
```

- 为什么背包倒序遍历？

注意，**背包的遍历是倒序的！是为了保证物品 i 只会被放入一次！！**

说明：假设下标为 0 的物品，`weight[0]=1, value[0]=15`

```go
// 正序遍历
dp[1] = dp[1 - weight[0]] + value[0] = 15
dp[2] = dp[2 - weight[0]] + value[0] = 30   // 这里下标 0 的物品价值被计算了两次！
// 倒序遍历
dp[2] = dp[2 - weight[0]] + value[0] = 15
dp[1] = dp[1 - weight[0]] + value[0] = 15
```

而之前二维数组中，由于 `dp[i][j]` 都是通过上一层，即 `dp[i-1][j]` 计算出来的，本层的 `dp[i][j]` 并不会被覆盖！

- 为什么不能先遍历背包再遍历物品？

因为一维数组时，背包必须倒序遍历。如果先遍历背包的话，那么每个 `dp[j]` 中就只会放入一个物品！

## 完全背包

- 题目：
    - 变体：518（排列数还是组合数？遍历顺序来决定）

[完全背包](https://mp.weixin.qq.com/s/akwyxlJ4TLvKcw26KB9uJw)

完全背包和01背包问题唯一不同的地方就是，每种物品有无限件，即可以被添加多次，所以在遍历背包的时候需要从小到大的遍历：

```go
// 先遍历物品后遍历背包
func wholeBag(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	for i, w := range weight {
		for j := w; j <= bagWeight; j++ {
			dp[j] = max(dp[j], dp[j - weight[i]] + value[i])
		}
	}
	return dp[bagWeight]
}
```

与 01背包(滚动数组)不同，完全背包中，背包和物品遍历顺序无所谓，除了上面先遍历物品后遍历背包外，还可以先遍历背包后遍历物品：

```go
// 先遍历背包后遍历物品
func wholeBag(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
    for j := 0; j <= bagWeight; j++ {
        for i, w := range weight {
        	if j - w >= 0 {
                dp[j] = max(dp[j], dp[j - weight[i]] + value[i])
            }
        }
    }
	return dp[bagWeight]
}
```


