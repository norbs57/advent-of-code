package adv2022

import (
	"bufio"
	"fmt"
	"os"

	"github.com/norbs57/advofcode/lib"
)

func Day17() {
	sc := bufio.NewScanner(os.Stdin)
	const maxCapacity = 1e5
	const maxTowerHeight = int(1e5)
	buf := make([]byte, maxCapacity)
	sc.Buffer(buf, maxCapacity)
	sc.Scan()
	line := sc.Bytes()
	patStr := [][]string{
		{"####"},
		{".#.", "###", ".#."},
		{"###", "..#", "..#"},
		{"#", "#", "#", "#"},
		{"##", "##"},
	}
	type shape struct {
		sqs    [][2]int
		width  int
		height int
		ty     string // used during debugging
	}
	type rock struct {
		shape
		pos [2]int
	}
	tys := [5]string{"-", "+", "⅃", "|", "■"}
	shapes := make([]shape, 5)
	for p, ps := range patStr {
		sh := &shapes[p]
		sh.ty = tys[p]
		sh.height = len(ps)
		for i, s := range ps {
			sh.width = max(sh.width, len(s))
			for j, r := range s {
				if r == '#' {
					sh.sqs = append(sh.sqs, [2]int{i, j})
				}
			}
		}
		// fmt.Println(shapes[p])
	}
	chamber := make([]lib.Bits, maxTowerHeight+1)
	chamber[0] = lib.Bits(127)
	const chamberWidth = 7
	canMove := func(r *rock, dir [2]int) bool {
		for _, sq := range r.sqs {
			row := r.pos[0] + sq[0] + dir[0]
			col := r.pos[1] + sq[1] + dir[1]
			if col < 0 || col >= chamberWidth || chamber[row].Has(col) {
				return false
			}
		}
		return true
	}
	Left := [2]int{0, -1}
	Right := [2]int{0, 1}
	Down := [2]int{-1, 0}
	move := func(r *rock, dir [2]int) {
		for i := range r.pos {
			r.pos[i] += dir[i]
		}
	}
	// numRocks is number of rocks in Part 1.
	// it is sufficiently large to determine the periodicity in Part 2.
	const numOfRocks = 2022
	// relevantRows was computed in previous program run using maxBelow
	const relevantRows = 43
	const rowGap = 3
	const startCol = 2
	th := make([]int, numOfRocks)
	var period, diff, p0 int
	dropRocks := func() {
		type state [relevantRows + 2]lib.Bits
		stateOf := func(r, mv, topRow int) state {
			st := state{}
			st[0] = lib.Bits(r % 5)
			st[1] = lib.Bits(mv)
			for i := 0; i+2 < relevantRows; i++ {
				if topRow-i >= 0 {
					st[i+2] = chamber[topRow-i]
				}
			}
			return st
		}
		table := make(map[state]int)
		// maxBelow was used to determine the number of relevant rows
		var inputPos, topRow, maxBelow int
		for r := 0; r < numOfRocks; r++ {
			rk := &rock{shapes[r%5], [2]int{topRow + rowGap + 1, startCol}}
			for {
				switch line[inputPos] {
				case '<':
					if canMove(rk, Left) {
						move(rk, Left)
					}
				case '>':
					if canMove(rk, Right) {
						move(rk, Right)
					}
				}
				inputPos = (inputPos + 1) % len(line)
				if canMove(rk, Down) {
					move(rk, Down)
				} else {
					break
				}
			}
			for _, sq := range rk.sqs {
				row, col := rk.pos[0]+sq[0], rk.pos[1]+sq[1]
				chamber[row] = chamber[row].Set(col)
			}
			topRow = max(topRow, rk.pos[0]+rk.height-1)
			maxBelow = max(maxBelow, topRow-rk.pos[0])
			th[r] = topRow
			st := stateOf(r, inputPos, topRow)
			r0, found := table[st]
			if found {
				period, diff, p0 = r-r0, th[r]-th[r0], r0
				return
			} else {
				table[st] = r
			}
		}
	}
	dropRocks()
	computeTH := func(n int) int {
		n--
		if n <= p0+period {
			return th[n]
		}
		n1 := p0 + (n-p0)%period
		return th[n1] + (n-n1)/period*diff
	}
	fmt.Println(computeTH(2022))
	fmt.Println(computeTH(1_000_000_000_000))
}
