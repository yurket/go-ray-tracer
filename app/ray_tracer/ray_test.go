package ray_tracer

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

	require.True(t, r.CalcPosition(0).Equal(newPoint(2, 3, 4)))
	require.True(t, r.CalcPosition(1).Equal(newPoint(3, 3, 4)))
	require.True(t, r.CalcPosition(-1).Equal(newPoint(1, 3, 4)))
	require.True(t, r.CalcPosition(2.5).Equal(newPoint(4.5, 3, 4)))
}

func TestRayIntersectsSphereAtTwoPoints(t *testing.T) {
	origin, direction := newPoint(0, 0, -5), newVector(0, 0, 1)
	r := newRay(origin, direction)
	s := newSphere("sphere_id")

	xs := r.Intersect(&s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, 4, xs[0].t)
	require.EqualValues(t, 6, xs[1].t)
}

func TestRayIntersectsSphereAtATangent(t *testing.T) {
	origin, direction := newPoint(0, 1, -5), newVector(0, 0, 1)
	r := newRay(origin, direction)
	s := newSphere("sphere_id")

	xs := r.Intersect(&s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, 5.0, xs[0].t)
	require.EqualValues(t, xs[0], xs[1])
}

func TestRayMissesSphere(t *testing.T) {
	origin, direction := newPoint(0, 2, -5), newVector(0, 0, 1)
	r := newRay(origin, direction)
	s := newSphere("sphere_id")

	xs := r.Intersect(&s)

	require.EqualValues(t, 0, len(xs))
}

// Ray extends *behind* the starting point, so we'll have 2 intersections
func TestRayOriginatesInsideSphere(t *testing.T) {
	origin, direction := newPoint(0, 0, 0), newVector(0, 0, 1)
	r := newRay(origin, direction)
	s := newSphere("sphere_id")

	xs := r.Intersect(&s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, -1.0, xs[0].t)
	require.EqualValues(t, 1.0, xs[1].t)
}

func TestSphereCompletelyBehindRay(t *testing.T) {
	origin, direction := newPoint(0, 0, 5), newVector(0, 0, 1)
	r := newRay(origin, direction)
	s := newSphere("sphere_id")

	xs := r.Intersect(&s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, -6.0, xs[0].t)
	require.EqualValues(t, -4.0, xs[1].t)
}

func TestAndIntersectionEncapsulatesTAndObject(t *testing.T) {
	s := newSphere("sphere_id")

	i := newIntersection(3.5, &s)

	require.EqualValues(t, 3.5, i.t)
	require.EqualValues(t, s, *i.object)
}

func TestIntersectSetsTheObjectOnTheIntersection(t *testing.T) {
	origin, direction := newPoint(0, 0, -5), newVector(0, 0, 1)
	r := newRay(origin, direction)
	s := newSphere("sphere_id")

	xs := r.Intersect(&s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, s, *xs[0].object)
	require.EqualValues(t, s, *xs[1].object)
}

func TestHitWithAllIntersectionsWithPositiveT(t *testing.T) {
	s := newSphere("sphere_id")
	i1 := newIntersection(1, &s)
	i2 := newIntersection(2, &s)
	xs := []Intersection{i1, i2}

	i, ok := Hit(xs)

	require.True(t, ok)
	require.EqualValues(t, i, i1)
}

func TestHitWithSomeIntersectionsWithNegativeT(t *testing.T) {
	s := newSphere("sphere_id")
	i1 := newIntersection(-1, &s)
	i2 := newIntersection(1, &s)
	xs := []Intersection{i1, i2}

	i, ok := Hit(xs)

	require.True(t, ok)
	require.EqualValues(t, i, i2)
}

func TestNoHitWhenAllIntersectionsHaveNegativeT(t *testing.T) {
	s := newSphere("sphere_id")
	i1 := newIntersection(-1, &s)
	i2 := newIntersection(-2, &s)
	xs := []Intersection{i1, i2}

	_, ok := Hit(xs)

	require.False(t, ok)
}

func TestHitIsAlwaysTheLowestNonnegativeIntersection(t *testing.T) {
	s := newSphere("sphere_id")
	i1 := newIntersection(5, &s)
	i2 := newIntersection(6, &s)
	i3 := newIntersection(-3, &s)
	i4 := newIntersection(2, &s)
	xs := []Intersection{i1, i2, i3, i4}

	i, ok := Hit(xs)

	require.True(t, ok)
	require.EqualValues(t, i, i4)
}

func TestTranslatingRay(t *testing.T) {
	origin, direction := newPoint(1, 2, 3), newVector(0, 1, 0)
	r := newRay(origin, direction)
	m := newTranslationMatrix(3, 4, 5)

	r2 := r.ApplyTransform(m)

	expectOrigin, expectDirection := newPoint(4, 6, 8), newVector(0, 1, 0)
	require.True(t, r2.origin.Equal(expectOrigin))
	require.True(t, r2.direction.Equal(expectDirection))
}

func TestScalingRay(t *testing.T) {
	origin, direction := newPoint(1, 2, 3), newVector(0, 1, 0)
	r := newRay(origin, direction)
	m := newScalingMatrix(2, 3, 4)

	r2 := r.ApplyTransform(m)

	expectOrigin, expectDirection := newPoint(2, 6, 12), newVector(0, 3, 0)
	require.True(t, r2.origin.Equal(expectOrigin))
	require.True(t, r2.direction.Equal(expectDirection))
}

func TestSpheresDefaultTransofrmationIsIdentity(t *testing.T) {
	s := newSphere("sphere_id")

	require.True(t, s.transform.Equal(newIdentityMatrix(4)))
}

func TestSettingSphereTransformation(t *testing.T) {
	s := newSphere("sphere_id")
	translation := newTranslationMatrix(2, 3, 4)

	s.SetTransform(translation)

	transform := s.Transform()
	require.True(t, transform.Equal(translation))
}

func TestIntersectingScaledSphereWithRay(t *testing.T) {
	origin, direction := newPoint(0, 0, -5), newVector(0, 0, 1)
	r := newRay(origin, direction)
	s := newSphere("sphere_id")
	s.SetTransform(newScalingMatrix(2, 2, 2))

	xs := r.Intersect(&s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, 3, xs[0].t)
	require.EqualValues(t, 7, xs[1].t)
}

func TestIntersectingTranslatedSphereWithRay(t *testing.T) {
	origin, direction := newPoint(0, 0, -5), newVector(0, 0, 1)
	r := newRay(origin, direction)
	s := newSphere("sphere_id")
	s.SetTransform(newTranslationMatrix(5, 0, 0))

	xs := r.Intersect(&s)

	require.EqualValues(t, 0, len(xs))
}

// TODO: test spheres creation - 2 spheres with same id?
