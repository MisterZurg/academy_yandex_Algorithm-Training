package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sequence := make([]int, n)
	for i := range sequence {
		fmt.Scan(&sequence[i])
	}

	token := true
	for i := 0; i < n-1; i++ {
		if sequence[i] > sequence[i+1] {
			token = false
		}
	}

	if token {
		fmt.Println(sequence[n-1] - sequence[0])
	} else {
		fmt.Println(-1)
	}
}
