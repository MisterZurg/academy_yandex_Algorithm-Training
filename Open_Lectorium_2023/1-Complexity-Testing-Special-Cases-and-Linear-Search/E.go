package main

import (
	"bufio"
	"fmt"
	"os"
)

func sol(s string) string {
	if len(s) == 1 {
		return ""
	}

	middle := len(s) / 2
	flag := false
	var ans string

	for i := 0; i < middle; i++ {
		if s[i] != 'a' {
			ans = s[:i] + "a" + s[i+1:]
			flag = true
			break
		}
	}

	if flag {
		return ans
	} else {
		return s[:len(s)-1] + "b"
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(sol(s))
}
