/*
Решите в целых числах уравнение:
sqrt(ax+b) = c
a, b, c – данные целые числа: найдите все решения или сообщите,
что решений в целых числах нет.

Формат ввода
Вводятся три числа a, b и c по одному в строке.

Формат вывода
Программа должна вывести все решения уравнения в порядке возрастания,
либо NO SOLUTION (заглавными буквами), если решений нет.
Если решений бесконечно много, вывести MANY SOLUTIONS.
*/
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	equationSolver(a, b, c)
}

func equationSolver(a, b, c int) {
	// Чтобы избавиться от знака корня в левой части уравнения,
	// возведем обе части в квадрат.
	// ax + b = c ^ 2
	// x := (c * c - b)  / a
	if c < 0 {
		fmt.Println("NO SOLUTION")
		return
	} else if a == b && b == c && c == 0 {
		fmt.Println("MANY SOLUTIONS")
		return
	} else {
		fmt.Println(c*c/a - b/a)
		return
	}
}
