package adv2022

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day22() {
	sc := bufio.NewScanner(os.Stdin)
	grid := make([]string, 0)
	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			break
		}
		grid = append(grid, line)
	}
	startOfRow, endOfRow := make([]int, len(grid)), make([]int, len(grid))
	for i := range grid {
		j := 0
		for ; j < len(grid[i]) && grid[i][j] == ' '; j++ {
		}
		startOfRow[i] = j
		j = len(grid[i]) - 1
		for ; j >= 0 && grid[i][j] == ' '; j-- {
		}
		endOfRow[i] = j
	}
	startOfCol, endOfCol := make([]int, len(grid[0])), make([]int, len(grid[0]))
	for j := range grid[0] {
		i := 0
		for ; i < len(grid) && j < startOfRow[i] || j > endOfRow[i] || grid[i][j] == ' '; i++ {
		}
		startOfCol[j] = i
		i = len(grid) - 1
		for ; i >= 0 && j < startOfRow[i] || j > endOfRow[i] || grid[i][j] == ' '; i-- {
		}
		endOfCol[j] = i
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	const (
		East = iota
		South
		West
		North
	)
	turn := func(dir int, r byte) int {
		switch r {
		case 'R':
			return (dir + 1) % len(dirs)
		case 'L':
			return (dir + (len(dirs) - 1)) % len(dirs)
		}
		return -1
	}
	type nextFun func(int, int, int) (int, int, int)
	next := func(r, c, dir int) (int, int, int) {
		d := dirs[dir]
		r, c = r+d[0], c+d[1]
		switch {
		case dir == West && c < startOfRow[r]:
			return r, endOfRow[r], dir
		case dir == East && c > endOfRow[r]:
			return r, startOfRow[r], dir
		case dir == North && r < startOfCol[c]:
			return endOfCol[c], c, dir
		case dir == South && r > endOfCol[c]:
			return startOfCol[c], c, dir
		default:
			return r, c, dir
		}
	}
	move := func(r, c, dir, n int, f nextFun) (int, int, int) {
		for ; n > 0; n-- {
			r1, c1, dir1 := f(r, c, dir)
			if grid[r1][c1] == '#' {
				break
			} else {
				r, c, dir = r1, c1, dir1
			}
		}
		return r, c, dir
	}
	sc.Scan()
	path := make([]byte, 0)
	for _, b := range sc.Bytes() {
		if '0' <= b && b <= '9' {
			path = append(path, b)
		} else {
			path = append(path, ' ', b, ' ')
		}
	}
	followPath := func(f nextFun) {
		r, c, dir := 0, startOfRow[0], 0
		items := strings.Fields(string(path))
		for _, str := range items {
			switch str {
			case "R", "L":
				dir = turn(dir, str[0])
			default:
				n, _ := strconv.Atoi(str)
				r, c, dir = move(r, c, dir, n, f)
			}
		}
		result := 1000*(r+1) + 4*(c+1) + dir
		fmt.Println(result)
	}
	followPath(next)
	n := math.MaxInt
	for i := range grid {
		n = min(n, endOfRow[i]-startOfRow[i])
	}
	n++
	topLeft := make([][2]int, 0)
	for r := 0; r < len(grid); r += n {
		for c := startOfRow[r]; c <= endOfRow[r] && grid[r][c] != ' '; c += n {
			topLeft = append(topLeft, [2]int{r, c})
		}
	}
	sqOf := func(r, c int) int {
		for i, p := range topLeft {
			pr, pc := p[0], p[1]
			if pr <= r && r < pr+n && pc <= c && c < pc+n {
				return i
			}
		}
		return -1
	}
	bds := make([][3]int, 0)
	var setBd func(r, c, dir int)
	setBd = func(r, c, dir int) {
		bds = append(bds, [3]int{r, c, dir})
		d := dirs[dir]
		r, c = r+(n-1)*d[0], c+(n-1)*d[1]
		if r == 0 && c == startOfRow[0] {
			return
		}
		dLeft := (dir + 3) % 4
		dl := dirs[dLeft]
		r1, c1 := r+d[0]+dl[0], c+d[1]+dl[1]
		if sqOf(r1, c1) != -1 {
			setBd(r1, c1, dLeft)
			return
		}
		r1, c1 = r+d[0], c+d[1]
		if sqOf(r1, c1) != -1 {
			setBd(r1, c1, dir)
			return
		}
		setBd(r, c, (dir+1)%4)
	}
	setBd(0, startOfRow[0], East)
	// glue describes the re-folding of the unfolded cube
	var glue []int
	if n == 4 {
		glue = []int{11, 4, 3, 2, 1, 10, 9, 8, 7, 6, 5, 0, 13, 12}
	} else {
		glue = []int{9, 8, 5, 4, 3, 2, 7, 6, 1, 0, 13, 12, 11, 10}
	}
	bNext := make(map[[3]int][3]int)
	for i, bd := range bds {
		d := dirs[bd[2]]
		left := (bd[2] + 3) % 4
		g := bds[glue[i]]
		dg := dirs[g[2]]
		gRight := (g[2] + 1) % 4
		for j := 0; j < n; j++ {
			r, c := bd[0]+j*d[0], bd[1]+j*d[1]
			r1, c1 := g[0]+(n-1-j)*dg[0], g[1]+(n-1-j)*dg[1]
			bNext[[3]int{r, c, left}] = [3]int{r1, c1, gRight}
		}
	}
	next3D := func(r, c, dir int) (int, int, int) {
		y, found := bNext[[3]int{r, c, dir}]
		if found {
			return y[0], y[1], y[2]
		} else {
			d := dirs[dir]
			return r + d[0], c + d[1], dir
		}
	}
	followPath(next3D)
}
