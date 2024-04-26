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

import (
	"fmt"
)

func main() {
	var lapW1, lapH1, lapW2, lapH2 int // Range:
	fmt.Scan(&lapW1, &lapH1, &lapW2, &lapH2)
	fmt.Println(tableParams(lapW1, lapH1, lapW2, lapH2))
	//fmt.Println(tableParams(10, 2, 2, 10))
}

func tableParams(a, b, c, d int) (tbW, tbH int) {
	m := [4]int{}
	n := [4]int{}

	if a > d {
		m[0] = a
	} else {
		m[0] = d
	}
	n[0] = b + c

	if a > c {
		m[1] = a
	} else {
		m[1] = c
	}
	n[1] = b + d

	if b > c {
		m[2] = b
	} else {
		m[2] = c
	}
	n[2] = a + d

	if b > d {
		m[3] = b
	} else {
		m[3] = d
	}
	n[3] = a + c

	min := 9223372036854775807 // math.MaxInt

	var height, width int
	for i := 0; i < 4; i++ {
		if m[i]*n[i] < min {
			min = m[i] * n[i]
			height = m[i]
			width = n[i]
		}
	}
	return height, width
}
