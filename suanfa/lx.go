package main

import "fmt"

func mac1(x int) {
	x = 100
}

func mac2(x *int) {
	*x = 100
}

func main() {
	//a := 10
	//b := &a
	//fmt.Printf("a=%d par =%p,%d", a, &a, b)
	//fmt.Println()
	//mac1(a)
	//fmt.Printf("a=%d par =%p,%d", a, &a, b)
	//fmt.Println()
	//mac2(b)
	//fmt.Printf("a=%d par =%p,%d", a, &a, b)
	//var s *string
	//fmt.Println(s)
	//if s == nil {
	//	fmt.Println("kong")
	//} else {
	//	fmt.Println(s)
	//}

	var a *int
	a = new(int)
	*a = 10
	fmt.Println(*a)
}
