package main

import (
	"fmt"
)

func changstring() {
	s1 := "hello"
	bytes1 := []byte(s1)
	bytes1[0] = 'H'
	fmt.Println(string(bytes1))
}

func main() {
	changstring()
	var arr0 [5]int = [5]int{1, 2, 3}
	//var arr1 [5]int = [5]int{1, 2, 3} append 不能作用再数组
	fmt.Println(arr0)
	slice1 := make([]int, 10, 20)
	slice2 := make([]int, 10, 20)
	slice1[3] = 1
	b := slice1[2:11]
	fmt.Println(slice1)
	fmt.Println(b)
	d := [5]struct {
		x int
	}{}
	var a int = len(d)
	fmt.Println(a)
	ace := append(slice1, slice2...)
	fmt.Println(ace)
	str := "hello world"
	s1 := str[0:5]
	fmt.Println(s1, len(s1))

	s2 := str[5:]
	fmt.Println(s2, len(s2))
}
