package ray_tracer

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

var COS45 = math.Sqrt(2) / 2.0

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

func TestAndIntersectionEncapsulatesTAndObject(t *testing.T) {
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

func TestSpheresDefaultTransofrmationIsIdentity(t *testing.T) {
	s := NewDefaultSphere()

	require.True(t, s.transform.Equal(NewIdentityMatrix(4)))
}

func TestSettingSphereTransformation(t *testing.T) {
	s := NewDefaultSphere()
	translation := NewTranslationMatrix(2, 3, 4)

	s.SetTransform(translation)

	transform := s.Transform()
	require.True(t, transform.Equal(translation))
}

func TestIntersectingScaledSphereWithRay(t *testing.T) {
	origin, direction := NewPoint(0, 0, -5), NewVector(0, 0, 1)
	r := NewRay(origin, direction)
	s := NewDefaultSphere()
	s.SetTransform(NewScalingMatrix(2, 2, 2))

	xs := r.Intersect(&s)

	require.EqualValues(t, 2, len(xs))
	require.EqualValues(t, 3, xs[0].time)
	require.EqualValues(t, 7, xs[1].time)
}

func TestIntersectingTranslatedSphereWithRay(t *testing.T) {
	origin, direction := NewPoint(0, 0, -5), NewVector(0, 0, 1)
	r := NewRay(origin, direction)
	s := NewDefaultSphere()
	s.SetTransform(NewTranslationMatrix(5, 0, 0))

	xs := r.Intersect(&s)

	require.EqualValues(t, 0, len(xs))
}

func TestNormalOnSphereX(t *testing.T) {
	s := NewDefaultSphere()
	n := s.NormalAt(NewPoint(1, 0, 0))

	expect := NewVector(1, 0, 0)
	require.True(t, n.Equal(expect))
}

func TestNormalOnSphereY(t *testing.T) {
	s := NewDefaultSphere()
	n := s.NormalAt(NewPoint(0, 1, 0))

	expect := NewVector(0, 1, 0)
	require.True(t, n.Equal(expect))

}
func TestNormalOnSphereZ(t *testing.T) {
	s := NewDefaultSphere()
	n := s.NormalAt(NewPoint(0, 0, 1))

	expect := NewVector(0, 0, 1)
	require.True(t, n.Equal(expect))
}

func TestNormalOnSphereNonAxial(t *testing.T) {
	s := NewDefaultSphere()
	v := math.Sqrt(3) / 3.0
	n := s.NormalAt(NewPoint(v, v, v))

	expect := NewVector(v, v, v)
	require.True(t, n.Equal(expect))
}

func TestNormalVectorsAreAlwaysNormalized(t *testing.T) {
	s := NewDefaultSphere()
	v := math.Sqrt(3) / 3.0
	n := s.NormalAt(NewPoint(v, v, v))

	expect := n.Normalize()
	require.True(t, n.Equal(expect))
}

func TestComputingNormalOnTranslatedSphere(t *testing.T) {
	s := NewDefaultSphere()
	s.SetTransform(NewTranslationMatrix(0, 1, 0))
	n := s.NormalAt(NewPoint(0, 1.70711, -0.70711))

	expect := NewVector(0, 0.70711, -0.70711)
	require.True(t, n.Equal(expect))
}

func TestComputingNormalOnTransformedSphere(t *testing.T) {
	s := NewDefaultSphere()
	transform := NewIdentityMatrix(4).RotateZ(math.Pi/5).Scale(1, 0.5, 1)
	s.SetTransform(transform)
	n := s.NormalAt(NewPoint(0, COS45, -COS45))

	expect := NewVector(0, 0.97014, -0.24254)
	require.True(t, n.Equal(expect))

}

func TestCreatingPointLight(t *testing.T) {
	pos := NewPoint(0, 0, 0)
	intensity := WHITE

	pl := NewPointLight(pos, intensity)

	require.True(t, intensity.Equal(pl.intensity))
	require.True(t, pos.Equal(pl.position))
}

func TestCreatingDefaultMaterial(t *testing.T) {
	m := NewDefaultMaterial()

	require.True(t, m.color.Equal(WHITE))
	require.EqualValues(t, 0.1, m.ambient)
	require.EqualValues(t, 0.9, m.diffuse)
	require.EqualValues(t, 0.9, m.specular)
	require.EqualValues(t, 200., m.shininess)
}

func TestMaterialArgumentsMustBePositive(t *testing.T) {
	require.Panics(t, func() { NewMaterial(WHITE, -2, 0, 0, 0) })
	require.Panics(t, func() { NewMaterial(WHITE, 0, -2, 0, 0) })
	require.Panics(t, func() { NewMaterial(WHITE, 0, 0, -2, 0) })
	require.Panics(t, func() { NewMaterial(WHITE, 0, 0, 0, -2) })
}

func TestCreatingSphereWithMaterial(t *testing.T) {
	m := NewMaterial(RED, 3, 4, 5, 6)

	s := NewSphere("sphere_id", m)

	require.Equal(t, m, s.material)
}

func TestLightingWithEyeBetweenLightAndSurface(t *testing.T) {
	m, pos := NewMaterial(WHITE, 0.1, 0.9, 0.9, 200.), NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), WHITE)

	i := 0.1 + 0.9 + 0.9
	expect := NewColor(i, i, i)
	res := CalcLighting(m, light, pos, eye, normal)

	require.True(t, res.Equal(expect))
}

func TestLightingWithEyeOffset45DegreesBetweenLightAndSurface(t *testing.T) {
	m, pos := NewMaterial(WHITE, 0.1, 0.9, 0.9, 200.), NewPoint(0, 0, 0)
	eye := NewVector(0, COS45, COS45)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), WHITE)

	i := 0.1 + 0.9 + 0
	expect := NewColor(i, i, i)
	res := CalcLighting(m, light, pos, eye, normal)

	require.True(t, res.Equal(expect))
}

func TestLightingWithEyeOppositeSurfaceAndLightOffset45Degrees(t *testing.T) {
	m, pos := NewMaterial(WHITE, 0.1, 0.9, 0.9, 200.), NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), WHITE)

	i := 0.1 + 0.9*COS45 + 0
	expect := NewColor(i, i, i)
	res := CalcLighting(m, light, pos, eye, normal)

	require.True(t, res.Equal(expect))
}

func TestLightingWithEyeInThePathOfReflectionVector(t *testing.T) {
	m, pos := NewMaterial(WHITE, 0.1, 0.9, 0.9, 200.), NewPoint(0, 0, 0)
	eye := NewVector(0, -COS45, -COS45)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), WHITE)

	i := 0.1 + 0.9*COS45 + 0.9
	expect := NewColor(i, i, i)
	res := CalcLighting(m, light, pos, eye, normal)

	require.True(t, res.Equal(expect))
}

func TestLightingWithLightBehindTheSurface(t *testing.T) {
	m, pos := NewMaterial(WHITE, 0.1, 0.9, 0.9, 200.), NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, 10), WHITE)

	i := 0.1 + 0 + 0
	expect := NewColor(i, i, i)
	res := CalcLighting(m, light, pos, eye, normal)

	require.True(t, res.Equal(expect))
}

// TODO: test spheres creation - 2 spheres with same id?
// TODO: use test fixtures to reduce code?
