package main

// 给你一个字符串 s，请你返回 两个相同字符之间的最长子字符串的长度 ，计算长度时不含这两个字符。如果不存在这样的子字符串，返回 -1 。
//
// 子字符串 是字符串中的一个连续字符序列。
// 1 <= s.length <= 300
// s 只含小写英文字母
func maxLengthBetweenEqualCharacters(s string) int {
	firstpor := make([]int, 26)
	for i := range firstpor {
		firstpor[i] = -1 //未出现
	}
	max := -1
	for i := 0; i < len(s); i++ {
		charindex := s[i] - 'a'
		if firstpor[charindex] < 0 {
			firstpor[charindex] = i
		} else {
			cur := i - firstpor[charindex] - 1
			if cur > max {
				max = cur
			}
		}
	}
	return max
} //创建一个有26位代表26个字母的二维数组，当检测到该元素的时候对数组对应位置输入其下标
// 如果存在数组元素超过两个则用中间变量存储其中最大值减最小值，向后继续检测，如果后面存在
// 更大数值则替代
func main() {

}
