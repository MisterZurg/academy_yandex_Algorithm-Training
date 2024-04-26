package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, t int // количество скульптур, идеальное количество льда в скульптуре и оставшееся количество минут до наступления праздника.
	fmt.Scan(&n, &x, &t)

	iceInSculptures := make([]int, n)
	for i := range iceInSculptures {
		fmt.Scan(&iceInSculptures[i])
	}

	// Sliding window
	memo := make([][]int, t)
	for i := 0; i < t; i++ {
		avaibleTime := t
		memo[i] = make([]int, 0)
		for j, sc := range iceInSculptures {
			if sc+avaibleTime == x || sc-avaibleTime == x {
				memo[i] = append(memo[i], j+1)
				avaibleTime -= int(math.Abs(float64(x - sc)))
			}
		}
	}
	var maxIdx int
	for i := range memo {
		maxIdx = max(maxIdx, len(memo[i]))
	}
	for _, pos := range memo[maxIdx] {
		fmt.Printf("%d ", pos)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
