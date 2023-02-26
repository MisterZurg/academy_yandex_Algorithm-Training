package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	// Избавляемся от повторов
	diegoCardSet := make(map[int]bool)
	var collection []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		if _, ok := diegoCardSet[num]; !ok {
			diegoCardSet[num] = true
			collection = append(collection, num)
		}
	}

	// Сортируем - получаем возрастающую последовательность
	sort.Ints(collection)

	var k int
	fmt.Scan(&k)
	buyers := make([]int, k)

	for i := range buyers {
		fmt.Scan(&buyers[i])
	}

	// fmt.Println(">", collection, buyers)

	// Используем бинпоиск
	for _, card := range buyers {
		fmt.Println(BinarySearch(collection, card))
	}
}

func BinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
