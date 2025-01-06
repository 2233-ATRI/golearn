package main

import "fmt"

func main() {
	var a string
	//pair<statictype:string,value:"">
	a = "abcd"
	//pair<statictype:string,value:"">
	var alltype interface{}
	alltype = a
	value, ok := alltype.(string)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("no")
	}
}
