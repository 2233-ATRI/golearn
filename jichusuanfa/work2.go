package main

import "fmt"

type mentodutils struct {
	str int
}

func (ac *mentodutils) add() {
	var i = 0
	var j = 0
	for j < 8 {
		for i < 10 {
			fmt.Printf("*")
			i++
		}
		i = 0
		j++
		fmt.Println()
	}
	return
}

func (ac *mentodutils) add2(aa int, ab int) {
	var i = 0
	var j = 0
	for j < aa {
		for i < ab {
			fmt.Printf("*")
			i++
		}
		i = 0
		j++
		fmt.Println()
	}
	return
}

func (ac *mentodutils) add3(aa int, bb int) int {
	return aa * bb
}
func main() {
	var ac mentodutils
	var aa = 3
	var ab = 5
	//ac.add2(aa, ab)
	var d = ac.add3(aa, ab)
	fmt.Println(d)
	//ac.add()
}
