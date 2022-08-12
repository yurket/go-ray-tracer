package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChapter01(t *testing.T) {
	require.NotPanics(t, func() { Chapter01Projectile() })
}

func TestChapter02(t *testing.T) {
	filename := "chapter02_test.ppm"

	require.NotPanics(t, func() { Chapter02DrawProjectilePpm(filename) })

	require.FileExists(t, filename)

	// test cleanup
	err := os.Remove(filename)
	if err != nil {
		panic("Can't remove file after testing Chapter02!")
	}
}

func TestChapter03(t *testing.T) {
	require.NotPanics(t, func() { Chapter03MatrixTransforms() })
}

func TestChapter04DrawsClock(t *testing.T) {
	filename := "chapter04_clock.ppm"
	require.NotPanics(t, func() { Chapter04DrawAnalogClock(filename) })

	require.FileExists(t, filename)

	// test cleanup
	err := os.Remove(filename)
	if err != nil {
		panic("Can't remove file after testing Chapter04!")
	}
}
