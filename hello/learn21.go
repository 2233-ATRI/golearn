package main

import "fmt"

// Writer 接口定义
type Writer interface {
	Writebook()
}

// Reader 接口定义
type Reader interface {
	Readbook()
}

// Book 结构体定义
type Book struct {
}

// Book 实现 Reader 接口的方法
func (this *Book) Readbook() {
	fmt.Println("Readbook")
}

// Book 实现 Writer 接口的方法
func (this *Book) Writebook() {
	fmt.Println("Writebook")
}

func main() {
	// 创建 Book 实例
	b := &Book{}

	// 将 Book 赋值给 Reader 接口
	var r Reader
	r = b
	r.Readbook()

	fmt.Println()

	// 类型断言，将 Reader 转为 Writer
	var w Writer
	w = r.(Writer) // 此时 Book 实现了 Writer 接口，可以安全断言
	w.Writebook()
}
