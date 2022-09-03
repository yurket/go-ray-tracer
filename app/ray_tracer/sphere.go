package ray_tracer

// Unit sphere (radius == 1), with a center in (0,0,0)
type Sphere struct {
	id        string
	origin    Tuple
	transform Matrix
	material  Material
}

func NewSphere(id string, material Material) Sphere {
	return Sphere{
		id:        id,
		origin:    NewPoint(0, 0, 0),
		transform: *NewIdentityMatrix(4),
		material:  material,
	}
}

func NewDefaultSphere() Sphere {
	return Sphere{
		id:        "sphere_id",
		origin:    NewPoint(0, 0, 0),
		transform: *NewIdentityMatrix(4),
		material:  NewDefaultMaterial(),
	}
}

func (s *Sphere) Equal(s2 *Sphere) bool {
	return s.id == s2.id &&
		s.origin.Equal(s2.origin) &&
		s.transform.Equal(&s2.transform)
}

func (s *Sphere) Transform() Matrix {
	return s.transform
}

func (s *Sphere) SetTransform(m *Matrix) {
	s.transform = *m
}

func (s *Sphere) NormalAt(worldPoint Tuple) Tuple {
	sphereCenter := NewPoint(0, 0, 0)

	objectPoint := s.transform.Inverse().MulTuple(worldPoint)
	objectNormal := objectPoint.Sub(sphereCenter)
	// Don't understand why transpose here added
	worldNormal := s.transform.Inverse().Transpose().MulTuple(objectNormal)
	worldNormal.w = 0

	if !worldNormal.IsVector() {
		panic("Normal vector must be a vector!")
	}
	return worldNormal.Normalize()
}
