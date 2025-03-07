package main

// 给你一个整数数组 nums ，返回全部为 0 的 子数组 数目。
//
// 子数组 是一个数组中一段连续非空元素组成的序列。
func zeroFilledSubarray(nums []int) int64 {
	var s int64 = 0 // 当前连续0的个数
	var t int64 = 0 // 总子数组数

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			s++
		} else if s != 0 { // 遇到非0且有累计的0
			t += (s * (s + 1)) / 2
			s = 0
		}
	}
	// 处理末尾的0段
	if s != 0 {
		t += (s * (s + 1)) / 2
	}

	return t
} //数学方法解决，直接统计连续多少个0，当出现非0则结束
// 统计，给中间变量+n*(n+1)/2然后继续向后寻找新0开始的
// 元素
func main() {

}
