package lib

import (
	"container/heap"
	"fmt"
)

// https://pkg.go.dev/container/heap

// Don't forget to use heap.Init, heap.Push, heap.Pop

type Item[T any] struct {
	Value T   // The value of the item; arbitrary.
	Cost  int // The cost of the item in the queue.
	// The Index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PrioQueue[T any] []*Item[T]

func (q PrioQueue[T]) Len() int { return len(q) }

func (q PrioQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the lowest cost
	return q[i].Cost < q[j].Cost
}

func (q PrioQueue[T]) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].Index = i
	q[j].Index = j
}

func (pq *PrioQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.Index = n
	*pq = append(*pq, item)
}

func (q *PrioQueue[T]) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*q = old[0 : n-1]
	return item
}

func (pq *PrioQueue[T]) Peek() *Item[T] {
	return (*pq)[0]
}

func (q *PrioQueue[T]) String() string {
	result := ""
	for i := range *q {
		result += fmt.Sprintf("%v ", *((*q)[i]))
	}
	return result
}

// update modifies the cost of an Item in the queue.
func (pq *PrioQueue[T]) UpdateCost(item *Item[T], cost int) {
	item.Cost = cost
	heap.Fix(pq, item.Index)
}
