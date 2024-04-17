package dataStructure

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

// 实现排序规则的接口, Len(), Less(), Swap()

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i // 调整index
	pq[i].index = j
}

// 实现 heap.interface的接口函数

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item) // 强制类型转换
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	old[n-1] = nil //  避免内存泄漏
	x.index = -1   // 安全检查
	*pq = old[0 : n-1]
	return x
}

// 更新操作，维护堆的性质

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority // 更新了priority这个值
	heap.Fix(pq, item.index)
}

func PriorityQueueUseExample() {
	items := map[string]int{
		"banana": 3,
		"apple":  2,
		"pear":   4,
	}
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{value, priority, i}
		i++
	}
	heap.Init(&pq)
	item := &Item{value: "orange", priority: 1}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s\n", item.priority, item.value)
	}
}
