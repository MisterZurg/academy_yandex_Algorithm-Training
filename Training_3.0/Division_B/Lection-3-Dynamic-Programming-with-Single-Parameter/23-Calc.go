package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n == 1 {
		fmt.Printf("0\n1")
		return
	}

	dp := make([]int, n+1)
	prev := make([]int, n+1)

	for i := 2; i <= n; i++ {
		minOp := dp[i-1]
		prev[i] = i - 1
		if i%2 == 0 {
			minOp = min(minOp, dp[i/2])
			prev[i] = i / 2
		}
		if i%3 == 0 {
			minOp = min(minOp, dp[i/3])
			prev[i] = i / 3
		}
		dp[i] = minOp + 1
	}

	// Show operations number
	fmt.Println(dp[n])

	// Prepare certificate number
	i := n
	ans := ""
	for i > 1 {
		if dp[i] == dp[i-1]+1 {
			i--
			ans = strconv.Itoa(i) + " " + ans
			continue
		}
		if i%2 == 0 && dp[i] == dp[i/2]+1 {
			i /= 2
			ans = strconv.Itoa(i) + " " + ans
			continue
		}
		i /= 3
		ans = strconv.Itoa(i) + " " + ans
	}
	ans += strconv.Itoa(n)
	// Show certificate
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
