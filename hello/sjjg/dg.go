package main

//func ac(n int) int {
//	if n == 0 {
//		return 1
//	}
//	return ac(n-1) * n
//}

//func ac(n int) int {
//
//	if n < 2 {
//
//		return 1
//
//	}
//
//	return ac(n-1) + ac(n-2)
//
//}

func erfcz(arr []int, tar int, l int, r int) int {
	if l > r {
		return 0
	}
	mid := (l + r) / 2
	if mid > tar {
		return erfcz(arr, tar, l, mid)
	} else if mid < tar {
		return erfcz(arr, tar, mid, r)
	} else if mid == tar {
		return mid
	}
	return 0
}

func main() {
	//fmt.Println(ac(3))

}
