package main

import "fmt"

func changstring() {
	s1 := "hello"
	bytes1 := []byte(s1)
	bytes1[0] = 'H'
	fmt.Println(string(bytes1))
}

func main() {
	changstring()
	var arr0 [5]int = [5]int{1, 2, 3}
	fmt.Println(arr0)
}
