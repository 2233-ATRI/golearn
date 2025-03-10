package main

import "fmt"

func combin(a, b int) int {
	if b > a-b {
		b = a - b
	}
	res := 1
	for i := 0; i < b; i++ {
		res = res * (a - i) / (i + 1)
	}
	return res
}

func main() {
	var k, m, n int
	fmt.Scan(&k)
	fmt.Scan(&m)
	fmt.Scan(&n)

	minstep := m + n

	if k < minstep {
		fmt.Println(0)
		return
	}

	pasth := combin(minstep, m)
	fmt.Println(pasth)

}
