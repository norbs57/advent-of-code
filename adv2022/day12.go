package adv2022

import (
	"bufio"
	"fmt"
	"os"
)

func Day12() {
	sc := bufio.NewScanner(os.Stdin)
	var src, tgt int
	grid := make([]int, 0)
	rows := 0
	zs := make([]int, 0)
	for i := 0; sc.Scan(); {
		line := sc.Text()
		for _, ch := range line {
			switch ch {
			case 'S':
				src = i
				ch = 'a'
			case 'E':
				tgt = i
				ch = 'z'
			}
			grid = append(grid, int(ch-'a'))
			if grid[i] == 0 {
				zs = append(zs, i)
			}
			i++
		}
		rows++
	}
	cols := len(grid) / rows
	nbs := func(i int) []int {
		col := i % cols
		result := make([]int, 0)
		add := func(j int) {
			result = append(result, j)
		}
		j := i - cols
		if j >= 0 {
			add(j)
		}
		j = i + cols
		if j < len(grid) {
			add(j)
		}
		if col > 0 {
			add(i - 1)
		}
		if col < cols-1 {
			add(i + 1)
		}
		return result
	}
	canGoUp := func(i, j int) bool {
		return grid[j] <= grid[i]+1
	}
	canGoDown := func(i, j int) bool {
		return canGoUp(j, i)
	}
	n := len(grid)
	_, distSrc := ShortestPathsLee(n, nbs, canGoUp, src)
	fmt.Println(distSrc[tgt])
	_, distTgt := ShortestPathsLee(n, nbs, canGoDown, tgt)
	minZDist := distTgt[zs[0]]
	for _, z := range zs {
		if distTgt[z] < minZDist {
			minZDist = distTgt[z]
		}
	}
	fmt.Println(minZDist)
}

func ShortestPathsLee(n int, nbs func(i int) []int, canGo func(i, j int) bool, src int) (parent []int, dist []int) {
	// for mazes, simpler than Dijkstra
	maxDist := n
	dist = make([]int, n)
	parent = make([]int, n)
	for i := range dist {
		dist[i] = maxDist
		parent[i] = -1
	}
	dist[src] = 0
	// Lee's Algorithm
	q := make([]int, 1, n)
	qNext := make([]int, 0, n)
	q[0] = (src)
	for len(q) > 0 {
		for _, v := range q {
			for _, w := range nbs(v) {
				if dist[w] < maxDist || !canGo(v, w) {
					continue
				}
				dist[w] = dist[v] + 1
				parent[w] = v
				qNext = append(qNext, w)
			}
		}
		q, qNext = qNext, q
		qNext = qNext[:0]
	}
	return
}
