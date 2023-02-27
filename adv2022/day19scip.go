package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Day19 Scip hack: generate LP files, run SCIP and parse results
// Apparently LP format is not much used anymore, but AMPL and GAMS are
// Google OR tools can read MPS files
func Day19Scip() {
	Day19WriteLPFiles()
	Day19ReadScipResults()
}

func Day19WriteLPFiles() {
	sc := bufio.NewScanner(os.Stdin)
	folder := "./geodes/"
	for n := 0; n < 30 && sc.Scan(); n++ {
		line := sc.Text()
		bp := Day19ParseBlueprint(line)
		Day19WriteLPFile(folder, n+1, 24, bp)
		if n < 3 {
			Day19WriteLPFile(folder, 31+n, 32, bp)
		}
	}
}

func Day19WriteLPFile(folder string, id int, minutes int, a [][]int) {
	idStr := strconv.Itoa(id)
	f, _ := os.Create(folder + "geode_" + idStr + ".lp")
	w := bufio.NewWriter(f)
	fmt.Fprintln(w, "Maximize")
	fmt.Fprintln(w, "  result")
	fmt.Fprintln(w, "Subject To")
	templ := func(i int, j int) string {
		result := ""
		min := 'a' + j
		r := 'r' + j
		result += fmt.Sprintf("%c%d - %c%d - %c%d", min, i, min, i-1, r, i-1)
		for k := 0; k < 4; k++ {
			if j < 3 && a[k][j] > 0 {
				rk := 'r' + k
				result += fmt.Sprintf(" + %d d%c%d", a[k][j], rk, i)
			}
		}
		result += " = 0"
		return result
	}
	fmt.Fprintln(w, "")
	for i := 2; i <= minutes; i++ {
		for j := 0; j < 4; j++ {
			fmt.Fprintln(w, templ(i, j))
			ch := rune('r' + j)
			fmt.Fprintf(w, "d%c%d + %c%d - %c%d = 0\n", ch, i, ch, i-1, ch, i)
			for k := 0; k < 3; k++ {
				if a[j][k] > 0 {
					ch1 := rune('a' + k)
					fmt.Fprintf(w, "%c%d - %d d%c%d >= 0\n", ch1, i-1, a[j][k], ch, i)
				}
			}
		}
		fmt.Fprintf(w, "d%c%d + d%c%d + d%c%d + d%c%d <= 1\n",
			'r', i, 's', i, 't', i, 'u', i)
		fmt.Fprintln(w)
	}
	fmt.Fprintf(w, "result - d%d = 0", minutes)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Bounds")
	fmt.Fprintln(w, "1 <= a1 <= 1")
	fmt.Fprintln(w, "1 <= r1 <= 1")
	for j := 1; j < 4; j++ {
		ch := rune('a' + j)
		fmt.Fprintf(w, "0 <= %c1 <= 0\n", ch)
		ch = rune('r' + j)
		fmt.Fprintf(w, "0 <= %c1 <= 0\n", ch)
	}
	fmt.Fprintln(w, "General")
	for i := 1; i <= minutes; i++ {
		for _, ch := range []rune{'a', 'r'} {
			for j := 0; j < 4; j++ {
				fmt.Fprintf(w, "%c%d ", rune(int(ch)+j), i)
			}
		}
	}
	w.Flush()
}

func Day19ReadScipResults() {
	sc := bufio.NewScanner(os.Stdin)
	result1 := 0
	result2 := 1
	for line := 1; sc.Scan(); line++ {
		fields := strings.Fields(sc.Text())
		n, _ := strconv.Atoi(fields[2])
		if line <= 30 {
			result1 += line * n
		} else {
			result2 *= n
		}
	}
	fmt.Println(result1)
	fmt.Println(result2)
}
