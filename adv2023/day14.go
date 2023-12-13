package adv2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day14() {
	sc := bufio.NewScanner(os.Stdin)
	grid := make([][]byte, 0)
	gridCopy := make([][]byte, len(grid))
	for sc.Scan() {
		row := make([]byte, len(sc.Bytes()))
		copy(row, sc.Bytes())
		grid = append(grid, row)
		gridCopy = append(gridCopy, row)
	}
	idx := lib.MkSlice2[int](4, 0)
	for i := range grid {
		idx[0] = append(idx[0], i)
		idx[2] = append(idx[2], len(grid)-i-1)
	}
	for i := range grid[0] {
		idx[1] = append(idx[1], i)
		idx[3] = append(idx[3], len(grid[0])-i-1)
	}
	// Part 1
	MoveNorthSouth(grid, idx[0], 1)
	fmt.Println(Load(grid))
	// Part 2
	fmt.Println(Day14P2(gridCopy, idx))
}

func MoveNorthSouth(grid [][]byte, rowIdx []int, dir int) {
	for col := range grid[0] {
		gap := 0
		for _, row := range rowIdx {
			switch grid[row][col] {
			case 'O':
				if gap > 0 {
					grid[row][col] = '.'
					grid[row-dir*gap][col] = 'O'
				}
			case '#':
				gap = 0
			default:
				gap++
			}
		}
	}
}

func MoveWestEast(grid [][]byte, colIdx []int, dir int) {
	for row := range grid {
		gap := 0
		for _, col := range colIdx {
			switch grid[row][col] {
			case 'O':
				if gap > 0 {
					grid[row][col] = '.'
					grid[row][col-dir*gap] = 'O'
				}
			case '#':
				gap = 0
			default:
				gap++
			}
		}
	}
}

func Load(grid [][]byte) int {
	res := 0
	for i, row := range grid {
		for _, ch := range row {
			if ch == 'O' {
				res += len(grid) - i
			}
		}
	}
	return res
}

func Day14P2(grid [][]byte, idx [][]int) int {
	stringOfGrid := func(gr [][]byte) string {
		var sb strings.Builder
		for i := range gr {
			sb.Write(grid[i])
		}
		return sb.String()
	}
	spin := func() {
		MoveNorthSouth(grid, idx[0], 1)
		MoveWestEast(grid, idx[1], 1)
		MoveNorthSouth(grid, idx[2], -1)
		MoveWestEast(grid, idx[3], -1)
	}
	loads := make([]int, 1)
	loads[0] = Load(grid)
	seen := make(map[string]int)
	seen[stringOfGrid(grid)] = 0
	const cycles = 1_000_000_000
	for i := 1; i <= cycles; i++ {
		spin()
		s := stringOfGrid(grid)
		lkp, ok := seen[s]
		if ok {
			period := i - lkp
			k := (cycles - lkp) % period
			return loads[lkp+k]
		}
		seen[s] = i
		loads = append(loads, Load(grid))
	}
	return -1
}
