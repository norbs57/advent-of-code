package adv2023

import (
	"bufio"
	"fmt"
	"os"
)

func Day16() {
	sc := bufio.NewScanner(os.Stdin)
	grid := make([][]byte, 0)
	for sc.Scan() {
		row := make([]byte, len(sc.Bytes()))
		copy(row, sc.Bytes())
		grid = append(grid, row)
	}
	n, m := len(grid), len(grid[0])
	inGrid := func(r, c int) bool {
		return 0 <= r && r < n && 0 <= c && c < m
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	refls := [][4]int{{3, 2, 1, 0}, {1, 0, 3, 2}}
	splits := [][][]int{
		{{1, 3}, {1}, {1, 3}, {3}},
		{{0}, {0, 2}, {2}, {0, 2}},
	}
	f := func(start [3]int) int {
		state := make([][][4]bool, len(grid))
		for i := range grid {
			state[i] = make([][4]bool, len(grid[i]))
		}
		visited := 0
		rays := [][3]int{}
		rays = append(rays, start)
		for len(rays) > 0 {
			next := make([][3]int, 0)
			for _, ray := range rays {
				r, c, d := ray[0], ray[1], ray[2]
				r, c = r+dirs[d][0], c+dirs[d][1]
				var ds []int
				if !inGrid(r, c) {
					continue
				}
				switch grid[r][c] {
				case '/':
					ds = []int{refls[0][d]}
				case '\\':
					ds = []int{refls[1][d]}
				case '|':
					ds = splits[0][d]
				case '-':
					ds = splits[1][d]
				default: // '.'
					ds = []int{d}
				}
				for _, d := range ds {
					if !state[r][c][d] {
						if state[r][c] == [4]bool{} {
							visited++
						}
						state[r][c][d] = true
						next = append(next, [3]int{r, c, d})
					}
				}
			}
			rays = next
		}
		return visited
	}
	part1 := f([3]int{0, -1, 0})
	fmt.Println(part1)
	part2 := 0
	for i := range grid {
		for _, start := range [][3]int{{i, -1, 0}, {i, m, 2}} {
			part2 = max(part2, f(start))
		}
	}
	for i := range grid[0] {
		for _, start := range [][3]int{{-1, i, 1}, {n, i, 3}} {
			part2 = max(part2, f(start))
		}
	}
	fmt.Println(part2)
}
