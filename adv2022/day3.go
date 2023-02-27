package adv2022

import (
	"bufio"
	"fmt"
	"os"
)

func Day3() {
	sc := bufio.NewScanner(os.Stdin)
	prio := func(c byte) int {
		switch {
		case 'a' <= c && c <= 'z':
			return int(c - 'a' + 1)
		case 'A' <= c && c <= 'Z':
			return int(c - 'A' + 27)
		default:
			panic("prio: invalid byte")
		}
	}
	sum1 := 0
	sum2 := 0
	linesRead := 0
	count := make([]int, 53)
	for sc.Scan() {
		line := sc.Text()
		// part 1
		seen := make([]bool, 53)
		for i := 0; i < len(line)/2; i++ {
			c := line[i]
			seen[prio(c)] = true
		}
		for i := len(line) / 2; i < len(line); i++ {
			c := line[i]
			if seen[prio(c)] {
				sum1 += prio(c)
				break
			}
		}
		// part 2
		linesRead++
		if linesRead < 3 {
			seen2 := make([]bool, 53)
			for i := range line {
				c := line[i]
				seen2[prio(c)] = true
			}
			for i := range seen2 {
				if seen2[i] {
					count[i]++
				}
			}
		} else {
			linesRead = 0
			for i := range line {
				c := line[i]
				if count[prio(c)] == 2 {
					sum2 += prio(c)
					break
				}
			}
			for i := range count {
				count[i] = 0
			}
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}
