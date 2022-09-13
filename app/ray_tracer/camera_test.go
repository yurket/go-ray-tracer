package ray_tracer

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConstructingACamera(t *testing.T) {
	hsize, vsize, fieldOfView := 160, 120, math.Pi/2
	c := NewCamera(hsize, vsize, fieldOfView)

	require.EqualValues(t, c.hSize, hsize)
	require.EqualValues(t, c.vSize, vsize)
	require.EqualValues(t, c.fieldOfView, fieldOfView)
	require.True(t, c.transform.Equal(NewIdentityMatrix(4)))
}

func TestCalculatedPixelSizeForAHorizontalCanvas(t *testing.T) {
	c := NewCamera(200, 125, math.Pi/2)

	require.EqualValues(t, c.pixelSize, 0.01)
}

func TestCalculatedPixelSizeForAVerticalCanvas(t *testing.T) {
	c := NewCamera(125, 200, math.Pi/2)

	require.EqualValues(t, c.pixelSize, 0.01)
}

func TestConstructingARayThroughTheCenterOfTheCanvas(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := c.CastRayIntoPixel(100, 50)

	require.True(t, r.origin.Equal(NewPoint(0, 0, 0)))
	require.True(t, r.direction.Equal(NewVector(0, 0, -1)))
}

func TestConstructingARayThorughACornerOfTheCanvas(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := c.CastRayIntoPixel(0, 0)

	require.True(t, r.origin.Equal(NewPoint(0, 0, 0)))
	require.True(t, r.direction.Equal(NewVector(0.66519, 0.33259, -0.66851)))
}

func TestConstructingARayWhenTheCameraIsTransformed(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	transform := NewRotationYMatrix(math.Pi / 4).MulMat(NewTranslationMatrix(0, -2, 5))
	c.transform = *transform
	r := c.CastRayIntoPixel(100, 50)

	require.True(t, r.origin.Equal(NewPoint(0, 2, -5)))
	require.True(t, r.direction.Equal(NewVector(COS45, 0, -COS45)))
}

func TestRenderingAWorldWithACamera(t *testing.T) {
	w := NewDefaultWorld()
	c := NewCamera(11, 11, math.Pi/2)
	from, to, up := NewPoint(0, 0, -5), NewPoint(0, 0, 0), NewVector(0, 1, 0)
	c.transform = *NewViewTranformation(from, to, up)
	image := Render(c, w)

	expect := NewColor(0.38066, 0.47583, 0.2855)
	res := image.PixelAt(5, 5)

	require.True(t, expect.Equal(res))
}
