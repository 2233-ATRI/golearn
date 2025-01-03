package main

import "fmt"

const (
	//iota
	beijing = 10 * iota
	xian
	baoji
)

func lengh(string2 string) int {

	fmt.Println("lengh", string2)
	return lengh(string2)
}
func foo(a string, b int) (int, int) {
	fmt.Println(b)
	var c int = lengh(a)
	fmt.Println(c)
	return c, b
}

func main() {
	const length int = 10
	var age int16 = 10
	fmt.Println(age)

}
