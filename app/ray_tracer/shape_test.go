package ray_tracer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSphapeDefaultTransofrmationIsIdentity(t *testing.T) {
	var s Shape = NewDefaultShape()

	require.True(t, s.transform.Equal(NewIdentityMatrix(4)))
}

func TestSettingShapeTransformation(t *testing.T) {
	var s Shape = NewDefaultShape()
	translation := NewTranslationMatrix(2, 3, 4)

	s.SetTransform(translation)

	transform := s.Transform()
	require.True(t, transform.Equal(translation))
}
