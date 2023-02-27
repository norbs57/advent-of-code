package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day10() {
	sc := bufio.NewScanner(os.Stdin)
	const rows = 6
	const cols = 40
	display := make([][]byte, rows)
	for i := range display {
		display[i] = make([]byte, cols)
	}
	cycle := 1
	x := 1
	draw := func() {
		col := (cycle - 1) % cols
		row := (cycle - 1) / cols
		if lib.Abs(x-col) <= 1 {
			display[row][col] = '#'
		} else {
			display[row][col] = '.'
		}
	}
	sum := 0
	next := 20
	inc := func() {
		cycle++
		if cycle == next {
			sigStr := next * x
			sum += sigStr
			next += 40
		}
	}
	// the sprite is 3 pixels wide, and the X register sets the horizontal position
	// of the middle of that sprite
	for sc.Scan() {
		items := strings.Fields(sc.Text())
		draw()
		inc()
		if len(items) == 2 {
			n, _ := strconv.Atoi(items[1])
			draw()
			x += n
			inc()
		}
	}
	fmt.Println(sum)

	// for _, row := range display {
	// 	fmt.Println(string(row))
	// }

	// letter recognizer for the "display" not coded :)
	fmt.Println("PAPJCBHP")
}
