package adv2023

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

type particle24 struct {
	pos [3]int
	vel [3]int
}

func Day24() {
	ps := make([]particle24, 0)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		halves := strings.Split(sc.Text(), "@")
		var p particle24
		for i := 0; i < 2; i++ {
			strs := strings.Split(halves[i], ",")
			for j, str := range strs {
				n, _ := strconv.Atoi(strings.TrimSpace(str))
				if i == 0 {
					p.pos[j] = n
				} else {
					p.vel[j] = n
				}
			}
		}
		ps = append(ps, p)
	}
	rlineOf := func(p particle24) lib.RLine {
		coords := [2][2]int{
			{p.pos[0], p.pos[1]},
			{p.pos[0] + p.vel[0], p.pos[1] + p.vel[1]},
		}
		r := lib.RLine{}
		for i := range r {
			for j := range r[i] {
				r[i][j] = big.NewRat(int64(coords[i][j]), 1)
			}
		}
		return r
	}
	const lwbEx = 7
	const upbEx = 27
	const lwbIn = 200_000_000_000_000
	const upbIn = 400_000_000_000_000
	lwb := big.NewRat(lwbIn, 1)
	upb := big.NewRat(upbIn, 1)
	inRegion := func(p [2]*big.Rat) bool {
		return lwb.Cmp(p[0]) <= 0 && p[0].Cmp(upb) <= 0 &&
			lwb.Cmp(p[1]) <= 0 && p[1].Cmp(upb) <= 0
	}
	// Part a
	coll := 0
	for i, pi := range ps {
		li := rlineOf(pi)
		for j := i + 1; j < len(ps); j++ {
			lj := rlineOf(ps[j])
			pij, ok := li.Intersect(lj)
			if ok && inRegion(pij) && lib.IsBetween(li[0], li[1], pij) &&
				lib.IsBetween(lj[0], lj[1], pij) {
				coll++
			}
		}
	}
	res := [2]int{coll, 0}
	
	// Part b 
	
	// Generate Sagemath equations for first five particles
	// See function writeSagemath

	/* Sagemath solution:
	[[t == 281427954234, u == 487736179331, v == 637228617556,
		w == 744158020102, x == 219726831090,
		a == 420851642592931, b == 273305746686315,	c == 176221626745613,
		d == -261, e == 15,	f == 233]]
	*/

	const (
		a = 420851642592931
		b = 273305746686315
		c = 176221626745613
	)
	res[1] = a + b + c
	fmt.Println(res[0], res[1])
}

func writeSagemath(ps []particle24) {
	timeVars := "tuvwx"
	posVars := "abc"
	velVars := "def"
	for _, ch := range posVars + velVars + timeVars {
		s := string(ch)
		fmt.Printf("%s = var('%s')\n", s, s)
	}
	fmt.Println("solve([")
	for i := 0; i < 5; i++ {
		p := ps[i]
		t := string(rune(timeVars[i]))
		for j := 0; j < 3; j++ {
			pv, vv := string(rune(posVars[j])), string(rune(velVars[j]))
			fmt.Printf("  %d + %s * %d == %s + %s * %s",
				p.pos[j], t, p.vel[j], pv, t, vv)
			if i < 4 || j < 2 {
				fmt.Println(",")
			} else {
				fmt.Println()
			}
		}
	}
	fmt.Println("], t, u, v, w, x, a, b, c, d, e, f)")
}
