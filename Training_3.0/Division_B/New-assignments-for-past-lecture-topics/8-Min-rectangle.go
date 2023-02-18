package main

import (
	"fmt"
)

/* Паттерн: Линейный поиск

Считывая данные, линейным поиском находим минимальные и максимальные абсциссы и ординаты точек.
Остается только вывести ответ на задачу - полученные координаты и будут координатами левого нижнего и верхнего правого углов искомого прямоугольника.
*/

func main() {
	var k int
	fmt.Scan(&k)

	var x, y, minX, minY, maxX, maxY int
	fmt.Scan(&x, &y)

	minX, minY, maxX, maxY = x, y, x, y
	for i := 1; i < k; i++ {
		fmt.Scan(&x, &y)
		minX = Min(x, minX)
		minY = Min(y, minY)
		maxX = Max(x, maxX)
		maxY = Max(y, maxY)
	}
	fmt.Println(minX, minY, maxX, maxY)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
