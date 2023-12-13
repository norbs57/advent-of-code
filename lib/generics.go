package lib

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"

	"golang.org/x/exp/constraints"
)

// Generics

type Integer interface {
	constraints.Integer
}
type Num interface {
	constraints.Integer | constraints.Float
}

type SignedNum interface {
	constraints.Signed | constraints.Float
}

// Ordered represents the set of types for which the '<' operator work.
type Ordered interface {
	constraints.Integer | constraints.Float | ~string
}

type Pair[U any, V any] struct {
	Fst U
	Snd V
}

func (p Pair[U, V]) Extract() (U, V) {
	return p.Fst, p.Snd
}

func MkPair[U any, V any](u U, v V) Pair[U, V] {
	return Pair[U, V]{u, v}
}

func Swap[T any](a [2]T) [2]T {
	return [2]T{a[1], a[0]}
}

func Abs[T SignedNum](a T) T {
	if a >= T(0) {
		return a
	} else {
		return -a
	}
}

func Sign[T SignedNum](a T) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return 1
	default:
		return 0
	}
}

func Append[T any](xs *[]T, el ...T) {
	*xs = append(*xs, el...)
}

func Last[E any](xs []E) E {
	return xs[len(xs)-1]
}

func Map[T, S any](f func(T) S, xs []T) []S {
	result := make([]S, len(xs))
	for i, t := range xs {
		result[i] = f(t)
	}
	return result
}

func MaxIdx[T Ordered](xs []T) int {
	if len(xs) == 0 {
		return -1
	}
	idx, max := 0, xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] > max {
			idx, max = i, xs[i]
		}
	}
	return idx
}

func PrefixSums[T Num](xs []T) []T {
	if len(xs) == 0 {
		return xs
	}
	res := make([]T, 0, len(xs))
	sum := T(0)
	for _, x := range xs {
		sum += x
		res = append(res, sum)
	}
	return res
}

func PrefixSumsFunc[T Num](f func(int) T, n int) []T {
	res := make([]T, 1, n)
	sum := T(0)
	for i := 0; i < n; i++ {
		sum += f(i)
		res = append(res, sum)
	}
	return res
}

func Prod[T Num](xs []T) T {
	result := T(1)
	for _, x := range xs {
		result *= x
	}
	return result
}

func Sum[T Num](xs []T) T {
	result := T(0)
	for _, x := range xs {
		result += x
	}
	return result
}

func Freq[E comparable](xs []E) map[E]int {
	res := make(map[E]int)
	for _, x := range xs {
		res[x]++
	}
	return res
}

func Filter[E any](xs []E, pred func(E) bool) []E {
	res := make([]E, 0, len(xs))
	for _, x := range xs {
		if pred(x) {
			res = append(res, x)
		}
	}
	return res
}

func Mod(a, b int) int {
	// result is positive if b is positive
	return (a%b + b) % b
}

func Dist2[T SignedNum](a, b [2]T) T {
	return Abs(b[0]-a[0]) + Abs(b[1]-a[1])
}

func Add2[T Num](p, q [2]T) [2]T {
	return [2]T{p[0] + q[0], p[1] + q[1]}
}

func Add3[T Num](p, q [3]T) [3]T {
	return [3]T{p[0] + q[0], p[1] + q[1], p[2] + q[2]}
}

func Greater[T Ordered](p, q T) bool {
	return p > q
}

func Greater2[T Ordered](p, q [2]T) bool {
	return p[0] > q[0] || p[0] == q[0] && p[1] > q[1]
}

func Greater3[T Ordered](a, b [3]T) bool {
	return a[0] > b[0] || a[0] == b[0] && a[1] > b[1] ||
		a[0] == b[0] && a[1] == b[1] && a[2] > b[2]
}

func Less2[T Ordered](a, b [2]T) bool {
	return a[0] < b[0] || a[0] == b[0] && a[1] < b[1]
}

func LessOrEqual2[T Ordered](a, b [2]T) bool {
	return a == b || a[0] < b[0] || a[0] == b[0] && a[1] < b[1]
}

func Less3[T Ordered](a, b [3]T) bool {
	return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] ||
		a[0] == b[0] && a[1] == b[1] && a[2] < b[2]
}

