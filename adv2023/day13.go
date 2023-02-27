package adv2023

import (
	"bufio"
	"fmt"
	"os"

	"github.com/norbs57/advofcode/lib"
)

func Day13() {
	sc := bufio.NewScanner(os.Stdin)

	ReadGrid := func() ([]string, bool) {
		result := make([]string, 0)
		for sc.Scan() {
			line := sc.Text()
			if len(line) > 0 {
				result = append(result, line)
			} else {
				return result, true
			}
		}
		return result, false
	}
	ids := func(xs []string) []lib.Bits {
		res := make([]lib.Bits, len(xs))
		for i, row := range xs {
			for j, ch := range row {
				if ch == '#' {
					res[i] = res[i].Set(j)
				}
			}
		}
		return res
	}
	sum := [2]int{}
	mult := [2]int{100, 1}
	for {
		grid, ok := ReadGrid()
		tGrid := lib.TransposeStrings(grid)
		idsGrids := [2][]lib.Bits{ids(grid), ids(tGrid)}
		// First part
		rfl := [2]int{}
		for i, ids := range idsGrids {
			for _, x := range Symmetries(ids) {
				rfl[i] = x
				break
			}
		}
		sum[0] += mult[0]*rfl[0] + mult[1]*rfl[1]
		// Second part
	loop:
		for k, ids := range idsGrids {
			otherIds := idsGrids[(k+1)%2]
			for i := 0; i < len(ids); i++ {
				for j := 0; j < len(otherIds); j++ {
					b := &idsGrids[k][i]
					*b = b.Toggle(j)
					for _, sr := range Symmetries(idsGrids[k]) {
						if sr != rfl[k] {
							sum[1] += mult[k] * sr
							break loop
						}
					}
					*b = b.Toggle(j)
				}
			}
		}
		if !ok {
			break
		}
	}
	fmt.Println(sum[0], sum[1])
}

func Symmetries[T comparable](xs []T) []int {
	res := []int{}
	for i := 1; i < len(xs); i++ {
		if xs[i-1] != xs[i] {
			continue
		}
		success := true
		for j := 1; true; j++ {
			if i-1-j < 0 || i+j >= len(xs) {
				break
			}
			if xs[i-1-j] != xs[i+j] {
				success = false
				break
			}
		}
		if success {
			res = append(res, i)
		}
	}
	return res
}
