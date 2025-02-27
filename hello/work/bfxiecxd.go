package main

import (
	"fmt"
	"time"
)

//func HU() {
//	time.Sleep(2 * time.Second)
//	fmt.Println("HU")
//}
//
//func main() {
//	go HU()
//	fmt.Println("main")
//	for {
//		time.Sleep(1000)
//	}
//}
func Hu(ch chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("huu")
	ch <- 1000
}

func main() {
	ch := make(chan int)
	go Hu(ch)
	fmt.Println("start")
	v := <-ch
	fmt.Println("end", v)

}
