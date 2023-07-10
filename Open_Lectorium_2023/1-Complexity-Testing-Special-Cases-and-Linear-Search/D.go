package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// TODO RE 3
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	timePointsStr := scanner.Text()
	timePointsArr := strings.Split(timePointsStr, " ")
	minutePoints := make([]int, n)

	for i, nowTime := range timePointsArr {
		hm := strings.Split(nowTime, ":")
		h, _ := strconv.Atoi(hm[0])
		m, _ := strconv.Atoi(hm[1])
		minutePoints[i] = h*60 + m
	}

	sort.Ints(minutePoints)
	minDist := 1440 + minutePoints[0] - minutePoints[n-1]

	for i := 1; i < n; i++ {
		dist := minutePoints[i] - minutePoints[i-1]
		if dist < minDist {
			minDist = dist
		}
	}

	fmt.Println(minDist)
}
