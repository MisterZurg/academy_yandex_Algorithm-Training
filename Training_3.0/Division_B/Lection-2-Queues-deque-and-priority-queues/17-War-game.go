package main

import (
	"container/list"
	"fmt"
)

// Patterns: Stack, Deq

const MAX_ATTEMPTS = 1_000_000

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

func main() {
	first := NewQueue()
	second := NewQueue()
	for i := 0; i < 5; i++ {
		var card int
		fmt.Scan(&card)
		first.Enqueue(card)
	}

	for i := 0; i < 5; i++ {
		var card int
		fmt.Scan(&card)
		second.Enqueue(card)
	}
	var attempt = 0
	for ; attempt < MAX_ATTEMPTS && !first.isEmpty() && !second.isEmpty(); attempt++ {
		firstTop, _ := first.PopFront()
		secondTop, _ := second.PopFront()

		// "шестерка берет туза"
		if firstTop == 0 && secondTop == 9 {
			first.Enqueue(firstTop)
			first.Enqueue(secondTop)
			continue
		} else if firstTop == 9 && secondTop == 0 {
			second.Enqueue(firstTop)
			second.Enqueue(secondTop)
			continue
		}

		if firstTop > secondTop {
			first.Enqueue(firstTop)
			first.Enqueue(secondTop)
		} else {
			second.Enqueue(firstTop)
			second.Enqueue(secondTop)
		}
	}

	if attempt == MAX_ATTEMPTS {
		fmt.Println("botva")
	} else if first.isEmpty() {
		fmt.Println("second", attempt)
	} else {
		fmt.Println("first", attempt)
	}
}

// Непонятно почему код ниже не работает из-за, поэтому пришлось явно очередь писать
// interface conversion: interface {} is *list.Element, not int
//type Hand struct {
//	*list.List
//}
//
//func main() {
//	first := Hand{list.New()}
//	second := Hand{list.New()}
//
//	//  номера карт первого игрока
//	for i := 0; i < 5; i++ {
//		var card int
//		fmt.Scan(&card)
//		first.PushBack(card)
//	}
//
//	// номера карт второго игрока
//	for i := 0; i < 5; i++ {
//		var card int
//		fmt.Scan(&card)
//		second.PushBack(card)
//	}
//	// if stack empty that one loses
//	var attempt = 0
//	for ; attempt < MAX_ATTEMPTS && first.Len() > 0 && second.Len() > 0; attempt++ {
//		firstTop := first.Front()
//		secondTop := second.Front()
//		// fmt.Println(firstTop.Value.(int), secondTop.Value.(int))
//		first.Remove(firstTop)
//		second.Remove(secondTop)
//
//		// "шестерка берет туза"
//		if firstTop.Value == 6 && secondTop.Value == 11 {
//			first.PushBack(firstTop)
//			first.PushBack(secondTop)
//			continue
//		} else if firstTop.Value == 11 && secondTop.Value == 6 {
//			second.PushBack(firstTop)
//			second.PushBack(secondTop)
//			continue
//		}
//
//		topValue := firstTop.Value.(int)
//		secondTopValue := secondTop.Value.(int)
//
//		if topValue > secondTopValue {
//			first.PushBack(firstTop)
//			first.PushBack(secondTop)
//		} else {
//			second.PushBack(firstTop)
//			second.PushBack(secondTop)
//		}
//	}
//
//	if attempt == MAX_ATTEMPTS {
//		fmt.Println("botva")
//	} else if first.Len() == 0 {
//		fmt.Println("second", attempt)
//	} else {
//		fmt.Println("first", attempt)
//	}
//}
