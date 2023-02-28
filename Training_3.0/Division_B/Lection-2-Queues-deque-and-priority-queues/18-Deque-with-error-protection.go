package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dequeue struct {
	dq *list.List
}

func NewDequeue() *Dequeue {
	return &Dequeue{
		dq: list.New(),
	}
}

// PushFront - Добавить (положить) в начало дека новый элемент. Программа должна вывести ok.
func (dq *Dequeue) PushFront(item interface{}) {
	dq.dq.PushFront(item)
}

// PushBack - Добавить (положить) в конец дека новый элемент. Программа должна вывести ok.
func (dq *Dequeue) PushBack(item interface{}) {
	dq.dq.PushBack(item)
}

// PopFront - Извлечь из дека первый элемент. Программа должна вывести его значение.
func (dq *Dequeue) PopFront() (interface{}, error) {
	if dq.isEmpty() {
		return 0, fmt.Errorf("error")
	}
	front := (*dq).dq.Front()
	// This will remove the allocated memory and avoid memory leaks
	(*dq).dq.Remove(front)
	return front.Value, nil
}

// PopBack - Извлечь из дека последний элемент. Программа должна вывести его значение.
func (dq *Dequeue) PopBack() (interface{}, error) {
	if dq.isEmpty() {
		return 0, fmt.Errorf("error")
	}
	back := (*dq).dq.Back()
	// This will remove the allocated memory and avoid memory leaks
	(*dq).dq.Remove(back)
	return back.Value, nil
}

// front - Узнать значение первого элемента (не удаляя его). Программа должна вывести его значение.
func (dq *Dequeue) Front() (interface{}, error) {
	if dq.isEmpty() {
		return 0, fmt.Errorf("error")
	}
	front := (*dq).dq.Front()
	return front.Value, nil
}

// back - Узнать значение последнего элемента (не удаляя его). Программа должна вывести его значение.
func (dq *Dequeue) Back() (interface{}, error) {
	if dq.isEmpty() {
		return 0, fmt.Errorf("error")
	}
	back := (*dq).dq.Back()
	// This will remove the allocated memory and avoid memory leaks
	return back.Value, nil
}

// size - Вывести количество элементов в деке.
func (dq *Dequeue) Size() int {
	return dq.dq.Len()
}

// clear - Очистить дек (удалить из него все элементы) и вывести ok.
func (dq *Dequeue) Clear() {
	*dq = *NewDequeue()
}

func (dq *Dequeue) isEmpty() bool {
	return dq.dq.Len() == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	dq := NewDequeue()
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		switch words[0] {
		case "push_front":
			n, _ := strconv.Atoi(words[1])
			dq.PushFront(n)
			fmt.Println("ok")
		case "push_back":
			n, _ := strconv.Atoi(words[1])
			dq.PushBack(n)
			fmt.Println("ok")
		case "pop_front":
			top, err := dq.PopFront()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(top) // Если нам нужно тайп кастить ЖОСКО top.(int)
			}
		case "pop_back":
			back, err := dq.PopBack()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(back) // Если нам нужно тайп кастить ЖОСКО back.(int)
			}
		case "front":
			top, err := dq.Front()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(top)
			}
		case "back":
			top, err := dq.Back()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(top)
			}
		case "size":
			fmt.Println(dq.Size())
		case "clear":
			dq.Clear()
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			return
		}
	}
}
