package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, q int
	fmt.Fscan(reader, &n, &q)

	a := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	pre := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + a[i]
	}

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		rest := a[l] - (pre[r] - pre[l]) - a[r]
		fmt.Fprintln(writer, rest)
	}
}
