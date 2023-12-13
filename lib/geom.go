package lib

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
)

// * 2D floats: Vect2
// * 3D floats: Vect3

func Det22[T Num, E ~[2]T](a, b E) T {
	return a[0]*b[1] - a[1]*b[0]
}

func Det32[T Num, E ~[2]T](a, b, c E) T {
	return a[0]*(b[1]-c[1]) + b[0]*(c[1]-a[1]) + c[0]*(a[1]-b[1])
}

func Colinear(a, b, c [2]float64) bool {
	det := Det32(a, b, c)
	return math.Abs(det) < EPSILON
}

type Vect2 [2]float64

func (v Vect2) Add(w [2]float64) Vect2 {
	return Vect2{v[0] + w[0], v[1] + w[1]}
}

func (v Vect2) Subtract(w [2]float64) Vect2 {
	return Vect2{v[0] - w[0], v[1] - w[1]}
}

// Angle between vector and x-axis. result range is -π to π
func (v Vect2) Angle() float64 {
	return math.Atan2(v[1], v[0])
}

// AngleBetween returns angle between vectors v and w at origin
func (v Vect2) AngleBetween(w [2]float64) float64 {
	dot := v[0]*w[0] + v[1]*w[1]
	det := v[0]*w[1] - v[1]*w[0]
	return math.Atan2(det, dot)
}

// AngleAt is angle between lines p.To(q) and p.To(r)
func (p Vect2) AngleAt(q, r [2]float64) float64 {
	return p.To(q).AngleBetween(p.To(r))
}

func (v Vect2) Dist(w [2]float64) float64 {
	return math.Hypot(v[0]-w[0], v[1]-w[1])
}

func (v Vect2) DistSquared(q [2]float64) float64 {
	dx := v[0] - q[0]
	dy := v[1] - q[1]
	return dx*dx + dy*dy
}

// Dot product (= "inner product")
func (v Vect2) Dot(w [2]float64) float64 {
	return v[0]*w[0] + v[1]*w[1]
}

func (v Vect2) Mag() float64 {
	return math.Hypot(v[0], v[1])
}

func (v Vect2) MagSquared() float64 {
	return v[0]*v[0] + v[1]*v[1]
}

func (v Vect2) Normal() Vect2 {
	return Vect2{-v[1], v[0]}
}

func (v Vect2) Normalize() Vect2 {
	mag := v.Mag()
	return Vect2{v[0] / mag, v[1] / mag}
}

func (a Vect2) Orientation(b, c [2]float64) int {
	det := Det32(a, b, c)
	switch {
	case det < 0:
		return -1 // clockwise
	case det > 0:
		return +1 // counter-clockwise
	default:
		return 0
	}
}

func PrintlnVect2s(vs []Vect2) {
	fmt.Print("[")
	for _, x := range vs {
		fmt.Print(x, " ")
	}
	fmt.Println("]")
}

// Rotate vector v around the origin by θ
func (v Vect2) Rotate(θ float64) Vect2 {
	x, y := v[0], v[1]
	x1 := math.Cos(θ)*x - math.Sin(θ)*y
	y1 := math.Sin(θ)*x + math.Cos(θ)*y
	return Vect2{x1, y1}
}

// RotateAround rotates point q around point p by θ
func (p Vect2) RotateAround(q [2]float64, θ float64) Vect2 {
	w := p.To(q).Rotate(θ)
	return p.Add(w)
}

func (v Vect2) Scale(λ float64) Vect2 {
	return Vect2{λ * v[0], λ * v[1]}
}

func ScanVect2(sc *bufio.Scanner) Vect2 {
	return Vect2{ScanFloat(sc), ScanFloat(sc)}
}

func ScanVect2s(sc *bufio.Scanner, n int) []Vect2 {
	result := make([]Vect2, n)
	for i := range result {
		result[i] = Vect2{ScanFloat(sc), ScanFloat(sc)}
	}
	return result
}

func (v Vect2) String() string {
	return fmt.Sprintf("[%.3f, %.3f]", v[0], v[1])
}

// To(p,q) is vector from p to q
func (p Vect2) To(q Vect2) Vect2 {
	return Vect2{q[0] - p[0], q[1] - p[1]}
}

func Vect2DOfPolar(d, θ float64) Vect2 {
	return Vect2{d * math.Cos(θ), d * math.Sin(θ)}
}

func (v Vect2) X() float64 {
	return v[0]
}

func (v Vect2) Y() float64 {
	return v[1]
}

// big.Rat 2D line intersections

type RLine [2][2]*big.Rat

func Det22Rat(a, b [2]*big.Rat) *big.Rat {
	c := &big.Rat{}
	d := &big.Rat{}
	c.Mul(a[0], b[1])
	d.Mul(a[1], b[0])
	return c.Sub(c, d)
}

func (l RLine) Intersect(s RLine) ([2]*big.Rat, bool) {
	zero := func() *big.Rat {
		return &big.Rat{}
	}
	x := [4]*big.Rat{l[0][0], l[1][0], s[0][0], s[1][0]}
	y := [4]*big.Rat{l[0][1], l[1][1], s[0][1], s[1][1]}
	detLS := [2]*big.Rat{Det22Rat(l[0], l[1]), Det22Rat(s[0], s[1])}
	dx := [2]*big.Rat{zero().Sub(x[0], x[1]), zero().Sub(x[2], x[3])}
	dy := [2]*big.Rat{zero().Sub(y[0], y[1]), zero().Sub(y[2], y[3])}
	a := Det22Rat(detLS, dx)
	b := Det22Rat(detLS, dy)
	d := Det22Rat(dx, dy)
	if d.Cmp(zero()) != 0 {
		a = zero().Quo(a, d)
		b = zero().Quo(b, d)
		return [2]*big.Rat{a, b}, true
	} else {
		return [2]*big.Rat{}, false
	}
}

func DistL1Rat(a, b [2]*big.Rat) *big.Rat {
	zero := &big.Rat{}
	minusOne := big.NewRat(-1, 1)
	c := (&big.Rat{}).Sub(b[0], a[0])
	d := (&big.Rat{}).Sub(b[1], a[1])
	if c.Cmp(zero) < 0 {
		c.Mul(c, minusOne)
	}
	if d.Cmp(zero) < 0 {
		d.Mul(d, minusOne)
	}
	return c.Add(c, d)
}

// IsBetween checks whether point b is between a and c
// It assumes all three points are on one line
func IsBetween(a, b, c [2]*big.Rat) bool {
	dab := DistL1Rat(a, b)
	dbc := DistL1Rat(b, c)
	dac := DistL1Rat(a, c)
	return dab.Add(dab, dbc).Cmp(dac) == 0
}
