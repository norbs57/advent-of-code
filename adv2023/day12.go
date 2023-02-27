package adv2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxStrLen = 104
const maxNumLen = 30

var Day12Table [maxStrLen + 1][maxNumLen + 1]int

func Day12() {
	sum := [2]int{}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		items := strings.Fields(sc.Text())
		numStrs := strings.Split(items[1], ",")
		nums := make([][]int, 2)
		nums[0] = make([]int, len(numStrs))
		for i := range numStrs {
			nums[0][i], _ = strconv.Atoi(numStrs[i])
		}
		strs := make([]string, 2)
		strs[0] = items[0]
		nums[1] = make([]int, 0, 5*len(nums))
		for i := 0; i < 5; i++ {
			strs[1] = strs[1] + strs[0]
			if i < 4 {
				strs[1] = strs[1] + "?"
			}
			nums[1] = append(nums[1], nums[0]...)
		}
		for i := 0; i < 2; i++ {
			for j := range Day12Table {
				for k := range Day12Table[i] {
					Day12Table[j][k] = -1
				}
			}
			str := strs[i]
			for len(str) > 0 && str[len(str)-1] == '.' {
				str = str[:len(str)-1]
			}
			sum[i] += Day12Solve(str, nums[i])
		}
	}
	fmt.Println(sum[0])
	fmt.Println(sum[1])
}

func Day12Match(str string, num int) bool {
	return num <= len(str) && !strings.ContainsRune(str[:num], '.') &&
		(len(str) == num || str[num] != '#')
}

func Day12Solve(str string, nums []int) int {
	res := &Day12Table[len(str)][len(nums)]
	if *res != -1 {
		return *res
	}
	if len(nums) == 0 {
		if strings.ContainsRune(str, '#') {
			*res = 0
		} else {
			*res = 1
		}
		return *res
	}
	if len(str) == 0 {
		*res = 0
		return *res
	}
	tryMatch := func() {
		if Day12Match(str, nums[0]) {
			str1 := str[nums[0]:]
			if len(str1) > 0 {
				str1 = str1[1:]
			}
			*res = Day12Solve(str1, nums[1:])
		} else {
			*res = 0
		}
	}
	switch str[0] {
	case '#':
		tryMatch()
	case '?':
		tryMatch()
		*res += Day12Solve(str[1:], nums)
	default:
		*res = Day12Solve(str[1:], nums)
	}
	return *res
}
