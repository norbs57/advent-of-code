package adv2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
	"golang.org/x/exp/maps"
)

func Day22() {
	type brick [][2]int
	bricks := []brick{}
	sc := bufio.NewScanner(os.Stdin)
	maxDims := [3]int{}
	maxTopHeight := 0
	for sc.Scan() {
		parts := strings.Split(sc.Text(), "~")
		br := make([][2]int, 3)
		for i, p := range parts {
			numStrs := strings.Split(p, ",")
			for j, numStr := range numStrs {
				br[j][i], _ = strconv.Atoi(numStr)
				maxDims[j] = max(maxDims[j], br[j][i])
			}
		} 
		maxTopHeight = max(maxTopHeight, br[2][1])
		bricks = append(bricks, br)
	}
	lib.SortSliceFunc(bricks, func(b brick) int {
		return b[2][0]
	}) 
	restsOn := lib.MkAdjMap(len(bricks))
	supports := lib.MkAdjMap(len(bricks)) 
	stack := lib.MkSlice3Filled[int](maxDims[0]+1, maxDims[1]+1, maxTopHeight+1, -1)
	heightAt := lib.MkSlice2Filled[int](maxDims[0]+1, maxDims[1]+1, -1)
	for bi, br := range bricks {
		h := 0
		for i := br[0][0]; i <= br[0][1]; i++ {
			for j := br[1][0]; j <= br[1][1]; j++ {
				h = max(h, heightAt[i][j])
			}
		}
		dh := br[2][1] - br[2][0]
		for i := br[0][0]; i <= br[0][1]; i++ {
			for j := br[1][0]; j <= br[1][1]; j++ {
				heightAt[i][j] = h + dh + 1
				for k := 0; k <= br[2][1]-br[2][0]; k++ {
					stack[i][j][h+k] = bi
				}
				if h > 0 {
					s := stack[i][j][h-1]
					if s != -1 {
						restsOn.Add(bi, s)
						supports.Add(s, bi)
					}
				}
			}
		} 
	}
	res := [2]int{}
	notRemovable := make(map[int]bool)
	for bi := range bricks {
		sp := restsOn[bi]
		if len(sp) == 1 {
			s0 := maps.Keys(sp)[0]
			notRemovable[s0] = true
		}
	}
	res[0] = len(bricks) - len(notRemovable)
	falling := func(br int) int {
		desIntegrated := make([]bool, len(bricks))
		q := make([]int, 1)
		q[0] = br
		desIntegrated[br] = true
		for len(q) > 0 {
			b := lib.PopFront(&q)
			for r := range supports[b] {
				supported := false
				for sr := range restsOn[r] {
					if !desIntegrated[sr] {
						supported = true
						break
					}
				}
				if !supported {
					q = append(q, r)
					desIntegrated[r] = true
				}
			}
		}
		return lib.Freq(desIntegrated)[true] - 1
	} 
	for i := range notRemovable {
		res[1] += falling(i)
	}
	fmt.Println(res[0], res[1])
}
