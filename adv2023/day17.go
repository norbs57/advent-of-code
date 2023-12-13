package adv2023

import (
	"bufio"
	"fmt"
	"os"

	"github.com/norbs57/advofcode/lib"
)

func Day17() {
	sc := bufio.NewScanner(os.Stdin)
	grid := make([][]int, 0)
	for sc.Scan() {
		m := len(sc.Bytes())
		if m > 0 {
			row := make([]int, m)
			for j, ch := range sc.Bytes() {
				row[j] = int(ch - '0')
			}
			grid = append(grid, row)
		}
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	n, m := len(grid), len(grid[0])
	inGrid := func(r, c int) bool {
		return 0 <= r && r < n && 0 <= c && c < m
	}
	type path struct {
		row, col int
		dir      int
		cost     int
	}
	h := lib.NewHeap[path](func(a, b path) bool {
		return a.cost < b.cost
	})
	cost := make([][][]int, n)
	const maxCost = 1e5
	for i := range cost {
		cost[i] = lib.MkSlice2Filled[int](m, 4, maxCost)
	}
	addPathsDir := func(p path, d, minSteps, maxSteps int) {
		newCost := 0
		for j := 1; j <= maxSteps; j++ {
			r1, c1 := p.row+j*dirs[d][0], p.col+j*dirs[d][1]
			if !inGrid(r1, c1) {
				return
			}
			newCost += grid[r1][c1]
			if j >= minSteps {
				cost1 := p.cost + newCost
				if cost[r1][c1][d] > cost1 {
					cost[r1][c1][d] = cost1
					h.Push(path{r1, c1, d, cost1})
				}
			}
		}
	}
	addPaths := func(p path, minSteps, maxSteps int) {
		for d := range dirs {
			if d != p.dir && d != (p.dir+2)%4 {
				addPathsDir(p, d, minSteps, maxSteps)
			}
		}
	}
	shortestPath := func(minSteps, maxSteps int) int {
		p := path{}
		for d := 0; d < 2; d++ {
			addPathsDir(p, d, minSteps, maxSteps)
		}
		tgt := [2]int{n - 1, m - 1}
		for h.Len() > 0 {
			v := h.Pop()
			r, c, d, cv := v.row, v.col, v.dir, v.cost
			if cv <= cost[r][c][d] {
				if r == tgt[0] && c == tgt[1] {
					return cv
				}
				addPaths(v, minSteps, maxSteps)
			}
		}
		return -1
	}
	sol := [2]int{}
	sol[0] = shortestPath(1, 3)
	h.Clear()
	for i := range cost {
		cost[i] = lib.MkSlice2Filled[int](m, 4, maxCost)
	}
	sol[1] = shortestPath(4, 10)
	fmt.Println(sol[0], sol[1])
}
