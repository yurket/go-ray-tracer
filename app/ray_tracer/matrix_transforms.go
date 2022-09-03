package ray_tracer

import "math"

func NewTranslationMatrix(x, y, z float64) *Matrix {
	translation := NewIdentityMatrix(4)
	translation.data[0][3] = x
	translation.data[1][3] = y
	translation.data[2][3] = z

	return translation
}

func NewScalingMatrix(x, y, z float64) *Matrix {
	scaling := NewIdentityMatrix(4)
	scaling.data[0][0] = x
	scaling.data[1][1] = y
	scaling.data[2][2] = z

	return scaling
}

// Left-handed coordinate system is used in all rotations
func NewRotationXMatrix(angleInRads float64) *Matrix {
	rot := NewIdentityMatrix(4)
	rot.data[1][1] = math.Cos(angleInRads)
	rot.data[1][2] = -math.Sin(angleInRads)
	rot.data[2][1] = math.Sin(angleInRads)
	rot.data[2][2] = math.Cos(angleInRads)
	return rot
}

func NewRotationYMatrix(angleInRads float64) *Matrix {
	rot := NewIdentityMatrix(4)
	rot.data[0][0] = math.Cos(angleInRads)
	rot.data[0][2] = math.Sin(angleInRads)
	rot.data[2][0] = -math.Sin(angleInRads)
	rot.data[2][2] = math.Cos(angleInRads)
	return rot
}

func NewRotationZMatrix(angleInRads float64) *Matrix {
	rot := NewIdentityMatrix(4)
	rot.data[0][0] = math.Cos(angleInRads)
	rot.data[0][1] = -math.Sin(angleInRads)
	rot.data[1][0] = math.Sin(angleInRads)
	rot.data[1][1] = math.Cos(angleInRads)
	return rot
}

func NewShearingMatrix(xToY, xToZ, yToX, yToZ, zToX, zToY float64) *Matrix {
	shearing := NewIdentityMatrix(4)
	shearing.data[0][1] = xToY
	shearing.data[0][2] = xToZ
	shearing.data[1][0] = yToX
	shearing.data[1][2] = yToZ
	shearing.data[2][0] = zToX
	shearing.data[2][1] = zToY
	return shearing
}

func (a *Matrix) Translate(x, y, z float64) *Matrix {
	return NewTranslationMatrix(x, y, z).MulMat(a)
}

func (a *Matrix) Scale(x, y, z float64) *Matrix {
	return NewScalingMatrix(x, y, z).MulMat(a)
}

func (a *Matrix) RotateX(angleInRads float64) *Matrix {
	return NewRotationXMatrix(angleInRads).MulMat(a)
}

func (a *Matrix) RotateY(angleInRads float64) *Matrix {
	return NewRotationYMatrix(angleInRads).MulMat(a)
}
func (a *Matrix) RotateZ(angleInRads float64) *Matrix {
	return NewRotationZMatrix(angleInRads).MulMat(a)
}

func (a *Matrix) Shear(xToY, xToZ, yToX, yToZ, zToX, zToY float64) *Matrix {
	return NewShearingMatrix(xToY, xToZ, yToX, yToZ, zToX, zToY).MulMat(a)
}
