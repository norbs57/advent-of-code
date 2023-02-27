package adv2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day4() {
	lines := lib.ReadLines()
	total := 0
	copies := make([]int, len(lines))
	for i := range copies {
		copies[i] = 1
	}
	for i, line := range lines {
		line = line[strings.IndexByte(line, ':')+1:]
		parts := strings.Split(line, "|")
		winning := make(map[int]bool)
		for _, str := range strings.Fields(parts[0]) {
			num, _ := strconv.Atoi(str)
			winning[num] = true
		}
		score := 0
		matches := 0
		for _, str := range strings.Fields(parts[1]) {
			num, _ := strconv.Atoi(str)
			if winning[num] {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
				matches++
			}
		}
		for j := 1; j <= matches; j++ {
			if i+j >= len(lines) {
				break
			} else {
				copies[i+j] += copies[i]
			}
		}
		total += score
	}
	fmt.Println(total)
	fmt.Println(lib.Sum(copies))
}
