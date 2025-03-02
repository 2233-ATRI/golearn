package main

// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
//
// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
//
// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
func mySqrt(x int) int {
	var t = 2 //尝试数字
	var z = 0
	for t < x {
		z++
		t = t * 2
	}
	t = 1
	for i := 0; i < z/2; i++ {
		t = t * 2
	}
	return t
}
func main() {

}
