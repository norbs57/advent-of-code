package adv2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day9() {
	// polynom fitting
	lines := lib.ReadLines()
	total := [2]int{}
	for _, line := range lines {
		items := strings.Fields(line)
		nums := make([]int, len(items))
		for i, item := range items {
			nums[i], _ = strconv.Atoi(item)
		}
		x := Extrapolate(nums)
		total[0] += x[0]
		total[1] += x[1]
	}
	fmt.Println(total[0])
	fmt.Println(total[1])
}

func Extrapolate(xs []int) [2]int {
	xss := make([][]int, 1)
	xss[0] = xs
	for len(xs) > 0 {
		next := make([]int, 0)
		allZero := true
		for i := 1; i < len(xs); i++ {
			di := xs[i] - xs[i-1]
			next = append(next, di)
			allZero = allZero && di == 0
		}
		if allZero {
			res := [2]int{}
			for len(xss) > 0 {
				ys := lib.PopBack(&xss)
				res[0] += ys[len(ys)-1]
				res[1] = ys[0] - res[1]
			}
			return res
		} else {
			xss = append(xss, next)
			xs = next
		}
	}
	return [2]int{}
}
