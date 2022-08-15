package main

import "math"

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

func (r *Ray) Position(time float64) Tuple {
	return r.origin.Add(r.direction.Mul(time))
}

func (r *Ray) Intersect(s Sphere) []float64 {
	sphereToRay := r.origin.Sub(newPoint(0, 0, 0))
	a := r.direction.Dot(r.direction)
	b := 2 * r.direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return []float64{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	return []float64{math.Min(t1, t2), math.Max(t1, t2)}
}

type Sphere struct {
	id     string
	origin Tuple
	radius float64
}

func newSphere(id string) Sphere {
	return Sphere{id, newPoint(0, 0, 0), 1.0}
}