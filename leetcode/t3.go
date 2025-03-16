package main

func checkPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	// 复制一份数组用于修改
	nums1 := make([]int, len(nums))
	nums2 := make([]int, len(nums))
	copy(nums1, nums)
	copy(nums2, nums)

	// 找到第一个违反非递减的地方
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			// 尝试两种修改方式
			// 方式 1：改 nums[i] 为 nums[i-1]
			nums1[i] = nums1[i-1]
			// 方式 2：改 nums[i-1] 为 nums[i]
			nums2[i-1] = nums2[i]

			// 检查两种修改后的数组是否非递减
			return isNonDecreasing(nums1) || isNonDecreasing(nums2)
		}
	}

	return true // 如果没有违反，已经是非递减
}

// 辅助函数：检查数组是否非递减
func isNonDecreasing(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
} // 在出现第一次违反了非递减的情况的时候将其修改为其左或者右的最大值，
// 这个时候在判断修改之后的是不是非递减数列，这个时候假如依然不是非递减数列则输出false，否则输出true
func main() {

}
