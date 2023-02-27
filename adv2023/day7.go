package adv2023

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day7() {
	countCards := 13
	cardStr := [2]string{"AKQJT98765432", "AKQT98765432J"}
	rank := func(ch rune) [2]int {
		res := [2]int{}
		for i, str := range cardStr {
			res[i] = len(str) - strings.IndexRune(str, ch)
		}
		return res
	}
	typeOfHand := func(hand []int) []int {
		freq := make([]int, countCards+1)
		for _, r := range hand {
			freq[r]++
		}
		res := make([]int, 0)
		for _, f := range freq {
			if f != 0 {
				res = append(res, f)
			}
		}
		lib.SortSliceDecr(res)
		return res
	}
	catOfHand := func(hand []int) int {
		ty := typeOfHand(hand)
		cats := [][]int{{5}, {4}, {3, 2}, {3}, {2, 2}, {2}}
		for i, cat := range cats {
			if slices.Equal(ty[:len(cat)], cat) {
				return 6 - i
			}
		}
		return 0
	}
	valueOfHand0 := func(hand []int) int {
		val := catOfHand(hand)
		for _, card := range hand {
			val = val*16 + card
		}
		return val
	}
	valueOfHand1 := func(hand []int) int {
		nonJokersMap := make(map[int]int)
		nonJokersCount := 0
		joker := rank('J')[1]
		for _, card := range hand {
			if card != joker {
				nonJokersMap[card]++
				nonJokersCount++
			}
		}
		if nonJokersCount == 0 || nonJokersCount == len(hand) {
			return valueOfHand0(hand)
		}
		maxVal := 0
		for card := range nonJokersMap {
			hand1 := make([]int, len(hand))
			copy(hand1, hand)
			for i := range hand {
				if hand[i] == joker {
					hand1[i] = card
				}
			}
			val := catOfHand(hand1)
			for _, card := range hand {
				val = val*16 + card
			}
			if val > maxVal {
				maxVal = val
			}
		}
		return maxVal
	}
	valueOf := [2]func([]int) int{valueOfHand0, valueOfHand1}

	lines := lib.ReadLines()
	vs := make([][][2]int, 2)
	for _, line := range lines {
		items := strings.Fields(line)
		mult, _ := strconv.Atoi(items[1])
		for j := 0; j < 2; j++ {
			cards := make([]int, 0, 5)
			for _, ch := range items[0] {
				cards = append(cards, rank(ch)[j])
			}
			v := valueOf[j](cards)
			vs[j] = append(vs[j], [2]int{v, mult})
		}
	}
	for _, xs := range vs {
		lib.SortSliceLessFunc(xs, lib.Less2)
		res := 0
		for i, v := range xs {
			res += v[1] * (i + 1)
		}
		fmt.Println(res)
	}
}
