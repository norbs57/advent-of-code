package lib

import (
	"math"
)

func FloydWarshallUnitEdges(g [][]int) (dist [][]int) {
	// Floyd-Warshall algorith for a graph where all edges have length one
	// assumes that graph has no self-loops
	const inf = math.MaxInt
	N := len(g)
	dist = make([][]int, N)
	for i := range dist {
		dist[i] = make([]int, N)
		for j := range dist[i] {
			if i != j {
				dist[i][j] = inf
			}
		}
	}
	for u := range g {
		for _, v := range g[u] {
			dist[u][v] = 1
		}
	}
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if dist[i][k] < inf && dist[k][j] < inf {
					dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
				}
			}
		}
	}
	return
}

func SecantSearch(f func(float64) float64, a, b, precision float64) float64 {
	// secant method for finding root of function on interval [a,b]
	// assumes that a and b are "sufficiently close" to the root
	// see https://en.wikipedia.org/wiki/Secant_method
	fa, fb := f(a), f(b)
	for math.Abs(b-a) >= precision {
		c := a - fa*(b-a)/(fb-fa)
		fc := f(c)
		if fc == 0 {
			return c
		}
		a, b = b, c
		fa, fb = fb, fc
	}
	return (a + b) / 2
}

// https://en.wikipedia.org/wiki/Stoer%E2%80%93Wagner_algorithm
// https://github.com/kth-competitive-programming/kactl
// No prio queue, hence runtime is O(V^3)

func GlobalMinCut(mat [][]int) (int, []int) {
	bestCutSize, bestCut := math.MaxInt, []int{}
	n := len(mat)
	co := make([][]int, n)
	for i := range co {
		co[i] = []int{i}
	}
	for ph := 1; ph < n; ph++ {
		w := make([]int, n)
		copy(w, mat[0])
		var s, t int
		for it := 0; it < n-ph; it++ { // O(V^2) -> O(E log V) with prio. queue
			w[t] = math.MinInt
			s = t
			t = MaxIdx(w)
			for i := 0; i < n; i++ {
				w[i] += mat[t][i]
			}
		}
		cutSize := w[t] - mat[t][t]
		if cutSize < bestCutSize {
			bestCutSize = cutSize
			bestCut = make([]int, len(co[t]))
			copy(bestCut, co[t])
		}
		co[s] = append(co[s], co[t]...)
		for i := 0; i < n; i++ {
			mat[s][i] += mat[t][i]
		}
		for i := 0; i < n; i++ {
			mat[i][s] = mat[s][i]
		}
		mat[0][t] = math.MinInt
	}
	return bestCutSize, bestCut
}
