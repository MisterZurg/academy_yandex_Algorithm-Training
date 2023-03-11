package main

import "fmt"

func main() {
	var i, n, k int
	fmt.Scan(&n, &k)
	// Обобщенные числа Фиббоначи
	dp := make([]int, 35) // По
	// БАЗА
	dp[0], dp[1], dp[2] = 0, 1, 1 // Перед Лестницей и I & II - Ступеньки
	// Считаем сколькими способами мы можем добраться до ступеньки
	for i = 3; i <= k; i++ {
		dp[i] = 2 * dp[i-1]
	}

	for i <= n {
		dp[i] = 2*dp[i-1] - dp[i-k-1]
		i++
	}

	fmt.Println(dp[n])
}
