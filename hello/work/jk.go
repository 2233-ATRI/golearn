package main

import (
	"fmt"
	"reflect"
)

type A interface {
	println()
}

type B interface {
	println()
	Printf() int
}

type AAA struct {
	data string
}

func (aaa *AAA) println() {
	fmt.Println(aaa.data)
}

type BBB struct {
	data string
}

func (bbb *BBB) Printf() int {
	//TODO implement me
	panic("implement me")
}

func (bbb *BBB) println() {
	fmt.Println(bbb.data)
}
func (bbb *BBB) Print() int {
	fmt.Println(bbb.data)
	return 0
}

func main() {
	//var a interface{}
	//a = 2
	//fmt.Printf("%T\n ,%v", a, a)
	//
	//Print(a)
	//Print("aaa")
	//
	//v, ok := a.(int)
	//if ok {
	//	fmt.Println(v)
	//}
	//switch a.(type) {
	//case int:
	//	fmt.Println("a is type int")
	//case string:
	//	fmt.Println("a is type string")
	//default:
	//	fmt.Println("a not type found type")
	//}
	//t := reflect.TypeOf(a)
	//fmt.Printf("type %T\n", t)
	//var a interface{}
	//a = "acs"
	//fmt.Printf("a type: %T\n", a)
	var a A
	// 将具体的结构体赋予该变量
	a = &AAA{data: "i love you"}
	// 调用接口的方法
	a.println()
	// 断言类型
	if v, ok := a.(*AAA); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not a A1")
	}
	fmt.Println(reflect.TypeOf(a).String())
	// 将具体的结构体赋予该变量
	a = &BBB{data: "i love you"}
	// 调用接口的方法
	a.println()
	// 断言类型
	if v, ok := a.(*AAA); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not a A1")
	}
	fmt.Println(reflect.TypeOf(a).String())
	// 定义一个B接口类型的变量
	var b B
	//b = &A1Instance{Data: "i love you"} // 不是 B 类型
	b = &BBB{data: "i love you"}
	fmt.Println(b.Printf())
}
