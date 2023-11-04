package main

import "fmt"

// LeetCode 242. Valid Anagram
func isAnagram(f, s string) string {
	fMp := make(map[rune]int)
	sMp := make(map[rune]int)
	for _, lt := range f {
		fMp[lt]++
	}

	for _, lt := range s {
		sMp[lt]++
	}

	for k, v := range fMp {
		if sMp[k] != v {
			return "NO"
		}
	}

	for k, v := range sMp {
		if fMp[k] != v {
			return "NO"
		}
	}

	return "YES"
}

func main() {
	var first, second string
	fmt.Scan(&first, &second)
	fmt.Println(isAnagram(first, second))
}
