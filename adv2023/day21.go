package adv2023

import (
	"bufio"
	"fmt"
	"os"

	"github.com/norbs57/advofcode/lib"
	"golang.org/x/exp/maps"
)

func Day21() {
	sc := bufio.NewScanner(os.Stdin)
	grid := make([][]byte, 0)
	var start [2]int
	for sc.Scan() {
		line := make([]byte, len(sc.Bytes()))
		copy(line, sc.Bytes())
		for j, ch := range line {
			if ch == 'S' {
				start = [2]int{len(grid), j}
			}
		}
		grid = append(grid, line)
	}
	res := [2]int{}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	n := len(grid)
	// assume grid is square
	inGrid := func(i, j int) bool {
		return 0 <= i && i < len(grid) && 0 <= j && j < n
	}
	// Part 1
	steps := 64
	sqs := [][2]int{start}
	for i := 0; i < steps; i++ {
		next := make(map[[2]int]bool)
		for _, sq := range sqs {
			for _, d := range dirs {
				x0, x1 := sq[0]+d[0], sq[1]+d[1]
				if inGrid(x0, x1) && grid[x0][x1] != '#' {
					next[[2]int{x0, x1}] = true
				}
			}
		}
		sqs = maps.Keys(next)
	}
	res[0] = len(sqs)
	// Part 2
	s := 26501365
	// k/n is even, hence l needs to be even too
	l := 2
	sections := make([][2]int, 0)
	sectionMap := make(map[[2]int]int)
	for i := -l; i <= l; i++ {
		for j := -l; j <= l; j++ {
			p := [2]int{i * n, j * n}
			sectionMap[p] = len(sections)
			sections = append(sections, p)
		}
	}
	cs := make([]int, len(sections))
	sqs = [][2]int{start}
	steps = l*n + s%n
	for i := 1; i <= steps; i++ {
		for i := range cs {
			cs[i] = 0
		}
		next := make(map[[2]int]bool)
		for _, sq := range sqs {
			for _, d := range dirs {
				x := [2]int{sq[0] + d[0], sq[1] + d[1]}
				y := x
				for i := range y {
					if y[i] < 0 {
						y[i] += (1 + (-y[i] / n)) * n
					}
					y[i] %= n
				}
				if grid[y[0]][y[1]] != '#' {
					_, ok := next[x]
					if !ok {
						z := [2]int{x[0] - lib.Mod(x[0], n), x[1] - lib.Mod(x[1], n)}
						cs[sectionMap[z]]++
					}
					next[x] = true
				}
			}
		}
		sqs = maps.Keys(next)
		if i%n == s%n && i/n == l {
			break
		}
	}
	csOf := func(p [2]int) int {
		idx := sectionMap[[2]int{p[0] * n, p[1] * n}]
		return cs[idx]
	}
	// see data/log21.txt for cs values
	f := func(n int) int {
		res := 0
		// tips
		for _, p := range [][2]int{{-2, 0}, {0, -2}, {0, 2}, {2, 0}} {
			res += csOf(p)
		}
		// diags
		for _, p := range [][2]int{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}} {
			res += n * csOf(p)
		}
		// between diags
		for _, p := range [][2]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}} {
			res += (n - 1) * csOf(p)
		}
		// dist one from center
		res += (n * n) * csOf([2]int{1, 0})
		// center
		res += (n - 1) * (n - 1) * csOf([2]int{})
		return res
	}
	res[1] = f(s / n)
	fmt.Println(res[0], res[1])
}
