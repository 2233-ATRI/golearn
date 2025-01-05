package main

import "fmt"

type Human struct {
	Name string
	Age  int
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("eat something")
}

func (this *Human) WALK() {
	fmt.Println("walk something")
}

type Superman struct {
	Human //继承human
	Level int
}

func (this *Superman) Eat() {
	fmt.Println("superman eat something")
}

func (this *Superman) fly() {
	fmt.Println("superman walk something")
}

func main() {
	h := Human{"Bob", 25, "male"}
	h.Eat()
	h.WALK()
	fmt.Println()
	s := Superman{Human{"wu", 25, "male"}, 5}
	s.Eat()
	s.WALK()
	s.fly()
}
