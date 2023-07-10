package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int //количество операций,
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)
	STRIZH := NewTrain()

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text() // Input: Sheldon Cooper 100 98 Physics\n
		words := strings.Split(line, " ")
		switch words[0] {
		case "add":
			STRIZH.Add(words[2], words[1])
		case "get":
			fmt.Println(STRIZH.Get(words[1]))
		case "delete":
			wagons2Delete, _ := strconv.Atoi(words[1])
			STRIZH.Delete(wagons2Delete)
		}
	}
}

type Train struct {
	wagons_stuff map[string]int
	wagons_names []string
}

func NewTrain() *Train {
	return &Train{
		wagons_stuff: map[string]int{},
		wagons_names: []string{},
	}
}

func (t *Train) Get(res string) int {
	return (*t).wagons_stuff[res]
}

func (t *Train) Add(res, wagons string) {
	num, _ := strconv.Atoi(wagons)
	(*t).wagons_stuff[res] += num
	for i := 0; i < num; i++ {
		(*t).wagons_names = append((*t).wagons_names, res)
	}
}

func (t *Train) Delete(num int) {
	for num != 0 {
		num--
		(*t).wagons_stuff[(*t).wagons_names[len((*t).wagons_names)-1]] -= 1
		(*t).wagons_names = (*t).wagons_names[:len((*t).wagons_names)-1]
	}
}
