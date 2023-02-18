package main

import "fmt"

func main() {
	var k int
	var line string
	fmt.Scan(&k, &line)

	letters := []rune(line)
	w := [26]int{}
	mx, ans, j := 0, 0, 0

	for i := range letters {
		w[letters[i]-'a']++
		mx = max(mx, w[letters[i]-'a'])
		if i-j+1-mx > k {
			w[letters[j]-'a']--
			j++
		}
		ans = max(ans, min(i-j+1, mx+k))
	}
	fmt.Print(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
