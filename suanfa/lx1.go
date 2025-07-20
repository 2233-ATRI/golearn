package main

import "fmt"

func traversalstring() {
	s := "wuqihao"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()
}

func main() {
	traversalstring()

}
