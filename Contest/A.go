package main

import "fmt"

func main() {
	keys1, keys2, text := parseInput()

	switches := 0
	//i := 0
	//j := 0
	for i := range text {
		hasSwitch := false
		// find letter in first
		for j := 1; j < len(keys1); j++ {
			if text[i-1] == keys2[i-1] && text[i] == keys1[i] {
				switches++
				hasSwitch = true
				break
			}
		}
		for j := 1; j < len(keys2); j++ {
			if !hasSwitch && text[i-1] == keys1[i-1] && text[i] == keys2[i] {
				switches++
				hasSwitch = true
				break
			}
		}
	}
	fmt.Println(switches)
}

func parseInput() ([]int, []int, []int) {
	var n int
	fmt.Scan(&n)

	keys1 := make([]int, n)
	for i := range keys1 {
		fmt.Scan(&keys1[i])
	}
	keys2 := make([]int, n)
	for i := range keys2 {
		fmt.Scan(&keys2[i])
	}

	var k int
	fmt.Scan(&n)

	text := make([]int, k)
	for i := range text {
		fmt.Scan(&text[i])
	}
	return keys1, keys2, text
}
