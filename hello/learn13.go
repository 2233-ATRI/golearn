package main

import "fmt"

type MyInt int

type book struct {
	name   string
	index  int16
	writer string
}

func changebook(book book) {
	book.index = 100
}
func changebook2(book *book) {
	book.index = 200
}

func main() {
	var a MyInt = 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println()
	var book1 book
	//var book2 book
	book1.name = "go"
	book1.writer = "wu"
	changebook(book1)
	fmt.Println(book1)

	changebook2(&book1)
	fmt.Println(book1)
}
