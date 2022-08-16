package ray_tracer

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

func (m *Matrix) Copy() *Matrix {
	newData := make([][]float64, m.rows)
	for i := 0; i < m.rows; i++ {
		newData[i] = make([]float64, 0, m.columns)
		newData[i] = append(newData[i], m.data[i]...)
	}
	return &Matrix{m.rows, m.columns, newData}
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

func (a *Matrix) MulMat(b *Matrix) *Matrix {
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

func (a *Matrix) Mul(num float64) *Matrix {
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.columns; j++ {
			a.data[i][j] *= num
		}
	}
	return a
}

func (a *Matrix) Div(num float64) *Matrix {
	return a.Mul(1.0 / num)
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
	res := a.MulMat(columnVector)
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

func (a *Matrix) panicIfWrongDimensions() {
	if !a.IsSquare() {
		panic(fmt.Sprintf("Matrix [%d, %d] should be square!", a.rows, a.columns))
	}
	if a.rows > 4 {
		panic(fmt.Sprintf("Matrix [%d, %d] should not exceed dimension 4x4!", a.rows, a.columns))
	}
}

func (a *Matrix) Determinant() float64 {
	a.panicIfWrongDimensions()

	if a.rows == 2 {
		return a.data[0][0]*a.data[1][1] - a.data[0][1]*a.data[1][0]
	}

	sum := 0.0
	for j := 0; j < a.columns; j++ {
		sum += a.data[0][j] * a.Cofactor(0, j)
	}
	return sum
}

func removeColumn(slice []float64, s int) []float64 {
	return append(slice[:s], slice[s+1:]...)
}
func removeRow(slice [][]float64, s int) [][]float64 {
	return append(slice[:s], slice[s+1:]...)
}

func (a *Matrix) Submatrix(rowToDelete int, colToDelete int) *Matrix {
	a.panicIfWrongDimensions()

	submatrix := a.Copy()

	for i := 0; i < a.rows; i++ {
		submatrix.data[i] = removeColumn(submatrix.data[i], colToDelete)
	}
	submatrix.data = removeRow(submatrix.data, rowToDelete)

	submatrix.rows = a.rows - 1
	submatrix.columns = a.columns - 1
	return submatrix
}

func (a *Matrix) Minor(row int, col int) float64 {
	a.panicIfWrongDimensions()

	return a.Submatrix(row, col).Determinant()
}

func (a *Matrix) Cofactor(row int, col int) float64 {
	a.panicIfWrongDimensions()

	isNegationNeeded := (row+col)%2 == 1
	if isNegationNeeded {
		return -a.Minor(row, col)
	}
	return a.Minor(row, col)
}

func (a *Matrix) IsInvertible() bool {
	return a.Determinant() != 0
}

func (a *Matrix) Inverse() *Matrix {
	if !a.IsInvertible() {
		panic("Trying to invert non-invertible matrix!")
	}

	cofactors := newZeroMatrix(a.rows, a.columns)
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.columns; j++ {
			cofactors.data[i][j] = a.Cofactor(i, j)
		}
	}

	cofactors = cofactors.Transpose()
	return cofactors.Div(a.Determinant())
}

func (a *Matrix) ToString() string {
	res := ""
	for i := 0; i < a.rows; i++ {
		res += "|"
		for j := 0; j < a.columns; j++ {
			ending := ", "
			if j == a.columns-1 {
				ending = " "
			}
			res += fmt.Sprintf("%6.2f%s", a.data[i][j], ending)
		}
		res += "|\n"
	}
	return res
}
