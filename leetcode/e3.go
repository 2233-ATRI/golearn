package main

// 给你一个 从 0 开始的排列 nums（下标也从 0 开始）。请你构建一个 同样长度 的数组 ans ，其中，对于每个 i（0 <= i < nums.length），都满足 ans[i] = nums[nums[i]] 。返回构建好的数组 ans 。
//
// 从 0 开始的排列 nums 是一个由 0 到 nums.length - 1（0 和 nums.length - 1 也包含在内）的不同整数组成的数组。
func buildArray(nums []int) []int {
	// 创建一个新数组来存储结果，避免修改原数组导致错误
	ans := make([]int, len(nums))

	// 按照题目要求构建 ans[i] = nums[nums[i]]
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}

func main() {

}
