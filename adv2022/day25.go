package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func Day25() {
	sc := bufio.NewScanner(os.Stdin)
	sum := 0
	for sc.Scan() {
		n := FromSnafu(sc.Text())
		sum += n
	}
	fmt.Println(ToSnafu(sum))
}

func ToSnafu(n int) string {
	if n == 0 {
		return "0"
	}
	result := make([]rune, 0)
	for n > 0 {
		d := n % 5
		switch d {
		case 0, 1, 2:
			result = append(result, rune('0'+d))
		case 3:
			result = append(result, '=')
			n += 5
		case 4:
			result = append(result, '-')
			n += 5
		}
		n /= 5
	}
	slices.Reverse(result)
	return string(result)
}

func FromSnafu(s string) int {
	power := 1
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		r := s[i]
		switch r {
		case '0', '1', '2':
			result += int(r-'0') * power
		case '-':
			result -= power
		case '=':
			result -= 2 * power
		}
		power *= 5
	}
	return result
}
