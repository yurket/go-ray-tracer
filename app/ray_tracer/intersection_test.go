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

	xs := r.Intersect(&s)

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
