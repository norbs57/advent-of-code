package adv2023

import (
	"fmt"
	"slices"

	"github.com/norbs57/advofcode/lib"
)

func Day10() {
	grid := lib.ReadLines()
	N, M := len(grid), len(grid[0])
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	north, east, south, west := 0, 1, 2, 3
	opposite := func(i int) int {
		return (i + 2) % 4
	}
	dirsOfChar := func(ch byte) [2]int {
		switch ch {
		case '|':
			return [2]int{north, south}
		case '-':
			return [2]int{east, west}
		case 'L':
			return [2]int{north, east}
		case 'J':
			return [2]int{north, west}
		case '7':
			return [2]int{south, west}
		case 'F':
			return [2]int{south, east}
		default:
			return [2]int{-1, -1}
		}
	}
	outDir := func(pos [2]int, inDir int) int {
		dirs := dirsOfChar(grid[pos[0]][pos[1]])
		switch opposite(inDir) {
		case dirs[0]:
			return dirs[1]
		case dirs[1]:
			return dirs[0]
		default:
			return -1
		}
	}
	findStartPos := func() [2]int {
		for row := range grid {
			for col := range grid[row] {
				if grid[row][col] == 'S' {
					return [2]int{row, col}
				}
			}
		}
		return [2]int{-1, -1}
	}
	nbsOf := func(pos [2]int, n, m int) [][2]int {
		res := make([][2]int, 0, 4)
		for _, v := range dirs {
			p := [2]int{pos[0] + v[0], pos[1] + v[1]}
			if 0 <= p[0] && p[0] < n &&
				0 <= p[1] && p[1] < m {
				res = append(res, [2]int{p[0], p[1]})
			}
		}
		return res
	}
	dirFromTo := func(p, q [2]int) int {
		q[0] -= p[0]
		q[1] -= p[1]
		return slices.Index(dirs, q)
	}
	start := findStartPos()
	findStartDir := func() int {
		nbs := nbsOf(start, N, M)
		for _, nb := range nbs {
			ch := grid[nb[0]][nb[1]]
			dch := dirsOfChar(ch)
			dToNb := dirFromTo(start, nb)
			if dch[0] == opposite(dToNb) || dch[1] == opposite(dToNb) {
				return dToNb
			}
		}
		return -1
	}
	pos, dir := start, findStartDir()
	steps := 1
	turns := make([][2]int, 0)
	for ; true; steps++ {
		v := dirs[dir]
		next := [2]int{pos[0] + v[0], pos[1] + v[1]}
		if next == start {
			break
		}
		nextDir := outDir(next, dir)
		if nextDir != dir {
			turns = append(turns, next)
		}
		pos, dir = next, nextDir
	}
	turns = append(turns, start)
	fmt.Println(steps / 2)
	fmt.Println(lib.PolygonArea(turns) - steps/2 + 1)
}

func Day10Floodfill() {
	// solution with expanded grid and using floodfill from borders to identify
	// inner points
	grid := lib.ReadLines()
	N, M := len(grid), len(grid[0])
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	north, east, south, west := 0, 1, 2, 3
	oppositeDir := func(i int) int {
		return (i + 2) % 4
	}
	dirsOfChar := func(ch byte) [2]int {
		switch ch {
		case '|':
			return [2]int{north, south}
		case '-':
			return [2]int{east, west}
		case 'L':
			return [2]int{north, east}
		case 'J':
			return [2]int{north, west}
		case '7':
			return [2]int{south, west}
		case 'F':
			return [2]int{south, east}
		default:
			return [2]int{-1, -1}
		}
	}
	outDir := func(pos [2]int, inDir int) int {
		dirs := dirsOfChar(grid[pos[0]][pos[1]])
		switch oppositeDir(inDir) {
		case dirs[0]:
			return dirs[1]
		case dirs[1]:
			return dirs[0]
		default:
			return -1
		}
	}
	findStartPos := func() [2]int {
		for row := range grid {
			for col := range grid[row] {
				if grid[row][col] == 'S' {
					return [2]int{row, col}
				}
			}
		}
		return [2]int{-1, -1}
	}
	nbsOf := func(pos [2]int, n, m int) [][2]int {
		res := make([][2]int, 0, 4)
		for _, v := range dirs {
			p := [2]int{pos[0] + v[0], pos[1] + v[1]}
			if 0 <= p[0] && p[0] < n &&
				0 <= p[1] && p[1] < m {
				res = append(res, [2]int{p[0], p[1]})
			}
		}
		return res
	}
	dirFromTo := func(p, q [2]int) int {
		q[0] -= p[0]
		q[1] -= p[1]
		return slices.Index(dirs, q)
	}
	charOfDir := func(d int) byte {
		if d%2 == 0 {
			return '|'
		} else {
			return '-'
		}
	}
	start := findStartPos()
	findStartDir := func() int {
		nbs := nbsOf(start, N, M)
		for _, nb := range nbs {
			ch := grid[nb[0]][nb[1]]
			dch := dirsOfChar(ch)
			dToNb := dirFromTo(start, nb)
			if dch[0] == oppositeDir(dToNb) || dch[1] == oppositeDir(dToNb) {
				return dToNb
			}
		}
		return -1
	}
	lN, lM := 2*N-1, 2*M-1
	lGrid := lib.MkSlice2Filled[byte](lN, lM, '.')
	steps := 1
	pos := start
	lGrid[2*pos[0]][2*pos[1]] = '1'
	dir := findStartDir()
	v := dirs[dir]
	for ; true; steps++ {
		next := [2]int{pos[0] + v[0], pos[1] + v[1]}
		lGrid[2*pos[0]+v[0]][2*pos[1]+v[1]] = charOfDir(dir)
		lGrid[2*next[0]][2*next[1]] = '*'
		if next == start {
			break
		}
		pos = next
		dir = outDir(pos, dir)
		v = dirs[dir]
	}
	for _, row := range lGrid {
		fmt.Println(string(row))
	}
	fmt.Println(steps / 2)
	borders := make(map[[2]int]bool)
	for i := range lGrid {
		borders[[2]int{i, 0}] = true
		borders[[2]int{i, lM - 1}] = true
	}
	for i := range lGrid[0] {
		borders[[2]int{0, i}] = true
		borders[[2]int{lN - 1, i}] = true
	}
	var fill func([2]int)
	fill = func(pos [2]int) {
		ch := lGrid[pos[0]][pos[1]]
		if ch != '.' {
			return
		}
		lGrid[pos[0]][pos[1]] = '0'
		for _, nb := range nbsOf(pos, len(lGrid), len(lGrid[0])) {
			fill(nb)
		}
	}
	for sq := range borders {
		fill(sq)
	}
	enclosed := 0
	for i := 0; i < lN; i += 2 {
		for j := 0; j < lM; j += 2 {
			if lGrid[i][j] == '.' {
				enclosed++
			}
		}
	}
	fmt.Println(enclosed)
}
