package ray_tracer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersectionEncapsulatesTimeAndObject(t *testing.T) {
	s := NewDefaultSphere()

	i := NewIntersection(3.5, &s)

	require.EqualValues(t, 3.5, i.time)
	require.EqualValues(t, s, *i.object)
}

func TestIntersectSetsTheObjectOnTheIntersection(t *testing.T) {
	origin, direction := NewPoint(0, 0, -5), NewVector(0, 0, 1)
	r := NewRay(origin, direction)
	s := NewDefaultSphere()

	xs := s.IntersectWith(&r)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, s, *xs[0].object)
	require.EqualValues(t, s, *xs[1].object)
}

func TestHitWithAllIntersectionsWithPositiveT(t *testing.T) {
	s := NewDefaultSphere()
	i1 := NewIntersection(1, &s)
	i2 := NewIntersection(2, &s)
	xs := []Intersection{i1, i2}

	i, ok := Hit(xs)

	require.True(t, ok)
	require.EqualValues(t, i, i1)
}

func TestHitWithSomeIntersectionsWithNegativeT(t *testing.T) {
	s := NewDefaultSphere()
	i1 := NewIntersection(-1, &s)
	i2 := NewIntersection(1, &s)
	xs := []Intersection{i1, i2}

	i, ok := Hit(xs)

	require.True(t, ok)
	require.EqualValues(t, i, i2)
}

func TestNoHitWhenAllIntersectionsHaveNegativeT(t *testing.T) {
	s := NewDefaultSphere()
	i1 := NewIntersection(-1, &s)
	i2 := NewIntersection(-2, &s)
	xs := []Intersection{i1, i2}

	_, ok := Hit(xs)

	require.False(t, ok)
}

func TestHitIsAlwaysTheLowestNonnegativeIntersection(t *testing.T) {
	s := NewDefaultSphere()
	i1 := NewIntersection(5, &s)
	i2 := NewIntersection(6, &s)
	i3 := NewIntersection(-3, &s)
	i4 := NewIntersection(2, &s)
	xs := []Intersection{i1, i2, i3, i4}

	i, ok := Hit(xs)

	require.True(t, ok)
	require.EqualValues(t, i, i4)
}

func TestPrecomputingTheStateOfAnIntersection(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewDefaultSphere()
	i := NewIntersection(4, &s)

	comps := PrepareIntersectionComputations(i, r)

	require.EqualValues(t, comps.intersectionTime, i.time)
	require.EqualValues(t, comps.intersectionObject, i.object)
	require.True(t, comps.intersectionPoint.Equal(NewPoint(0, 0, -1)))
	require.True(t, comps.eyev.Equal(NewVector(0, 0, -1)))
	require.True(t, comps.objectNormalv.Equal(NewVector(0, 0, -1)))
}

func TestTheHitWithOutsideIntersection(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewDefaultSphere()
	i := NewIntersection(4, &s)

	comps := PrepareIntersectionComputations(i, r)

	require.False(t, comps.insideHit)
}

func TestTheHitWithInsideIntersection(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	s := NewDefaultSphere()
	i := NewIntersection(1, &s)

	comps := PrepareIntersectionComputations(i, r)

	require.True(t, comps.intersectionPoint.Equal(NewPoint(0, 0, 1)))
	require.True(t, comps.eyev.Equal(NewVector(0, 0, -1)))
	require.True(t, comps.insideHit)
	// normal would have been (0, 0, 1), but is inverted!
	require.True(t, comps.objectNormalv.Equal(NewVector(0, 0, -1)))
}

// Test, that "acne effect" can be successfully overcome
func TestTheHitShouldOffsetThePoint(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewDefaultSphere()
	s.SetTransform(NewTranslationMatrix(0, 0, 1))
	i := NewIntersection(5, &s)

	comps := PrepareIntersectionComputations(i, r)

	require.Less(t, comps.overPoint.z, -EPSILON/2)
	require.Greater(t, comps.intersectionPoint.z, comps.overPoint.z)
}
