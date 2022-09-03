package ray_tracer

import "sort"

type Intersection struct {
	time   float64
	object *Sphere
}

func NewIntersection(t float64, object *Sphere) Intersection {
	return Intersection{t, object}
}

func Hit(intersections []Intersection) (Intersection, bool) {
	sort.Slice(intersections, func(i, j int) bool { return intersections[i].time < intersections[j].time })

	for _, intersection := range intersections {
		if intersection.time > 0 {
			return intersection, true
		}
	}

	dummy := NewIntersection(0, nil)
	return dummy, false
}
