package ray_tracer

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_newCanvas(t *testing.T) {
	c := NewCanvas(10, 20)

	require.Equal(t, c.width, 10)
	require.Equal(t, c.height, 20)

	BLACK_PIXEL := Color{0, 0, 0}
	for x := 0; x < c.width; x++ {
		for y := 0; y < c.height; y++ {
			require.True(t, c.PixelAt(2, 3).Equal(BLACK_PIXEL))
		}
	}
}

func TestWritingPixelsToACanvas(t *testing.T) {
	c := NewCanvas(10, 20)

	c.WritePixel(2, 3, RED)

	require.True(t, c.PixelAt(2, 3).Equal(RED))
}

func TestWritingPixelsOutOfCanvasBoundsHasNoEffectOnCanvas(t *testing.T) {
	c := NewCanvas(100, 100)

	require.NotPanics(t, func() { c.WritePixel(101, 99, GREEN) })
	require.NotPanics(t, func() { c.WritePixel(99, 101, GREEN) })

	// TODO: test, that canvas are left intact. For that, a canvas.DeepCopy() function should implemented first
}

func TestGettingPixelValueAtOutOfBoundCanvasCooridnatePanics(t *testing.T) {
	c := NewCanvas(2, 2)

	require.Panics(t, func() { c.PixelAt(3, 2) })
	require.Panics(t, func() { c.PixelAt(2, 3) })
}

func TestCorrectPpmHeader(t *testing.T) {
	c := NewCanvas(5, 3)

	ppmHeader := c.ppmHeader()
	expect := `P3
5 3
255
`

	require.Equal(t, ppmHeader, expect)
}

func TestPpmPixelData(t *testing.T) {
	canvas := NewCanvas(5, 3)

	c1 := Color{1.5, 0, 0}
	c2 := Color{0, 0.5, 0}
	c3 := Color{-0.5, 0, 1}
	canvas.WritePixel(0, 0, c1)
	canvas.WritePixel(2, 1, c2)
	canvas.WritePixel(4, 2, c3)

	pixelData := canvas.ppmPixelData()
	expect := `255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`

	require.Equal(t, pixelData, expect)
}

func TestSplittingLongLinesInPpmData(t *testing.T) {
	canvas := NewCanvas(10, 2)
	c := Color{1, 0.8, 0.6}
	canvas.Fill(c)

	pixelData := canvas.ppmPixelData()
	expect := `255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
`

	require.Equal(t, expect, pixelData)
}

func TestPpmFilesTerminatedByNewLineCharacter(t *testing.T) {
	canvas := NewCanvas(6, 7)
	ppmData := canvas.PpmData()

	require.True(t, strings.HasSuffix(ppmData, "\n"))
}

func TestCreatingAndSavingLargePpmFileShouldBeFast(t *testing.T) {
	testFilename := "testLargeCanvas.ppm"
	timeout := time.After(2 * time.Second)
	done := make(chan bool)
	go func() {
		canvas := NewCanvas(500, 500)
		canvas.SavePpm(testFilename)

		done <- true
	}()

	select {
	case <-timeout:
		os.Remove(testFilename)
		t.Fatal("Test didn't finish in time!")

	case <-done:
		require.FileExists(t, testFilename)
		os.Remove(testFilename)
	}
}
