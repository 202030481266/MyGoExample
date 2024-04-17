package dataStructure

import (
	"container/heap"
	"fmt"
)

// 最小堆

type IntHeap []int

// 实现指定的排序接口

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 实现 heap.interface 接口中的方法
// Push方法和Pop方法调用后都会调用heap.Fix()方法维护堆的性质

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// 这里的Pop方法实际上会首先将最小的元素和最后一个元素调换位置

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func IntHeapUseExample() {
	h := &IntHeap{6, 2, 1, 5, 8, 4, 3}
	heap.Init(h)
	heap.Push(h, 7)
	for h.Len() > 0 {
		fmt.Printf("minimum values is %d\n", heap.Pop(h))
	}
}
