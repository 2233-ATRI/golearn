package main

import "fmt"

func main() {
	c := make(chan int, 3) //三个管道缓冲区
	fmt.Println(len(c), cap(c))
	go func() {
		defer fmt.Println("a defer")
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("c", i)
			fmt.Println("cap", cap(c))
			fmt.Println("len", len(c))
		}
		fmt.Println()
	}()

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println(num)
	}
	fmt.Println("main")
}
