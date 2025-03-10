package main

import "fmt"

func count(n int, sequene []string) []int {
	result := make([]int, len(sequene))
	for i, seq := range sequene {
		if len(seq) == 0 {
			result[i] = 0
			continue
		}
		paragragh := make(map[rune]int)
		prev := rune(seq[0])
		paragragh[prev] = 1
		for _, ch := range seq[1:] {
			if ch != prev {
				paragragh[ch]++
				prev = ch
			}
		}
		count := 0
		for _, para := range paragragh {
			if para-1 >= n {
				count++
			}
		}
		result[i] = count

	}
	return result
}
func main() {
	var n, m int
	fmt.Scan(&n, &m)
	sequenec := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&sequenec[i])
	}
	results := count(n, sequenec)
	for _, result := range results {
		fmt.Println(result)
	}
}
