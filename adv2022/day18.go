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

func Day18() {
	sc := bufio.NewScanner(os.Stdin)
	cubeMap := make(map[[3]int]bool)
	mins, maxs := [3]int{}, [3]int{}
	for i := range mins {
		mins[i] = math.MaxInt
		maxs[i] = math.MinInt
	}
	for sc.Scan() {
		items := strings.Split(sc.Text(), ",")
		cube := [3]int{}
		for i, item := range items {
			cube[i], _ = strconv.Atoi(item)
			mins[i] = min(mins[i], cube[i])
			maxs[i] = max(maxs[i], cube[i])
		}
		cubeMap[cube] = true
	}
	surfaceArea := func(cMap map[[3]int]bool) int {
		result := len(cMap) * 6
		for cube := range cMap {
			cube1 := cube
			for i := range cube {
				for _, d := range [2]int{-1, +1} {
					cube1[i] = cube[i] + d
					if cMap[cube1] {
						result--
					}
				}
				cube1[i] = cube[i]
			}
		}
		return result
	}
	fmt.Println(surfaceArea(cubeMap))
	size := [3]int{}
	for i := range mins {
		mins[i]--
		maxs[i]++
		size[i] = maxs[i] - mins[i] + 1
	}
	inBounds := func(cube [3]int) bool {
		for i := range cube {
			if cube[i] < mins[i] || cube[i] > maxs[i] {
				return false
			}
		}
		return true
	}
	q := make([][3]int, 1)
	q[0] = mins
	next := make([][3]int, 0)
	seen := make(map[[3]int]bool)
	seen[q[0]] = true
	for len(q) > 0 {
		for _, cube := range q {
			cube1 := cube
			for i := range cube {
				for _, d := range [2]int{-1, +1} {
					cube1[i] = cube[i] + d
					if inBounds(cube1) && !cubeMap[cube1] && !seen[cube1] {
						cpCube1 := cube1
						next = append(next, cpCube1)
						seen[cpCube1] = true
					}
				}
				cube1[i] = cube[i]
			}
		}
		q, next = next, q
		next = next[:0]
	}
	surfCuboid := lib.SurfaceOfCuboid(size[0], size[1], size[2])
	fmt.Println(surfaceArea(seen) - surfCuboid)
}
