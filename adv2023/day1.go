package adv2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day1() {
	sc := bufio.NewScanner(os.Stdin)
	fs := []func(string) []int{Digits, DigitsWithWords}
	res := [2]int{}
	for sc.Scan() {
		line := sc.Text()
		for i, f := range fs {
			digits := f(line)
			res[i] += 10*digits[0] + lib.Last(digits)
		}
	}
	fmt.Println(res[0], res[1])
}

func Digits(str string) []int {
	digits := make([]int, 0, len(str))
	for _, ch := range str {
		if lib.IsDigit(ch) {
			n := int(ch - '0')
			if len(digits) < 2 {
				digits = append(digits, n)
			} else {
				digits[1] = n
			}
		}
	}
	return digits
}

func DigitsWithWords(str string) []int {
	words := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	for i := 0; i <= 9; i++ {
		words[string(byte('0'+i))] = i
	}
	digits := make([]int, 0, len(str))
	for i := range str {
		for w, v := range words {
			if strings.HasSuffix(str[:i+1], w) {
				if len(digits) < 2 {
					digits = append(digits, v)
				} else {
					digits[1] = v
				}
				break
			}
		}
	}
	return digits
}
