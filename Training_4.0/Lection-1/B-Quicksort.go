// TL 15 change reader
package main

import (
	"fmt"
)

// ImprovingQuickSort returns given sequence sorted in non-decreasing order.
func ImprovingQuickSort(sequence []int, l, r int) {
	// Here we have to implement 3-way partition
	if l >= r {
		return
	}
	m1, m2 := Partition3(sequence, l, r)

	ImprovingQuickSort(sequence, l, m1-1)
	ImprovingQuickSort(sequence, m2+1, r)
}

func Partition3(sequence []int, l, r int) (int, int) {
	pivot := sequence[l]
	m1 := l // We initiate m1 to be the part that is less than the pivot
	m2 := r // The part that is greater than the pivot
	i := l
	for i <= m2 {
		if sequence[i] < pivot {
			sequence[m1], sequence[i] = sequence[i], sequence[m1]
			m1++
			i++
		} else if sequence[i] > pivot {
			sequence[m2], sequence[i] = sequence[i], sequence[m2]
			m2--
		} else {
			i++
		}
	}

	return m1, m2
}

func main() {
	var n int
	var a []int

	fmt.Scan(&n)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	ImprovingQuickSort(a, 0, len(a)-1)
	// QuickSort(a, 0, len(a))
	for _, elem := range a {
		fmt.Printf("%d ", elem)
	}
}
