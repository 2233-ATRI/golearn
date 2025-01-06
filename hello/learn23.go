package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) call() {
	fmt.Println("call user")
	fmt.Printf("%T\n", this)
}

func Dofileandmethod(input interface{}) {

	inputtype := reflect.TypeOf(input)
	fmt.Println("inputtype:", inputtype)

	inputval := reflect.ValueOf(input)
	fmt.Println("inputval:", inputval)
	fmt.Println()
	for i := 0; i < inputval.NumField(); i++ {
		field := inputval.Field(i)
		value := inputval.Field(i).Interface()
		fmt.Println("field:", field, "value:", value)
	}
}

func main() {
	user := User{1, "wu", 18}
	Dofileandmethod(user)
}
