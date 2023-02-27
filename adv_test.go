package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"testing"

	adv2022 "github.com/norbs57/advofcode/adv2022"
	"github.com/norbs57/advofcode/lib"
)

const ext = ".txt"
const dataFolder = "../../go/advofcode/adv2022/data"

var dayFuns []func()

func init() {
	dayFuns = []func(){
		adv2022.Day1,
		adv2022.Day2,
		adv2022.Day3,
		adv2022.Day4,
		adv2022.Day5,
		adv2022.Day6,
		adv2022.Day7,
		adv2022.Day8,
		adv2022.Day9,
		adv2022.Day10,
		adv2022.Day11,
		adv2022.Day12,
		adv2022.Day13,
		adv2022.Day14,
		adv2022.Day15,
		adv2022.Day16,
		adv2022.Day17,
		adv2022.Day18,
		adv2022.Day19,
		adv2022.Day20,
		adv2022.Day21,
		adv2022.Day22,
		adv2022.Day23,
		adv2022.Day24,
		adv2022.Day25,
	}
}

func TestAll(t *testing.T) {
	for i, f := range dayFuns {
		AdvOneTest(t, i+1, f)
	}
}

func TestOne(t *testing.T) {
	const n = 25
	AdvOneTest(t, n, dayFuns[n-1])
}

func AdvOneTest(t *testing.T, d int, fn func()) {
	day := strconv.Itoa(d)
	in := dataFolder + "/input" + day + ext
	ans := dataFolder + "/ans" + day + ext
	defer lib.SetStdin(in).Close()
	out := lib.CaptureStdout(fn)
	sc := bufio.NewScanner(bytes.NewReader(out))
	output := lib.ReadLinesFromScanner(sc)
	fmt.Println(d, output)
	expected := lib.ReadLinesFromFile(ans)
	lib.AssertSlicesEqual(t, output, expected)
}

func TestSnafu(t *testing.T) {
	for i := 0; i < 10000; i++ {
		s := adv2022.ToSnafu(i)
		j := adv2022.FromSnafu(s)
		if i != j {
			t.Errorf("Expected %d, got %d, toSnafu= %v", i, j, s)
		}
	}
}
