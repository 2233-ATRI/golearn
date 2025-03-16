package main

import "fmt"

func main() {
	var l0, r0, lp, rp int
	var o, p string
	fmt.Scan(&l0, &r0)
	fmt.Scan(&lp, &rp)
	fmt.Scan(&o, &p)

	var oo int
	if o == "L" {
		oo = r0
	} else {
		oo = l0
	}

	var pp int
	if p == "L" {
		pp = rp
	} else {
		pp = lp
	}

	if pp == oo {
		fmt.Println("D")
	} else if (oo == 0 && pp == 2) || oo == 2 && pp == 5 || oo == 5 && pp == 0 {
		fmt.Println("O")
	} else {
		fmt.Println("P")
	}
}
