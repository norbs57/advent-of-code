package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day2() {
	sc := bufio.NewScanner(os.Stdin)
	intOf := func(s byte) int {
		if s <= 'C' {
			return int(s - 'A')
		} else {
			return int(s - 'X')
		}
	}
	const draw, win = 3, 6
	const loss2, win2 = 0, 2
	val := []int{1, 2, 3}
	scores2 := []int{0, 3, 6}
	score := func(a, b int) int {
		result := val[b]
		switch b {
		case a:
			result += draw
		case (a + 1) % 3:
			result += win
		}
		return result
	}
	result1, result2 := 0, 0
	for sc.Scan() {
		items := strings.Fields(sc.Text())
		// part 1
		result1 += score(intOf(items[0][0]), intOf(items[1][0]))
		// part 2
		opp, outcome := intOf(items[0][0]), intOf(items[1][0])
		result2 += scores2[outcome]
		we := opp
		switch outcome {
		case loss2:
			we = (opp + 2) % 3
		case win2:
			we = (opp + 1) % 3
		}
		result2 += val[we]
	}
	fmt.Println(result1)
	fmt.Println(result2)
}
