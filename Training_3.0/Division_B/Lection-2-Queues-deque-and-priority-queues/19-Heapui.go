package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PriorityQueue struct {
	MaxHeap  []int
	HeapSize int
}

func (pq *PriorityQueue) Insert(value int) {
	pq.HeapSize += 1
	pq.MaxHeap = append(pq.MaxHeap, value)
	// После добавления элемнта в кучу, возможно нарушение её главного свойства
	// (каждый элемент больше своих потомков). Поэтому нужно пофиксить эего используя
	// функцию просеивания вверх siftUp pq.siftUp(pq.HeapSize - 1)
	pq.siftUp(pq.HeapSize - 1)
}

// O(log n)
func (pq *PriorityQueue) siftUp(i int) {
	// Двигаемся до того момента,
	// пока не станет выполнятся свойство кучи   // MaxHeap[(i - 1) / 2] < i[i]
	for pq.MaxHeap[i] > pq.MaxHeap[(i-1)/2] {
		// Swap them
		pq.Swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}

// ExtractMax returns the maximum element and removes it from the PriorityQueue
func (pq *PriorityQueue) ExtractMax() int {
	if pq.HeapSize == 0 {
		return 0
	}
	max := pq.MaxHeap[0]
	// Put last elem; on poped elem place
	pq.MaxHeap[0] = pq.MaxHeap[pq.HeapSize-1]
	pq.HeapSize -= 1
	// Crop ast elem;
	pq.MaxHeap = pq.MaxHeap[:pq.HeapSize]

	pq.siftDown(0)
	return max
}

func (pq *PriorityQueue) siftDown(i int) {
	for 0 < pq.maxChild(i) && pq.MaxHeap[i] < pq.MaxHeap[pq.maxChild(i)] {
		maxChild := pq.maxChild(i)
		pq.Swap(i, maxChild)
		i = maxChild
	}
}

//     11              11
//    /  \             /  \      //  34    25           4    25

//        11      //     11    10      //  11   11      // 11      // Берем максимальный из двух и меняем их местами}

func (pq *PriorityQueue) maxChild(i int) int {
	var child int

	if 2*i+1 < pq.HeapSize {
		child = 2*i + 1
	}

	if 2*i+2 < pq.HeapSize && pq.MaxHeap[child] < pq.MaxHeap[2*i+2] {
		child = 2*i + 2
	}

	return child
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.MaxHeap[i], pq.MaxHeap[j] = pq.MaxHeap[j], pq.MaxHeap[i]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	heap := &PriorityQueue{}

	var commandsNum int
	fmt.Scan(&commandsNum)

	for i := 0; i < commandsNum; i++ {
		scanner.Scan()
		line := scanner.Text()
		words := strings.Split(line, " ")
		switch words[0] {
		case "0": // Insert
			n, _ := strconv.Atoi(words[1])
			heap.Insert(n)
			// fmt.Println("ok")
		case "1": // Extract
			fmt.Println(heap.ExtractMax())
		}
		// fmt.Println(heap)
	}
}

// Linked List Version Time Limit 11 testcase
/*
type HeapNode struct {
	val    int
	parent *HeapNode
	left   *HeapNode
	right  *HeapNode
}

type MaxHeap struct {
	root *HeapNode
}

func (mh *MaxHeap) getLastparent() *HeapNode {
	q := list.New()
	q.PushBack(mh.root)

	var nodeElem *HeapNode
	for q.Len() > 0 {
		node := q.Front()
		nodeElem = node.Value.(*HeapNode)
		q.Remove(node)

		if node.Value.(*HeapNode).left != nil && node.Value.(*HeapNode).right != nil {
			q.PushBack(node.Value.(*HeapNode).left)
			q.PushBack(node.Value.(*HeapNode).right)
		} else {
			break
		}
	}
	return nodeElem
}

func (mh *MaxHeap) SiftUp(node *HeapNode) {
	if node.parent == nil {
		return
	}
	if node.parent.val < node.val {
		node.parent.val, node.val = node.val, node.parent.val
		mh.SiftUp(node.parent)
	}
}

func (mh *MaxHeap) Insert(elem int) {
	node := &HeapNode{val: elem}
	if mh.root == nil {
		mh.root = node
		return
	}
	lPar := mh.lastParent()

	if lPar.left == nil {
		lPar.left = node
		node.parent = lPar
	} else {
		lPar.right = node
		node.parent = lPar
	}
	mh.SiftUp(node)
}

func (mh *MaxHeap) Extract() int {
	if mh.isEmpty() {
		fmt.Println("Heap is empty")
		return 0
	}
	// "Deleting: "root.data
	retn := mh.root.val

	last := mh.lastNode()
	if last == mh.root {
		last = nil
		mh.root = nil
		return retn
	}

	mh.root.val, last.val = last.val, mh.root.val
	parent := last.parent

	if parent.left == last {
		last = nil
		parent.left = nil
	} else {
		last = nil
		parent.right = nil
	}

	mh.SiftDown(mh.root)
	return retn
}

//
func (mh *MaxHeap) isEmpty() bool {
	return mh.root == nil
}

func (mh *MaxHeap) SiftDown(node *HeapNode) {
	largest := node
	if node.left != nil && node.left.val > largest.val {
		largest = node.left
	}

	if node.right != nil && node.right.val > largest.val {
		largest = node.right
	}

	if largest != node {
		largest.val, node.val = node.val, largest.val
		mh.SiftDown(largest)
	}
}

//
func (mh *MaxHeap) lastParent() *HeapNode {
	q := list.New()
	q.PushBack(mh.root)

	var nodeElem *HeapNode
	for q.Len() > 0 {
		node := q.Front()
		nodeElem = node.Value.(*HeapNode)
		q.Remove(node)

		if node.Value.(*HeapNode).left != nil && node.Value.(*HeapNode).right != nil {
			q.PushBack(node.Value.(*HeapNode).left)
			q.PushBack(node.Value.(*HeapNode).right)
		} else {
			break
		}
	}
	return nodeElem
}

//
func (mh *MaxHeap) lastNode() *HeapNode {
	q := list.New()
	q.PushBack(mh.root)

	var nodeElem *HeapNode
	for q.Len() > 0 {
		node := q.Front()
		nodeElem = node.Value.(*HeapNode)
		q.Remove(node)

		if node.Value.(*HeapNode).left != nil {
			q.PushBack(node.Value.(*HeapNode).left)
		}
		if node.Value.(*HeapNode).right != nil {
			q.PushBack(node.Value.(*HeapNode).right)
		}
	}
	return nodeElem
}
*/
