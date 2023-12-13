package adv2022

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day15() {
	// The program computes intervals of covered locations for each line. It
	// finds the first point not covered. Part 2 take about 3.1 seconds on my
	// machine. Faster solutions would be possible making use of the fact that
	// "there is only one such point" but that feels rather artificial.
	const yRow, upb = 2000000, 4000000
	// const yRow, upb = 10, 20
	const multiplier = 4000000
	sensors, dists, bsRow := Day15ReadInput(yRow)
	ivs := Day15Ivs(sensors, dists, yRow, math.MinInt, math.MaxInt)

	SumItvSizes := func(xs [][2]int) int {
		// assumes intervals are disjoint and non-empty
		result := 0
		for _, x := range xs {
			result += x[1] - x[0] + 1
		}
		return result
	}

	fmt.Println(SumItvSizes(ivs) - bsRow)
	for row := 0; row <= upb; row++ {
		ivs = Day15Ivs(sensors, dists, row, 0, upb)
		if len(ivs) > 1 {
			switch {
			case ivs[0][0] > 0:
				fmt.Println(row)
			case ivs[len(ivs)-1][1] < upb:
				fmt.Println(row + upb*multiplier)
			default:
				for i := 0; i < len(ivs)-1; i++ {
					col := ivs[i][1] + 1
					if col < ivs[i+1][0] {
						fmt.Println(row + col*multiplier)
					}
				}
			}
			break
		}
	}
}

func Day15ReadInput(y int) ([][2]int, []int, int) {
	sc := bufio.NewScanner(os.Stdin)
	// coordinates are switched: x=col, y=row,
	sensors := make([][2]int, 0)
	beacons := make(map[int]bool)
	dists := make([]int, 0)
	for sc.Scan() {
		splitF := func(r rune) bool {
			return r != '-' && (r < '0' || r > '9')
		}
		items := strings.FieldsFunc(sc.Text(), splitF)
		sensor := [2]int{}
		beacon := [2]int{}
		for i := 0; i < 2; i++ {
			sensor[1-i], _ = strconv.Atoi(items[i])
			beacon[1-i], _ = strconv.Atoi(items[i+2])
		}
		sensors = append(sensors, sensor)
		if beacon[0] == y {
			beacons[beacon[1]] = true
		}
		dist := lib.Dist2(sensor, beacon)
		dists = append(dists, dist)
	}
	return sensors, dists, len(beacons)
}

func Day15Ivs(sensors [][2]int, dists []int, row, lwb, upb int) [][2]int {
	cutIntv := [2]int{lwb, upb}
	intvs := make([][2]int, 0)
	for i, s := range sensors {
		dCol := dists[i] - lib.Abs(s[0]-row)
		if dCol > 0 {
			intv := [2]int{s[1] - dCol, s[1] + dCol}
			intv = lib.IntersectItvs(intv, cutIntv)
			if !lib.IsEmptyItv(intv) {
				intvs = append(intvs, intv)
			}
		}
	}
	lib.SortLess2[int](intvs)
	return lib.MergeItvs(intvs)
}
