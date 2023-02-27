package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IntervalIntersect(a0, a1, b0, b1 int) (int, int) {
	return max(a0, b0), min(a1, b1)
}

func Day4() {
	sc := bufio.NewScanner(os.Stdin)
	sum1, sum2 := 0, 0
	for sc.Scan() {
		split := func(r rune) bool {
			return r == '-' || r == ','
		}
		items := strings.FieldsFunc(sc.Text(), split)
		nums := make([]int, len(items))
		for i := range nums {
			nums[i], _ = strconv.Atoi(items[i])
		}
		a, b := IntervalIntersect(nums[0], nums[1], nums[2], nums[3])
		if a == nums[0] && b == nums[1] ||
			a == nums[2] && b == nums[3] {
			sum1++
		}
		if a <= b {
			sum2++
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}
