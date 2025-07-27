package main

import "fmt"

//func test(x, y int, s string) (int, string) {
//	n := x + y
//	return n, fmt.Sprintf("%s%d", s, n)
//}
//
//func main() {
//	x := 1
//	y := 2
//	s := "wqh"
//	r, s := test(x, y, s)
//	fmt.Println(r, s)
//}

func test(fn func() int) int {
	return fn()
}

type formatfunc func(s string, x, y int) string

func format(fn formatfunc, s string, x, y int) string {
	return fn(s, x, y)
}

func swap(x, y *int) {
	var tamp int
	tamp = *x
	*x = *y
	*y = tamp
}

func tests(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf("%s%d", s, x)
}

func main() {
	s1 := test(func() int { return 1 })
	s2 := format(func(s string, x, y int) string {
		return s + fmt.Sprintf("%d", x+y)
	}, "wqh", 1, 2)
	fmt.Println(s1, s2)

	x := 10
	y := 20
	swap(&x, &y)
	fmt.Println(x, y)
	fmt.Println(tests("wqh", 1, 2, 3, 4, 5))
}
