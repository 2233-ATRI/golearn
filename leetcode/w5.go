package main

// 给你一个正整数 n 。n 中的每一位数字都会按下述规则分配一个符号：
//
// 最高有效位 上的数字分配到 正 号。
// 剩余每位上数字的符号都与其相邻数字相反。
// 返回所有数字及其对应符号的和。
// 1 <= n <= 10^9
//
//	func alternateDigitSum(n int) int {
//		m := make([]int, 9)
//		var t = 0
//		for i := 0; i < 9; i++ {
//			m[i] = n % (10 * (i + 1))
//		}
//		var i = 0
//		for {
//			if m[i] != 0 {
//				var x = 1
//				for {
//					m[i] = x * m[i]
//					x = -1 * x
//				}
//				break
//			} else {
//				i++
//			}
//		}
//		for i = 0; i < 9; i++ {
//			t = t + m[i]
//		}
//		return t
//	} //创建一个九位的数组，对其中的元素进行取余，然后放入，从最高非0 开始进行数位反转
//
// 然后对数组求和
func alternateDigitSum(n int) int {
	m := make([]int, 10) // 改为10位，适应10^9
	var t = 0

	temp := n
	i := 0
	for temp > 0 && i < 10 { // 限制在10以内
		m[i] = temp % 10
		temp /= 10
		i++
	}
	digitCount := i

	if digitCount == 0 {
		return 0
	}
	sign := 1 // 最高位为正
	for j := digitCount - 1; j >= 0; j-- {
		m[j] = m[j] * sign
		sign = -sign
	}

	for j := 0; j < digitCount; j++ {
		t += m[j]
	}

	return t
}
func main() {

}
