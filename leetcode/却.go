package main

// 给你一个下标从 0 开始的整数数组 nums ，同时给你一个整数 key ，它在 nums 出现过。
//
// 统计 在 nums 数组中紧跟着 key 后面出现的不同整数 target 的出现次数。换言之，target 的出现次数为满足以下条件的 i 的数目：
//
// 0 <= i <= n - 2
// nums[i] == key 且
// nums[i + 1] == target 。
// 请你返回出现 最多 次数的 target 。测试数据保证出现次数最多的 target 是唯一的。
// 2 <= nums.length <= 1000
// 1 <= nums[i] <= 1000
// 测试数据保证答案是唯一的。
func mostFrequent(nums []int, key int) int {
	ac := make([]int, 1001)

	// 只遍历到len(nums)-2，因为要访问i+1
	for i := 0; i < len(nums)-1; i++ { // 从前往后遍历也可以解决问题
		if nums[i] == key {
			ac[nums[i+1]]++
		}
	}

	maxCount := 0
	result := 0
	for target, count := range ac {
		if count > maxCount {
			maxCount = count
			result = target // 记录出现次数最多的target
		}
	}

	return result
} //创建一个数组长度1000，当对应的元素紧跟着 key 后面出现该位置++最后遍历
// 找到最大的元素下标并记录
func main() {

}
