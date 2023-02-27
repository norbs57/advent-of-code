package adv2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"container/list"

	"github.com/norbs57/advofcode/lib"
)

func Day20() {
	// input contain duplicates but 0 is unique
	sc := bufio.NewScanner(os.Stdin)
	nums := list.New()
	elems := make([]*list.Element, 0)
	var zero *list.Element
	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		elem := nums.PushBack(n)
		elems = append(elems, elem)
		if n == 0 {
			zero = elem
		}
	}
	for _, e := range elems {
		ListMoveCyclic(nums, e, e.Value.(int))
	}
	xs := []int{1000, 2000, 3000}
	for i := range xs {
		xs[i] %= nums.Len()
	}
	fmt.Println(ListSumElems(nums, xs, zero))
	const decryptKey = 811589153
	nums = list.New()
	for _, e := range elems {
		nums.PushBack(decryptKey * e.Value.(int))
	}
	elems = elems[:0]
	for e := nums.Front(); e != nil; e = e.Next() {
		elems = append(elems, e)
		if e.Value == 0 {
			zero = e
		}
	}
	for i := 0; i < 10; i++ {
		for _, e := range elems {
			ListMoveCyclic(nums, e, e.Value.(int))
		}
	}
	fmt.Println(ListSumElems(nums, xs, zero))
}

func ListSumElems(nums *list.List, xs []int, start *list.Element) int {
	sum := 0
	for _, x := range xs {
		mark := start
		for i := 0; i < x; i++ {
			mark = mark.Next()
			if mark == nil {
				mark = nums.Front()
			}
		}
		sum += mark.Value.(int)
	}
	return sum
}

// Cyclic moves of elements for advent ofcode 2022, day 20
func ListMoveCyclic(l *list.List, e *list.Element, n int) {
	if n == 0 {
		return
	}
	// important: limit steps to avoid issues with moving e past itself
	steps := lib.Abs(n) % (l.Len() - 1)
	mark := e
	if n >= 0 {
		for i := 0; i < steps; i++ {
			mark = mark.Next()
			if mark == nil {
				mark = l.Front()
			}
		}
		l.MoveAfter(e, mark)
	} else {
		for i := 0; i < steps; i++ {
			mark = mark.Prev()
			if mark == nil {
				mark = l.Back()
			}
		}
		l.MoveBefore(e, mark)
	}
}
