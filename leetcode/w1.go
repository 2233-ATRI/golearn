package main

// 给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
//
// 请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
//
// 假设每一种面额的硬币有无限个。
//
// 题目数据保证结果符合 32 位带符号整数。

func change(amount int, coins []int) int {
	// 创建 dp 数组，长度为 amount + 1
	dp := make([]int, amount+1)
	// 初始化：金额 0 的组合数为 1
	dp[0] = 1

	// 遍历每种硬币
	for _, coin := range coins {
		// 对于当前硬币，更新从 coin 到 amount 的所有金额的组合数
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}

// 我的思路是将里面的所有元素都视为最小的金额的倍数，然后看能否得到
// 这种情况下的组合，当存在时则存在这么一个序列
// 直接对总金额对所有的硬币面额进行除法，得到该面额最多需要多少枚，这样就变成了一个有限的问题了
func main() {

}
