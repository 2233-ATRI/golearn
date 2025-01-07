package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		defer fmt.Println("go over")
		fmt.Println("gocoutine going")
		c <- 123 //123写入c当中
	}()
	num := <-c //从c当中读出数据并赋值
	fmt.Println(num)
	fmt.Println("main over")
}
