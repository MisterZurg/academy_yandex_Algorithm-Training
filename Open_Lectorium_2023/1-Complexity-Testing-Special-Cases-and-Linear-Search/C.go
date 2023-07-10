package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO: TC 9 WA
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	costStr := scanner.Text()
	costArr := make([]int, n)
	costStrArr := strings.Split(costStr, " ")
	for i, s := range costStrArr {
		costArr[i], _ = strconv.Atoi(s)
	}

	bestBuyDay := 0
	bestSellDay := 0
	minCostDay := 0

	for i := 1; i < n; i++ {
		if costArr[bestSellDay]*costArr[minCostDay] < costArr[bestBuyDay]*costArr[i] {
			bestBuyDay = minCostDay
			bestSellDay = i
		}
		if costArr[i] < costArr[minCostDay] {
			minCostDay = i
		}
	}

	if bestSellDay == 0 && bestBuyDay == 0 {
		fmt.Println(0, 0)
	} else {
		fmt.Println(bestBuyDay+1, bestSellDay+1)
	}
}
