package main

import "fmt"

func testyc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(1, err.(string))
		}
	}()
	panic("wqh")
}

func main() {
	testyc()
}
