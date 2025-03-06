package main

// 给你一个仅由 0 和 1 组成的二进制字符串 s 。
//
// 如果子字符串中 所有的 0 都在 1 之前 且其中 0 的数量等于 1 的数量，则认为 s 的这个子字符串是平衡子字符串。请注意，空子字符串也视作平衡子字符串。
//
// 返回  s 中最长的平衡子字符串长度。
//
// 子字符串是字符串中的一个连续字符序列。
func findTheLongestBalancedSubstring(s string) int {
	maxlen := 0
	for i := 0; i < len(s); i++ {
		m := 0 // 0和1的计数差
		n := 0 // 当前子串长度
		for j := i; j < len(s); j++ {
			if s[j] == '0' {
				n++
				m++
			} else if s[j] == '1' {
				m--
				n++
			}
			if m == 0 && j-i+1 > maxlen {
				zer := 0
				ones := 0
				valid := true
				for k := i; k <= j; k++ { // 修正：包含j
					if s[k] == '0' {
						if ones > 0 { // 1出现在0之前，无效
							valid = false
							break
						}
						zer++
					} else {
						ones++
					}
				}
				if valid { // m==0已保证zer==ones，无需再检查
					maxlen = j - i + 1
				}
			}
			if m < 0 { // 1比0多，不可能平衡
				break
			}
		}
	}
	return maxlen
} //先创建两个中间变量m=0,n=0，当从开始检测到0的时候m++，n++，当遇到
// 1的时候m--，当在m!=n的时候进行++则从该位置重新开始，如果在运行期间遇到
// m=0则使用中间元素暂存，如果下面再次m=0时则将n和该中间变量对比，更大的存在
// 该中间变量，直到遍历结束s，返回2*中间变量
func main() {

}
