package main

// 给你一个长度为偶数下标从 0 开始的二进制字符串 s 。
//
// 如果可以将一个字符串分割成一个或者更多满足以下条件的子字符串，那么我们称这个字符串是 美丽的 ：
//
// 每个子字符串的长度都是 偶数 。
// 每个子字符串都 只 包含 1 或 只 包含 0 。
// 你可以将 s 中任一字符改成 0 或者 1 。
//
// 请你返回让字符串 s 美丽的 最少 字符修改次数。
//
//	func minChanges(s string) int {
//		changes := 0
//		a := 1 // 当前连续相同字符长度，从1开始
//		for i := 1; i < len(s); i++ {
//			if s[i] == s[i-1] {
//				a++
//			} else {
//				if a%2 == 0 { // 偶数段，符合条件
//					a = 1 // 从当前字符开始新段
//				} else { // 奇数段，修改前一个字符
//					changes++
//					// 假设改为当前字符，使前段变偶数
//					a = 1 // 从当前字符开始新段
//				}
//			}
//		}
//		// 处理最后一段
//		if a%2 != 0 {
//			changes++
//		}
//		return changes
//	}
func minChanges(s string) int {
	changes := 0
	for i := 0; i < len(s)-1; i += 2 { // 强制每段长度2
		if s[i] != s[i+1] {
			changes++
		}
	}
	return changes
}

//从头开始遍历，将得到的相同数记录a++，当出现不同时进行判断，如果是
//偶数则a=0从该位继续++，如果奇数则对上一位更改，本位从a=0开始++

func main() {

}
