package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	fre := make(map[rune]int)
	for _, chars := range s {
		fre[chars]++
	}
	deleted := 0
	oddcount := 0
	for _, count := range fre {
		if count%2 != 0 {
			oddcount++
		}
	}
	if oddcount > 0 {
		deleted = oddcount - 1
	}
	var half []rune
	for chars := 'a'; chars <= 'z'; chars++ {
		if count, exists := fre[chars]; exists {
			for i := 0; i < count/2; i++ {
				half = append(half, chars)
			}
		}
	}

	sort.Slice(half, func(i, j int) bool {
		return half[i] < half[j]
	})

	var result []rune
	for _, char := range half {
		result = append(result, char)
	}
	midd := rune(0)

	for chars := 'a'; chars <= 'z'; chars++ {
		if fre[chars]%2 != 0 {
			midd = chars
			break
		}
	}

	if midd != 0 {
		result = append(result, midd)
	}
	for i := len(half) - 1; i >= 0; i-- {
		result = append(result, half[i])
	}
	fmt.Println(deleted)
	fmt.Println(string(result))

} //统计所有字符出现的次数，当为奇数则需要减至偶数，当该字符是回文的元素时，后面应该有一个对应的元素，当没有的话需要对比大小,思路问题，是重新排序不是保持原本的
