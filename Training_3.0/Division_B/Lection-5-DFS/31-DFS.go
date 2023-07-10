package main

import "fmt"

func main() {
	var n, m int // количество вершин и ребер в графе
	fmt.Scan(&n, &m)
	graph := make([][2]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&graph[i][0], &graph[i][1])
	}
}

func DFS(graph, viseted, curr) {

}
