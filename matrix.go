package main

import "fmt"

type Matrix struct {
	rows    int
	columns int
	data    [][]float64
}

func newMatrix(data [][]float64) *Matrix {
	rows := len(data)
	if rows == 0 {
		panic("Invalid matrix with 0 rows!")
	}
	cols := len(data[0])
	if cols == 0 {
		panic("Invalid matrix with 0 columns!")
	}
	return &Matrix{rows, cols, data}
}

func newZeroMatrix(rows int, cols int) *Matrix {
	zeroes := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		zeroes[i] = make([]float64, cols)
	}
	return &Matrix{rows, cols, zeroes}
}

// Only square matrices may be identity
func newIdentityMatrix(rowsAndCols int) *Matrix {
	m := newZeroMatrix(rowsAndCols, rowsAndCols)
	for i := 0; i < rowsAndCols; i++ {
		m.data[i][i] = 1.0
	}
	return m
}

func (m *Matrix) At(row int, col int) float64 {
	return m.data[row][col]
}

func (m *Matrix) Shape() (int, int) {
	return m.rows, m.columns
}

func (a *Matrix) Equal(b *Matrix) bool {
	if (a.rows != b.rows) || (a.columns != b.columns) {
		return false
	}
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.columns; j++ {
			if !equal_fp(a.At(i, j), b.At(i, j)) {
				return false
			}
		}
	}
	return true
}

func (a *Matrix) Mul(b *Matrix) *Matrix {
	if a.rows == 0 || a.columns == 0 || b.rows == 0 || b.columns == 0 {
		panic("Matrix with one of the dimensions == 0!")
	}
	if a.columns != b.rows {
		msg := fmt.Sprintf("Can't multiply matrices with different number of columns and rows [%d != %d]!", a.columns, b.rows)
		panic(msg)
	}

	res := newZeroMatrix(a.rows, b.columns)
	for i := 0; i < a.rows; i++ {
		for j := 0; j < b.columns; j++ {
			sum := 0.0
			for k := 0; k < a.columns; k++ {
				sum += a.At(i, k) * b.At(k, j)
			}
			res.data[i][j] = sum
		}
	}
	return res
}

// Converst column-vector to a tuple
func (a *Matrix) ToTuple() Tuple {
	panicMsg := fmt.Sprintf("Can't convert matrix with dimensions [%d, %d] to a tuple!", a.rows, a.columns)
	if a.columns != 1 || (a.rows < 1 || a.rows > 4) {
		panic(panicMsg)
	}

	switch a.rows {
	case 1:
		return newTuple(a.At(0, 0), 0, 0, 0)
	case 2:
		return newTuple(a.At(0, 0), a.At(1, 0), 0, 0)
	case 3:
		return newTuple(a.At(0, 0), a.At(1, 0), a.At(2, 0), 0)
	case 4:
		return newTuple(a.At(0, 0), a.At(1, 0), a.At(2, 0), a.At(3, 0))
	default:
		panic(panicMsg)
	}
}

func (a *Matrix) IsSquare() bool {
	return a.rows == a.columns
}

func (a *Matrix) MulTuple(t Tuple) Tuple {
	columnVector := t.ToMatrix()
	res := a.Mul(columnVector)
	return res.ToTuple()
}

func (a *Matrix) Transpose() *Matrix {
	transposed := newZeroMatrix(a.columns, a.rows)
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.columns; j++ {
			transposed.data[j][i] = a.data[i][j]
		}
	}
	return transposed
}

func (a *Matrix) Det2() float64 {
	if !a.IsSquare() || a.rows > 2 {
		panic("Not implemented!")
	}
	return a.data[0][0]*a.data[1][1] - a.data[0][1]*a.data[1][0]
}
