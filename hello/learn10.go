package main

import "fmt"

func main() {
	var number = make([]int, 3, 5) //创建5个元素，初始化三个
	fmt.Println(len(number), cap(number))
	number = append(number, 1) //初始化多一个，但容量不变
	//cap满了之后不可以再使用append使用，但继续使用的话go会自动增加cap的大小
	number = append(number, 1)
	number = append(number, 1)
	fmt.Println(len(number), cap(number))

	//////////////////////
	s := []int{1, 2, 3, 4, 5}
	s1 := s[0:2] //左闭右开的一个区间，所以只有1，2
	fmt.Println(s1, len(s1), cap(s1))

	//////////////////////
	s2 := make([]int, 3)
	copy(s2, s1)
	fmt.Println(s2, len(s2), cap(s2))

}
