package main

import (
	"fmt"
	"reflect"
)

func relectnum(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(arg))
}

func main() {
	var num float32 = 3.1415926
	relectnum(num)
}
