package ray_tracer

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

var COS45 = math.Sqrt(2) / 2.0

func TestCreatingPointLight(t *testing.T) {
	pos := NewPoint(0, 0, 0)
	intensity := WHITE

	pl := NewPointLight(pos, intensity)

	require.True(t, intensity.Equal(pl.intensity))
	require.True(t, pos.Equal(pl.position))
}

func TestLightingWithEyeBetweenLightAndSurface(t *testing.T) {
	m, pos := NewMaterial(WHITE, 0.1, 0.9, 0.9, 200.), NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), WHITE)

	i := 0.1 + 0.9 + 0.9
	expect := NewColor(i, i, i)
	res := CalcLighting(m, light, pos, eye, normal, false)

	require.True(t, res.Equal(expect))
}

func TestLightingWithEyeOffset45DegreesBetweenLightAndSurface(t *testing.T) {
	m, pos := NewMaterial(WHITE, 0.1, 0.9, 0.9, 200.), NewPoint(0, 0, 0)
	eye := NewVector(0, COS45, COS45)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), WHITE)

	i := 0.1 + 0.9 + 0
	expect := NewColor(i, i, i)
	res := CalcLighting(m, light, pos, eye, normal, false)

	require.True(t, res.Equal(expect))
}

func TestLightingWithEyeOppositeSurfaceAndLightOffset45Degrees(t *testing.T) {
	m, pos := NewMaterial(WHITE, 0.1, 0.9, 0.9, 200.), NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), WHITE)

	i := 0.1 + 0.9*COS45 + 0
	expect := NewColor(i, i, i)
	res := CalcLighting(m, light, pos, eye, normal, false)

	require.True(t, res.Equal(expect))
}

func TestLightingWithEyeInThePathOfReflectionVector(t *testing.T) {
	m, pos := NewMaterial(WHITE, 0.1, 0.9, 0.9, 200.), NewPoint(0, 0, 0)
	eye := NewVector(0, -COS45, -COS45)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), WHITE)

	i := 0.1 + 0.9*COS45 + 0.9
	expect := NewColor(i, i, i)
	res := CalcLighting(m, light, pos, eye, normal, false)

	require.True(t, res.Equal(expect))
}

func TestLightingWithLightBehindTheSurface(t *testing.T) {
	m, pos := NewMaterial(WHITE, 0.1, 0.9, 0.9, 200.), NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, 10), WHITE)

	i := 0.1 + 0 + 0
	expect := NewColor(i, i, i)
	res := CalcLighting(m, light, pos, eye, normal, false)

	require.True(t, res.Equal(expect))
}

func TestShadingAnIntersectionFromTheOutside(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s1 := w.Sphere("s1")
	i := NewIntersection(4, s1)
	comps := PrepareIntersectionComputations(i, r)

	expect := NewColor(0.38066, 0.47583, 0.2855)
	res := ShadeHit(w, &comps)

	require.True(t, expect.Equal(res))
}

func TestShadingAnIntersectionFromTheInside(t *testing.T) {
	w := NewDefaultWorld()
	w.SetLight(NewPointLight(NewPoint(0, 0.25, 0), WHITE))
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))

	s2 := w.Sphere("s2")
	i := NewIntersection(0.5, s2)
	comps := PrepareIntersectionComputations(i, r)

	expect := NewColor(0.90498, 0.90498, 0.90498)
	res := ShadeHit(w, &comps)

	require.True(t, expect.Equal(res))
}

func TestOneSphereShadowingPointOfIntersectionWithOtherSphere(t *testing.T) {
	world := NewWorld()
	pl := NewPointLight(NewPoint(0, 0, -10), WHITE)
	world.SetLight(pl)

	default_material := NewDefaultMaterial()
	s1 := NewSphere("s1", default_material)
	world["s1"] = &s1

	s2 := NewSphere("s2", default_material)
	s2.SetTransform(NewTranslationMatrix(0, 0, 10))
	world["s2"] = &s2

	r := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
	unit_radius := 1.
	distance_to_s2 := (10. - unit_radius) - r.origin.z
	i := NewIntersection(distance_to_s2, &s2)
	comps := PrepareIntersectionComputations(i, r)
	color_at_intersection := ShadeHit(world, &comps)
	expected_ambient_color := NewColor(0.1, 0.1, 0.1)

	require.Equal(t, expected_ambient_color, color_at_intersection)
}

func TestLightingWithTheSurfaceInShadow(t *testing.T) {
	ambientColor := 0.1
	m, pos := NewMaterial(WHITE, ambientColor, 0.9, 0.9, 200.), NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), WHITE)
	inShadow := true

	expect := NewColor(ambientColor, ambientColor, ambientColor)
	res := CalcLighting(m, light, pos, eye, normal, inShadow)

	require.True(t, expect.Equal(res))
}

// For all shadow tests light is at (-10, 10, -10)
func TestPointIsNotShadowedAndNotCollinear(t *testing.T) {
	w := NewDefaultWorld()

	p := NewPoint(0, 10, 0)

	require.False(t, IsShadowed(w, p))
}

func TestPointIsShadowedBySphere(t *testing.T) {
	w := NewDefaultWorld()
	p := NewPoint(10, -10, 10)

	require.True(t, IsShadowed(w, p))
}

func TestPointIsNotShadowedAndBehindTheLight(t *testing.T) {
	w := NewDefaultWorld()
	p := NewPoint(-20, 20, -20)

	require.False(t, IsShadowed(w, p))
}

func TestPointIsNotShadowedAndBetweenTheLightAndSphere(t *testing.T) {
	w := NewDefaultWorld()
	p := NewPoint(-2, 2, -2)

	require.False(t, IsShadowed(w, p))
}
