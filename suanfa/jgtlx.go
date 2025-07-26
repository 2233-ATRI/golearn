package main

import "fmt"

type str struct {
	name, city string
	age        int
}

type person struct {
	name string
	city string
	age  int8
}

func newperson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}

}

func (p person) dream() {
	fmt.Printf("%s dream is learn c \n", p.name)
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

	var p2 = new(person)
	p2.name = "wuqihao"
	p2.age = 18
	p2.city = "beijing"
	fmt.Printf("%#v\n", p2)
	fmt.Println()
	var p4 person
	p4.name = "wuqihao"
	fmt.Printf("%#v\n", p4)
	fmt.Println()
	fmt.Println()
	type test struct {
		a int
		b int
		c int
		d int
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
	fmt.Println()
	var p5 = newperson("wuqihao", "beijing", 18)
	fmt.Println(p5)
	p8 := newperson("wqh", "beijing", 18)
	p8.dream()

}
