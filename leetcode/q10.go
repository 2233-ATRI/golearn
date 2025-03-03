package main

// 给你一个二进制字符串 s（仅由 '0' 和 '1' 组成的字符串）。
//
// 返回所有字符都为 1 的子字符串的数目。
//
// 由于答案可能很大，请你将它对 10^9 + 7 取模后返回。

//const mod = 1000000007
//
//var count int
//
//func numSub(s string) int {
//	count = 0
//	for i := 0; i < len(s); i++ {
//		if s[i] == '1' {
//			ddf(s, i, 0)
//		}
//	}
//	return count % mod
//}
//
//func ddf(s string, i int, j int) {
//	if i >= len(s) || s[i] == '0' {
//		return
//	}
//	j++
//	count = (count + 1) % mod
//	ddf(s, i+1, j)
//}

// 在函数外有一个全局变量设置一个递归在每次递归当中设置一个计数器
// ，对每次出现的1的子串进行统计，每一次递归都对匹配子串长度+1
// ，当出现没有匹配上的时候返回全局变量

// 纯数学办法：
const mod = 1000000007

func numSub(s string) int {
	count := 0
	ones := 0 // 当前连续 1 的个数

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			ones++
		} else {
			// 遇到 '0' 时，计算前一段连续 1 的子字符串数量
			count = (count + (ones*(ones+1))/2) % mod
			ones = 0 // 重置计数
		}
	}
	// 处理末尾的连续 1
	if ones > 0 {
		count = (count + (ones*(ones+1))/2) % mod
	}
	return count
}
func main() {

}
