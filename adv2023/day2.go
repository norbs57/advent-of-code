package adv2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day2() {
	// FP-style combinators
	lines := lib.ReadLines()
	fs := []func(string) int{GamePossible, MinPossible}
	for _, f := range fs {
		fmt.Println(lib.Sum(lib.Map(f, lines)))
	}
}

func MaxColor(str string) [3]int {
	dict := map[byte]int{'r': 0, 'g': 1, 'b': 2}
	maxColor := [3]int{}
	for _, grab := range strings.Split(str, ";") {
		for _, item := range strings.Split(grab, ",") {
			parts := strings.Fields(item)
			n, _ := strconv.Atoi(parts[0])
			i := dict[parts[1][0]]
			maxColor[i] = max(maxColor[i], n)
		}
	}
	return maxColor
}

func GamePossible(str string) int {
	limits := [3]int{12, 13, 14}
	game := strings.Split(str, ":")
	id, _ := strconv.Atoi(game[0][5:])
	mx := MaxColor(game[1])
	for i, l := range limits {
		if mx[i] > l {
			return 0
		}
	}
	return id
}

func MinPossible(str string) int {
	game := strings.Split(str, ":")
	mx := MaxColor(game[1])
	return lib.Prod(mx[:])
}
