package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	storage []int
}

// push n Добавить в стек число n (значение n задается после команды). Программа должна вывести ok.
func (s *Stack) Push(n int) {
	s.storage = append(s.storage, n)
}

// pop Удалить из стека последний элемент. Программа должна вывести его значение.
func (s *Stack) Pop() (int, error) {
	last := len(s.storage) - 1
	if last <= -1 { // check the size
		return 0, errors.New("error") // and return error
	}

	value := s.storage[last]     // save the value
	s.storage = s.storage[:last] // remove the last element

	return value, nil // return saved value and nil error
}

// back Программа должна вывести значение последнего элемента, не удаляя его из стека.
func (s *Stack) Back() (int, error) {
	last := len(s.storage) - 1
	if last <= -1 { // check the size
		return 0, errors.New("error") // and return error
	}

	return s.storage[last], nil // return saved value and nil error
}

// size Программа должна вывести количество элементов в стеке.
func (s *Stack) Size() int {
	return len(s.storage)
}

// clear Программа должна очистить стек и вывести ok.
func (s *Stack) Clear() {
	s.storage = []int{}
}

// exit  Программа должна вывести bye и завершить работу.

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	st := Stack{}
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		switch words[0] {
		case "push":
			n, _ := strconv.Atoi(words[1])
			st.Push(n)
			fmt.Println("ok")
		case "pop":
			top, err := st.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(top)
			}
		case "back":
			top, err := st.Back()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(top)
			}
		case "size":
			fmt.Println(st.Size())
		case "clear":
			st.Clear()
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			return
		}
	}
}
