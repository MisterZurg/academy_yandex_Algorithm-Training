// TL 14
package main

import "fmt"

// 88. Merge Sorted Array
func Merge(as, bs []int) []int {
	merged := make([]int, len(as)+len(bs))
	i, j, m := 0, 0, 0

	for i < len(as) && j < len(bs) {
		if as[i] < bs[j] {
			merged[m] = as[i]
			i++
		} else {
			merged[m] = bs[j]
			j++
		}
		m++
	}
	for i < len(as) {
		merged[m] = as[i]
		i++
		m++
	}
	for j < len(bs) {
		merged[m] = bs[j]
		j++
		m++
	}
	return merged
}

func main() {
	var a, b int
	fmt.Scan(&a)
	as := make([]int, a)
	for i := range as {
		fmt.Scan(&as[i])
	}

	fmt.Scan(&b)
	bs := make([]int, b)
	for i := range bs {
		fmt.Scan(&bs[i])
	}

	merged := Merge(as, bs)
	for _, el := range merged {
		fmt.Printf("%d ", el)
	}
}
