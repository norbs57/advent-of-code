package adv2022

import (
	"bufio"
	"fmt"
	"os"

	"github.com/norbs57/advofcode/lib"
)

func Day24() {
	sc := bufio.NewScanner(os.Stdin)
	grid := make([]string, 0)
	sc.Scan()
	for sc.Scan() {
		line := sc.Text()
		if line[1] == '#' {
			break
		}
		grid = append(grid, line[1:len(line)-1])
	}
	R, C := len(grid), len(grid[0])
	inGrid := func(r, c int) bool {
		return 0 <= r && r < R && 0 <= c && c < C
	}
	isBlizz := func(r, c, t int) bool {
		r0 := lib.ModInt(r-t, R)
		r1 := lib.ModInt(r+t, R)
		c0 := lib.ModInt(c-t, C)
		c1 := lib.ModInt(c+t, C)
		return grid[r0][c] == 'v' || grid[r1][c] == '^' ||
			grid[r][c0] == '>' || grid[r][c1] == '<'
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {0, 0}}
	nbs := func(r, c, t int) [][3]int {
		result := [][3]int{}
		for _, d := range dirs {
			a, b := r+d[0], c+d[1]
			if inGrid(a, b) && !isBlizz(a, b, t+1) {
				lib.Append(&result, [3]int{a, b, t + 1})
			}
		}
		switch {
		case r == 0 && c == 0:
			lib.Append(&result, [3]int{-1, 0, t + 1})
		case r == R-1 && c == C-1:
			lib.Append(&result, [3]int{R, c, t + 1})
		case r == -1 || r == R:
			lib.Append(&result, [3]int{r, c, t + 1})
		}
		return result
	}
	type Item struct {
		Value [3]int
		Cost  int
	}
	push := func(q *lib.Heap[*Item], p [3]int, tgt [2]int) {
		d := lib.Dist2([2]int{p[0], p[1]}, tgt)
		item := Item{p, p[2] + d}
		q.Push(&item)
	}
	f := func(src [3]int, tgt [2]int) int {
		q := lib.NewHeap[*Item](func(a, b *Item) bool {
			return a.Cost < b.Cost
		})
		seen := make(map[[3]int]bool)
		seen[src] = true
		push(q, src, tgt)
		for q.Len() > 0 {
			p := q.Pop().Value
			// fmt.Println("p=", p)
			r, c, t := p[0], p[1], p[2]
			if r == tgt[0] && c == tgt[1] {
				return t
			}
			for _, nb := range nbs(r, c, t) {
				if !seen[nb] {
					seen[nb] = true
					push(q, nb, tgt)
				}
			}
		}
		return -1
	}
	atTime := func(p [2]int, t int) [3]int {
		return [3]int{p[0], p[1], t}
	}
	src := [2]int{-1, 0}
	tgt := [2]int{R, C - 1}
	t1 := f(atTime(src, 0), tgt)
	fmt.Println(t1)
	t2 := f(atTime(tgt, t1), [2]int(src))
	t3 := f(atTime(src, t2), tgt)
	fmt.Println(t3)
}
