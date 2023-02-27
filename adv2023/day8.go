package adv2023

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/norbs57/advofcode/lib"
	"golang.org/x/exp/maps"
)

func Day8() {
	locs, next, N := Day8Input()
	Day8a(next)
	Day8b(locs, next, N)
}

func Day8Input() ([]string, func(string, int) string, int) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	instructions := lib.ScanText(sc)
	succ := make(map[string][2]string)
	for sc.Scan() {
		loc := sc.Text()
		sc.Scan()
		left := lib.ScanText(sc)
		left = left[1 : len(left)-1]
		right := lib.ScanText(sc)
		right = right[:len(right)-1]
		succ[loc] = [2]string{left, right}
	}
	next := func(loc string, steps int) string {
		switch instructions[steps%len(instructions)] {
		case 'L':
			return succ[loc][0]
		case 'R':
			return succ[loc][1]
		}
		return ""
	}
	locs := maps.Keys(succ)
	sort.Strings(locs)
	return locs, next, len(instructions)
}

func Day8a(next func(string, int) string) {
	steps := 0
	loc := "AAA"
	end := "ZZZ"
	for {
		loc = next(loc, steps)
		steps++
		if loc == end {
			fmt.Println(steps)
			break
		}
	}
}

func Day8b(locs []string, next func(string, int) string, N int) {
	aLocs := make([]string, 0)
	for _, loc := range locs {
		if strings.HasSuffix(loc, "A") {
			aLocs = append(aLocs, loc)
		}
	}
	period := make(map[string]int)
	for _, aLoc := range aLocs {
		seen := make([]map[string]int, N)
		for i := range seen {
			seen[i] = make(map[string]int)
		}
		seen[0][aLoc] = 0
		steps := 0
		zLocs := make(map[string][]int)
		loc := aLoc
		for {
			loc = next(loc, steps)
			steps++
			if strings.HasSuffix(loc, "Z") {
				zLocs[loc] = append(zLocs[loc], steps)
			}
			s := steps % N
			lookup, ok := seen[s][loc]
			if ok {
				period[aLoc] = steps - lookup
				break
			}
			seen[s][loc] = steps
		}
	}

	// AAA: steps=18025, period=18023, state=(2 VXH) zLocs= map[ZZZ:[18023]]
	// BFA: steps=19642, period=19637, state=(5 XND) zLocs= map[MGZ:[19637]]
	// DFA: steps=21255, period=21251, state=(4 KKR) zLocs= map[TQZ:[21251]]
	// QJA: steps=11569, period=11567, state=(2 TSH) zLocs= map[BNZ:[11567]]
	// SBA: steps=14265, period=14257, state=(8 JXQ) zLocs= map[SSZ:[14257]]
	// XFA: steps=16413, period=16409, state=(4 PGX) zLocs= map[BKZ:[16409]]

	lcm := 1
	for _, aLoc := range aLocs {
		lcm = lib.Lcm(lcm, period[aLoc])
	}
	fmt.Println(lcm)

}
