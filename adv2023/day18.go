package adv2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day18() {
	sc := bufio.NewScanner(os.Stdin)
	dirOfLetter := map[byte]int{'R': 0, 'D': 1, 'L': 2, 'U': 3}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	pts := make([][][2]int, 2)
	steps := [2]int{}
	rows, cols := [2]int{}, [2]int{}
	for sc.Scan() {
		items := strings.Fields(sc.Text())
		ds, stps := [2]int{}, [2]int{}
		ds[0] = dirOfLetter[items[0][0]]
		stps[0], _ = strconv.Atoi(items[1])
		ds[1], _ = strconv.Atoi(items[2][7:8])
		stps1, _ := strconv.ParseInt(items[2][2:7], 16, 64)
		stps[1] = int(stps1)
		for i := 0; i < 2; i++ {
			dv := dirs[ds[i]]
			rows[i] += stps[i] * dv[0]
			cols[i] += stps[i] * dv[1]
			pts[i] = append(pts[i], [2]int{rows[i], cols[i]})
			steps[i] += stps[i]
		}
	}
	for i := 0; i < 2; i++ {
		fmt.Println(lib.PolygonArea(pts[i]) + steps[i]/2 + 1)
	}

}
