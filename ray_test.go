package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatingAndQueringRays(t *testing.T) {
	origin := newPoint(1, 2, 3)
	direction := newVector(4, 5, 6)
	r := newRay(origin, direction)

	require.True(t, r.origin.Equal(origin))
	require.True(t, r.direction.Equal(direction))
}

func TestComputingRayPositionAfterElapsedTimeT(t *testing.T) {
	origin, direction := newPoint(2, 3, 4), newVector(1, 0, 0)
	r := newRay(origin, direction)

	require.True(t, r.Position(0).Equal(newPoint(2, 3, 4)))
	require.True(t, r.Position(1).Equal(newPoint(3, 3, 4)))
	require.True(t, r.Position(-1).Equal(newPoint(1, 3, 4)))
	require.True(t, r.Position(2.5).Equal(newPoint(4.5, 3, 4)))
}

func TestRayIntersectsSphereAtTwoPoints(t *testing.T) {
	origin, direction := newPoint(0, 0, -5), newVector(0, 0, 1)
	r := newRay(origin, direction)
	s := newSphere("sphere_id")
	xs := r.Intersect(s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, 4, xs[0])
	require.EqualValues(t, 6, xs[1])
}

func TestRayIntersectsSphereAtATangent(t *testing.T) {
	origin, direction := newPoint(0, 1, -5), newVector(0, 0, 1)
	r := newRay(origin, direction)
	s := newSphere("sphere_id")
	xs := r.Intersect(s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, 5.0, xs[0])
	require.True(t, xs[0] == xs[1])
}

func TestRayMissesSphere(t *testing.T) {
	origin, direction := newPoint(0, 2, -5), newVector(0, 0, 1)
	r := newRay(origin, direction)
	s := newSphere("sphere_id")
	xs := r.Intersect(s)

	require.EqualValues(t, 0, len(xs))
}

// Ray extends *behind* the starting point, so we'll have 2 intersections
func TestRayOriginatesInsideSphere(t *testing.T) {
	origin, direction := newPoint(0, 0, 0), newVector(0, 0, 1)
	r := newRay(origin, direction)
	s := newSphere("sphere_id")
	xs := r.Intersect(s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, -1.0, xs[0])
	require.EqualValues(t, 1.0, xs[1])
}

func TestSphereCompletelyBehindRay(t *testing.T) {
	origin, direction := newPoint(0, 0, 5), newVector(0, 0, 1)
	r := newRay(origin, direction)
	s := newSphere("sphere_id")
	xs := r.Intersect(s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, -6.0, xs[0])
	require.EqualValues(t, -4.0, xs[1])

}

// TODO: test spheres creation - 2 spheres with same id?
