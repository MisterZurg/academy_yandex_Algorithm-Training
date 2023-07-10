package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func simplifyPath(path string) string {
	flag := true
	path += "/"
	for flag {
		if strings.Contains(path, "/./") {
			path = strings.ReplaceAll(path, "/./", "/")
		} else if strings.Contains(path, "//") {
			path = strings.ReplaceAll(path, "//", "/")
		} else if strings.Contains(path, "/../") {
			pos := strings.Index(path, "/../")
			if pos == 0 {
				path = path[3:]
			} else {
				prevSlashPos := strings.LastIndex(path[:pos-1], "/")
				path = path[:prevSlashPos] + path[pos+3:]
			}
		} else {
			flag = false
		}
	}
	if strings.HasSuffix(path, "/") && path != "/" {
		path = path[:len(path)-1]
	}
	return path
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(simplifyPath(s))
}
