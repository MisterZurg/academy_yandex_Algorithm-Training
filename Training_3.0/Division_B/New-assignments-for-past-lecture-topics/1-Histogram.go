package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Задачи городской олимпиады по информатике 1995 года
// Сортировка подсчетом a.k.a Counting sort

func main() {
	sc := bufio.NewScanner(os.Stdin)
	text := []string{}

	for sc.Scan() {
		line := sc.Text()
		text = append(text, line)
	}

	countingSort(text)
}

func countingSort(text []string) {
	freq := make(map[string]int)
	maxFreq := 0
	for _, line := range text {
		for _, lt := range line {
			if string(lt) == " " {
				continue
			}
			if _, ok := freq[string(lt)]; !ok {
				freq[string(lt)] = 0
			}
			freq[string(lt)]++
			maxFreq = max(maxFreq, freq[string(lt)])
		}
	}

	sortedKeys := make([]string, len(freq))
	i := 0
	for k := range freq {
		sortedKeys[i] = k
		i++
	}
	sort.Strings(sortedKeys)

	for row := maxFreq; row > 0; row-- {
		for _, lt := range sortedKeys {
			if freq[lt] >= row {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	for _, lt := range sortedKeys {
		fmt.Printf("%s", lt)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func testcases() {

}

/*
def printChart(s):
    symcount = {}
    maxsymcount = 0
    for sym in s:
        if sym not in symcount:
            symcount[sym] = 0
        symcount[sym] += 1
        maxsymcount = max(maxsymcount, symcount[sym])
    sorteduniqsyms = sorted(symcount.keys())
    for row in range(maxsymcount, 0, -1):
        for sym in sorteduniqsyms:
            if symcount[sym] >= row:
                print('#', end='')
            else:
                print(' ', end='')
        print()
    print(''.join(sorteduniqsyms))
*/
