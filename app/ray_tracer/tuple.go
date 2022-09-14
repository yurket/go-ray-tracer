package ray_tracer

import (
	"math"
)

// w represents whether it's a vector or a point:
// w == 1 for point, w == 0 for vector
type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

func NewZeroTuple() Tuple {
	return Tuple{0, 0, 0, 0}
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func (t *Tuple) IsPoint() bool {
	return equal_fp(t.w, 1.0)
}

func (t *Tuple) IsVector() bool {
	return equal_fp(t.w, 0.0)
}

const EPSILON = 1e-5

func equal_fp(a float64, b float64) bool {
	diff := a - b
	return math.Abs(diff) < EPSILON
}

func (a Tuple) Equal(b Tuple) bool {
	return equal_fp(a.x, b.x) && (equal_fp(a.y, b.y) && equal_fp(a.z, b.z) && equal_fp(a.w, b.w))
}

func (a Tuple) Add(b Tuple) Tuple {
	if a.IsPoint() && b.IsPoint() {
		panic("Can't add two points!")
	}

	return Tuple{a.x + b.x, a.y + b.y, a.z + b.z, a.w + b.w}
}

func (a Tuple) Sub(b Tuple) Tuple {
	w := a.w - b.w
	if w < 0 {
		panic("Can't subtract point from vector!")
	}
	return Tuple{a.x - b.x, a.y - b.y, a.z - b.z, w}
}

func (t Tuple) Negate() Tuple {
	return Tuple{-t.x, -t.y, -t.z, -t.w}
}

func (t Tuple) Mul(c float64) Tuple {
	return Tuple{t.x * c, t.y * c, t.z * c, t.w * c}
}

func (t Tuple) Div(c float64) Tuple {
	if equal_fp(c, 0) {
		panic("Can't divide by zero!")
	}
	return Tuple{t.x / c, t.y / c, t.z / c, t.w / c}
}

func (t Tuple) Magnitude() float64 {
	if !t.IsVector() {
		panic("Magnitude isn't applicable to non-Vectors!")
	}

	// the book proposes also to add up w*w term (error?)
	return math.Sqrt(t.x*t.x + t.y*t.y + t.z*t.z)
}

func (t Tuple) Normalize() Tuple {
	m := t.Magnitude()
	return Tuple{t.x / m, t.y / m, t.z / m, t.w / m}
}

// Dot product tells "how much co-directed" 2 vectors.
// 0 to +1: pointing "same-ish" directions
// 0: perpendicular to each other
// -1 to 0: pointing "opposit-ish" directions
func (a Tuple) Dot(b Tuple) float64 {
	if !a.IsVector() || !b.IsVector() {
		panic("Dot product isn't applicable to non-Vectors!")
	}

	return a.x*b.x + a.y*b.y + a.z*b.z
}

func (a Tuple) Cross(b Tuple) Tuple {
	if !a.IsVector() || !b.IsVector() {
		panic("Cross product isn't applicable to non-Vectors!")
	}

	return NewVector(a.y*b.z-a.z*b.y, a.z*b.x-a.x*b.z, a.x*b.y-a.y*b.x)
}

// returns a so called column-vector
func (t Tuple) ToMatrix() *Matrix {
	m := NewMatrix([][]float64{
		{t.x},
		{t.y},
		{t.z},
		{t.w},
	})

	return m
}

func (t Tuple) ReflectAround(normal Tuple) Tuple {
	if !t.IsVector() {
		panic("Point can not be reflected!")
	}

	if !normal.IsVector() {
		panic("Normal must be a vector!")
	}

	// It won't work when a vector perpendicular to a normal,
	// but let's stick to the book's implementation
	negTerm := normal.Mul(2).Mul(t.Dot(normal))
	return t.Sub(negTerm)
}
