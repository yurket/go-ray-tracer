package ray_tracer

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChapter01(t *testing.T) {
	require.NotPanics(t, func() { Chapter01Projectile() })
}

func cleanup(filename string) {
	err := os.Remove(filename)
	if err != nil {
		panic(fmt.Sprintf("Can't remove file %q!", filename))
	}
}

func TestChapter02(t *testing.T) {
	filename := "chapter02_test.ppm"

	require.NotPanics(t, func() { Chapter02DrawProjectilePpm(filename) })

	require.FileExists(t, filename)

	cleanup(filename)
}

func TestChapter03(t *testing.T) {
	require.NotPanics(t, func() { Chapter03MatrixTransforms() })
}

func TestChapter04DrawsClock(t *testing.T) {
	filename := "chapter04_clock.ppm"
	require.NotPanics(t, func() { Chapter04DrawAnalogClock(filename) })

	require.FileExists(t, filename)

	cleanup(filename)
}

func TestChapter05(t *testing.T) {
	filename := "chapter05_sphere_projection.ppm"

	require.NotPanics(t, func() { Chapter05(filename) })
	require.FileExists(t, filename)

	cleanup(filename)
}

func TestChapter06LightAndShading(t *testing.T) {
	filename := "chapter06_lighted_sphere.ppm"

	require.NotPanics(t, func() { Chapter06LightAndShading(filename) })
	require.FileExists(t, filename)

	cleanup(filename)
}

func TestChapter07MakingAScene(t *testing.T) {
	filename := "chapter07_scene.ppm"

	require.NotPanics(t, func() { Chapter07MakingAScene(filename) })
	require.FileExists(t, filename)

	cleanup(filename)
}
