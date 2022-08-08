package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCreateAndAccess4x4MatrixElements(t *testing.T) {
	m := newMatrix([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5}})

	require.EqualValues(t, m.At(0, 0), 1)
	require.EqualValues(t, m.At(1, 0), 5.5)
	require.EqualValues(t, m.At(2, 3), 12)
	require.EqualValues(t, m.At(3, 1), 14.5)
}

func TestCreatingMatricesWithArbitrarySizes(t *testing.T) {
	m1x3 := newMatrix([][]float64{{1, 1, 1}})
	rows, cols := m1x3.Shape()
	require.EqualValues(t, rows, 1)
	require.EqualValues(t, cols, 3)
	require.EqualValues(t, m1x3.At(0, 2), 1)

	m3x1 := newMatrix([][]float64{{2}, {2}, {2}})
	rows, cols = m3x1.Shape()
	require.EqualValues(t, rows, 3)
	require.EqualValues(t, cols, 1)
	require.EqualValues(t, m3x1.At(2, 0), 2)

	m2x2 := newMatrix([][]float64{{3, 3}, {3, 3}})
	rows, cols = m2x2.Shape()
	require.EqualValues(t, rows, 2)
	require.EqualValues(t, cols, 2)
	require.EqualValues(t, m2x2.At(1, 1), 3)

	m3x3 := newMatrix([][]float64{{4, 4, 4}, {4, 4, 4}, {4, 4, 4}})
	rows, cols = m3x3.Shape()
	require.EqualValues(t, rows, 3)
	require.EqualValues(t, cols, 3)
	require.EqualValues(t, m3x3.At(0, 2), 4)
}

func TestCreatingZeroMatrix(t *testing.T) {
	m := newZeroMatrix(2, 4)

	require.EqualValues(t, m.At(0, 0), 0)
	require.EqualValues(t, m.At(1, 3), 0)
}

func TestCreatingIdentityMatrix(t *testing.T) {
	m := newIdentityMatrix(4)

	require.EqualValues(t, m.At(0, 0), 1)
	require.EqualValues(t, m.At(0, 1), 0)
	require.EqualValues(t, m.At(1, 3), 0)
	require.EqualValues(t, m.At(1, 1), 1)
	require.EqualValues(t, m.At(3, 3), 1)
}

func TestComparingMatricesWithDifferentShapes(t *testing.T) {
	m1x1 := newMatrix([][]float64{{1}})
	m1x2 := newMatrix([][]float64{{1, 2}})

	require.False(t, m1x1.Equal(m1x2))
}

func TestMatrixComparisonIsCommutative(t *testing.T) {
	m1 := newMatrix([][]float64{{1, 2}, {3, 4}})
	m2 := newMatrix([][]float64{{1, 2}, {3, 4}})

	require.True(t, m1.Equal(m2))
	require.True(t, m2.Equal(m1))
}

func TestMatrixComparisonIsApproximate(t *testing.T) {
	m1 := newMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})
	m2 := newMatrix([][]float64{{1, 2, 3}, {3.999999999999, 5.0, 6.0000000000001}})

	require.True(t, m1.Equal(m2))
}

func TestComparingMatricesWithDifferentValues(t *testing.T) {
	m1 := newMatrix([][]float64{{1, 2, 3}, {4, 5, 6}})
	m2 := newMatrix([][]float64{{1, 2, 3}, {0, 5, 6}})

	require.False(t, m1.Equal(m2))
}

func TestMatrixMultiplication(t *testing.T) {
	a := newMatrix([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	b := newMatrix([][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	})

	expect := newMatrix([][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	})

	res := a.Mul(b)
	require.EqualValues(t, res, expect)
}

func TestNonSquareMatrixMultiplication(t *testing.T) {
	a := newMatrix([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	b := newMatrix([][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	})

	expect := newMatrix([][]float64{
		{22, 28},
		{49, 64},
	})
	res := a.Mul(b)

	require.EqualValues(t, res, expect)
}

func TestMultiplicationOnNilMatrix(t *testing.T) {
	a := newMatrix([][]float64{{1, 2}})
	var b *Matrix = nil

	require.Panics(t, func() { a.Mul(b) })
}

func TestMatrixMultipliedByATupleGivesTuple(t *testing.T) {
	a := newMatrix([][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	})
	tup := newTuple(1, 2, 3, 1)

	expect := newTuple(18, 24, 33, 1)
	res := a.MulTuple(tup)

	require.True(t, res.Equal(expect))
}

func TestMatrixMultipliedByIdentityGivesTheSameMatrix(t *testing.T) {
	m := newMatrix([][]float64{
		{0, 1, 2, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 12},
		{3, 5, 7, 8},
	})
	I := newIdentityMatrix(4)

	res := m.Mul(I)

	require.True(t, m.Equal(res))
}

func TestIdentityMatrixMultipliedByTupleGivesTheSameTuple(t *testing.T) {
	a := newTuple(1, 2, 3, 4)
	I := newIdentityMatrix(4)

	res := I.MulTuple(a)

	require.True(t, res.Equal(a))
}

func TestTranposingSquareMatrix(t *testing.T) {
	m := newMatrix([][]float64{
		{0, 1, 2, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 12},
		{3, 5, 7, 8},
	})

	expect := newMatrix([][]float64{
		{0, 1, 2, 3},
		{1, 2, 4, 5},
		{2, 4, 8, 7},
		{4, 8, 12, 8},
	})
	res := m.Transpose()

	require.True(t, res.Equal(expect))
}

func TestTransposingNonSquareMatrix(t *testing.T) {
	m := newMatrix([][]float64{
		{1, 3, 5},
		{2, 4, 6},
	})

	exect := newMatrix([][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	})
	res := m.Transpose()

	require.True(t, res.Equal(exect))
}

func TestTransposingIdentityMatrixResultsInIdentityMatrix(t *testing.T) {
	I := newIdentityMatrix(4)

	require.True(t, I.Equal(I.Transpose()))
}

func TestComputingDeterminantOf2x2Matrix(t *testing.T) {
	m := newMatrix([][]float64{
		{1, 5},
		{-3, 2},
	})

	require.EqualValues(t, m.Det2(), 17)
}
