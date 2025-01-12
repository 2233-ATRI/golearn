package main

//	func longestCommonPrefix(strs []string) string {
//		base := strs[0] //1)panic自排错不会死循环
//		for i := 0; i < len(base); i++ {
//			for _, str := range strs[1:] {
//				if i >= len(str) || str[i] != base[i] {
//					return base[:i]
//				}
//			}
//		}
//		return base
//	}
func longestCommonPrefix(strs []string) string {
	min := len(strs[0])
	tip := 0
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < min {
			min = len(strs[i])
			tip = i
		}
	}
	for i := 1; i < min; i++ {
		if strs[0][i] == strs[1][i] && strs[0][i] == strs[2][i] {
			continue
		}
		if strs[0][i] != strs[1][i] || strs[1][i] != strs[2][i] || strs[2][i] != strs[0][i] {
			if i == 0 {
				return strs[0][:0]
			}
			return strs[0][:i]
		}
	}
	return strs[tip]
}
func main() {

}
