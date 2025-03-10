package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getcount(woods []int, lengt int) int {
	dp := make([]int, lengt+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for _, wood := range woods {
		for l := wood; l <= lengt; l++ {
			if dp[l-wood] != math.MaxInt32 {
				dp[l] = min(dp[l], dp[l-wood]+1)
			}
		}
	}
	if dp[lengt] != math.MaxInt32 {
		return dp[lengt]
	}
	return -1
}

func parsewood(woodstr string) []int {
	parts := strings.Split(woodstr, ",")
	woods := make([]int, len(parts))
	for i, part := range parts {
		woods[i], _ = strconv.Atoi(part)
	}
	return woods
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	woodstr, _ := reader.ReadString('\n')
	woodstr = strings.TrimSpace(woodstr)
	woodstr = strings.TrimPrefix(woodstr, "[")
	woodstr = strings.TrimSuffix(woodstr, "]")
	woodspart := strings.Split(woodstr, ",")
	woods := make([]int, len(woodspart))
	for i, part := range woodspart {
		part = strings.TrimSpace(part)
		woods[i], _ = strconv.Atoi(part)
	}
	lengstr, _ := reader.ReadString('\n')
	lengstr = strings.TrimSpace(lengstr)
	length, _ := strconv.Atoi(lengstr)

	result := getcount(woods, length)
	fmt.Println(result)
}
