package adv2023

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day6() {
	lines := lib.ReadLines()
	items := make([][]string, len(lines))
	bigrace := [2]int{}
	for i := range bigrace {
		items[i] = strings.Fields(lines[i])[1:]
		str := ""
		for _, s := range items[i] {
			str += s
		}
		bigrace[i], _ = strconv.Atoi(str)
	}
	n := len(items[0])
	times := make([]int, n)
	dists := make([]int, n)
	for i := range times {
		times[i], _ = strconv.Atoi(items[0][i])
		dists[i], _ = strconv.Atoi(items[1][i])
	}
	solutions := func(t, d int) int {
		tf, df := float64(t), float64(d)
		xf0, xf1, _ := lib.QuadraticSolve(1.0, -tf, df)
		x0, x1 := int(math.Ceil(xf0)), int(math.Floor(xf1))
		return x1 - x0 + 1
	}
	ways := 1
	for i := range times {
		ways *= solutions(times[i], dists[i])
	}
	fmt.Println(ways)
	fmt.Println(solutions(bigrace[0], bigrace[1]))
}
