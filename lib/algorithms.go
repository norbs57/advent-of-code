package lib

import "math"

func MergeIntvs(intvs [][2]int) [][2]int {
	// assume intvs have been lex sorted
	if len(intvs) <= 1 {
		return intvs
	}
	result := make([][2]int, 1)
	result[0] = intvs[0]
	intvs = intvs[1:]
	last := 0
	for _, intv := range intvs {
		if intv[0] <= result[last][1] {
			result[last][1] = max(intv[1], result[last][1])
		} else {
			result = append(result, intv)
			last++
		}
	}

	return result
}

func FloydWarshallEdgesOne(g [][]int) (dist [][]int) {
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
