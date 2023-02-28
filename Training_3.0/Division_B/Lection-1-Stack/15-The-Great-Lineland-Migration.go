package main

import (
	"container/list"
	"fmt"
)

type Pair struct {
	val int
	idx int
}

// PairStack  содержит возрастающую последовательность
type PairStack struct {
	*list.List
}

// Достаем из стека числа до пока != или меньше
func main() {
	var n int
	fmt.Scan(&n)
	seq := make([]int, n)
	for i := range seq {
		fmt.Scan(&seq[i])
	}

	ans := make([]int, n)
	ps := &PairStack{list.New()}

	for i, v := range seq {
		if ps.Len() == 0 {
			ps.PushBack(Pair{val: v, idx: i})
			continue
		}

		for ps.Len() > 0 && v < ps.Back().Value.(Pair).val {
			poped := ps.Back()
			ps.Remove(poped)
			ans[poped.Value.(Pair).idx] = i
		}
		ps.PushBack(Pair{val: v, idx: i})
	}

	for i, v := range ans {
		if v == 0 {
			ans[i] = -1
		}
		fmt.Printf("%d ", ans[i])
	}
}
