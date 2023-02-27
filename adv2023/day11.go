package adv2023

import (
	"fmt"

	"github.com/norbs57/advofcode/lib"
)

func Day11() {
	grid := lib.ReadLines()
	noStarRows := lib.MkSliceFilled[int](len(grid), 1)
	noStarCols := lib.MkSliceFilled[int](len(grid[0]), 1)
	stars := make([][2]int, 0)
	for i, row := range grid {
		for j, ch := range row {
			if ch == '#' {
				noStarRows[i] = 0
				noStarCols[j] = 0
				stars = append(stars, [2]int{i, j})
			}
		}
	}
	pRows := lib.PrefixSums(noStarRows)
	pCols := lib.PrefixSums(noStarCols)
	for _, mult := range []int{2, 1_000_000} {
		d := 0
		for i := range stars {
			ri, ci := stars[i][0], stars[i][1]
			for j := i + 1; j < len(stars); j++ {
				rj, cj := stars[j][0], stars[j][1]
				minR, maxR := min(ri, rj), max(ri, rj)
				minC, maxC := min(ci, cj), max(ci, cj)
				d += maxR - minR + maxC - minC
				noStars := pRows[maxR] - pRows[minR] + pCols[maxC] - pCols[minC]
				d += (mult - 1) * noStars
			}
		}
		fmt.Println(d)
	}
}
