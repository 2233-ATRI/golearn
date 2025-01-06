package main

import (
	"fmt"
)

func myfunc(arg interface{}) {
	fmt.Println(arg)

	value, ok := arg.(string)
	if !ok {
		fmt.Println("error!")
	} else {
		fmt.Println(value)
	}
}

//type Book struct {
//	auth string
//}

//func main() {
//	book := Book{"go"}
//	myfunc(book.auth)
//}
