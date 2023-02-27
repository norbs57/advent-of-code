package lib

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"strings"
	"testing"

	"golang.org/x/exp/constraints"
)

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

// Generics

type Integer interface {
	constraints.Integer
}
type Num interface {
	constraints.Integer | constraints.Float
}

type SignedNum interface {
	constraints.Signed | constraints.Float
}

// Ordered represents the set of types for which the '<' operator work.
type Ordered interface {
	constraints.Integer | constraints.Float | ~string
}

type Pair[U any, V any] struct {
	Fst U
	Snd V
}

func (p Pair[U, V]) Extract() (U, V) {
	return p.Fst, p.Snd
}

func MkPair[U any, V any](u U, v V) Pair[U, V] {
	return Pair[U, V]{u, v}
}

func Swap[T any](a [2]T) [2]T {
	return [2]T{a[1], a[0]}
}

func Abs[T SignedNum](a T) T {
	if a >= T(0) {
		return a
	} else {
		return -a
	}
}

func Sign[T SignedNum](a T) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return 1
	default:
		return 0
	}
}

func Append[T any](xs *[]T, el ...T) {
	*xs = append(*xs, el...)
}

func Last[E any](xs []E) E {
	return xs[len(xs)-1]
}

func Map[T, S any](f func(T) S, xs []T) []S {
	result := make([]S, len(xs))
	for i, t := range xs {
		result[i] = f(t)
	}
	return result
}

func PrefixSums[T Num](xs []T) []T {
	if len(xs) == 0 {
		return xs
	}
	res := make([]T, 0, len(xs))
	sum := T(0)
	for _, x := range xs {
		sum += x
		res = append(res, sum)
	}
	return res
}

func PrefixSumsFunc[T Num](f func(int) T, n int) []T {
	res := make([]T, 1, n)
	sum := T(0)
	for i := 0; i < n; i++ {
		sum += f(i)
		res = append(res, sum)
	}
	return res
}

func Prod[T Num](xs []T) T {
	result := T(1)
	for _, x := range xs {
		result *= x
	}
	return result
}

func Sum[T Num](xs []T) T {
	result := T(0)
	for _, x := range xs {
		result += x
	}
	return result
}

func ModInt(a, b int) int {
	// result is positive if b is positivew
	return (a%b + b) % b
}

func Dist2[T SignedNum](a, b [2]T) T {
	return Abs(b[0]-a[0]) + Abs(b[1]-a[1])
}

func Add2[T Num](p, q [2]T) [2]T {
	return [2]T{p[0] + q[0], p[1] + q[1]}
}

func Add3[T Num](p, q [3]T) [3]T {
	return [3]T{p[0] + q[0], p[1] + q[1], p[2] + q[2]}
}

func Greater[T Ordered](p, q T) bool {
	return p > q
}

func Greater2[T Ordered](p, q [2]T) bool {
	return p[0] > q[0] || p[0] == q[0] && p[1] > q[1]
}

func Greater3[T Ordered](a, b [3]T) bool {
	return a[0] > b[0] || a[0] == b[0] && a[1] > b[1] ||
		a[0] == b[0] && a[1] == b[1] && a[2] > b[2]
}

func Less2[T Ordered](a, b [2]T) bool {
	return a[0] < b[0] || a[0] == b[0] && a[1] < b[1]
}

func LessOrEqual2[T Ordered](a, b [2]T) bool {
	return a == b || a[0] < b[0] || a[0] == b[0] && a[1] < b[1]
}

func Less3[T Ordered](a, b [3]T) bool {
	return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] ||
		a[0] == b[0] && a[1] == b[1] && a[2] < b[2]
}

func MkNumsUpto[E Num](n int) []E {
	result := make([]E, n)
	for i := range result {
		result[i] = E(i)
	}
	return result
}

func MkSliceFilled[E any](rows int, t E) []E {
	result := make([]E, rows)
	for i := range result {
		result[i] = t
	}
	return result
}

func MkSlice2[E any](rows, cols int) [][]E {
	result := make([][]E, rows)
	for i := range result {
		result[i] = make([]E, cols)
	}
	return result
}

func MkSlice2Filled[E any](rows, cols int, t E) [][]E {
	result := make([][]E, rows)
	for i := range result {
		result[i] = make([]E, cols)
		for j := range result[i] {
			result[i][j] = t
		}
	}
	return result
}

func PopBack[E any](x *[]E) E {
	n := len(*x) - 1
	res := (*x)[n]
	*x = (*x)[:n]
	return res
}

func PopFront[E any](x *[]E) E {
	res := (*x)[0]
	*x = (*x)[1:]
	return res
}

func PushBack[E any](x *[]E, e E) {
	*x = append((*x), e)
}

func PushFront[E any](x *[]E, e E) {
	n := len(*x)
	*x = append((*x), e)
	copy((*x)[1:], (*x)[:n])
	(*x)[0] = e
}

func Reverse[E any](x []E) {
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
}

func SetMax[T Ordered](a *T, b T) {
	if b > *a {
		*a = b
	}
}

func SetMin[T Ordered](a *T, b T) {
	if b < *a {
		*a = b
	}
}

func SlicesEqual[E comparable](s1, s2 []E) bool {
	// https://pkg.go.dev/golang.org/x/exp/slices#Equal
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func SortLess2[T Ordered](s [][2]int) {
	sort.Slice(s, func(i, j int) bool {
		return Less2(s[i], s[j])
	})
}

func SortSliceDecr[T Ordered](xs []T) {
	sort.Slice(xs, func(i, j int) bool {
		return xs[i] > xs[j]
	})
}

func SortSliceFunc[E any, T Ordered](xs []E, f func(E) T) {
	sort.Slice(xs, func(i, j int) bool {
		return f(xs[i]) < f(xs[j])
	})
}

func SortSliceLessFunc[E any](xs []E, less func(a, b E) bool) {
	sort.Slice(xs, func(i, j int) bool {
		return less(xs[i], xs[j])
	})
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

func Lcm(a, b int) int {
	gab := Gcd(a, b)
	return (a / gab) * b
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

func SurfaceOfCuboid(a, b, c int) int {
	return 2 * (a*b + a*c + b*c)
}

func VolumeOfCuboid(a, b, c int) int {
	return a * b * c
}

// Intersection of half-open intervals
func Intersect[T Ordered](x, y [2]T) [2]T {
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

func SumItvSizes(xs [][2]int) int {
	// assumes intervals are disjoint and non-empty
	result := 0
	for _, x := range xs {
		result += x[1] - x[0] + 1
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
