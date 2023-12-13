package adv2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
	"golang.org/x/exp/maps"
)

func Day15() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 100_000), 100_000)
	sc.Scan()
	strs := strings.Split(sc.Text(), ",")

	Hash := func(str string) int {
		hash := 0
		for i := range str {
			ch := int(str[i])
			hash += ch
			hash *= 17
			hash %= 256
		}
		return hash
	}
	// Part 1
	total := [2]int{}
	for _, str := range strs {
		total[0] += Hash(str)
	}
	// Part 2
	mps := make([]map[string][2]int, 256)
	for i := range mps {
		mps[i] = make(map[string][2]int)
	}
	for i, str := range strs {
		label, lenStr, equalSign := strings.Cut(str, "=")
		if !equalSign {
			label, lenStr, _ = strings.Cut(str, "-")
		}
		h := Hash(label)
		lens, _ := strconv.Atoi(lenStr)
		if equalSign {
			lookup, ok := mps[h][label]
			if ok {
				mps[h][label] = [2]int{lookup[0], lens}
			} else {
				mps[h][label] = [2]int{i, lens}
			}
		} else {
			delete(mps[h], label)
		}
	}
	for i, mp := range mps {
		lenses := maps.Values(mp)
		lib.SortSliceLessFunc(lenses, lib.Less2)
		for j, lens := range lenses {
			total[1] += (i + 1) * (j + 1) * lens[1]
		}
	}
	fmt.Println(total[0], total[1])

}
