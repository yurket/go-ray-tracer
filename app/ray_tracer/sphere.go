package ray_tracer

import "math"

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

// Finds intersection of a ray going through the center of the sphere with a unit radius
func (s *Sphere) IntersectWith(r *Ray) []Intersection {
	// Inverse-transform the ray instead of transforming the sphere.
	// It makes the math easier.
	t := s.Transform()
	transformedRay := r.ApplyTransform(t.Inverse())

	sphereOrigin := NewPoint(0, 0, 0)
	sphereToRay := transformedRay.origin.Sub(sphereOrigin)
	a := transformedRay.direction.Dot(transformedRay.direction)
	b := 2 * transformedRay.direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return []Intersection{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	return []Intersection{
		NewIntersection(math.Min(t1, t2), s),
		NewIntersection(math.Max(t1, t2), s),
	}
}

func (s *Sphere) NormalAt(worldPoint Tuple) Tuple {
	sphereCenter := NewPoint(0, 0, 0)

	objectSpacePoint := s.transform.Inverse().MulTuple(worldPoint)
	objectSpaceNormal := objectSpacePoint.Sub(sphereCenter)
	// For usual point we could just multiply by a sphere's transformation matrix to
	// transform vector from Object space to World space. But for normals it doesn't work,
	// because it transforms them in undesired way (e.g. squishing normals along with squishing
	// the object)
	normalInWorldSpace := s.transform.Inverse().Transpose().MulTuple(objectSpaceNormal)
	normalInWorldSpace = normalInWorldSpace.AsVector().Normalize()

	return normalInWorldSpace
}