func MatrixOfAdjList(g AdjList) [][]int {
	res := MkSlice2[int](len(g), len(g))
	for i, gi := range g {
		for _, j := range gi {
			res[i][j] = 1
		}
	}
	return res
}
func MkNumsUpto[E Num](n int) []E {
	result := make([]E, n)
	for i := range result {
		result[i] = E(i)
	}
	return result
}

func MkSliceFilled[E any](rows int, t E) []E {
	result := make([]E, rows)
	for i := range result {
		result[i] = t
	}
	return result
}

func MkSlice2[E any](rows, cols int) [][]E {
	result := make([][]E, rows)
	for i := range result {
		result[i] = make([]E, cols)
	}
	return result
}

func MkSlice2Filled[E any](rows, cols int, t E) [][]E {
	result := make([][]E, rows)
	for i := range result {
		result[i] = make([]E, cols)
		for j := range result[i] {
			result[i][j] = t
		}
	}
	return result
}

func MkSlice3[E any](rows, cols, height int) [][][]E {
	result := make([][][]E, rows)
	for i := range result {
		result[i] = make([][]E, cols)
		for j := range result[i] {
			result[i][j] = make([]E, height)
		}
	}
	return result
}

func MkSlice3Filled[E any](rows, cols, height int, e E) [][][]E {
	result := make([][][]E, rows)
	for i := range result {
		result[i] = make([][]E, cols)
		for j := range result[i] {
			result[i][j] = make([]E, height)
			for k := range result[i][j] {
				result[i][j][k] = e
			}
		}
	}
	return result
}

func PopBack[E any](x *[]E) E {
	n := len(*x) - 1
	res := (*x)[n]
	*x = (*x)[:n]
	return res
}

func PopFront[E any](x *[]E) E {
	res := (*x)[0]
	*x = (*x)[1:]
	return res
}

func PushBack[E any](x *[]E, e E) {
	*x = append((*x), e)
}

func PushFront[E any](x *[]E, e E) {
	n := len(*x)
	*x = append((*x), e)
	copy((*x)[1:], (*x)[:n])
	(*x)[0] = e
}

func Reverse[E any](x []E) {
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
}

func SetMax[T Ordered](a *T, b T) {
	if b > *a {
		*a = b
	}
}

func SetMin[T Ordered](a *T, b T) {
	if b < *a {
		*a = b
	}
}

func Shuffle[T any](xs []T) {
	rand.Shuffle(len(xs), func(i, j int) {
		xs[i], xs[j] = xs[j], xs[i]
	})
}

