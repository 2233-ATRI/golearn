package main

import "fmt"

//稀疏数组

type valuenod struct {
	long   int
	height int
	value  int
}

func main() {
	var chassmap [11][11]int
	chassmap[1][2] = 1
	chassmap[2][3] = 2
	//for _, v := range chassmap {
	//	for _, w := range v {
	//		fmt.Printf("%d", w)
	//	}
	//	fmt.Println()
	//}

	var spr []valuenod
	//valuenod := valuenod{
	//	long:   11,
	//	height: 11,
	//	value:  0,
	//}
	for i, v := range chassmap {
		for j, w := range v {
			if w != 0 {
				valuenod := valuenod{
					long:   i,
					height: j,
					value:  w,
				}
				spr = append(spr, valuenod)
			}
		}
	}
	for i, v := range spr {
		fmt.Printf("%d\t%d\t%d\t%d\n", i, v.long, v.height, v.value)
	}
}
