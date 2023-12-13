package lib

import (
	"container/list"
	"fmt"
	"math"
	"strings"
	"testing"
)

const EPSILON = 1e-12

var AocFunMap map[int][]func()

func init() {
	AocFunMap = make(map[int][]func())
}

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

// Maths

func Gcd(a, b int) int {
	b = Abs(b)
	if a == 0 {
		return b
	}
	a = Abs(a)
	for b != 0 {
		remainder := a % b
		a = b
		b = remainder
	}
	return a
}

func Lcm(xs ...int) int {
	if len(xs) == 1 {
		return xs[0]
	}
	lcm := func(a, b int) int {
		gab := Gcd(a, b)
		return (a / gab) * b
	}
	res := xs[0]
	for i := 1; i < len(xs); i++ {
		res = lcm(res, xs[i])
	}
	return res
}

func PolygonArea(pts [][2]int) int {
	area := 0
	for i := range pts {
		p, q := pts[i], pts[(i+1)%len(pts)]
		area += (p[0] - q[0]) * (p[1] + q[1])
	}
	return Abs(area / 2)
}

func QuadraticSolve(a, b, c float64) (float64, float64, bool) {
	discrim := b*b - 4*a*c
	if discrim < 0 {
		return 0.0, 0.0, false
	}
	x1 := (-b - math.Sqrt(discrim)) / (2 * a)
	x2 := (-b + math.Sqrt(discrim)) / (2 * a)
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	return x1, x2, true
}

func SurfaceOfCuboid3(width, height, depth int) int {
	return 2 * (width*height + width*depth + height*depth)
}

func TernarySearch(f func(float64) float64, a, b, precision float64) float64 {
	// find minimum of f on interval [a,b] https://en.wikipedia.org/wiki/Ternary_search
	for math.Abs(b-a) >= precision {
		left_third := a + (b-a)/3
		right_third := b - (b-a)/3
		if f(left_third) > f(right_third) {
			a = left_third
		} else {
			b = right_third
		}
	}
	return (a + b) / 2
}

func IntersectCuboids[T Ordered](xs, ys [][2]T) [][2]T {
	res := make([][2]T, len(xs))
	for i := range res {
		res[i] = IntersectItvs(xs[i], ys[i])
	}
	return res
}

// Intersection of two (closed) intervals
func IntersectItvs[T Ordered](x, y [2]T) [2]T {
	a, b := x[0], x[1]
	c, d := y[0], y[1]
	var result [2]T
	if b <= c || d <= a {
		result = [2]T{}
	} else {
		result = [2]T{max(a, c), min(b, d)}
	}
	return result
}

// IsEmpty returns true if the interval is empty
func IsEmptyItv[T Ordered](a [2]T) bool {
	return a[0] > a[1]
}

func MergeItvs[T Ordered](itvs [][2]T) [][2]T {
	// assume itvs is lex sorted
	if len(itvs) <= 1 {
		return itvs
	}
	result := make([][2]T, 1)
	result[0] = itvs[0]
	itvs = itvs[1:]
	last := 0
	for _, itv := range itvs {
		if itv[0] <= result[last][1] {
			result[last][1] = max(itv[1], result[last][1])
		} else {
			result = append(result, itv)
			last++
		}
	}

	return result
}

// Miscellaneous

func IsDigit[T rune | byte](ch T) bool {
	return '0' <= ch && ch <= '9'
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

func TransposeStrings(xs []string) []string {
	res := make([]string, -0, len(xs[0]))
	for i := 0; i < len(xs[0]); i++ {
		col := make([]byte, len(xs))
		for j := 0; j < len(xs); j++ {
			col[j] = xs[j][i]
		}
		res = append(res, string(col))
	}
	return res
}

type Bits uint64

func (b Bits) SetFlag(flag Bits) Bits    { return b | flag }
func (b Bits) ClearFlag(flag Bits) Bits  { return b &^ flag }
func (b Bits) ToggleFlag(flag Bits) Bits { return b ^ flag }
func (b Bits) HasFlag(flag Bits) bool    { return b&flag != 0 }

func (b Bits) Set(pos int) Bits    { return b.SetFlag(1 << pos) }
func (b Bits) Clear(pos int) Bits  { return b.ClearFlag(1 << pos) }
func (b Bits) Toggle(pos int) Bits { return b.ToggleFlag(1 << pos) }
func (b Bits) Has(pos int) bool    { return b.HasFlag(1 << pos) }
