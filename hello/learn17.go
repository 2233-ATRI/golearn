package main

import "fmt"

type Animal interface {
	sleep()
	Getcolor() string
	Gettype() string
} //本质是一个指针

type Dog struct {
	color string
}

func (this *Dog) sleep() {
	fmt.Println("dog sleep")
}

func (this *Dog) Getcolor() string {
	return this.color
}

func (this *Dog) Gettype() string {
	return "Dog"
}

func showanimal(animal Animal) {
	animal.sleep()
	fmt.Println(animal.Getcolor())
	fmt.Println(animal.Gettype())
}

func main() {
	var animal Animal
	//animal = &Dog{"green"}

	animal.sleep()
	//showanimal(animal)
	//animal := Dog{"green"} // 创建一个 Dog 类型的实例
	//showanimal(&animal)    // 传递该实例的指针

}
