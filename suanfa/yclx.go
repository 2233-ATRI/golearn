package main

import "fmt"

//func testyc() {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println(1, err.(string))
//		}
//	}()
//	panic("wqh")
//}

func testyc() {
	defer func() {
		fmt.Println(recover())
	}()
	//defer func() {
	//	panic("wqh")
	//}()
	panic("wqh1")
}

func main() {
	testyc()
}
