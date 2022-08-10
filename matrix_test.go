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

	require.EqualValues(t, 1, m.At(0, 0))
	require.EqualValues(t, 5.5, m.At(1, 0))
	require.EqualValues(t, 12, m.At(2, 3))
	require.EqualValues(t, 14.5, m.At(3, 1))
}

func TestCreatingMatricesWithArbitrarySizes(t *testing.T) {
	m1x3 := newMatrix([][]float64{{1, 1, 1}})
	rows, cols := m1x3.Shape()
	require.EqualValues(t, 1, rows)
	require.EqualValues(t, 3, cols)
	require.EqualValues(t, 1, m1x3.At(0, 2))

	m3x1 := newMatrix([][]float64{{2}, {2}, {2}})
	rows, cols = m3x1.Shape()
	require.EqualValues(t, 3, rows)
	require.EqualValues(t, 1, cols)
	require.EqualValues(t, 2, m3x1.At(2, 0))

	m2x2 := newMatrix([][]float64{{3, 3}, {3, 3}})
	rows, cols = m2x2.Shape()
	require.EqualValues(t, 2, rows)
	require.EqualValues(t, 2, cols)
	require.EqualValues(t, 3, m2x2.At(1, 1))

	m3x3 := newMatrix([][]float64{{4, 4, 4}, {4, 4, 4}, {4, 4, 4}})
	rows, cols = m3x3.Shape()
	require.EqualValues(t, 3, rows)
	require.EqualValues(t, 3, cols)
	require.EqualValues(t, 4, m3x3.At(0, 2))
}

func TestCreatingZeroMatrix(t *testing.T) {
	m := newZeroMatrix(2, 4)

	require.EqualValues(t, 0, m.At(0, 0))
	require.EqualValues(t, 0, m.At(1, 3))
}

func TestCreatingIdentityMatrix(t *testing.T) {
	m := newIdentityMatrix(4)

	require.EqualValues(t, 1, m.At(0, 0))
	require.EqualValues(t, 0, m.At(0, 1))
	require.EqualValues(t, 0, m.At(1, 3))
	require.EqualValues(t, 1, m.At(1, 1))
	require.EqualValues(t, 1, m.At(3, 3))
}

func TestCopyingMatrix(t *testing.T) {
	m := newMatrix([][]float64{{1, 2}, {3, 4}})

	mCopy := m.Copy()

	require.True(t, m.Equal(mCopy))
}

func TestCopyingMatrixIsDeepCopy(t *testing.T) {
	m := newMatrix([][]float64{{1, 2}, {3, 4}})
	mCopy := m.Copy()

	mCopy.data[1][1] = 999

	require.False(t, m.Equal(mCopy))
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

	res := a.MulMat(b)
	require.EqualValues(t, expect, res)
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
	res := a.MulMat(b)

	require.EqualValues(t, expect, res)
}

func TestMultiplicationOnNilMatrix(t *testing.T) {
	a := newMatrix([][]float64{{1, 2}})
	var b *Matrix = nil

	require.Panics(t, func() { a.MulMat(b) })
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

	res := m.MulMat(I)

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

	require.EqualValues(t, 17, m.Determinant())
}

func TestGettingSubmatrixOf3x3Matrix(t *testing.T) {
	m := newMatrix([][]float64{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	})

	expect := newMatrix([][]float64{
		{-3, 2},
		{0, 6},
	})
	res := m.Submatrix(0, 2)

	require.True(t, res.Equal(expect))
}

func TestGettingSubmatrixOf4x4Matrix(t *testing.T) {
	m := newMatrix([][]float64{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	})

	expect := newMatrix([][]float64{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	})
	res := m.Submatrix(2, 1)

	require.True(t, res.Equal(expect))
}

func Test3x3MatrixMinor(t *testing.T) {
	m := newMatrix([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})

	require.EqualValues(t, 25, m.Minor(1, 0))
}

func Test3x3MatrixCofactor(t *testing.T) {
	m := newMatrix([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})

	require.EqualValues(t, -12, m.Minor(0, 0))
	require.EqualValues(t, -12, m.Cofactor(0, 0))

	require.EqualValues(t, 25, m.Minor(1, 0))
	require.EqualValues(t, -25, m.Cofactor(1, 0))
}

func TestComputingDeterminantOf3x3Matrix(t *testing.T) {
	m := newMatrix([][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	})

	require.EqualValues(t, 56, m.Cofactor(0, 0))
	require.EqualValues(t, 12, m.Cofactor(0, 1))
	require.EqualValues(t, -46, m.Cofactor(0, 2))
	require.EqualValues(t, -196, m.Determinant())
}

func TestComputingDeterminantOf4x4Matrix(t *testing.T) {
	m := newMatrix([][]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	})

	require.EqualValues(t, 690, m.Cofactor(0, 0))
	require.EqualValues(t, 447, m.Cofactor(0, 1))
	require.EqualValues(t, 210, m.Cofactor(0, 2))
	require.EqualValues(t, 51, m.Cofactor(0, 3))

	require.EqualValues(t, -4071, m.Determinant())
}

func TestIsMatrixInvertible(t *testing.T) {
	m1 := newMatrix([][]float64{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	})
	require.True(t, m1.IsInvertible())

	m2 := newMatrix([][]float64{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{-2, 1, -1, -1.5},
	})
	require.False(t, m2.IsInvertible())
}

func TestMatrixInverse(t *testing.T) {
	m := newMatrix([][]float64{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	})

	mInv := m.Inverse()
	expect := newMatrix([][]float64{
		{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639},
	})

	require.True(t, mInv.Equal(expect))

	// Additional tests from the book
	require.EqualValues(t, 532, m.Determinant())
	require.EqualValues(t, -160, m.Cofactor(2, 3))
	require.EqualValues(t, -160.0/532, mInv.At(3, 2))
}

func TestMatrixInverse2(t *testing.T) {
	m := newMatrix([][]float64{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	})

	mInv := m.Inverse()
	expect := newMatrix([][]float64{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308},
	})

	require.True(t, mInv.Equal(expect))
}

func TestMatrixInverse3(t *testing.T) {
	m := newMatrix([][]float64{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	})

	mInv := m.Inverse()
	expect := newMatrix([][]float64{
		{-0.04074, -0.07778, 0.14444, -0.22222},
		{-0.07778, 0.03333, 0.36667, -0.33333},
		{-0.02901, -0.14630, -0.10926, 0.12963},
		{0.17778, 0.06667, -0.26667, 0.33333},
	})

	require.True(t, mInv.Equal(expect))
}

func TestMatrixInverseProperty(t *testing.T) {
	a := newMatrix([][]float64{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	})

	b := newMatrix([][]float64{
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5},
	})

	c := a.MulMat(b)
	shouldBeA := c.MulMat(b.Inverse())

	require.True(t, shouldBeA.Equal(a))
}
