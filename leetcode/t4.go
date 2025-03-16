package main

import "sort"

// 给你两个正整数数组 nums1 和 nums2 ，数组的长度都是 n 。
//
// 数组 nums1 和 nums2 的 绝对差值和 定义为所有 |nums1[i] - nums2[i]|（0 <= i < n）的 总和（下标从 0 开始）。
//
// 你可以选用 nums1 中的 任意一个 元素来替换 nums1 中的 至多 一个元素，以 最小化 绝对差值和。
//
// 在替换数组 nums1 中最多一个元素 之后 ，返回最小绝对差值和。因为答案可能很大，所以需要对 109 + 7 取余 后返回。
//
// |x| 定义为：
//
// 如果 x >= 0 ，值为 x ，或者
// 如果 x <= 0 ，值为 -x
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	const MOD = 1000000007
	n := len(nums1)

	// 第一步：计算初始绝对差值和，并找到最大差值
	sum := 0
	for i := 0; i < n; i++ {
		diff := abs(nums1[i] - nums2[i])
		sum = (sum + diff) % MOD
	}

	// 如果初始和为 0，无需替换
	if sum == 0 {
		return 0
	}

	// 第二步：对 nums1 排序，用于快速查找最接近 nums2[i] 的值
	sortedNums1 := make([]int, n)
	copy(sortedNums1, nums1)
	sort.Ints(sortedNums1)

	// 第三步：遍历每个位置，计算替换后能减少的最大差值
	maxGain := 0 // 记录能减少的最大差值
	for i := 0; i < n; i++ {
		if nums1[i] == nums2[i] {
			continue // 如果当前差值为 0，无需替换
		}

		// 当前位置的差值
		currDiff := abs(nums1[i] - nums2[i])

		// 二分查找在 sortedNums1 中找到最接近 nums2[i] 的值
		idx := binarySearch(sortedNums1, nums2[i])
		newDiff := abs(sortedNums1[idx] - nums2[i])

		// 如果 idx 不是第一个元素，还要检查前一个值
		if idx > 0 {
			newDiff = min(newDiff, abs(sortedNums1[idx-1]-nums2[i]))
		}
		// 如果 idx 不是最后一个元素，还要检查后一个值
		if idx < n-1 {
			newDiff = min(newDiff, abs(sortedNums1[idx+1]-nums2[i]))
		}

		// 计算当前替换能减少的差值
		gain := currDiff - newDiff
		maxGain = max(maxGain, gain)
	}

	// 第四步：计算最终结果并取模
	result := (sum - maxGain) % MOD
	if result < 0 {
		result += MOD
	}
	return result
}

// 辅助函数：计算绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 辅助函数：取最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 辅助函数：取最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 辅助函数：二分查找找到最接近 target 的值的索引
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// 进行三次遍历，第一次找到差值最大的一对数，第二次找到在两个数组中除了这两个数以外和这两个数差值最小的数和该数
// 对应的数替换，最后遍历取余
func main() {

}
