package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	q *list.List
}

func NewQueue() *Queue {
	return &Queue{
		q: list.New(),
	}
}

// Enqueue
func (q *Queue) Enqueue(item interface{}) {
	q.q.PushBack(item)
}

// Dequue
func (q *Queue) PopFront() (int, error) {
	if q.isEmpty() {
		return 0, fmt.Errorf("error")
	}
	front := (*q).q.Front()
	// This will remove the allocated memory and avoid memory leaks
	(*q).q.Remove(front)
	return front.Value.(int), nil
}

// Dequue
func (q *Queue) getFront() (int, error) {
	if q.isEmpty() {
		return 0, fmt.Errorf("error")
	}

	return (*q).q.Front().Value.(int), nil
}

func (q *Queue) isEmpty() bool {
	return (*q).q.Len() == 0
}

func (q *Queue) Size() int {
	return (*q).q.Len()
}

func (q *Queue) Clear() {
	*q = *NewQueue()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	q := NewQueue()
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		switch words[0] {
		case "push":
			n, _ := strconv.Atoi(words[1])
			q.Enqueue(n)
			fmt.Println("ok")
		case "pop":
			top, err := q.PopFront()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(top)
			}
		case "front":
			top, err := q.getFront()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(top)
			}
		case "size":
			fmt.Println(q.Size())
		case "clear":
			q.Clear()
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			return
		}
	}
}
