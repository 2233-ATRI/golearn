package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	name string `info:"name" doc:"myname"`
	sex  string `info:"sex"`
}

func findtag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagstring := t.Field(i).Tag.Get("info")
		fmt.Println(tagstring)
	}
}

func main() {
	// 创建一个 resume 实例
	r := resume{name: "John Doe"}
	findtag(&r)
	// 输出结构体内容
	fmt.Printf("Resume: %+v\n", r)
}
