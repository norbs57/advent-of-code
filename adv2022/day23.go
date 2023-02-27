package adv2022

import (
	"bufio"
	"fmt"
	"os"

	"github.com/norbs57/advofcode/lib"
)

func Day23() {
	sc := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	maxW := 0
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
		maxW = max(maxW, len(line))
	}
	const rounds = 100
	elves := []*[2]int{}
	grid := make([][]bool, len(lines)+2*rounds)
	for i := range grid {
		grid[i] = make([]bool, len(grid))
	}
	for i, s := range lines {
		for j, ch := range s {
			if ch == '#' {
				r, c := rounds+i, rounds+j
				grid[r][c] = true
				lib.Append(&elves, &[2]int{r, c})
			}
		}
	}
	dir := [][2]int{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			lib.Append(&dir, [2]int{i, j})
		}
	}
	dirs := [][][2]int{
		{dir[0], dir[1], dir[2]},
		{dir[6], dir[7], dir[8]},
		{dir[0], dir[3], dir[6]},
		{dir[2], dir[5], dir[8]},
	}
	hasNbs := func(r, c int) bool {
		for _, d := range dir {
			if d[0] != 0 || d[1] != 0 {
				r1, c1 := r+d[0], c+d[1]
				if grid[r1][c1] {
					return true
				}
			}
		}
		return false
	}
	propOfElf := func(r, c, startDir int) [2]int {
	loop:
		for i := 0; i < 4; i++ {
			d := (startDir + i) % 4
			for _, p := range dirs[d] {
				r1, c1 := r+p[0], c+p[1]
				if grid[r1][c1] {
					continue loop
				}
			}
			return [2]int{r + dirs[d][1][0], c + dirs[d][1][1]}
		}
		return [2]int{r, c}
	}
	part1 := func() {
		minR, minC := len(grid), len(grid)
		maxR, maxC := -1, -1
		for _, e := range elves {
			r, c := e[0], e[1]
			minR = min(minR, r)
			minC = min(minC, c)
			maxR = max(maxR, r)
			maxC = max(maxC, c)
		}
		result := (maxR-minR+1)*(maxC-minC+1) - len(elves)
		fmt.Println(result)
	}
	for round := 0; true; round++ {
		if round == 10 {
			part1()
		}
		startDir := round % 4
		finished := true
		props := make(map[[2]int][]*[2]int)
		for _, elf := range elves {
			r, c := elf[0], elf[1]
			if hasNbs(r, c) {
				finished = false
				pe := propOfElf(r, c, startDir)
				if pe != *elf {
					props[pe] = append(props[pe], elf)
				}
			}
		}
		for p, es := range props {
			if len(es) == 1 {
				e := es[0]
				r, c := p[0], p[1]
				grid[r][c] = true
				grid[e[0]][e[1]] = false
				e[0] = r
				e[1] = c
			}
		}
		if finished {
			fmt.Println(round + 1)
			return
		}
	}
}
