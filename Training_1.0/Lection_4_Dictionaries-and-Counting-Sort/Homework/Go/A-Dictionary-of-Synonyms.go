package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	num, _ := strconv.Atoi(s)

	first := make(map[string]string)
	second := make(map[string]string)

	for i := 0; i < num; i++ {
		scanner.Scan()
		line := scanner.Text()
		kv := strings.Split(line, " ")

		k, v := kv[0], kv[1]

		first[k] = v
		second[v] = k
	}

	scanner.Scan()
	find := scanner.Text()
	if r1, ok1 := first[find]; ok1 {
		fmt.Println(r1)
	} else if r2, ok2 := second[find]; ok2 {
		fmt.Println(r2)
	}
}
