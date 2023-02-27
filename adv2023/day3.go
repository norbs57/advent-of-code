package adv2023

import (
	"fmt"

	"github.com/norbs57/advofcode/lib"
)

func Day3() {
	grid := lib.ReadLines()
	n := len(grid)
	m := len(grid[0])
	for i := range grid {
		grid[i] = grid[i] + "."
	}
	dirs := make([][2]int, 0, 8)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				dirs = append(dirs, [2]int{i, j})
			}
		}
	}
	inGrid := func(i, j int) bool {
		return 0 <= i && i < n && 0 <= j && j < m
	}
	nbs := func(i, j int) [][2]int {
		res := [][2]int{}
		for _, dir := range dirs {
			pos := [2]int{i + dir[0], j + dir[1]}
			if inGrid(pos[0], pos[1]) {
				res = append(res, pos)
			}
		}
		return res
	}
	isSymbol := func(ch rune) bool {
		return ch != '.' && !lib.IsDigit(ch)
	}
	nums := make([]int, 0, n*m)
	idOf := make([][]int, n)
	symbols := make([][2]int, 0, n*m)
	for i := range grid {
		num := 0
		start := -1
		idOf[i] = make([]int, len(grid[i]))
		for j, ch := range grid[i] {
			idOf[i][j] = -1
			if lib.IsDigit(ch) {
				num = 10*num + int(ch-'0')
				if start == -1 {
					start = j
				}
				continue
			}
			if isSymbol(ch) {
				idOf[i][j] = len(symbols)
				symbols = append(symbols, [2]int{i, j})
			}
			if start != -1 {
				for k := start; k < j; k++ {
					idOf[i][k] = len(nums)
				}
				nums = append(nums, num)
				num = 0
				start = -1
			}
		}
	}
	sum := 0
	seen := make([]bool, len(nums))
	numMaps := make([]map[int]bool, len(symbols))
	for i := range numMaps {
		numMaps[i] = make(map[int]bool)
	}
	for i, symb := range symbols {
		for _, sq := range nbs(symb[0], symb[1]) {
			id := idOf[sq[0]][sq[1]]
			if id != -1 {
				numMaps[i][id] = true
				if !seen[id] {
					seen[id] = true
					sum += nums[id]
				}
			}
		}
	}
	fmt.Println(sum)
	gearRatios := 0
	for i := range symbols {
		if len(numMaps[i]) == 2 {
			prod := 1
			for id := range numMaps[i] {
				prod *= nums[id]
			}
			gearRatios += prod
		}
	}
	fmt.Println(gearRatios)
}
