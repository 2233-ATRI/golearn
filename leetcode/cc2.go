package main

import "fmt"
import "sort"

func main() {
	var s, n, m int
	fmt.Scan(&s, &n, &m)

	oppoent := make([][]int, s)
	for i := 0; i < s; i++ {
		oppoent[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&oppoent[i][j])
		}
	}

	sorts := make([][]int, n)
	for i := 0; i < n; i++ {
		sorts[i] = make([]int, s)
		for j := 0; j < s; j++ {
			sorts[i][j] = oppoent[j][i]
		}
		sort.Ints(sorts[i])
	}

	dp := make([]int, m+1)
	for i := 0; i < n; i++ {
		for j := m; j >= 0; j-- {
			for b := 0; b <= j; b++ {
				score := 0
				thre := b / 2
				count := sort.Search(s, func(k int) bool {
					return sorts[i][k] > thre
				})
				score = (i + 1) * count

				if dp[j-b]+score > dp[j] {
					dp[j] = dp[j-b] + score
				}
			}
		}
	}
	fmt.Println(dp[m])

}
