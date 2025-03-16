package main

import "fmt"

func ab(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var x1, y1, x2, y2 int
		fmt.Scan(&x1, &y1, &x2, &y2)

		if x1 == x2 && y1 == y2 {
			fmt.Println(0)
		} else if x1 == x2 || y1 == y2 || ab(x1-x2) == ab(y1-y2) {
			fmt.Println(1)
		} else {
			fmt.Println(2)
		}

	}

}
