package main

import (
	"fmt"
	"strings"
)

func txz(t int, testc [][]string) []string {
	result := make([]string, 0, t)
	for _, testc := range testc {
		var x, y int
		fmt.Scanf("%d %d	", &x, &y)

		moves := testc[1]

		for _, m := range moves {
			switch m {
			case 'W':
				y++
			case 'S':
				y--
			case 'A':
				x--
			case 'D':
				x++
			}
		}
		if x == 0 && y == 0 {
			result = append(result, "YES")

		} else {
			result = append(result, "NO")
		}
	}
	return result

}

func main() {
	var t int
	fmt.Scanf("%d", &t)

	testc := make([][]string, t)
	for i := 0; i < t; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		var moves string
		fmt.Scan(&moves)
		testc[i] = []string{fmt.Sprintf("%d %d", x, y), moves}

	}
	result := txz(t, testc)

	fmt.Println(strings.Join(result, "\n"))
}
