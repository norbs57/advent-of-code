package adv2023

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/norbs57/advofcode/lib"
	"golang.org/x/exp/maps"
)

func Day25() {
	g := Day25ReadInput()
	mat := lib.MatrixOfAdjList(g)
	cutSize, cut0 := lib.GlobalMinCut(mat)
	if cutSize != 3 {
		fmt.Println("Incorrect cut size", cutSize)
		return
	}
	n := len(g)
	m := len(cut0)
	res := m * (n - m)
	fmt.Println(res)
}

func Day25ReadInput() lib.AdjList {
	sc := bufio.NewScanner(os.Stdin)
	gStr := make(map[string]map[string]int)
	for sc.Scan() {
		strs := strings.Fields(sc.Text())
		a := strs[0][:len(strs[0])-1]
		for _, b := range strs[1:] {
			for _, elem := range []string{a, b} {
				if gStr[elem] == nil {
					gStr[elem] = make(map[string]int)
				}
			}
			gStr[a][b]++
			gStr[b][a]++
		}
	}
	strs := maps.Keys(gStr)
	sort.Strings(strs)
	intOfStr := make(map[string]int)
	for i, str := range strs {
		intOfStr[str] = i
	}
	g := make(lib.AdjList, len(gStr))
	for a, bs := range gStr {
		i := intOfStr[a]
		for b := range bs {
			g.Add(i, intOfStr[b])
		}
	}
	return g
}
