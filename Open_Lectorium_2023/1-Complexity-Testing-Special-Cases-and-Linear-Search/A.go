package main

import "fmt"

func slow(n int) int {
	nowinrow := 1
	rows := 0
	for n >= nowinrow {
		n -= nowinrow
		rows += 1
		nowinrow += 1
	}
	return rows
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(slow(n))
}
