package main

import "fmt"

func add(a int, b int) int {
	c := a + b
	return c
}

func calc1(a, b int) (int, int) {
	sun := a + b
	avg := (a + b) / 2
	return sun, avg
}

func main() {
	var a, b int = 1, 2
	c := add(a, b)
	//d, e := calc(a, b)
	f, g := calc1(a, b)
	println(c, f, g)

	yt := func(a float64, b float64) float64 {
		return float64(a + b)
	}
	fmt.Println(yt(1.1, 2.2))

}
