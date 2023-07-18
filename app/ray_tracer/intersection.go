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

type IntersectionComputations struct {
	intersectionTime   float64
	intersectionObject *Sphere
	intersectionPoint  Tuple
	overPoint          Tuple
	eyev               Tuple
	objectNormalv      Tuple
	insideHit          bool
}

func PrepareIntersectionComputations(i Intersection, r Ray) IntersectionComputations {
	comps := IntersectionComputations{
		intersectionTime:   i.time,
		intersectionObject: i.object,
		insideHit:          false,
	}
	comps.intersectionPoint = r.CalcPosition(i.time)
	comps.eyev = r.direction.Mul(-1)
	comps.objectNormalv = i.object.NormalAt(comps.intersectionPoint)

	if comps.eyev.Dot(comps.objectNormalv) < 0 {
		comps.insideHit = true
		comps.objectNormalv = comps.objectNormalv.Mul(-1)
	}

	// A point, very close to the intersection point, but adjusted a bit into the
	// direction of a normal. Used to fight the "acne effect", while testing for shadowing
	comps.overPoint = comps.intersectionPoint.Add(comps.objectNormalv.Mul(EPSILON))

	return comps
}
