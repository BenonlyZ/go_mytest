package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{100, 16, 4, 8, 70, 2, 36, 22, 5, 12}
	heap.Init(h)
	fmt.Println(h)
	heap.Push(h, 10000)
	fmt.Println(h)
	heap.Remove(h, 1)
	//heap.Push(h, 5)
	fmt.Println(h)
	heap.Fix(h, 1)
	fmt.Println(h)

}
