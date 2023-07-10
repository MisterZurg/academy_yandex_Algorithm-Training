package main

import (
	"fmt"
	"math"
)

func main() {
	food := parseInput()
	// for _, row := range food {
	//	fmt.Println(row)
	// }
	fmt.Println(calcCosts(food))
}

func calcCosts(food [][]int) int {
	N, M := len(food), len(food[0])
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
		for j := 0; j < M; j++ {
			// Наполняем бесконечностью 0 столбец | строку
			dp[i][j] = math.MaxInt // 1e9
			if i == 0 && j == 0 {
				dp[i][j] = food[i][j]
			} else {
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j])
				}
				if j > 0 {
					dp[i][j] = min(dp[i][j], dp[i][j-1])
				}
				dp[i][j] += food[i][j]
			}
		}
	}

	return dp[N-1][M-1]
}

func parseInput() [][]int {
	var n, m int
	fmt.Scan(&n, &m)

	food := make([][]int, n)
	for i := range food {
		food[i] = make([]int, m)
	}

	// Заполняем с первой тстроки
	for i := 0; i < len(food); i++ {
		for j := 0; j < len(food[0]); j++ {
			fmt.Scan(&food[i][j])
		}
	}
	return food
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
