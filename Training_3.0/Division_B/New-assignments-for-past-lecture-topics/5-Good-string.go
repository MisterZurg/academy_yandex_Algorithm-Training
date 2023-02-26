package main

// Паттерны Рекурсия : Одномерные массивы

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	alphabet := [26]int{}

	for i := 0; i < num; i++ {
		fmt.Scan(&alphabet[i])
	}

	// Хорошесть строки - количество букв за которой идет следующая по алфавиту.
	// Максимизировать хорошесть
	goodness := 0
	for i := 1; i < num; i++ {
		goodness += min(alphabet[i-1], alphabet[i])
	}
	fmt.Println(goodness)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
