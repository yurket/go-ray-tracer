package ray_tracer

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatingZeroTuple(t *testing.T) {
	point := NewZeroTuple()
	if point.x != 0 || point.y != 0 || point.z != 0 || point.w != 0 {
		t.Error("All components should be zero!")
		t.Log("point: ", point)
	}
}

func TestVectorCreation(t *testing.T) {
	v := NewVector(1, 2, 3)

	require.True(t, v.IsVector())
}

func TestIsPoint(t *testing.T) {
	p := Tuple{4.3, -4.2, 3.1, 0.99999}
	require.True(t, p.IsPoint(), "tuple %v must be a point, because w == 1.0!", p)
}

func TestIsVector(t *testing.T) {
	v := Tuple{4.3, -4.2, 3.1, 0.0}
	require.True(t, v.IsVector(), "tuple %v must be considered vector, because w == 0.0!", v)
}

// I'm not sure if such a tuple should be erroneus, so for now let it be
func TestNeitherVectorNorPoint(t *testing.T) {
	s := Tuple{4, 3, 2, 0.5}

	require.False(t, s.IsVector() || s.IsPoint(), "tuple with w == %v should be neither vector nor point!", s.w)
}

func TestTupleEquality(t *testing.T) {
	t1 := NewPoint(1, 2, 3)
	t2 := NewPoint(1, 2, 3)

	require.True(t, t1.Equal(t2))
}

func TestAddition(t *testing.T) {
	p := Tuple{3, -2, 5, 1}
	v := Tuple{-2, 3, 1, 0}

	res := p.Add(v)
	expect := Tuple{1, 1, 6, 1}

	require.True(t, res.Equal(expect))
	require.True(t, res.IsPoint())
}

func TestSubtractingTwoPointsGivesVector(t *testing.T) {
	p1 := NewPoint(3, 2, 1)
	p2 := NewPoint(5, 6, 7)

	res := p1.Sub(p2)
	expect := NewVector(-2, -4, -6)

	require.True(t, res.Equal(expect), "res %v != expected %v", res, expect)
}

func TestSubtractingTwoVectorsGivesVector(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(5, 6, 7)

	res := v1.Sub(v2)
	expect := NewVector(-4, -4, -4)

	require.True(t, res.Equal(expect))
}

func TestSubtractingVectorFromPointGivesPoint(t *testing.T) {
	p := NewPoint(3, 2, 1)
	v := NewVector(5, 6, 7)

	res := p.Sub(v)
	expect := NewPoint(-2, -4, -6)

	require.True(t, res.Equal(expect))
}

func TestSubtractingPointFromVectorDoesntMakeSense(t *testing.T) {
	v := NewVector(5, 6, 7)
	p := NewPoint(3, 2, 1)

	require.Panics(t, func() { v.Sub(p) })
}

func TestNegation(t *testing.T) {
	s := Tuple{1, 2, -3, -4}

	res := s.Negate()
	expect := Tuple{-1, -2, 3, 4}

	require.True(t, res.Equal(expect))
}

func TestMul(t *testing.T) {
	s := Tuple{1, 2, -3, -4}

	res := s.Mul(4)
	require.True(t, res.Equal(Tuple{4, 8, -12, -16}))

	res = s.Mul(0.5)
	require.True(t, res.Equal(Tuple{0.5, 1, -1.5, -2}))

}

func TestDiv(t *testing.T) {
	s := Tuple{1, 2, -3, -4}

	res := s.Div(2)
	expect := Tuple{0.5, 1, -1.5, -2}
	require.True(t, res.Equal(expect))
}

func TestZeroDivisionPanics(t *testing.T) {
	s := Tuple{1, 2, -3, -4}

	require.Panics(t, func() { s.Div(0) })
}

func TestMagnitude(t *testing.T) {
	unitMagnitude := 1.0

	require.True(t, equal_fp(NewVector(1, 0, 0).Magnitude(), unitMagnitude))
	require.True(t, equal_fp(NewVector(0, 1, 0).Magnitude(), unitMagnitude))
	require.True(t, equal_fp(NewVector(0, 0, 1).Magnitude(), unitMagnitude))

	require.True(t, equal_fp(NewVector(1, 2, 3).Magnitude(), math.Sqrt(14)))
	require.True(t, equal_fp(NewVector(-1, -2, -3).Magnitude(), math.Sqrt(14)))
}

func TestMagnitudPanicsOnNonVectors(t *testing.T) {
	p := NewPoint(1, 2, 3)
	nonV := Tuple{1, 1, 1, 4}

	require.Panics(t, func() { p.Magnitude() })
	require.Panics(t, func() { nonV.Magnitude() })
}

func TestNormalization(t *testing.T) {
	v := NewVector(4, 0, 0)
	vNorm := NewVector(1, 0, 0)
	require.True(t, v.Normalize().Equal(vNorm))

	v = NewVector(1, 2, 3)
	vNorm = NewVector(0.26726, 0.53452, 0.80178)
	require.True(t, v.Normalize().Equal(vNorm))
}

func TestNormalizationPanicsOnNonVectors(t *testing.T) {
	p := NewPoint(1, 2, 3)
	nonV := Tuple{1, 1, 1, 4}

	require.Panics(t, func() { p.Normalize() })
	require.Panics(t, func() { nonV.Normalize() })
}

func TestDotProduct(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(-1, 0, 0)

	res := v1.Dot(v2)
	expect := -1.0

	require.True(t, equal_fp(res, expect))
}

func TestCrossProduct(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)

	res := v1.Cross(v2)
	expect := NewVector(-1, 2, -1)
	require.True(t, res.Equal(expect))

	res = v2.Cross(v1)
	expect = NewVector(1, -2, 1)
	require.True(t, res.Equal(expect))
}

func TestReflectingVectorApprochingAt45Degrees(t *testing.T) {
	v := NewVector(1, -1, 0)
	n := NewVector(0, 1, 0)

	res := v.ReflectAround(n)
	expect := NewVector(1, 1, 0)

	require.True(t, res.Equal(expect))
}

func TestReflectingVectorOffSlantedSurface(t *testing.T) {
	v := NewVector(0, -1, 0)
	n := NewVector(math.Sqrt(2)/2., math.Sqrt(2)/2., 0)

	res := v.ReflectAround(n)
	expect := NewVector(1, 0, 0)

	require.True(t, res.Equal(expect))
}
