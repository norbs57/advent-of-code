package adv2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule19 struct {
	op   byte // <, >
	num  int
	idx  int
	dest string
}

func (r Rule19) String() string {
	return fmt.Sprintf("(%s %d %d %s)", string(r.op), r.num, r.idx, r.dest)
}

func Day19() {
	const accepted = "A"
	const rejected = "R"
	const inflow = "in"
	type wf struct {
		rules       []Rule19
		defaultDest string
	}
	sc := bufio.NewScanner(os.Stdin)
	wflMap := make(map[string]wf)
	fields := "amsx"
	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			break
		}
		wfName, rules, defaultDest := ParseWf19(line, fields)
		wflMap[wfName] = wf{rules, defaultDest}
	}
	check := func(r Rule19, part [4]int) bool {
		x := part[r.idx]
		if r.op == '<' {
			return x < r.num
		} else {
			return x > r.num
		}
	}
	processWf := func(wf string, part [4]int) string {
		w := wflMap[wf]
		for _, r := range w.rules {
			if check(r, part) {
				return r.dest
			}
		}
		return w.defaultDest
	}
	res := [2]int{}
	for sc.Scan() {
		line := sc.Text()
		part := ParsePart19(line, fields)
		wf := "in"
		for wf != accepted && wf != rejected {
			wf = processWf(wf, part)
		}
		if wf == accepted {
			for _, a := range part {
				res[0] += a
			}
		}
	}
	VolumeOf := func(cub [][2]int) int {
		res := 1
		for _, itv := range cub {
			res *= itv[1] - itv[0] + 1
		}
		return res
	}

	var f func(string, [4][2]int) int
	f = func(wf string, cub [4][2]int) int {
		if wf == accepted {
			return VolumeOf(cub[:])
		}
		if wf == rejected {
			return 0
		}
		w := wflMap[wf]
		res := 0
		for _, r := range w.rules {
			i, num := r.idx, r.num
			switch r.op {
			case '<':
				if cub[i][0] < num {
					cubSuccess := cub
					cubSuccess[i][1] = min(cub[i][1], num-1)
					res += f(r.dest, cubSuccess)
				}
				if cub[i][1] >= num {
					cub[i][0] = max(num, cub[i][0])
				} else {
					return res
				}
			case '>':
				if num < cub[i][1] {
					cubSuccess := cub
					cubSuccess[i][0] = max(cub[i][0], num+1)
					res += f(r.dest, cubSuccess)
				}
				if cub[i][0] <= num {
					cub[i][1] = min(num, cub[i][1])
				} else {
					return res
				}
			}
		}
		res += f(w.defaultDest, cub)
		return res
	}
	maxCuboid := [4][2]int{}
	for i := range maxCuboid {
		maxCuboid[i] = [2]int{1, 4000}
	}
	res[1] = f("in", maxCuboid)
	fmt.Println(res[0], res[1])

}

func ParseWf19(s string, fields string) (string, []Rule19, string) {
	rules := make([]Rule19, 0)
	items := strings.Split(s, "{")
	items[1] = strings.TrimSuffix(items[1], "}")
	for _, rs := range strings.Split(items[1], ",") {
		ps := strings.Split(rs, ":")
		if len(ps) == 1 {
			return items[0], rules, ps[0]
		}
		cond, tgt := ps[0], ps[1]
		num, _ := strconv.Atoi(cond[2:])
		idx := strings.IndexByte(fields, cond[0])
		rules = append(rules, Rule19{cond[1], num, idx, tgt})
	}
	return "", nil, ""
}

func ParsePart19(s string, fields string) [4]int {
	res := [4]int{}
	for _, item := range strings.Split(s[1:len(s)-1], ",") {
		idx := strings.IndexByte(fields, item[0])
		num, _ := strconv.Atoi(item[2:])
		res[idx] = num
	}
	return res
}
