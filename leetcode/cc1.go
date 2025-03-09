// package main
//
// import "fmt"
//
// const modd = 998244353
//
//	func main() {
//		var T, n int
//		fmt.Scan(&T)
//		for t := 0; t < T; t++ {
//			fmt.Scan(&n)
//			p := make([]int, n)
//			for i := 0; i < n; i++ {
//				fmt.Scan(&p[i])
//			}
//			origin := count(p)
//			dp := make([]int, n+1)
//			dp[0] = 1
//
//			for i := 1; i <= n; i++ {
//				dp[i] = (dp[i] + dp[i-1]) % modd
//			}
//			fmt.Println(dp[n])
//		}
//	}
//
//	func count(arr []int) int {
//		if len(arr) <= 1 {
//			return 0
//		}
//		mid := len(arr) / 2
//		left := count(arr[:mid])
//		right := count(arr[mid:])
//		inver := count(left) + count(right)
//
//		merge := make([]int, 0, len(arr))
//
//		i, j := 0, 0
//		for i < len(left) && j < len(right) {
//			if left[i] <= right[j] {
//				merge = append(merge, left[i])
//				i++
//			} else {
//				merge = append(merge, right[j])
//				j++
//				inver += len(left) - i
//			}
//		}
//
// }
package main
