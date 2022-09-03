package ray_tracer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatingDefaultMaterial(t *testing.T) {
	m := NewDefaultMaterial()

	require.True(t, m.color.Equal(WHITE))
	require.EqualValues(t, 0.1, m.ambient)
	require.EqualValues(t, 0.9, m.diffuse)
	require.EqualValues(t, 0.9, m.specular)
	require.EqualValues(t, 200., m.shininess)
}

func TestMaterialArgumentsMustBePositive(t *testing.T) {
	require.Panics(t, func() { NewMaterial(WHITE, -2, 0, 0, 0) })
	require.Panics(t, func() { NewMaterial(WHITE, 0, -2, 0, 0) })
	require.Panics(t, func() { NewMaterial(WHITE, 0, 0, -2, 0) })
	require.Panics(t, func() { NewMaterial(WHITE, 0, 0, 0, -2) })
}
