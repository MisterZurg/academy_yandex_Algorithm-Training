/*
В школе решили на один прямоугольный стол поставить два прямоугольных ноутбука.
Ноутбуки нужно поставить так, чтобы их стороны были параллельны сторонам стола.
Определите, какие размеры должен иметь стол, чтобы оба ноутбука на него поместились,
и площадь стола была минимальна.

Формат ввода
Вводится четыре натуральных числа, первые два задают размеры одного ноутбука,
а следующие два — размеры второго. Числа не превышают 1000.

Формат вывода
Выведите два числа — размеры стола. Если возможно несколько ответов,
выведите любой из них (но только один).
*/
package main

import "fmt"

func main() {
	//var lapW1, lapH1, lapW2, lapH2 int8 // Range:
	//fmt.Scan(&lapW1, &lapH1, &lapW2, &lapH2)
	fmt.Println(tableParams(10, 2, 2, 10))
}

func tableParams(lapW1, lapH1, lapW2, lapH2 int8) (tbW, tbH int8) {
	width := make([]int8, 4)
	height := [4]int8{
		lapH1 + lapW2,
		lapH1 + lapH2,
		lapW1 + lapW2,
		lapW1 + lapH2,
	}
	if lapW1 < lapH2 {
		width[0] = lapW1
	} else {
		width[0] = lapH2
	}

	if lapW1 < lapW2 {
		width[1] = lapW1
	} else {
		width[1] = lapW2
	}

	if lapH1 < lapW2 {
		width[2] = lapH1
	} else {
		width[2] = lapW2
	}

	if lapH1 < lapH2 {
		width[3] = lapH1
	} else {
		width[3] = lapH2
	}
	min := width[0] * height[0]
	tbW = width[0]
	tbH = height[0]
	for i := 1; i < 4; i++ {
		if min > width[i]*height[i] {
			min = width[i] * height[i]
			tbW = width[i]
			tbH = height[i]
		}
	}
	return tbW, tbH
}
