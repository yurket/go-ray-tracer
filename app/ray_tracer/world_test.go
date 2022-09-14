package ray_tracer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatingEmptyWorld(t *testing.T) {
	w := NewWorld()

	require.Len(t, w, 0)
}

func TestDefaultWorldContainsPointLightAndTwoSpheres(t *testing.T) {
	w := NewDefaultWorld()

	obj1, ok := w["s1"]
	require.True(t, ok)
	s1, ok := obj1.(*Sphere)
	require.True(t, ok)
	require.True(t, s1.material.color.Equal(NewColor(0.8, 1, 0.6)))
	require.EqualValues(t, s1.material.ambient, 0.1)
	require.EqualValues(t, s1.material.diffuse, 0.7)
	require.EqualValues(t, s1.material.specular, 0.2)
	require.EqualValues(t, s1.material.shininess, 200)

	obj2, ok := w["s2"]
	require.True(t, ok)
	s2, ok := obj2.(*Sphere)
	require.True(t, ok)
	require.True(t, s2.material.color.Equal(WHITE))

	obj3, ok := w["light"]
	require.True(t, ok)
	light, ok := obj3.(*PointLight)
	require.True(t, ok)
	require.True(t, light.intensity.Equal(WHITE))
}

func TestObjectInWorldCanBeChangedThroughtTheReference(t *testing.T) {
	w := NewDefaultWorld()
	newId := "new_test_id"

	obj1 := w["s1"]
	s1 := obj1.(*Sphere)
	s1.id = newId

	sameObj := w["s1"]
	sameS1 := sameObj.(*Sphere)

	require.Equal(t, newId, sameS1.id)
}

func TestIntersectionsWithWorldReturnedInAscendingOrder(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	xs := w.IntersectWith(&r)

	require.Len(t, xs, 4)
	require.EqualValues(t, xs[0].time, 4)
	require.EqualValues(t, xs[1].time, 4.5)
	require.EqualValues(t, xs[2].time, 5.5)
	require.EqualValues(t, xs[3].time, 6)
}

func TestTheColorWhenRayMissesIsBlack(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))

	expect := BLACK
	res := w.ColorAtIntersection(r)

	require.True(t, expect.Equal(res))
}

func TestTheColorWhenRayHitsTheOuterSphere(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	expect := NewColor(0.38066, 0.47583, 0.2855)
	res := w.ColorAtIntersection(r)

	require.True(t, expect.Equal(res))
}

func TestTheColorWhenRayHitsInnerSphere(t *testing.T) {
	w := NewDefaultWorld()
	outer := w.Sphere("s1")
	outer.material.ambient = 1
	inner := w.Sphere("s2")
	inner.material.ambient = 1
	r := NewRay(NewPoint(0, 0, 0.75), NewVector(0, 0, -1))

	expect := inner.material.color
	res := w.ColorAtIntersection(r)

	require.True(t, expect.Equal(res))
}
