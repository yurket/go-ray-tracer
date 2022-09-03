package ray_tracer

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatingSphereWithMaterial(t *testing.T) {
	m := NewMaterial(RED, 3, 4, 5, 6)

	s := NewSphere("sphere_id", m)

	require.Equal(t, m, s.material)
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
