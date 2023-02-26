package main

import "fmt"

// Pattern : Пересечение множеств
// Московская командная олимпиада по программированию. tho
// Лига b. 11 октября 2009 года
// TODO: реализовать бинарный поиск для получения сложности O(N*log(N))

type Segment struct {
	left      int
	right     int
	isPresent bool
}

func main() {
	var M, N int
	fmt.Scan(&M, &N)

	if N == 0 {
		fmt.Println(0)
		return
	}

	segments := make([]Segment, N)
	cnt := 0 // Кол-во установленных ОС

	var left, right int
	fmt.Scan(&left, &right)
	segments[0].left = left
	segments[0].right = right
	segments[0].isPresent = true
	cnt++

	for i := 1; i < N; i++ {
		fmt.Scan(&left, &right)
		segments[i].left = left
		segments[i].right = right
		segments[i].isPresent = true
		cnt++

		for j := 0; j < i; j++ {
			// fmt.Println(segments[i], "vs", segments[j])
			if segments[j].isPresent &&
				(segments[i].left <= segments[j].right && segments[j].left <= segments[i].right) {
				segments[j].isPresent = false
				cnt--
			}
		}
	}
	// fmt.Println(segments)
	fmt.Println(cnt)
}
