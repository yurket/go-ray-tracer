package ray_tracer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColorCreation(t *testing.T) {
	c := Color{-0.5, 0.4, 1.7}

	require.Equal(t, c.r, -0.5)
	require.Equal(t, c.b, 1.7)
}

func TestAddingColors(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}

	res := c1.Add(c2)
	expect := Color{1.6, 0.7, 1.0}

	require.True(t, res.Equal(expect))
}

func TestSubtractingColors(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}

	res := c1.Sub(c2)
	expect := Color{0.2, 0.5, 0.5}

	require.True(t, res.Equal(expect))
}

func TestScalarMultiplyingColors(t *testing.T) {
	c1 := Color{0.2, 0.3, 0.4}
	n := 2.0

	res := c1.MultScalar(n)
	expect := Color{0.4, 0.6, 0.8}

	require.True(t, res.Equal(expect))
}

func TestHadamarMultiplyingColors(t *testing.T) {
	c1 := Color{1, 0.2, 0.4}
	c2 := Color{0.9, 1, 0.1}

	res := c1.MultHadamar(c2)
	expect := Color{0.9, 0.2, 0.04}

	require.True(t, res.Equal(expect))

}

func TestColorToRgbComponentsConvertion(t *testing.T) {
	c1 := Color{1, 0.8, 0.6}
	colorWithNegativeNumber := Color{-0.5, 0, 0.5}
	colorWithOutOfBoundNumber := Color{0, 0, 3}

	require.Equal(t, c1.ToScaledRgbComponents(), [3]int{255, 204, 153})
	require.Equal(t, colorWithNegativeNumber.ToScaledRgbComponents(), [3]int{0, 0, 128})
	require.Equal(t, colorWithOutOfBoundNumber.ToScaledRgbComponents(), [3]int{0, 0, 255})
}
