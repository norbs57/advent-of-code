package main

import (
	"bufio"
	"bytes"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/norbs57/advofcode/adv2022"
	"github.com/norbs57/advofcode/adv2023"
	"github.com/norbs57/advofcode/lib"
)

func init() {
	adv2022.AssignFuns()
	adv2023.AssignFuns()
}

func Test23(t *testing.T) {
	for year := 2023; year <= 2023; year++ {
		answers := ReadAnswers(year)
		for day := range lib.AocFunMap[year] {
			f := lib.AocFunMap[year][day]
			AdvOneTest(t, year, day+1, answers, f)
		}
	}
}

func TestAll(t *testing.T) {
	for year := 2022; year <= 2023; year++ {
		answers := ReadAnswers(year)
		for day := range lib.AocFunMap[year] {
			f := lib.AocFunMap[year][day]
			AdvOneTest(t, year, day+1, answers, f)
		}
	}
}

func TestOne(t *testing.T) {
	const year = 2023
	const day = 12
	answers := ReadAnswers(year)
	f := lib.AocFunMap[year][day-1]
	AdvOneTest(t, year, day, answers, f)
}

func AdvOneTest(t *testing.T, year, day int, answers []string, fn func()) {
	in := InputFile(year, day)
	defer lib.SetStdin(in).Close()
	out := lib.CaptureStdout(fn)
	sc := bufio.NewScanner(bytes.NewReader(out))
	sc.Split(bufio.ScanWords)
	output := lib.ReadStringsFromScanner(sc)
	fmt.Println(day, output)
	expected := strings.Fields(answers[day-1])
	lib.AssertSlicesEqual(t, output, expected)
}

func ReadAnswers(year int) []string {
	yearStr := strconv.Itoa(year)
	ans := filepath.Join("./adv"+yearStr, "data", "answers.txt")
	return lib.ReadLinesFromFile(ans)
}

func InputFile(year, day int) string {
	yearStr, dayStr := strconv.Itoa(year), strconv.Itoa(day)
	return filepath.Join("./adv"+yearStr, "data", "in"+dayStr+".txt")
}
