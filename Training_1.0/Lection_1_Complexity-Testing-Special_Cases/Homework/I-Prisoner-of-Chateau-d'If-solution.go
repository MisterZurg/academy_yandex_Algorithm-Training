/*
За многие годы заточения узник замка Иф проделал в стене прямоугольное отверстие размером D × E.
Замок Иф сложен из кирпичей, размером A × B × C.
Определите, сможет ли узник выбрасывать кирпичи в море через это отверстие,
если стороны кирпича должны быть параллельны сторонам отверстия.

Формат ввода
Программа получает на вход числа A, B, C, D, E.

Формат вывода
Программа должна вывести слово YES или NO.

Пример 1
Ввод
1
1
1
1
1
Вывод
YES

Пример 2
Ввод
2
2
2
1
1
Вывод
NO
*/
package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	fmt.Println(vlezet(a, b, c, d, e))
}

func vlezet(a, b, c, d, e int) string {
	a, b = sort2(a, b)
	b, c = sort2(b, c)
	a, b = sort2(a, b)
	d, e = sort2(d, e)
	if a <= d && b <= e {
		return "YES"
	}
	return "NO"
}

func sort2(first, second int) (int, int) {
	if first <= second {
		return first, second
	}
	return second, first
}
