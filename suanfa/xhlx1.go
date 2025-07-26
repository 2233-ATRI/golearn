package main

import (
	"fmt"
)

func length(s string) int {
	fmt.Println("length")
	return len(s)
}

func main() {
	a := "abcd"
	for i, n := 0, len(a); i < n; i++ {
		fmt.Println(a[i])
	}
	a1 := len(a)
	for a1 > 0 {
		a1--
		fmt.Println(a[a1])

	}

	s := "abcd"
	for i, n := 0, length(s); i < n; i++ {
		fmt.Println(s[i])
	}
}
