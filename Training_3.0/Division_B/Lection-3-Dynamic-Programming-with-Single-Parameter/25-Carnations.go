package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	carnations := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&carnations[i])
	}
	sort.Ints(carnations)

	dp := make([]int, n+1)
	dp[2] = carnations[2] - carnations[1]

	if n > 2 { // Если не проверить то RE on TC 16
		dp[3] = carnations[3] - carnations[1]

		for i := 4; i <= n; i++ {
			dp[i] = min(dp[i-2], dp[i-1]) + carnations[i] - carnations[i-1]
		}
	}
	fmt.Println(dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
