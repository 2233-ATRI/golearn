package main

import "fmt"

func changevaule(p int) {
	p = 100
}
func changevaule1(p *int) {
	*p = 100
}
func main() {
	var a int = 10
	changevaule(a)
	fmt.Println(a)
	fmt.Println(&a)
	changevaule1(&a)
	fmt.Println(a)
	fmt.Println(&a)

}
