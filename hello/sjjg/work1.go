package main

import "fmt"

//稀疏数组

func main() {
	var chassmap [11][11]int
	chassmap[1][2] = 1
	chassmap[2][3] = 2
	for _, v := range chassmap {
		for _, w := range v {
			fmt.Printf("%d", w)
		}
		fmt.Println()
	}
}
