package lib

import (
	"container/list"
	"fmt"
	"sort"
	"strings"
	"testing"

	"golang.org/x/exp/constraints"
)

// Assertions

func AssertEqual[E comparable](t *testing.T, got, want E) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertSlicesEqual[E comparable](t *testing.T, got, want []E) {
	if len(got) != len(want) {
		t.Errorf("length different %v %v", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("line %v, got %v want %v", i, got[i], want[i])
		}
	}
}

// Miscellaneous functions

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func AbsInt(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func SignInt(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return 1
	default:
		return 0
	}
}

func ModInt(a, b int) int {
	// result is positive if b is positivew
	return (a%b + b) % b
}

func DistInt2(a, b [2]int) int {
	return AbsInt(b[0]-a[0]) + AbsInt(b[1]-a[1])
}

func AddInt2(a, b, c, d int) (int, int) {
	return a + c, b + d
}

func IntersectInt2(a, b [2]int) [2]int {
	l, r := Max(a[0], b[0]), Min(a[1], b[1])
	return [2]int{l, r}
}

// IsEmpty returns true if a is an empty interval
func IsEmptyInt2(a [2]int) bool {
	return a[0] > a[1]
}

func SwapInt2(a [2]int) [2]int {
	return [2]int{a[1], a[0]}
}

func LexInt2(a, b [2]int) bool {
	return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
}

func SortLex2(s [][2]int) {
	sort.Slice(s, func(i, j int) bool {
		return LexInt2(s[i], s[j])
	})
}

func LexInt3(a, b [3]int) bool {
	return a[0] < b[0] ||
		a[0] == b[0] && a[1] < b[1] ||
		a[0] == b[0] && a[1] == b[1] && a[2] < b[2]
}

func SurfaceOfCuboid(a, b, c int) int {
	return 2 * (a*b + a*c + b*c)
}

func VolumeOfCuboid(a, b, c int) int {
	return a * b * c
}

func Append[T any](xs *[]T, el ...T) {
	*xs = append(*xs, el...)
}

func SortSlice[T any](x []T, f func(x, y T) bool) {
	sort.Slice(x, func(i, j int) bool {
		return f(x[i], x[j])
	})
}

func StringOfPointerSlice[T any](l []*T) string {
	var sb strings.Builder
	sb.WriteRune('[')
	for _, x := range l {
		sb.WriteString(fmt.Sprintf("%v", *x))
		sb.WriteString(" ")
	}
	sb.WriteRune(']')
	return sb.String()
}

func StringOfList(l *list.List) string {
	var sb strings.Builder
	sb.WriteRune('[')
	for e := l.Front(); e != nil; e = e.Next() {
		sb.WriteString(fmt.Sprintf("%v", e.Value))
		sb.WriteString(" ")
	}
	sb.WriteRune(']')
	return sb.String()
}

func SumIntvSizes(xs [][2]int) int {
	// assumes intervals are disjoint and non-empty
	result := 0
	for _, x := range xs {
		result += x[1] - x[0] + 1
	}
	return result
}

func ReverseRunes(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

type Bits uint64

func Set(b Bits, flag Bits) Bits    { return b | flag }
func Clear(b Bits, flag Bits) Bits  { return b &^ flag }
func Toggle(b Bits, flag Bits) Bits { return b ^ flag }
func Has(b Bits, flag Bits) bool    { return b&flag != 0 }

func SetBit(b Bits, pos int) Bits    { return Set(b, 1<<pos) }
func ClearBit(b Bits, pos int) Bits  { return Clear(b, 1<<pos) }
func ToggleBit(b Bits, pos int) Bits { return Toggle(b, 1<<pos) }
func HasBit(b Bits, pos int) bool    { return Has(b, 1<<pos) }
