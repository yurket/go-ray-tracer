// TODO: Rename file or move out Sphere, Intersection in separate files
package ray_tracer

import (
	"math"
	"sort"
)

// TODO: Can I do new types Point and Vector inherited from Tuple?
type Ray struct {
	origin    Tuple
	direction Tuple
}

func newRay(origin, direction Tuple) Ray {
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
	// It makes easy the math.
	t := s.Transform()
	transformedRay := r.ApplyTransform(t.Inverse())

	sphereOrigin := newPoint(0, 0, 0)
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
		newIntersection(math.Min(t1, t2), s),
		newIntersection(math.Max(t1, t2), s),
	}
}

func (r *Ray) ApplyTransform(m *Matrix) Ray {
	return newRay(m.MulTuple(r.origin), m.MulTuple(r.direction))
}

type Sphere struct {
	id        string
	origin    Tuple
	transform Matrix
}

func newSphere(id string) Sphere {
	return Sphere{id, newPoint(0, 0, 0), *newIdentityMatrix(4)}
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

type Intersection struct {
	t      float64
	object *Sphere
}

func newIntersection(t float64, object *Sphere) Intersection {
	return Intersection{t, object}
}

func Hit(intersections []Intersection) (Intersection, bool) {
	sort.Slice(intersections, func(i, j int) bool { return intersections[i].t < intersections[j].t })

	for _, intersection := range intersections {
		if intersection.t > 0 {
			return intersection, true
		}
	}

	dummy := newIntersection(0, nil)
	return dummy, false
}
