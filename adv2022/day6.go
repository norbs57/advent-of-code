package adv2022

import (
	"bufio"
	"fmt"
	"os"
)

func Day6() {
	// no need for bufio.NewReaderSize - default buffersize is just large enough
	// Part 1: L = 4, Part 2: L = 14
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := sc.Text()
	for _, L := range [2]int{4, 14} {
		occ := make([]int, 128)
		diffChars := 0
		for i := 0; i < L; i++ {
			ch := line[i]
			if occ[ch] == 0 {
				diffChars++
			}
			occ[ch]++
		}
		for i := L; i < len(line); i++ {
			hdCh := line[i-L]
			occ[hdCh]--
			if occ[hdCh] == 0 {
				diffChars--
			}
			ch := line[i]
			if occ[ch] == 0 {
				diffChars++
			}
			occ[ch]++
			if diffChars == L {
				fmt.Println(i + 1)
				break
			}
		}
	}
}
