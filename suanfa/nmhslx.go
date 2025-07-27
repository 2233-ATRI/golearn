package main

import "fmt"

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}
func main() {
	//fn := func() { fmt.Println("wqh") }
	//fn()
	//
	//fns := [](func(x int) int){
	//	func(x int) int { return x * x },
	//	func(x int) int { return x * x * x },
	//}
	//fmt.Println(fns[0](100))
	//d := struct {
	//	fn func() string
	//	bn func() string
	//}{
	//	fn: func() string { return "wqh" },
	//	bn: func() string { return "wqh" },
	//}
	//fmt.Println(d.fn())
	//
	//fv := make(chan func() string, 2)
	//fv <- func() string { return "wqh" }
	//
	//fmt.Println((<-fv)())
	c := a()
	c()
	c()
	c()
}
