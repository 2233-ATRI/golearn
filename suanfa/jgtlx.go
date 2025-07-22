package main

import "fmt"

type str struct {
	name, city string
	age        int
}

func main() {
	var s str
	s.name = "wuqihao"
	s.age = 18
	s.city = "beijing"
	fmt.Println(s)
	var user struct {
		name string
		age  int
	}
	user.name = "wuqihao"
	user.age = 18
	fmt.Println(user)
}
