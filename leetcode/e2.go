package main

// 给你一个整数数组，返回它的某个 非空 子数组（连续元素）在执行一次可选的删除操作后，所能得到的最大元素总和。换句话说，你可以从原数组中选出一个子数组，并可以决定要不要从中删除一个元素（只能删一次哦），（删除后）子数组中至少应当有一个元素，然后该子数组（剩下）的元素总和是所有子数组之中最大的。
//
// 注意，删除一个元素后，子数组 不能为空。
func maximumSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}

	// 检查是否有正数，并找最小负数
	hasPositive := false
	minNeg := -1
	for _, num := range arr {
		if num > 0 {
			hasPositive = true
			break
		}
		if num < 0 && (minNeg == -1 || num > minNeg) {
			minNeg = num
		}
	}

	// 全是非正数情况
	if !hasPositive {
		maxVal := arr[0]
		for _, num := range arr {
			if num > maxVal {
				maxVal = num
			}
		}
		return maxVal // 返回最大单元素（0或最小负数）
	}

	// 有正数，用DP
	dp0 := arr[0]
	dp1 := arr[0]
	maxSum := arr[0]
	for i := 1; i < len(arr); i++ {
		dp0New := max(arr[i], dp0+arr[i])
		dp1New := max(dp0, dp1+arr[i])
		dp0 = dp0New
		dp1 = dp1New
		maxSum = max(maxSum, max(dp0, dp1))
	}
	return maxSum
} //设置一个变量存当前的最大值，一个变量是当前队列当中最小
// 的负数，从第一个正数开始，向后遍历，每次加入新元素，将当前遍历和
// -最小负数并更新最小负数，和当前最大值对比，更大则更换，当把数组遍历结束则
// 从下一个正数开始，当遍历结束返回最大值
func main() {

}
