package main

import (
	"fmt"
	"math"
)

func main() {
	var n int // количество покупателей в очереди
	fmt.Scan(&n)

	buyers := make([][3]int, n+3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			buyers[i][j] = math.MaxInt
		}
	}

	for i := 3; i < len(buyers); i++ {
		fmt.Scan(&buyers[i][0], &buyers[i][1], &buyers[i][2])
	}

	fmt.Println(BuyingTickets(buyers))
}

// BuyingTickets() реалищует динамику с 57:00
// Тренировки по алгоритмам 3.0. Лекция 3: «Динамическое программирование с одним параметром»
func BuyingTickets(buyers [][3]int) int {
	// fmt.Println(buyers)
	a, b, c := 0, 1, 2
	dp := make([]int, len(buyers))

	for i := 3; i < len(buyers); i++ {
		dp[i] = min(dp[i-1]+buyers[i][a], min(dp[i-2]+buyers[i-1][b], dp[i-3]+buyers[i-2][c]))
		// fmt.Println(i-3, dp[i])
	}
	// fmt.Println(dp)
	return dp[len(buyers)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//2
//5 10 15
//2 10 15
