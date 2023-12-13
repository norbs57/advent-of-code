package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day16() {
	// Part 2 uses a branch-and-bound algorithm with memoization
	// It could be made faster by adding heuristics for selecting moves
	flows, tunnels, _, start := Day16Input()
	dist := lib.FloydWarshallUnitEdges(tunnels)
	nonZeros := make([]int, 0)
	for i := range flows {
		if flows[i] > 0 {
			nonZeros = append(nonZeros, i)
		}
	}
	var f func(int, int, lib.Bits) int
	f = func(v, t int, open lib.Bits) int {
		best := 0
		for _, w := range nonZeros {
			if !open.Has(w) && dist[v][w] < t-1 {
				tw := t - dist[v][w] - 1
				release := tw * flows[w]
				release += f(w, tw, open.Set(w))
				if release > best {
					best = release
				}
			}
		}
		return best
	}
	best1 := f(start, 30, 0)
	fmt.Println(best1)
	totalBest := 0
	table := make(map[[5]int]int)
	// var cutoffs, lookups int
	var g func([2]int, [2]int, lib.Bits, int) int
	g = func(v, t [2]int, open lib.Bits, current int) int {
		if v[0] > v[1] {
			return g(lib.Swap(v), lib.Swap(t), open, current)
		}
		args := [5]int{v[0], v[1], t[0], t[1], int(open)}
		lookup, found := table[args]
		if found {
			// lookups++
			return lookup
		}
		potential := current
		tRem := t[0] + t[1]
		for _, v := range nonZeros {
			if !open.Has(v) {
				tRem -= 2
				potential += flows[v] * tRem
			}
		}
		if potential < totalBest {
			// cutoffs++
			return 0
		}
		best := 0
		// could add heuristics for selecting moves...
		for _, w := range nonZeros {
			if !open.Has(w) {
				for i := range v {
					if dist[v[i]][w] < t[i]-1 {
						tw := t[i] - dist[v[i]][w] - 1
						release := tw * flows[w]
						v1 := v
						v1[i] = w
						t1 := t
						t1[i] = tw
						release += g(v1, t1, open.Set(w), current+release)
						if release > best {
							best = release
						}
					}
				}
			}
		}
		if best > totalBest {
			totalBest = best
		}
		table[args] = best
		return best
	}
	best2 := g([2]int{start, start}, [2]int{26, 26}, 0, 0)
	// fmt.Println("cutoffs=", cutoffs, "lookups=", lookups) 19451478 7072644
	fmt.Println(best2)
}

func Day16Input() ([]int, [][]int, []string, int) {
	intOfValve := make(map[string]int)
	flows := make([]int, 0)
	tunnels := make([][]int, 0)
	start := -1
	sc := bufio.NewScanner(os.Stdin)
	lookupPlus := func(s string) int {
		_, found := intOfValve[s]
		if !found {
			intOfValve[s] = len(intOfValve)
			flows = append(flows, 0)
			tunnels = append(tunnels, []int{})
			if start == -1 && s == "AA" {
				start = intOfValve[s]
			}
		}
		return intOfValve[s]
	}
	for sc.Scan() {
		line := sc.Text()
		splitF := func(r rune) bool {
			return r == ' ' || r == ',' || r == '=' || r == ';'
		}
		items := strings.FieldsFunc(line, splitF)
		v := lookupPlus(items[1])
		flows[v], _ = strconv.Atoi(items[5])
		for _, s := range items[10:] {
			w := lookupPlus(s)
			tunnels[v] = append(tunnels[v], w)
		}
	}
	valveOfString := make([]string, len(intOfValve))
	for s, v := range intOfValve {
		valveOfString[v] = s
	}
	return flows, tunnels, valveOfString, start
}
