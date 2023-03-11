package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(seqCnt(n))
}

// Последовательности где единицы 3 единицы не вместе
// 000	dp[i-1]
// 001	dp[i-2]
// 010	dp[i-1]
// 011	dp[i-3]
// 100	dp[i-1]
// 101	dp[i-2]
// 110	dp[i-1]
// НЕТ
// dp[i-1] = 4 | dp[i-2] = 2| dp[i-3] = 1
func seqCnt(n int) int {
	dp := make([]int, n+2)
	dp[0], dp[1], dp[2] = 1, 2, 4

	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}
