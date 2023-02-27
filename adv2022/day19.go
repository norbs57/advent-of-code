package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day19() {
	Day19ReadScipResults()
}

func Day19TooSlow() {
	sc := bufio.NewScanner(os.Stdin)
	bps := make([][][]int, 0)
	for sc.Scan() {
		bp := Day19ParseBlueprint(sc.Text())
		bps = append(bps, bp)
	}
	sum1 := 0
	for i, bp := range bps {
		mg := Day19MaxGeodes(bp, 24)
		/// fmt.Println(mg)
		sum1 += mg * (i + 1)
	}
	fmt.Println(sum1)
	sum2 := 1
	for i := 0; i < min(3, len(bps)); i++ {
		sum2 *= Day19MaxGeodes(bps[i], 32)
	}
	fmt.Println(sum2)
}

func Day19ParseBlueprint(line string) [][]int {
	spf := func(r rune) bool {
		return r < '0' || r > '9'
	}
	fields := strings.FieldsFunc(line, spf)
	nums := make([]int, len(fields))
	for i := range fields {
		nums[i], _ = strconv.Atoi(fields[i])
	}
	bp := [][]int{{nums[1], 0, 0}, {nums[2], 0, 0}, {nums[3], nums[4], 0}, {nums[5], 0, nums[6]}}
	return bp
}

func Day19MaxGeodes(a [][]int, rounds int) int {
	type state = [9]int
	seen := make(map[state]bool)
	maxG := 0
	// lookups := 0
	var f func(state)
	f = func(st state) {
		if seen[st] {
			//	lookups++
			return
		}
		seen[st] = true
		var canBuild lib.Bits
	rLoop:
		for r := 0; r < 4; r++ {
			for k := 0; k < 3; k++ {
				if a[r][k] > st[k] {
					continue rLoop
				}
			}
			canBuild = canBuild.Set(r)
		}
		t := st[8]
		tRem := rounds - t + 1
		tRem3 := tRem
		if !canBuild.Has(3) && tRem3 > 0 {
			tRem3--
		}
		maxAch := st[3] + tRem*st[7] + (tRem3*(tRem3-1))/2
		if maxAch <= maxG {
			return
		}
		for r := 0; r < 4; r++ {
			st[r] += st[4+r]
		}
		if st[3] > maxG {
			maxG = st[3]
		}
		if t == rounds {
			return
		}
		for r := 3; r >= 0; r-- {
			if canBuild.Has(r) {
				st1 := st
				for k := 0; k < 3; k++ {
					st1[k] -= a[r][k]
				}
				st1[4+r]++
				st1[8]++
				f(st1)
				if r == 3 {
					break
				}
			}
		}
		st[8]++
		f(st)
	}
	var st0 state
	st0[0] = 0
	st0[4] = 1
	st0[8] = 1
	f(st0)
	// fmt.Println("lookups=", lookups)
	return maxG
}
