package adv2023

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/norbs57/advofcode/lib"
)

func Day5() {
	// assumption: conversion domains are disjoint
	sc := bufio.NewScanner(os.Stdin)
	values := make([]int, 0)
	sc.Scan()
	items := strings.Fields(sc.Text())
	for i := 1; i < len(items); i++ {
		num, _ := strconv.Atoi(items[i])
		values = append(values, num)
	}
	mps := make([][][3]int, 0)
	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			continue
		}
		if line[len(line)-1] == ':' {
			mps = append(mps, make([][3]int, 0))
			continue
		}
		conv := [3]int{}
		items := strings.Fields(line)
		for i := range conv {
			conv[i], _ = strconv.Atoi(items[i])
		}
		lib.Append(&mps[len(mps)-1], conv)
	}
	for _, mp := range mps {
		lib.SortSliceLessFunc(mp, lib.Less3[int])
	}
	Day5a(values, mps)
	Day5b(values, mps)
}

func Day5a(values []int, mps [][][3]int) {
	transform := func(mp [][3]int, v int) int {
		for _, conv := range mp {
			dest, src, n := conv[0], conv[1], conv[2]
			if src <= v && v < src+n {
				return v + dest - src
			}
		}
		return v
	}
	for _, mp := range mps {
		next := make([]int, len(values))
		for i, v := range values {
			next[i] = transform(mp, v)
		}
		values = next
	}
	fmt.Println(slices.Min(values))
}

func Day5b(values []int, mps [][][3]int) {
	transformItv := func(mp [][3]int, itv [2]int) [][2]int {
		covered := make([][2]int, 0)
		mapped := make([][2]int, 0)
		for _, conv := range mp {
			dest, src, n := conv[0], conv[1], conv[2]
			dom := [2]int{src, src + n}
			offset := dest - src
			its := lib.Intersect(itv, dom)
			if its != [2]int{} {
				img := [2]int{its[0] + offset, its[1] + offset}
				mapped = append(mapped, img)
				covered = append(covered, its)
			}
		}
		if len(covered) == 0 {
			lib.Append(&mapped, itv)
			return mapped
		}
		lib.SortSliceLessFunc(covered, lib.Less2)
		start := itv[0]
		for _, itv := range covered {
			if itv[0] != start {
				lib.Append(&mapped, [2]int{start, itv[0]})
			}
			start = itv[1]
		}
		endCov := covered[len(covered)-1][1]
		if endCov < itv[1] {
			lib.Append(&mapped, [2]int{endCov, itv[1]})
		}
		return mapped
	}
	itvs := make([][2]int, 0)
	for i := 0; i < len(values); i += 2 {
		a, l := values[i], values[i+1]
		itvs = append(itvs, [2]int{a, a + l})
	}
	for _, mp := range mps {
		next := make([][2]int, 0)
		for _, itv := range itvs {
			next = append(next, transformItv(mp, itv)...)
		}
		lib.SortSliceLessFunc(next, lib.Less2)
		lib.MergeIntvs(next)
		itvs = next
	}
	fmt.Println(itvs[0][0])
}
