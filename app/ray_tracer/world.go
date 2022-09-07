package ray_tracer

import (
	"fmt"
	"sort"
)

type World map[string]SceneObject

func NewWorld() World {
	return World{}
}

// Default world is hardcoded and contains 2 spheres "s1" and "s2" and point light "light"
// Spheres' origins in the (0,0,0) and s2 is 2 times smaller than s1. Hence s1 may be conidered
// as an outer sphere, and s1 is an inner sphere
func NewDefaultWorld() World {
	light := NewPointLight(NewPoint(-10, 10, -10), WHITE)

	lightGreen := NewColor(0.8, 1, 0.6)
	m := NewDefaultMaterial()
	m.color = lightGreen
	m.diffuse = 0.7
	m.specular = 0.2

	s1 := NewSphere("sphere_id", m)

	s2 := NewDefaultSphere()
	s2.SetTransform(NewScalingMatrix(0.5, 0.5, 0.5))

	return World{"light": &light, "s1": &s1, "s2": &s2}
}

func (w World) Light() PointLight {
	object, ok := w["light"]
	if !ok {
		panic("World has no light in it =(")
	}

	light, ok := object.(*PointLight)
	if !ok {
		panic("The world is corrupted! world[\"light\"] type is not PointLight!")
	}

	return *light
}

func (w World) SetLight(pl PointLight) {
	w["light"] = &pl
}

func (w World) Sphere(objectName string) *Sphere {
	obj, ok := w[objectName]
	if !ok {
		panic(fmt.Sprintf("No object with name %q in the world!\n", objectName))
	}

	s, ok := obj.(*Sphere)
	if !ok {
		panic(fmt.Sprintf("Type of the object %q is not \"*Sphere\"!", objectName))
	}

	return s
}

func (w World) IntersectWith(r *Ray) []Intersection {
	allIntersections := []Intersection{}

	for _, obj := range w {
		switch obj := obj.(type) {
		case *Sphere:
			xs := r.Intersect(obj)
			allIntersections = append(allIntersections, xs...)
		case *PointLight:
			continue
		default:
			fmt.Printf("Intersection with type %T is not supported\n", obj)
		}
	}

	sort.Slice(allIntersections, func(i, j int) bool { return allIntersections[i].time < allIntersections[j].time })
	return allIntersections
}

// Returns BLACK if ray doesn't intersect with any objects in the world
func (w World) ColorAtIntersection(ray Ray) Color {
	intersections := w.IntersectWith(&ray)
	hit, ok := Hit(intersections)
	if !ok {
		return BLACK
	}

	comps := PrepareIntersectionComputations(hit, ray)
	return ShadeHit(w, &comps)
}
