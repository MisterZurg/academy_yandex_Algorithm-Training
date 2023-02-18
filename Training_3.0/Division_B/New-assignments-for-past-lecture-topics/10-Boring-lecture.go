package main

import (
	"fmt"
)

// https://site.ada.edu.az/~medv/acm/Docs%20e-olimp/Volume%2049/4859.htm

func main() {
	var word string
	fmt.Scan(&word)
	// fmt.Println("SUS")
	// const MAX = 100010
	m := [27]int{}
	for i := 0; i < len(word); i++ {
		m[word[i]-'a'] += 1 * (i + 1) * (len(word) - i)
	}

	for i := 0; i < 26; i++ {
		if m[i] > 0 {
			fmt.Printf("%s: %d\n", string(i+'a'), m[i])
		}
	}
}
