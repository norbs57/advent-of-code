package lib

import (
	"fmt"
	"strings"
)

// Generic heap based on pkg.go.dev/container/heap

type Heap[T any] struct {
	data []T
	less func(a, b T) bool
	swap func(i, j int)
}

func NewHeap[T any](less func(a, b T) bool, data ...T) *Heap[T] {
	h := &Heap[T]{data: data, less: less}
	h.swap = func(i, j int) {
		h.data[i], h.data[j] = h.data[j], h.data[i]
	}
	h.Init()
	return h
}

func NewCostFunHeap[T any, V Ordered](cost func(T) V, data ...T) *Heap[T] {
	less := func(p, q T) bool {
		return cost(p) < cost(q)
	}
	return NewHeap(less, data...)
}

func NewMinHeap[T Ordered](data ...T) *Heap[T] {
	less := func(p, q T) bool {
		return p < q
	}
	return NewHeap(less, data...)
}

func NewMaxHeap[T Ordered](data ...T) *Heap[T] {
	f := func(p, q T) bool {
		return p > q
	}
	return NewHeap(f, data...)
}

func NewIntCostHeap[T Ordered](cost []T, data ...int) *Heap[int] {
	f := func(a, b int) bool {
		return cost[a] < cost[b]
	}
	return NewHeap(f, data...)
}

func (h *Heap[T]) String() string {
	return fmt.Sprintf("%v", h.data)
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) Data() []T {
	return h.data
}

func (h *Heap[T]) Clear() {
	h.data = h.data[:0]
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func (h *Heap[T]) Init() {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Push(x T) {
	h.data = append(h.data, x)
	h.up(h.Len() - 1)
}

// Pop removes and returns the minimum element (according to lessFun) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to h.Remove(0).
func (h *Heap[T]) Pop() T {
	n := h.Len() - 1
	h.swap(0, n)
	h.down(0, n)
	x := h.data[n]
	h.data[n] = *new(T)
	h.data = h.data[0:n]
	return x
}

func (h *Heap[T]) Peek() T {
	return h.data[0]
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len()
func (h *Heap[T]) Remove(i int) T {
	t := h.data[i]
	n := h.Len() - 1
	if n != i {
		h.swap(i, n)
		if !h.down(i, n) {
			h.up(i)
		}
	}
	h.data[n] = *new(T)
	h.data = h.data[:n]
	return t
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling h.Remove(i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Fix(i int) {
	if !h.down(i, h.Len()) {
		h.up(i)
	}
}

func (h *Heap[T]) lessOnData(i, j int) bool {
	return h.less(h.data[i], h.data[j])
}

func (h *Heap[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.lessOnData(j, i) {
			break
		}
		h.swap(i, j)
		j = i
	}
}

func (h *Heap[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.lessOnData(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.lessOnData(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}
	return i > i0
}

// PQInt is an int-heap with an index slice
// Elements are positive integers up to len(index)-1
// Pop/peek removes/returns min element (according to lessFun)

type PQInt[T Integer] struct {
	Heap[T]
	index []int
}

func NewPQInt[T Integer](less func(a, b T) bool, num ...int) *PQInt[T] {
	var n, m int
	if len(num) > 0 {
		n = num[0]
		m = n
	}
	if len(num) > 1 {
		m = num[1]
		if m < n {
			m = n
		}
	}
	q := &PQInt[T]{
		index: make([]int, m),
		Heap: Heap[T]{
			less: less,
			data: make([]T, n, m),
		},
	}
	q.swap = func(i, j int) {
		q.data[i], q.data[j] = q.data[j], q.data[i]
		q.index[q.data[i]] = i
		q.index[q.data[j]] = j
	}
	for i := range q.data {
		q.data[i] = T(i)
		q.index[i] = i
	}
	for i := n; i < len(q.index); i++ {
		q.index[i] = -1
	}
	q.Init()
	return q
}

func (h *PQInt[T]) String() string {
	return fmt.Sprintf("heap %v, index %v", h.data, h.index)
}

func (q *PQInt[T]) Push(v T) {
	q.index[v] = len(q.data)
	q.Heap.Push(v)

}

func (q *PQInt[T]) Pop() T {
	v := q.Heap.Pop()
	q.index[v] = -1
	return v
}

// Remove element v
func (q *PQInt[T]) Remove(v int) {
	i := q.index[v]
	q.Heap.Remove(i)
	q.index[i] = -1
}

// Contains tells whether v is in the queue.
func (q *PQInt[T]) Contains(v int) bool {
	return q.index[v] >= 0
}

func (q *PQInt[T]) Fix(v T) {
	q.Heap.Fix(q.index[v])
}

func NewMinCostQG[T Ordered, U Integer](cost []T, num ...int) *PQInt[U] {
	n := 0
	if len(num) > 0 {
		n = num[0]
	}
	return NewPQInt[U](
		func(a, b U) bool {
			return cost[a] < cost[b]
		},
		n, len(cost))
}

func NewMaxPrioQG[T Ordered, U Integer](prio []T, num ...int) *PQInt[U] {
	n := 0
	if len(num) > 0 {
		n = num[0]
	}
	return NewPQInt[U](
		func(a, b U) bool {
			return prio[a] > prio[b]
		},
		n, len(prio))
}

func NewMinCostQ[T Ordered](cost []T, num ...int) *PQInt[int] {
	return NewMinCostQG[T, int](cost, num...)
}

func NewMaxPrioQ[T Ordered](prio []T, num ...int) *PQInt[int] {
	return NewMaxPrioQG[T, int](prio, num...)
}

func (q *PQInt[T]) IndexLen() int {
	return len(q.index)
}

// Resize the index so that the queue can hold larger elements
func (q *PQInt[T]) Resize(n int) {
	if n <= len(q.index) {
		return
	}
	tmp := make([]int, n)
	copy(tmp, q.index)
	for i := len(q.index); i < len(tmp); i++ {
		tmp[i] = -1
	}
	q.index = tmp
}

// PrioQueue is a generic heap of pointers to int-indexed values

type PQItem[T any] struct {
	Value T
	Index int
}

type PQueue[T any] struct {
	Heap[*PQItem[T]]
}

func NewPQueue[T any](less func(a, b T) bool) *PQueue[T] {
	q := &PQueue[T]{}
	q.less = func(a, b *PQItem[T]) bool {
		return less(a.Value, b.Value)
	}
	q.swap = func(i, j int) {
		q.data[i], q.data[j] = q.data[j], q.data[i]
		q.data[i].Index = i
		q.data[j].Index = j
	}
	q.Init()
	return q
}

func (q *PQueue[T]) String() string {
	var sb strings.Builder
	for _, item := range q.data {
		s := fmt.Sprintf("%v", *item)
		sb.WriteString(s)
		sb.WriteRune(' ')
	}
	return sb.String()
}

func (q *PQueue[T]) Push(x *PQItem[T]) {
	x.Index = len(q.data)
	q.Heap.Push(x)
}

func (q *PQueue[T]) Peek() T {
	return q.Heap.Peek().Value
}

func (q *PQueue[T]) Pop() T {
	return q.Heap.Pop().Value
}

func (q *PQueue[T]) Remove(i int) T {
	return q.Heap.Remove(i).Value
}
