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

func main() {

}
