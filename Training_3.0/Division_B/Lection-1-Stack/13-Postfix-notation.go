package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// LeetCode Medium 150. Evaluate Reverse Polish Notation

func main() {
	fmt.Println(evalRPN())
}

var operations = map[string]func(x, y int) int{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
	"*": func(x, y int) int { return x * y },
	"/": func(x, y int) int { return x / y },
}

type StackNode struct {
	val  int
	next *StackNode
}

type Stack struct {
	top *StackNode
}

func (s *Stack) Pop() int {
	val := s.top.val
	s.top = s.top.next

	return val
}

func (s *Stack) Push(val int) {
	newNode := &StackNode{val: val, next: s.top}
	s.top = newNode
}

func evalRPN() int {
	stack := &Stack{}
	reader := bufio.NewReader(os.Stdin)
	for {
		tk, _, err := reader.ReadRune()
		if err == io.EOF || tk == '\n' {
			break
		}
		if tk == ' ' {
			continue
		}
		t := string(tk)
		switch t {
		case "+", "-", "*", "/":
			rhs, lhs := stack.Pop(), stack.Pop()
			stack.Push(evaluate(lhs, rhs, t))
		default:
			i, _ := strconv.Atoi(t)
			stack.Push(i)
		}
	}

	return stack.Pop()
}

func evaluate(lhs, rhs int, operation string) int {
	switch operation {
	case "+":
		return lhs + rhs
	case "-":
		return lhs - rhs
	case "*":
		return lhs * rhs
	case "/":
		return lhs / rhs
	}

	return -1
}
