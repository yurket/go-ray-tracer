package ray_tracer

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiplyingByATranslationMatrixMovesPoint(t *testing.T) {
	translation := NewTranslationMatrix(5, -3, 2)
	p := NewPoint(-3, 4, 5)

	expect := NewPoint(2, 1, 7)
	res := translation.MulTuple(p)

	require.True(t, res.Equal(expect))
}

func TestMultiplyingByInverseTranslationMatrixMovesPointInReverse(t *testing.T) {
	translation := NewTranslationMatrix(5, -3, 2)
	p := NewPoint(-3, 4, 5)

	expect := NewPoint(-8, 7, 3)
	res := translation.Inverse().MulTuple(p)

	require.True(t, res.Equal(expect))
}

func TestMultiplyingTranslationMatrixByVectorDoesntEffectVector(t *testing.T) {
	translation := NewTranslationMatrix(5, -3, 2)
	v := NewVector(-3, 4, 5)

	require.True(t, translation.MulTuple(v).Equal(v))
}

func TestScalingAPoint(t *testing.T) {
	scaling := NewScalingMatrix(2, 3, 4)
	p := NewPoint(-4, 6, 8)

	expect := NewPoint(-8, 18, 32)
	res := scaling.MulTuple(p)

	require.True(t, res.Equal(expect))
}

func TestScalingAVector(t *testing.T) {
	scaling := NewScalingMatrix(2, 3, 4)
	v := NewVector(-4, 6, 8)

	expect := NewVector(-8, 18, 32)
	res := scaling.MulTuple(v)

	require.True(t, res.Equal(expect))
}

func TestMultiplyingByInversedScalingMatrix(t *testing.T) {
	scaling := NewScalingMatrix(2, 3, 4)
	v := NewVector(-4, 6, 8)

	expect := NewVector(-2, 2, 2)
	res := scaling.Inverse().MulTuple(v)

	require.True(t, res.Equal(expect))
}

func TestReflectionIsScalingByNegativeValue(t *testing.T) {
	scaling := NewScalingMatrix(-1, 1, 1)
	p := NewPoint(2, 3, 4)

	expect := NewPoint(-2, 3, 4)
	res := scaling.MulTuple(p)

	require.True(t, res.Equal(expect))
}

func TestRotatingPointAroundXAxis(t *testing.T) {
	p := NewPoint(0, 1, 0)
	halfQuarter := NewRotationXMatrix(math.Pi / 4)
	fullQuarter := NewRotationXMatrix(math.Pi / 2)

	expect := NewPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2)
	require.True(t, halfQuarter.MulTuple(p).Equal(expect))

	expect = NewPoint(0, 0, 1)
	require.True(t, fullQuarter.MulTuple(p).Equal(expect))
}

func TestInverseOfXRotationRotatesInOppositeDirection(t *testing.T) {
	p := NewPoint(0, 1, 0)

	fullInverseQuarter := NewRotationXMatrix(math.Pi / 2).Inverse()

	expect := NewPoint(0, 0, -1)
	require.True(t, fullInverseQuarter.MulTuple(p).Equal(expect))
}

func TestRotatingPointAroundYAxis(t *testing.T) {
	p := NewPoint(0, 0, 1)
	halfQuarter := NewRotationYMatrix(math.Pi / 4)
	fullQuarter := NewRotationYMatrix(math.Pi / 2)

	expect := NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)
	require.True(t, halfQuarter.MulTuple(p).Equal(expect))

	expect = NewPoint(1, 0, 0)
	require.True(t, fullQuarter.MulTuple(p).Equal(expect))
}

func TestRotatingPointAroundZAxis(t *testing.T) {
	p := NewPoint(0, 1, 0)
	halfQuarter := NewRotationZMatrix(math.Pi / 4)
	fullQuarter := NewRotationZMatrix(math.Pi / 2)

	expect := NewPoint(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)
	require.True(t, halfQuarter.MulTuple(p).Equal(expect))

	expect = NewPoint(-1, 0, 0)
	require.True(t, fullQuarter.MulTuple(p).Equal(expect))
}

func TestShearingXToY(t *testing.T) {
	p := NewPoint(2, 3, 4)
	transform := NewShearingMatrix(1, 0, 0, 0, 0, 0)

	expect := NewPoint(5, 3, 4)
	require.True(t, transform.MulTuple(p).Equal(expect))
}

func TestShearingXToZ(t *testing.T) {
	p := NewPoint(2, 3, 4)
	transform := NewShearingMatrix(0, 1, 0, 0, 0, 0)

	expect := NewPoint(6, 3, 4)
	require.True(t, transform.MulTuple(p).Equal(expect))
}

