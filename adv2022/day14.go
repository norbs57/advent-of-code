package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day14() {
	sc := bufio.NewScanner(os.Stdin)
	height := 1001
	width := 1001
	const sandCol = 500
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	maxRock := 0
	for sc.Scan() {
		items := strings.Split(sc.Text(), "->")
		coords := make([][2]int, 0)
		for _, item := range items {
			coord := [2]int{}
			coordStr := strings.Split(item, ",")
			for i, s := range coordStr {
				coord[i], _ = strconv.Atoi(strings.TrimSpace(s))
			}
			// switch coordinates to row first, col second
			coord[0], coord[1] = coord[1], coord[0]
			coords = append(coords, coord)
			if coord[0] > maxRock {
				maxRock = coord[0]
			}
		}
		c := coords[0]
		grid[c[0]][c[1]] = true
		for i := 1; i < len(coords); i++ {
			d := coords[i]
			var delta [2]int
			for i := range delta {
				delta[i] = lib.Sign(c[i] - d[i])
			}
			for x := d; x != c; {
				grid[x[0]][x[1]] = true
				x[0] += delta[0]
				x[1] += delta[1]
			}
			c = d
		}
	}
	ground := maxRock + 2
	drop := func() [2]int {
		sand := [2]int{0, 500}
	loop:
		for sand[0] < ground-1 {
			sand[0]++
			for _, i := range [3]int{0, -1, 1} {
				if !grid[sand[0]][sand[1]+i] {
					sand[1] += i
					continue loop
				}
			}
			sand[0]--
			break
		}
		grid[sand[0]][sand[1]] = true
		return sand
	}
	resting := 0
	part := 1
	for {
		sand := drop()
		if part == 1 && sand[0] > maxRock {
			fmt.Println(resting)
			part = 2
		}
		resting++
		if sand[0] == 0 {
			fmt.Println(resting)
			break
		}
	}
}
