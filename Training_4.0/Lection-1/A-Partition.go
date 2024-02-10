// TL 14 change reader
package main

import "fmt"

func printPredicats(partitioned []int, pivot int) {
	var before, after int
	for i := range partitioned {
		if partitioned[i] == pivot {
			before = i
		}
	}

	for i := len(partitioned) - 1; i >= 0; i-- {
		if partitioned[i] == pivot {
			after = i
			break
		}
	}
	fmt.Println(len(partitioned[:before]))
	fmt.Println(len(partitioned[after:]))
}

// Partition разбивает набор элементов на две части относительно заданного предиката.
func Partition(as []int, pivot int) []int {
	// 1 9 4 2 3
	// 3
	equal := 0 // pointer on pivot
	greater := 0
	// now := 0

	for now := range as {
		// Если новый элемент меньше Опорного, тогда его нужно добавить в группу где все меньше Опорного
		if as[now] < pivot {
			// сохраним текущий
			// гритер ставим в текущий
			// равный в гритер
			// сохраненный в равный
			// equal++
			// greater++
			tmp := as[now]
			as[now] = as[greater]
			as[greater] = as[equal]
			as[equal] = tmp
			// as[equal], as[greater], as[now] = as[now], as[equal], as[greater]
			equal++
			greater++
		} else if as[now] == pivot {
			// свапнуть гритер и текущий
			// гритер++
			as[greater], as[now] = as[now], as[greater]
			greater++
		} else { // уже больший
			continue
		}
	}
	return as
}

func main() {
	var n int // — количество элементов массива
	fmt.Scan(&n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}
	var pivot int // опорный элемент
	fmt.Scan(&pivot)

	partitioned := Partition(as, pivot)
	printPredicats(partitioned, pivot)
}

// C++
/*
#include <algorithm>
#include <cctype>
#include <cmath>
#include <cstdint>
#include <cstring>
#include <fstream>
#include <iostream>
#include <list>
#include <map>
#include <queue>
#include <set>
#include <sstream>
#include <stack>
#include <string>
#include <tuple>
#include <unordered_map>
#include <unordered_set>
#include <vector>

uint32_t partition(int32_t predicat, std::vector<int32_t>& vect, uint32_t i, uint32_t j) {
	  uint32_t eq = i, gt = i, now = i;
	  int32_t buf;
	  while (now < j) {
			if (vect[now] < predicat) {
				  buf = vect[now];
				  vect[now] = vect[gt];
				  vect[gt] = vect[eq];
				  vect[eq] = buf;
				  ++gt;
				  ++eq;
			} else if (vect[now] == predicat) {
				  std::swap(vect[now], vect[gt]);
				  ++gt;
			}
			++now;
	  }
	  return eq;
}

int main() {
	  uint32_t n, ans;
	  int32_t x;

	  std::cin >> n;
	  std::vector<int32_t> vect(n);
	  for (uint32_t i = 0; i < n; ++i) {
		std::cin >> vect[i];
	  }
	  std::cin >> x;

	  ans = partition(x, vect, 0, n);
	  std::cout << ans << std::endl << n - ans << std::endl;
}
*/
