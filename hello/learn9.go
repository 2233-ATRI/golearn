package main

import "fmt"

func printfarray(arr []int) []int {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println("%T", arr)
	printfarray(arr)
	var slice []int = make([]int, 3) //分配空间
	fmt.Println("%T", slice)
	fmt.Println("len = ", len(slice))
	if slice == nil {
		fmt.Println("nil slice")
	} else {
		fmt.Println(" slice")
	}
}
