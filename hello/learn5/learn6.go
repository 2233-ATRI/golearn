package main

import "fmt"

func changevaule(p int) {
	p = 100
}
func changevaule1(p *int) {
	*p = 100
}

func jiaohuan(a *int, b *int) (a1 int, b1 int) {
	var ac int = 0
	ac = *a
	*a = *b
	*b = ac
	return *a, *b
}

func main() {
	var a int = 10
	changevaule(a)
	fmt.Println(a)
	fmt.Println(&a)
	changevaule1(&a)
	fmt.Println(a)
	fmt.Println(&a)
	var b int = 20
	jiaohuan(&b, &a)
	fmt.Println(b, a)

}
