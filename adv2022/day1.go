package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	sc := bufio.NewScanner(os.Stdin)
	max := 0
	sum := 0
	sums := make([]int, 0, 4)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) == 0 {
			if sum > max {
				max = sum
			}
			sums = append(sums, sum)
			if len(sums) > 3 {
				sort.Ints(sums)
				sums = sums[1:]
			}
			sum = 0
			continue
		}
		cals, _ := strconv.Atoi(line)
		sum += cals
	}
	fmt.Println(max)
	fmt.Println(sums[0] + sums[1] + sums[2])
}
