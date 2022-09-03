package ray_tracer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatingAndQueringRays(t *testing.T) {
	origin := NewPoint(1, 2, 3)
	direction := NewVector(4, 5, 6)

	r := NewRay(origin, direction)

	require.True(t, r.origin.Equal(origin))
	require.True(t, r.direction.Equal(direction))
}

func TestComputingRayPositionAfterElapsedTimeT(t *testing.T) {
	origin, direction := NewPoint(2, 3, 4), NewVector(1, 0, 0)
	r := NewRay(origin, direction)

	require.True(t, r.CalcPosition(0).Equal(NewPoint(2, 3, 4)))
	require.True(t, r.CalcPosition(1).Equal(NewPoint(3, 3, 4)))
	require.True(t, r.CalcPosition(-1).Equal(NewPoint(1, 3, 4)))
	require.True(t, r.CalcPosition(2.5).Equal(NewPoint(4.5, 3, 4)))
}

func TestRayIntersectsSphereAtTwoPoints(t *testing.T) {
	origin, direction := NewPoint(0, 0, -5), NewVector(0, 0, 1)
	r := NewRay(origin, direction)
	s := NewDefaultSphere()

	xs := r.Intersect(&s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, 4, xs[0].time)
	require.EqualValues(t, 6, xs[1].time)
}

func TestRayIntersectsSphereAtATangent(t *testing.T) {
	origin, direction := NewPoint(0, 1, -5), NewVector(0, 0, 1)
	r := NewRay(origin, direction)
	s := NewDefaultSphere()

	xs := r.Intersect(&s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, 5.0, xs[0].time)
	require.EqualValues(t, xs[0], xs[1])
}

func TestRayMissesSphere(t *testing.T) {
	origin, direction := NewPoint(0, 2, -5), NewVector(0, 0, 1)
	r := NewRay(origin, direction)
	s := NewDefaultSphere()

	xs := r.Intersect(&s)

	require.EqualValues(t, 0, len(xs))
}

// Ray extends *behind* the starting point, so we'll have 2 intersections
func TestRayOriginatesInsideSphere(t *testing.T) {
	origin, direction := NewPoint(0, 0, 0), NewVector(0, 0, 1)
	r := NewRay(origin, direction)
	s := NewDefaultSphere()

	xs := r.Intersect(&s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, -1.0, xs[0].time)
	require.EqualValues(t, 1.0, xs[1].time)
}

func TestSphereCompletelyBehindRay(t *testing.T) {
	origin, direction := NewPoint(0, 0, 5), NewVector(0, 0, 1)
	r := NewRay(origin, direction)
	s := NewDefaultSphere()

	xs := r.Intersect(&s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, -6.0, xs[0].time)
	require.EqualValues(t, -4.0, xs[1].time)
}

func TestTranslatingRay(t *testing.T) {
	origin, direction := NewPoint(1, 2, 3), NewVector(0, 1, 0)
	r := NewRay(origin, direction)
	m := NewTranslationMatrix(3, 4, 5)

	r2 := r.ApplyTransform(m)

	expectOrigin, expectDirection := NewPoint(4, 6, 8), NewVector(0, 1, 0)
	require.True(t, r2.origin.Equal(expectOrigin))
	require.True(t, r2.direction.Equal(expectDirection))
}

func TestScalingRay(t *testing.T) {
	origin, direction := NewPoint(1, 2, 3), NewVector(0, 1, 0)
	r := NewRay(origin, direction)
	m := NewScalingMatrix(2, 3, 4)

	r2 := r.ApplyTransform(m)

	expectOrigin, expectDirection := NewPoint(2, 6, 12), NewVector(0, 3, 0)
	require.True(t, r2.origin.Equal(expectOrigin))
	require.True(t, r2.direction.Equal(expectDirection))
}

// TODO: test spheres creation - 2 spheres with same id?
// TODO: use test fixtures to reduce code?
