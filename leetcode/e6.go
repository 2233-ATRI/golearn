package main

import (
	"fmt"
)

func resever(s string) string {
	ru := []rune(s)
	for i, j := 0, len(ru)-1; i < j; i, j = i+1, j-1 {
		ru[i], ru[j] = ru[j], ru[i]
	}
	return string(ru)
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)

		t := ""
		p := 0
		for _, i := range s {
			if i >= '0' && i <= '9' {
				x := int(i - '0')
				if p == 0 {
					p = x
				} else {
					p = p*10 + x
				}
			} else {
				if p > 0 {
					if p > len(t) {
						p = p % len(t)
					}
					t = t[p:] + t[:p]
					p = 0
				}
				if i == 'R' {
					t = resever(t)
				} else {
					t += string(i)
				}
			}

		}
		fmt.Println(t)
	}
}
