// generic priority queue class which allows for a custom comparator
// author: Nora Cam
package pq

import (
	"container/heap"
)

type Comparator[T any] func(a, b T) bool

type PriorityQueue[T any] struct {
	items []T
	less  Comparator[T]
}

func NewPriorityQueue[T any](less Comparator[T]) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{less: less}
	heap.Init(pq)
	return pq
}

func NewPriorityQueueFromSlice[T any](slice []T, less Comparator[T]) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{items: slice, less: less}
	heap.Init(pq)
	return pq
}

func (pq *PriorityQueue[T]) Len() int           { return len(pq.items) }
func (pq *PriorityQueue[T]) Less(i, j int) bool { return pq.less(pq.items[i], pq.items[j]) }
func (pq *PriorityQueue[T]) Swap(i, j int)      { pq.items[i], pq.items[j] = pq.items[j], pq.items[i] }

// Returns the internal slice of items
func (pq *PriorityQueue[T]) GetItems() []T {
	return pq.items
}

// required -- do not use
func (pq *PriorityQueue[T]) Push(x any) {
	pq.items = append(pq.items, x.(T))
}

// required -- do not use
func (pq *PriorityQueue[T]) Pop() any {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) Peek() T {
	return pq.items[0]
}

func (pq *PriorityQueue[T]) PushItem(x T) {
	heap.Push(pq, x)
}

func (pq *PriorityQueue[T]) PopItem() T {
	return heap.Pop(pq).(T)
}
