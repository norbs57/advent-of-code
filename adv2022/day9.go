package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day9() {
	sc := bufio.NewScanner(os.Stdin)
	dist := func(a, b [2]int) int {
		d0, d1 := lib.Abs(b[0]-a[0]), lib.Abs(b[1]-a[1])
		return max(d0, d1)
	}
	fromTo := func(a, b [2]int) [2]int {
		return [2]int{b[0] - a[0], b[1] - a[1]}
	}
	tMove := func(d [2]int) [2]int {
		for i := range d {
			if lib.Abs(d[i]) == 2 {
				d[i] = d[i] / 2
			}
		}
		return d
	}
	add := func(a [2]int, b [2]int) [2]int {
		return [2]int{a[0] + b[0], a[1] + b[1]}
	}
	input := make([]string, 0)
	for sc.Scan() {
		input = append(input, sc.Text())
	}
	for _, knots := range []int{2, 10} {
		tail := knots - 1
		rope := make([][2]int, knots)
		seen := make(map[[2]int]bool)
		seen[rope[tail]] = true
		result := 1
		for _, line := range input {
			move := strings.Fields(line)
			dir := [2]int{}
			switch move[0][0] {
			// orientation: rows/cols going up/right
			case 'R':
				dir[1] = 1
			case 'U':
				dir[0] = 1
			case 'L':
				dir[1] = -1
			case 'D':
				dir[0] = -1
			}
			n, _ := strconv.Atoi(move[1])
			for i := 0; i < n; i++ {
				rope[0] = add(rope[0], dir)
				for k := 1; k < knots; k++ {
					prev, knot := rope[k-1], rope[k]
					if dist(knot, prev) == 2 {
						v := fromTo(knot, prev)
						tmv := tMove(v)
						rope[k] = add(knot, tmv)
					}
				}
				if !seen[rope[tail]] {
					result++
					seen[rope[tail]] = true
				}
			}
		}
		fmt.Println(result)
	}
}
