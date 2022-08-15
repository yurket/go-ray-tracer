package main

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

func (r *Ray) Position(time float64) Tuple {
	return r.origin.Add(r.direction.Mul(time))
}

func (r *Ray) Intersect(s Sphere) []Intersection {
	sphereToRay := r.origin.Sub(newPoint(0, 0, 0))
	a := r.direction.Dot(r.direction)
	b := 2 * r.direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return []Intersection{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	// TODO: here the Sphere object s is being copied. It may need be fixed
	return []Intersection{
		newIntersection(math.Min(t1, t2), s),
		newIntersection(math.Max(t1, t2), s),
	}
}

type Sphere struct {
	id     string
	origin Tuple
	radius float64
}

func newSphere(id string) Sphere {
	return Sphere{id, newPoint(0, 0, 0), 1.0}
}

type Intersection struct {
	t      float64
	object Sphere
}

func newIntersection(t float64, object Sphere) Intersection {
	return Intersection{t, object}
}

func Hit(intersections []Intersection) (Intersection, bool) {
	sort.Slice(intersections, func(i, j int) bool { return intersections[i].t < intersections[j].t })

	for _, intersection := range intersections {
		if intersection.t > 0 {
			return intersection, true
		}
	}

	// TODO: doesn't look right
	dummy := newIntersection(0, newSphere(""))
	return dummy, false
}