func TestShearingYToX(t *testing.T) {
	p := NewPoint(2, 3, 4)
	transform := NewShearingMatrix(0, 0, 1, 0, 0, 0)

	expect := NewPoint(2, 5, 4)
	require.True(t, transform.MulTuple(p).Equal(expect))
}

func TestShearingYToZ(t *testing.T) {
	p := NewPoint(2, 3, 4)
	transform := NewShearingMatrix(0, 0, 0, 1, 0, 0)

	expect := NewPoint(2, 7, 4)
	require.True(t, transform.MulTuple(p).Equal(expect))
}

func TestShearingZToX(t *testing.T) {
	p := NewPoint(2, 3, 4)
	transform := NewShearingMatrix(0, 0, 0, 0, 1, 0)

	expect := NewPoint(2, 3, 6)
	require.True(t, transform.MulTuple(p).Equal(expect))
}

func TestShearingZToY(t *testing.T) {
	p := NewPoint(2, 3, 4)
	transform := NewShearingMatrix(0, 0, 0, 0, 0, 1)

	expect := NewPoint(2, 3, 7)
	require.True(t, transform.MulTuple(p).Equal(expect))
}

func TestIndividualVsChainedTransformationsOrder(t *testing.T) {
	p := NewPoint(1, 0, 1)
	rot := NewRotationXMatrix(math.Pi / 2)
	scale := NewScalingMatrix(5, 5, 5)
	translate := NewTranslationMatrix(10, 5, 7)

	// Individual transformations should be applied in sequence
	pSeq := rot.MulTuple(p)
	pSeq = scale.MulTuple(pSeq)
	pSeq = translate.MulTuple(pSeq)

	// "Chained" transformations must be applied in reverse order
	pChained := translate.MulMat(scale.MulMat(rot)).MulTuple(p)

	require.True(t, pChained.Equal(pSeq))

	// The same but more prettier
	pChained2 := NewIdentityMatrix(4).RotateX(math.Pi/2).Scale(5, 5, 5).Translate(10, 5, 7).MulTuple(p)
	require.True(t, pChained2.Equal(pChained))
}

// Transformations transform all the space relative to the center of the coorditates, not relative
// to the center of the object. So when scaling occurs after the translation, it scales the coordinates
// of object as well, basically moving the center of the object.
func TestTransformationOrderMattersWhenTranslationInvolved(t *testing.T) {
	t1 := NewIdentityMatrix(4).Scale(5, 5, 5).Translate(1, 2, 3)
	t2 := NewIdentityMatrix(4).Translate(1, 2, 3).Scale(5, 5, 5)

	require.False(t, t1.Equal(t2))
}

func TestScaleAndRotationOrderDoesntMatter(t *testing.T) {
	t1 := NewIdentityMatrix(4).Scale(5, 5, 5).RotateX(1)
	t2 := NewIdentityMatrix(4).RotateX(1).Scale(5, 5, 5)

	require.True(t, t1.Equal(t2))
}

func TestViewTransformationForDefaultOrientationIsIdentity(t *testing.T) {
	from, to, up := NewPoint(0, 0, 0), NewPoint(0, 0, -1), NewVector(0, 1, 0)

	viewTransform := NewViewTranformation(from, to, up)

	require.True(t, viewTransform.Equal(NewIdentityMatrix(4)))
}

func TestViewTransformWhenLookingBack(t *testing.T) {
	from, to, up := NewPoint(0, 0, 0), NewPoint(0, 0, +1), NewVector(0, 1, 0)

	expect := NewIdentityMatrix(4).Scale(-1, 1, -1)
	viewTransform := NewViewTranformation(from, to, up)

	require.True(t, expect.Equal(viewTransform))
}

func TestViewTransformationMovesTheWorldAndNotTheEye(t *testing.T) {
	from, to, up := NewPoint(0, 0, 8), NewPoint(0, 0, 0), NewVector(0, 1, 0)

	// the whole world is moved 8 units away from the eye positioned at the origin
	expect := NewTranslationMatrix(0, 0, -8)
	viewTransform := NewViewTranformation(from, to, up)

	require.True(t, expect.Equal(viewTransform))
}

func TestArbitraryViewTransform(t *testing.T) {
	from, to, up := NewPoint(1, 3, 2), NewPoint(4, -2, 8), NewVector(1, 1, 0)

	expect := NewMatrix([][]float64{
		{-0.50709, 0.50709, 0.67612, -2.36643},
		{0.76772, 0.60609, 0.12122, -2.82843},
		{-0.35857, 0.59761, -0.71714, 0.00000},
		{0, 0, 0, 1},
	})
	viewTransform := NewViewTranformation(from, to, up)

	require.True(t, expect.Equal(viewTransform))
}
