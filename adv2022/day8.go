package adv2022

import (
	"bufio"
	"fmt"
	"os"
)

func Day8() {
	sc := bufio.NewScanner(os.Stdin)
	grid := make([][]int, 0)
	for sc.Scan() {
		line := sc.Text()
		nums := make([]int, len(line))
		for i := range nums {
			nums[i] = int(line[i] - '0')
		}
		grid = append(grid, nums)
	}
	rows := len(grid)
	cols := len(grid[0])
	visible := make([][]bool, rows)
	for i := range visible {
		visible[i] = make([]bool, cols)
	}
	markVis := func(x0, x1, d0, d1, n int) {
		maxH := grid[x0][x1]
		visible[x0][x1] = true
		for i := 1; i < n; i++ {
			x0, x1 = x0+d0, x1+d1
			h := grid[x0][x1]
			if h > maxH {
				maxH = h
				visible[x0][x1] = true
			}
		}
	}
	for i := 0; i < rows; i++ {
		markVis(i, 0, 0, 1, cols)
		markVis(i, cols-1, 0, -1, cols)
	}
	for i := 0; i < cols; i++ {
		markVis(0, i, 1, 0, rows)
		markVis(rows-1, i, -1, 0, rows)
	}
	result := 0
	for _, row := range visible {
		for _, b := range row {
			if b {
				result++
			}
		}
	}
	visFromInDir := func(x0, x1, d0, d1, n int) int {
		result := 0
		maxH := grid[x0][x1]
		for i := 0; i < n; i++ {
			x0, x1 = x0+d0, x1+d1
			result++
			if grid[x0][x1] >= maxH {
				break
			}
		}
		// fmt.Println("=", result)
		return result
	}
	calcScenic := func(x0, x1 int) int {
		result := 1
		result *= visFromInDir(x0, x1, 0, 1, cols-x1-1)
		result *= visFromInDir(x0, x1, 1, 0, rows-x0-1)
		result *= visFromInDir(x0, x1, 0, -1, x1)
		result *= visFromInDir(x0, x1, -1, 0, x0)
		return result
	}
	fmt.Println(result)
	maxScenic := 1
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			sc := calcScenic(i, j)
			if sc > maxScenic {
				maxScenic = sc
			}
		}
	}
	fmt.Println(maxScenic)
}
