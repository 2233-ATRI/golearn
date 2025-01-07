package main

import (
	"fmt"
	//"html/template"
	"time"
)

func main() {

	go func(a int, b int) bool {
		fmt.Println(a)
		fmt.Println(b)
		return false
	}(10, 20)

	//go func() {
	//	defer fmt.Println("a world")
	//
	//	func() {
	//		defer fmt.Println("b world")
	//		runtime.Goexit()
	//		fmt.Println("a hello")
	//	}()
	//
	//	fmt.Println("Hello World")
	//}()
	for {
		time.Sleep(time.Second)
	}
}
