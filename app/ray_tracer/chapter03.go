package ray_tracer

import "fmt"

func Chapter03MatrixTransforms() {
	I := NewIdentityMatrix(3)
	fmt.Println("1. Inverse of Identity matrix is Identity matrix:")
	fmt.Printf("Before inversion: \n%s\nAfter inversion: \n%s\n", I, I.Inverse())

	A := NewMatrix([][]float64{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	})
	shouldBeI := A.MulMat(A.Inverse())
	fmt.Println("2. A * A_inv should be equal Identity matrix: ")
	fmt.Printf("A: \n%s\n A_inv: \n%s\n A * A_inv: \n%s\n", A, A.Inverse(), shouldBeI)

	fmt.Printf("3. Transpose of the inverse: \n%s\n Inverse of the transpose: \n%s\n", A.Inverse().Transpose(), A.Transpose().Inverse())

	changedI := NewIdentityMatrix(4)
	changedI.data[0][1] = 3.3
	tup := NewTuple(1, 2, 3, 4)
	changedTup := changedI.MulTuple(tup)
	fmt.Println("4. Multiplying by not Identity results in chaged tuple:")
	fmt.Printf("changedI: \n%s\n tuple: \n%v\n changedI * tuple = \n%v\n",
		changedI, tup, changedTup)
}
