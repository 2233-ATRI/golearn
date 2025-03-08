package main

const MOD = 1000000007

func numberOfGoodSubarraySplits(nums []int) int {
	ones := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			ones = append(ones, i)
		}
	}

	// 如果没有1，返回0
	if len(ones) == 0 {
		return 0
	}
	// 如果只有1个1，返回1
	if len(ones) == 1 {
		return 1
	}

	// 计算每对相邻1之间0的个数
	result := 1
	for i := 0; i < len(ones)-1; i++ {
		zeros := ones[i+1] - ones[i] - 1 // 两1间0的个数
		// 每段的方法数 = 0的个数 + 1
		result = (result * (zeros + 1)) % MOD
	}

	return result
}

// 数学方法，先找到每个1所在的位置，当得到之后可以知道两个1之间的差距大小，不需要在意第一个1之前的0的数量
// 中间每次都是0的个数 + 1
func main() {

}
