package ray_tracer

import (
	"math"
)

type SceneObject interface{}

// TODO: Can I do new types Point and Vector inherited from Tuple?
type Ray struct {
	origin    Tuple
	direction Tuple
}

func NewRay(origin, direction Tuple) Ray {
	if !origin.IsPoint() {
		panic("Origin must be a point!")
	}
	if !direction.IsVector() {
		panic("Direction must be a vector!")
	}
	return Ray{origin, direction}
}

func (r *Ray) CalcPosition(time float64) Tuple {
	return r.origin.Add(r.direction.Mul(time))
}

// Finds intersection of a ray going through the center of the sphere
// with a unit radius
func (r *Ray) Intersect(s *Sphere) []Intersection {
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

func (r *Ray) ApplyTransform(m *Matrix) Ray {
	return NewRay(m.MulTuple(r.origin), m.MulTuple(r.direction))
}
