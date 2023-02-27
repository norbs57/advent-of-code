package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day11() {
	sc := bufio.NewScanner(os.Stdin)
	items := make([][]int, 0)
	type monkey struct {
		worry func(int) int
		throw func(int) int
	}
	monkeys := make([]*monkey, 0)
	setStartItems := func(s string) {
		m := len(items) - 1
		numStrs := strings.Fields(strings.Split(s, ":")[1])
		for _, str := range numStrs {
			if str[len(str)-1] == ',' {
				str = str[:len(str)-1]
			}
			n, _ := strconv.Atoi(str)
			items[m] = append(items[m], n)
		}
	}
	setOp := func(m *monkey, s string) {
		str := strings.Fields(strings.Split(s, "=")[1])
		opOfStr := map[string]func(int, int) int{
			"+": func(x, y int) int { return x + y },
			"*": func(x, y int) int { return x * y },
		}
		op := opOfStr[str[1]]
		switch {
		case str[0] == "old" && str[2] == "old":
			m.worry = func(x int) int { return op(x, x) }
		case str[0] == "old":
			y, _ := strconv.Atoi(str[2])
			m.worry = func(x int) int { return op(x, y) }
		case str[2] == "old":
			y, _ := strconv.Atoi(str[0])
			m.worry = func(x int) int { return op(x, y) }
		}
	}
	setTest := func(m *monkey, args []int) {
		m.throw = func(x int) int {
			if x%args[0] == 0 {
				return args[1]
			} else {
				return args[2]
			}
		}
	}
	prod := 1
	for sc.Scan() {
		m := &monkey{}
		items = append(items, []int{})
		setStartItems(lib.ReadTextFromScanner(sc))
		setOp(m, lib.ReadTextFromScanner(sc))
		nums := make([]int, 3)
		for i := range nums {
			s := lib.ReadTextFromScanner(sc)
			fields := strings.Fields(s)
			nums[i], _ = strconv.Atoi(fields[len(fields)-1])
		}
		prod *= nums[0]
		setTest(m, nums)
		sc.Scan()
		monkeys = append(monkeys, m)
	}
	itemsCopy := make([][]int, len(items))
	for i, nums := range items {
		itemsCopy[i] = make([]int, len(nums))
		copy(itemsCopy[i], nums)
	}
	for part, rounds := range []int{20, 10000} {
		if part == 1 {
			items = itemsCopy
		}
		active := make([]int, len(monkeys))
		for i := 0; i < rounds; i++ {
			for m, monkey := range monkeys {
				active[m] += len(items[m])
				for _, item := range items[m] {
					w := monkey.worry(item)
					switch part {
					case 0:
						w /= 3
					case 1:
						w %= prod
					}
					mt := monkey.throw(w)
					items[mt] = append(items[mt], w)
				}
				items[m] = items[m][:0]
			}
		}
		sort.Ints(active)
		last := len(active) - 1
		fmt.Println(active[last] * active[last-1])
	}
}
