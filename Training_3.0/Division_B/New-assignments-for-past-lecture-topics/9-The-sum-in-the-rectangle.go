package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Pattern : Дерево отрезков, RSQ, RMQ, Динамическое программирование на таблицах

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()
	words := strings.Split(line, " ")

	n, _ := strconv.Atoi(words[0])
	m, _ := strconv.Atoi(words[1])
	k, _ := strconv.Atoi(words[2])

	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		scanner.Scan()
		xStr := scanner.Text()
		xSl := strings.Split(xStr, " ")
		for j := 1; j <= m; j++ {
			x, _ := strconv.Atoi(xSl[j-1])
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1] - matrix[i-1][j-1] + x
		}
	}

	for i := 0; i < k; i++ {
		scanner.Scan()
		qr := scanner.Text()
		qrs := strings.Split(qr, " ")

		i1, _ := strconv.Atoi(qrs[0])
		j1, _ := strconv.Atoi(qrs[1])
		i2, _ := strconv.Atoi(qrs[2])
		j2, _ := strconv.Atoi(qrs[3])

		fmt.Println(matrix[i2][j2] - matrix[i1-1][j2] - matrix[i2][j1-1] + matrix[i1-1][j1-1])
	}
}
