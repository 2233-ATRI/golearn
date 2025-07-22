package main

import "fmt"

func mac1(x int) {
	x = 100
}

func mac2(x *int) {
	*x = 100
}

func main() {
	//a := 10
	//b := &a
	//fmt.Printf("a=%d par =%p,%d", a, &a, b)
	//fmt.Println()
	//mac1(a)
	//fmt.Printf("a=%d par =%p,%d", a, &a, b)
	//fmt.Println()
	//mac2(b)
	//fmt.Printf("a=%d par =%p,%d", a, &a, b)
	//var s *string
	//fmt.Println(s)
	//if s == nil {
	//	fmt.Println("kong")
	//} else {
	//	fmt.Println(s)
	//}

	var a *int
	a = new(int)
	*a = 10

	fmt.Println(*a) //new怼数值有用处，make更多是用在slice等操作
	var b map[string]int
	b = make(map[string]int, 10)
	b["a"] = 10
	b["b"] = 20
	fmt.Println(b)
	//程序定义一个int变量num的地址并打印
	//	将num的地址赋给指针ptr，并通过ptr去修改num的值

	var num int
	var prt *int = &num

	fmt.Println(prt)
	num = 10
	*prt = 20
	fmt.Println(num)

}
