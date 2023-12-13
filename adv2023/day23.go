package adv2023

import (
	"fmt"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day23() {
	grid := lib.ReadByteSlices()
	rows := len(grid)
	cols := len(grid[0])
	start := [2]int{}
	tgt := [2]int{rows - 1, 0}
	for i := range grid[0] {
		if grid[0][i] == '.' {
			start[1] = i
		}
		if grid[rows-1][i] == '.' {
			tgt[1] = i
		}
	}
	const slopes = "v>^<"
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	nbsOf := func(i, j int) [][2]int {
		res := [][2]int{}
		ch := grid[i][j]
		idx := strings.IndexByte(slopes, ch)
		if idx != -1 {
			dv := dirs[idx]
			res = append(res, [2]int{i + dv[0], j + dv[1]})
		} else {
			for _, dv := range dirs {
				r, c := i+dv[0], j+dv[1]
				if 0 <= r && r < rows && 0 <= c && c < cols && grid[r][c] != '#' {
					res = append(res, [2]int{r, c})
				}
			}
		}
		return res
	}
	jnctns := make([][2]int, 0)
	jnctns = append(jnctns, start)
	jnctsOf := lib.MkSlice2Filled[int](rows, cols, -1)
	for i := range grid {
		for j, ch := range grid[i] {
			if ch != '.' {
				continue
			}
			nbs := nbsOf(i, j)
			isJunction := true
			for _, nb := range nbs {
				chNb := rune(grid[nb[0]][nb[1]])
				if !strings.ContainsRune(slopes, chNb) {
					isJunction = false
					break
				}
			}
			if isJunction {
				jnctsOf[i][j] = len(jnctns)
				jnctns = append(jnctns, [2]int{i, j})
			}
		}
	}
	jnctns = append(jnctns, tgt)
	jnctsOf[start[0]][start[1]] = 0
	jnctsOf[tgt[0]][tgt[1]] = len(jnctns) - 1
	for _, jc := range jnctns {
		grid[jc[0]][jc[1]] = 'o'
	}
	nextJnctn := func(i, j int) (int, int, int) {
		seen := make(map[[2]int]bool)
		seen[[2]int{i, j}] = true
		d := 0
		for {
			ch := grid[i][j]
			if ch == 'o' {
				return i, j, d
			}
			nbs := lib.Filter(nbsOf(i, j), func(p [2]int) bool {
				return !seen[p]
			})
			i, j = nbs[0][0], nbs[0][1]
			seen[[2]int{i, j}] = true
			d++
		}
	}
	tgtInt := len(jnctns) - 1
	g := lib.MkMapGraph[int](len(jnctns))
	for ji, jc := range jnctns[:len(jnctns)-1] {
		for _, nb := range nbsOf(jc[0], jc[1]) {
			ch := grid[nb[0]][nb[1]]
			idx := strings.IndexByte(slopes, ch)
			if idx != -1 {
				d := dirs[idx]
				if [2]int{nb[0] + d[0], nb[1] + d[1]} == jc {
					continue
				}
			}
			i, j, d := nextJnctn(nb[0], nb[1])
			g.Add(ji, jnctsOf[i][j], d+1)
		}
	}
	dist := g.LongestPathsAcyclicNP(0, 1e6)
	res := [2]int{}
	res[0] = dist[tgtInt]
	// part 2
	gUnd := lib.MkMapGraph[int](len(g))
	for i := range g {
		for j, d := range g[i] {
			gUnd.AddBoth(i, j, d)
		}
	}
	g = gUnd
	type path struct {
		lastVertex int
		seen       lib.Bits
		length     int
	}
	paths := make([]path, 1, 1e6)
	paths[0] = path{0, lib.Bits(0).Set(0), 0}
	maxLenTgt := 0
	for k := 0; len(paths) > 0; k++ {
		p := lib.PopFront(&paths)
		if p.lastVertex == tgtInt && p.length > maxLenTgt {
			maxLenTgt = p.length
		}
		for q, d := range g[p.lastVertex] {
			if !p.seen.Has(q) {
				p1 := path{q, p.seen.Set(q), p.length + d}
				paths = append(paths, p1)
			}
		}
	}
	res[1] = maxLenTgt
	fmt.Println(res[0], res[1])
}