func SlicesEqual[E comparable](s1, s2 []E) bool {
	// https://pkg.go.dev/golang.org/x/exp/slices#Equal
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func SortLess2[T Ordered](s [][2]int) {
	sort.Slice(s, func(i, j int) bool {
		return Less2(s[i], s[j])
	})
}

func SortSliceDecr[T Ordered](xs []T) {
	sort.Slice(xs, func(i, j int) bool {
		return xs[i] > xs[j]
	})
}

func SortSliceFunc[E any, T Ordered](xs []E, f func(E) T) {
	sort.Slice(xs, func(i, j int) bool {
		return f(xs[i]) < f(xs[j])
	})
}

func SortSliceLessFunc[E any](xs []E, less func(a, b E) bool) {
	sort.Slice(xs, func(i, j int) bool {
		return less(xs[i], xs[j])
	})
}

func SumOverMap[T comparable, E Num](m map[T]E) E {
	res := E(0)
	for _, f := range m {
		res += f
	}
	return res
}

// AdjList
type AdjList [][]int

func (g AdjList) Add(v, w int) {
	g[v] = append(g[v], w)
}

func (g AdjList) AddBoth(v, w int) {
	g.Add(v, w)
	if v != w {
		g.Add(w, v)
	}
}

// AdjMap
type AdjMap []map[int]bool

func MkAdjMap(n int) AdjMap {
	res := make(AdjMap, n)
	for i := range res {
		res[i] = make(map[int]bool)
	}
	return res
}

func (g AdjMap) Add(v, w int) {
	g[v][w] = true
}

func (g AdjMap) AddBoth(v, w int) {
	g[v][w] = true
	g[w][v] = true
}

func EdgesOf[T comparable](g map[T]map[T]int) map[[2]T]int {
	res := make(map[[2]T]int)
	for a := range g {
		for b, freq := range g[a] {
			k := [2]T{a, b}
			_, ok := res[[2]T{b, a}]
			if !ok {
				res[k] = freq
			}
		}
	}
	return res
}

// MapGraph: slice of maps

const initialMapSize = 8

type MapGraph[T Num] []map[int]T

func MkMapGraph[T Num](n int) MapGraph[T] {
	res := make(MapGraph[T], n)
	for i := range res {
		res[i] = make(map[int]T)
	}
	return res
}

// Add inserts a directed edge from v to w with weight c.
// It overwrites the previous weight if this edge already exists.
func (g MapGraph[T]) Add(u, v int, weight T) {
	// Make sure not to break internal state.
	if v < 0 || v >= len(g) {
		panic("vertex out of range: " + strconv.Itoa(v))
	}
	if g[u] == nil {
		g[u] = make(map[int]T, initialMapSize)
	}
	g[u][v] = weight
	// fmt.Println("  Add", u, v, weight)
}

func (g MapGraph[T]) AddNoChecks(u, v int, weight T) {
	g[u][v] = weight
}

// AddBoth inserts edges with weigth c between v and w.
// It overwrites the previous weights if these edges already exist.
func (g MapGraph[T]) AddBoth(u, v int, weight T) {
	g.Add(u, v, weight)
	if u != v {
		g.Add(v, u, weight)
	}
}

func (g MapGraph[T]) AddBothNoChecks(u, v int, weight T) {
	g[u][v] = weight
	g[v][u] = weight
}

func (g MapGraph[T]) HasEdge(u, v int) bool {
	if u < 0 || u >= len(g) {
		return false
	}
	_, ok := g[u][v]
	return ok
}

func (g MapGraph[T]) HasEdgeNoChecks(u, v int) bool {
	_, ok := g[u][v]
	return ok
}

func (g MapGraph[T]) Weight(u, v int) T {
	if u < 0 || u >= len(g) {
		return 0
	}
	return g[u][v]
}

// Delete removes an edge from v to w.
func (g MapGraph[T]) Delete(u, v int) {
	delete(g[u], v)
}

// DeleteBoth removes all edges between v and w.
func (g MapGraph[T]) DeleteBoth(u, v int) {
	g.Delete(u, v)
	if u != v {
		g.Delete(v, u)
	}
}

func (g MapGraph[T]) ShortestPathNP(src, tgt int, MAX T) T {
	n := len(g)
	dist := make([]T, n)
	for i := range dist {
		dist[i] = MAX
	}
	dist[src] = 0
	// Dijkstra's algorithm
	q := NewMinCostQ(dist)
	q.Push(src)
	for q.Len() > 0 {
		v := q.Pop()
		if v == tgt {
			return dist[tgt]
		}
		fmt.Println("v=", v, ", dist=", dist[v])
		for w, d := range g[v] {
			alt := dist[v] + d
			switch {
			case dist[w] == MAX:
				dist[w] = alt
				q.Push(w)
			case alt < dist[w]:
				dist[w] = alt
				q.Fix(w)
			}
		}
	}
	return MAX
}

// TopoSort assumes that the graph is acyclic
func (g MapGraph[T]) TopoSort() []int {
	result := make([]int, 0, len(g))
	seen := make([]bool, len(g))
	var f func(v int)
	f = func(u int) {
		seen[u] = true
		for v := range g[u] {
			if !seen[v] {
				f(v)
			}
		}
		result = append(result, u)
	}
	for v := range g {
		if !seen[v] {
			f(v)
		}
	}
	Reverse(result)
	return result
}

func (g MapGraph[T]) LongestPathsAcyclicNP(src int, MAX T) []T {
	n := len(g)
	dist := make([]T, n)
	for i := range dist {
		dist[i] = -MAX
	}
	dist[src] = 0
	xs := g.TopoSort()
	for _, u := range xs {
		if dist[u] != -MAX {
			for v, d := range g[u] {
				if dist[v] < dist[u]+d {
					dist[v] = dist[u] + d
				}
			}
		}
	}
	return dist
}
