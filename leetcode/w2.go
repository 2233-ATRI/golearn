package main

//给你一个正整数数组 nums ，对 nums 所有元素求积之后，找出并返回乘积中 不同质因数 的数目。
//
//注意：
//
//质数 是指大于 1 且仅能被 1 及自身整除的数字。
//如果 val2 / val1 是一个整数，则整数 val1 是另一个整数 val2 的一个因数。
//func distinctPrimeFactors(nums []int) int {
//	var s int =1
//	//for p:=0; p<len(nums);p++ {
//	//	s =s*nums[p]//得到元素积
//	//}
//	var i =2
//	for {
//		if i ==
//	}
//
//} //函数进行循环，当不可以被2整除的时候递归到3，然后判断除数是否可以被比其小的数不可整除即可

// 给你一个正整数数组 nums ，对 nums 所有元素求积之后，找出并返回乘积中 不同质因数 的数目。
//
// 注意：
//
// 质数 是指大于 1 且仅能被 1 及自身整除的数字。
// 如果 val2 / val1 是一个整数，则整数 val1 是另一个整数 val2 的一个因数。
func distinctPrimeFactors(nums []int) int {
	var ans int
	// 复制数组避免修改原输入
	arr := make([]int, len(nums))
	copy(arr, nums)

	var zs = 2 // 当前检查的数字，从2开始

	for {
		// 检查数组是否全为1
		allOnes := true
		for i := 0; i < len(arr); i++ {
			if arr[i] > 1 {
				allOnes = false
				break
			}
		}
		if allOnes {
			break // 如果全为1，结束循环
		}

		// 寻找质数
		var s = 2
		isPrime := true
		for s < zs { // 修改：只需要检查到zs-1
			if zs%s == 0 {
				isPrime = false
				break
			}
			s++
		}

		if isPrime {
			divided := false
			// 对每个元素尝试除以当前质数
			for i := 0; i < len(arr); i++ {
				if arr[i] != 1 {
					for arr[i]%zs == 0 {
						arr[i] /= zs
						divided = true
					}
				}
			}
			if divided {
				ans++ // 如果有任何数字被除，增加计数
			}
		}
		zs++ // 检查下一个数字
	}

	return ans
}

// 直接计算每个元素的质数解，每次除法之前先对该元素判断==1则跳过，当所有的元素除该质数有余数则跳至下一个质数
// 则只需要找出所有的质数即可
func main() {

}
