package main

import "fmt"

func deferword() int {
	fmt.Println("hello defer")
	return 0
}

func returnword() int {
	fmt.Println("hello return")
	return 0
}

func returndefer() int {
	defer deferword()
	return returnword()
}

func main() {
	defer fmt.Println("world1")
	defer fmt.Println("world2")
	fmt.Println("hello")
	fmt.Println("hello2")
	returndefer()
}
