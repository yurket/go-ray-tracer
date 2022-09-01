package ray_tracer

import (
	"math"
)

// Each component in represented by a number from 0 to 1
type Color struct {
	r float64
	g float64
	b float64
}

var BLACK Color = Color{0, 0, 0}
var RED Color = Color{1, 0, 0}
var GREEN Color = Color{0, 1, 0}
var BLUE Color = Color{0, 0, 1}
var WHITE Color = Color{1, 1, 1}

func newColor(r, g, b float64) Color {
	return Color{r, g, b}
}

func (a Color) Equal(b Color) bool {
	return equal_fp(a.r, b.r) && equal_fp(a.g, b.g) && equal_fp(a.b, b.b)
}

func (a Color) Add(b Color) Color {
	return Color{a.r + b.r, a.g + b.g, a.b + b.b}
}

func (a Color) Sub(b Color) Color {
	return Color{a.r - b.r, a.g - b.g, a.b - b.b}
}

func (a Color) MultScalar(n float64) Color {
	return Color{a.r * n, a.g * n, a.b * n}
}

func (a Color) MultHadamar(b Color) Color {
	return Color{a.r * b.r, a.g * b.g, a.b * b.b}
}

func scaleColorComponent(c float64) int {
	res := c * MAX_COLORS
	res = math.Max(res, 0)
	res = math.Min(res, 255)
	return int(math.Round(res))
}

func (a Color) ToScaledRgbComponents() [3]int {
	return [3]int{scaleColorComponent(a.r), scaleColorComponent(a.g), scaleColorComponent(a.b)}
}
