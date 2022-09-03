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
