package main

// 给你一个以二进制形式表示的数字 s 。请你返回按下述规则将其减少到 1 所需要的步骤数：
//
// 如果当前数字为偶数，则将其除以 2 。
//
// 如果当前数字为奇数，则将其加上 1 。
//
// 题目保证你总是可以按上述规则将测试用例变为 1 。
func numSteps(s string) int {
	steps := 0
	tybes := []byte(s)
	for len(tybes) > 1 {
		if tybes[len(tybes)-1] == 0 {
			tybes = tybes[:len(tybes)-1]
		} else {
			i := len(tybes) - 1
			for i >= 0 && tybes[i] == '1' {
				tybes[i] = '0'
				i--
			}
			if i >= 0 {
				tybes[i] = '1'
			} else {
				tybes = append([]byte{'0'}, tybes[:i]...)
			}
		}
		steps++
	}
	return steps
}

// 二进制形式表示的数字s可以直接对二进制更改，需要对奇数加一则可以对二进制
// 结尾是1时结尾+1，结尾是0时去掉最后一位然后统计递归次数
func main() {

}
