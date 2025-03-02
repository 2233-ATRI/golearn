package main

//func BubbleSort(ac []int) {
//	n := len(ac)
//	var c = 0
//	for i := 0; i < n; i++ {
//		for j := 0; j < n-i; j++ {
//			if ac[j] > ac[i] {
//				c = i
//				i = j
//				j = c
//
//			}
//		}
//	}
//}

func bubble(ac []int) {
	n := len(ac)
	//var c =0
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if ac[j] > ac[j+1] {
				ac[j], ac[j+1] = ac[j+1], ac[j]
			}
		}
	}
}

func selection(ac []int) {
	n := len(ac)
	//var q =0//交换中间值
	var w = 0 //交换的序数

	for i := 0; i < n; i++ {
		var min = ac[0] //当前最小值
		for j := i; j < n; j++ {
			if ac[j] < min {
				min = ac[j]
				w = j
			}
		}
		//q = ac[i]
		ac[w] = ac[i]
		ac[i] = min
	}
}

func main() {

}
