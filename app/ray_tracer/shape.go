package ray_tracer

type Shape struct {
	transform Matrix
	material  Material
}

func NewDefaultShape() Shape {
	return Shape{
		transform: *NewIdentityMatrix(4),
		material:  NewDefaultMaterial(),
	}
}

func (s *Shape) Transform() Matrix {
	return s.transform
}

func (s *Shape) SetTransform(m *Matrix) {
	s.transform = *m
}
