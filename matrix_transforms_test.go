package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiplyingByATranslationMatrixMovesPoint(t *testing.T) {
	translation := newTranslationMatrix(5, -3, 2)
	p := Point(-3, 4, 5)

	expect := Point(2, 1, 7)
	res := translation.MulTuple(p)

	require.True(t, res.Equal(expect))
}

func TestMultiplyingByInverseTranslationMatrixMovesPointInReverse(t *testing.T) {
	translation := newTranslationMatrix(5, -3, 2)
	p := Point(-3, 4, 5)

	expect := Point(-8, 7, 3)
	res := translation.Inverse().MulTuple(p)

	require.True(t, res.Equal(expect))
}

func TestMultiplyingTranslationMatrixByVectorDoesntEffectVector(t *testing.T) {
	translation := newTranslationMatrix(5, -3, 2)
	v := Vector(-3, 4, 5)

	require.True(t, translation.MulTuple(v).Equal(v))
}

func TestScalingAPoint(t *testing.T) {
	scaling := newScalingMatrix(2, 3, 4)
	p := Point(-4, 6, 8)

	expect := Point(-8, 18, 32)
	res := scaling.MulTuple(p)

	require.True(t, res.Equal(expect))
}

func TestScalingAVector(t *testing.T) {
	scaling := newScalingMatrix(2, 3, 4)
	v := Vector(-4, 6, 8)

	expect := Vector(-8, 18, 32)
	res := scaling.MulTuple(v)

	require.True(t, res.Equal(expect))
}

func TestMultiplyingByInversedScalingMatrix(t *testing.T) {
	scaling := newScalingMatrix(2, 3, 4)
	v := Vector(-4, 6, 8)

	expect := Vector(-2, 2, 2)
	res := scaling.Inverse().MulTuple(v)

	require.True(t, res.Equal(expect))
}

func TestReflectionIsScalingByNegativeValue(t *testing.T) {
	scaling := newScalingMatrix(-1, 1, 1)
	p := Point(2, 3, 4)

	expect := Point(-2, 3, 4)
	res := scaling.MulTuple(p)

	require.True(t, res.Equal(expect))
}

func TestRotatingPointAroundXAxis(t *testing.T) {
	p := Point(0, 1, 0)
	halfQuarter := newRotationXMatrix(math.Pi / 4)
	fullQuarter := newRotationXMatrix(math.Pi / 2)

	expect := Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2)
	require.True(t, halfQuarter.MulTuple(p).Equal(expect))

	expect = Point(0, 0, 1)
	require.True(t, fullQuarter.MulTuple(p).Equal(expect))
}

func TestInverseOfXRotationRotatesInOppositeDirection(t *testing.T) {
	p := Point(0, 1, 0)

	fullInverseQuarter := newRotationXMatrix(math.Pi / 2).Inverse()

	expect := Point(0, 0, -1)
	require.True(t, fullInverseQuarter.MulTuple(p).Equal(expect))
}

func TestRotatingPointAroundYAxis(t *testing.T) {
	p := Point(0, 0, 1)
	halfQuarter := newRotationYMatrix(math.Pi / 4)
	fullQuarter := newRotationYMatrix(math.Pi / 2)

	expect := Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)
	require.True(t, halfQuarter.MulTuple(p).Equal(expect))

	expect = Point(1, 0, 0)
	require.True(t, fullQuarter.MulTuple(p).Equal(expect))
}

func TestRotatingPointAroundZAxis(t *testing.T) {
	p := Point(0, 1, 0)
	halfQuarter := newRotationZMatrix(math.Pi / 4)
	fullQuarter := newRotationZMatrix(math.Pi / 2)

	expect := Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)
	require.True(t, halfQuarter.MulTuple(p).Equal(expect))

	expect = Point(-1, 0, 0)
	require.True(t, fullQuarter.MulTuple(p).Equal(expect))
}

func TestShearingXToY(t *testing.T) {
	p := Point(2, 3, 4)
	transform := newShearingMatrix(1, 0, 0, 0, 0, 0)

	expect := Point(5, 3, 4)
	require.True(t, transform.MulTuple(p).Equal(expect))
}

func TestShearingXToZ(t *testing.T) {
	p := Point(2, 3, 4)
	transform := newShearingMatrix(0, 1, 0, 0, 0, 0)

	expect := Point(6, 3, 4)
	require.True(t, transform.MulTuple(p).Equal(expect))
}

func TestShearingYToX(t *testing.T) {
	p := Point(2, 3, 4)
	transform := newShearingMatrix(0, 0, 1, 0, 0, 0)

	expect := Point(2, 5, 4)
	require.True(t, transform.MulTuple(p).Equal(expect))
}

func TestShearingYToZ(t *testing.T) {
	p := Point(2, 3, 4)
	transform := newShearingMatrix(0, 0, 0, 1, 0, 0)

	expect := Point(2, 7, 4)
	require.True(t, transform.MulTuple(p).Equal(expect))
}

func TestShearingZToX(t *testing.T) {
	p := Point(2, 3, 4)
	transform := newShearingMatrix(0, 0, 0, 0, 1, 0)

	expect := Point(2, 3, 6)
	require.True(t, transform.MulTuple(p).Equal(expect))
}

func TestShearingZToY(t *testing.T) {
	p := Point(2, 3, 4)
	transform := newShearingMatrix(0, 0, 0, 0, 0, 1)

	expect := Point(2, 3, 7)
	require.True(t, transform.MulTuple(p).Equal(expect))
}

func TestIndividualVsChainedTransformationsOrder(t *testing.T) {
	p := Point(1, 0, 1)
	rot := newRotationXMatrix(math.Pi / 2)
	scale := newScalingMatrix(5, 5, 5)
	translate := newTranslationMatrix(10, 5, 7)

	// Individual transformations should be applied in sequence
	pSeq := rot.MulTuple(p)
	pSeq = scale.MulTuple(pSeq)
	pSeq = translate.MulTuple(pSeq)

	// "Chained" transformations must be applied in reverse order
	pChained := translate.MulMat(scale.MulMat(rot)).MulTuple(p)

	require.True(t, pChained.Equal(pSeq))

	// The same but more prettier
	pChained2 := newIdentityMatrix(4).RotateX(math.Pi/2).Scale(5, 5, 5).Translate(10, 5, 7).MulTuple(p)
	require.True(t, pChained2.Equal(pChained))
}

// TODO: rename Point, Vector -> newPoint, newVector
