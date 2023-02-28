package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	storage []int
}

func (s *Stack) Push(n int) {
	s.storage = append(s.storage, n)
}

func (s *Stack) Pop() (int, error) {
	last := len(s.storage) - 1
	if last <= -1 { // check the size
		return 0, errors.New("error") // and return error
	}

	value := s.storage[last]     // save the value
	s.storage = s.storage[:last] // remove the last element

	return value, nil // return saved value and nil error
}

func (s *Stack) Peek() (int, error) {
	last := len(s.storage) - 1
	if last <= -1 { // check the size
		return 0, errors.New("error") // and return error
	}

	return s.storage[last], nil // return saved value and nil error
}

func (s *Stack) Size() int {
	return len(s.storage)
}

// clear Программа должна очистить стек и вывести ok.
func (s *Stack) Clear() {
	s.storage = []int{}
}

func main() {
	var n int
	fmt.Scan(&n)
	way2 := &Stack{}
	deadEnd := &Stack{}
	for i := 0; i < n; i++ {
		var wagonNum int
		fmt.Scan(&wagonNum)
		deadEnd.Push(wagonNum)

		val, _ := deadEnd.Peek()
		for deadEnd.Size() != 0 && val-1 == way2.Size() {
			poped, _ := deadEnd.Pop()
			way2.Push(poped)
			val, _ = deadEnd.Peek()
		}
	}

	if deadEnd.Size() == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
