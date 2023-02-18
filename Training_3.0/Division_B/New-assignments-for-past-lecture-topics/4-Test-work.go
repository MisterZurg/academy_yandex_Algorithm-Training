package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k, row, column int
	fmt.Scan(&n, &k, &row, &column)

	pos1 := (row-1)*2 + column - k
	pos2 := (row-1)*2 + column + k
	row1 := (pos1 + 1) / 2
	row2 := (pos2 + 1) / 2

	if pos1 > 0 && (pos2 > n || abs(row-row1) < abs(row-row2)) {
		fmt.Println(row1, 2-pos1%2)
	} else if pos2 <= n {
		fmt.Println(row2, 2-pos2%2)
	} else {
		fmt.Println(-1)
	}
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
