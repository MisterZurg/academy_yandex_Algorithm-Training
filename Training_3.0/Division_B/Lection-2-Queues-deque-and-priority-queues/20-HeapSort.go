package main

import "fmt"

// HeapSort firsly builds MaxHeap, and afterwards pop it to rebuild tree
func HeapSort(sl []int) {
	n := len(sl)
	for i := n/2 - 1; i >= 0; i-- {
		Heapify(sl, i, n)
	}

	for i := n - 1; i >= 0; i-- {
		sl[0], sl[i] = sl[i], sl[0]
		Heapify(sl, 0, i)
	}
}

func Heapify(arr []int, i, n int) {
	largest := i
	left := getLeft(i)
	right := getRight(i)

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		Heapify(arr, largest, n)
	}
}

func getLeft(i int) int {
	return 2*i + 1
}

func getRight(i int) int {
	return 2*i + 2
}

func parseInput(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}
	return nums
}

func main() {
	var broken []int
	var n int
	fmt.Scan(&n)
	broken = parseInput(n)
	HeapSort(broken)

	for _, elem := range broken {
		fmt.Printf("%d ", elem)
	}
}
