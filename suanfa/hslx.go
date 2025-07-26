package main

import "fmt"

func test(x, y int, s string) (int, string) {
	n := x + y
	return n, fmt.Sprintf("%s%d", s, n)
}

func main() {
	x := 1
	y := 2
	s := "wqh"
	r, s := test(x, y, s)
	fmt.Println(r, s)
}
