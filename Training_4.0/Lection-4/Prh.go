package main

import "fmt"

func generateParenthesis(n int) []string {
	stack := ""
	res := []string{}
	backtrack(0, 0, n, stack, &res)
	return res
}

func backtrack(open, close, n int, stack string, res *[]string) {
	if open == close && close == n {
		*res = append(*res, stack)
		return
	}

	// If open < close; we can add either "(" or ")"
	if open < n {
		backtrack(open+1, close, n, stack+"(", res)
	}

	if open > close {
		backtrack(open, close+1, n, stack+")", res)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	for _, pr := range generateParenthesis(n) {
		fmt.Println(pr)
	}
}
