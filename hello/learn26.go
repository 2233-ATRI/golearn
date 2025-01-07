package main

import (
	"fmt"
	"time"
)

// 子协程
func nowtesk() {
	i := 0
	for {
		i++
		fmt.Printf("func i=%d\n", i)
		time.Sleep(time.Second)
	}
}

// 主协程
func main() {
	go nowtesk()

	fmt.Println("main goroutine")
	//i := 0
	//for {
	//	i++
	//	fmt.Printf("main %d\n", i)
	//	time.Sleep(time.Second)
	//}

}
