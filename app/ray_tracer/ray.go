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
	// It makes the math easier.
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

// Unit sphere (radius == 1), with a center in (0,0,0)
type Sphere struct {
	id        string
	origin    Tuple
	transform Matrix
	material  Material
}

func newSphere(id string, material Material) Sphere {
	return Sphere{
		id:        id,
		origin:    newPoint(0, 0, 0),
		transform: *newIdentityMatrix(4),
		material:  material,
	}
}

func newDefaultSphere() Sphere {
	return Sphere{
		id:        "sphere_id",
		origin:    newPoint(0, 0, 0),
		transform: *newIdentityMatrix(4),
		material:  newDefaultMaterial(),
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
	sphereCenter := newPoint(0, 0, 0)

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

type Intersection struct {
	time   float64
	object *Sphere
}

func newIntersection(t float64, object *Sphere) Intersection {
	return Intersection{t, object}
}

func Hit(intersections []Intersection) (Intersection, bool) {
	sort.Slice(intersections, func(i, j int) bool { return intersections[i].time < intersections[j].time })

	for _, intersection := range intersections {
		if intersection.time > 0 {
			return intersection, true
		}
	}

	dummy := newIntersection(0, nil)
	return dummy, false
}

type PointLight struct {
	position  Tuple
	intensity Color
}

func newPointLight(position Tuple, intensity Color) PointLight {
	return PointLight{position: position, intensity: intensity}
}

type Material struct {
	color     Color
	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
}

func newMaterial(color Color, ambient, diffuse, specular, shininess float64) Material {
	if ambient < 0 || diffuse < 0 || specular < 0 || shininess < 0 {
		panic("All Material's attribues must be nonnegative!")
	}

	return Material{color, ambient, diffuse, specular, shininess}
}

func newDefaultMaterial() Material {
	return Material{WHITE, 0.1, 0.9, 0.9, 200.}
}

func lighting(material Material, light PointLight, position, eyeV, normalV Tuple) Color {
	if !position.IsPoint() {
		panic("Position must be a point!")
	}

	if !eyeV.IsVector() || !normalV.IsVector() {
		panic("Eye and Normal must be vectors!")
	}

	effectiveColor := material.color.MultHadamar(light.intensity)
	ligthV := light.position.Sub(position).Normalize()
	ambient := effectiveColor.MultScalar(material.ambient)

	// negative dot product means the light is on the other side of the surface
	// and should not contribute to the final lighting
	lightDotNormal := ligthV.Dot(normalV)
	diffuse, specular := BLACK, BLACK
	if lightDotNormal >= 0 {
		diffuse = effectiveColor.MultScalar(material.diffuse).MultScalar(lightDotNormal)

		reflectV := ligthV.Mul(-1).ReflectAround(normalV)
		// the same for reflected light. If negative -> reflected light doesn't contribute
		// final intensity (specular == Black)
		reflectDotEye := reflectV.Dot(eyeV)
		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, material.shininess)
			specular = light.intensity.MultScalar(material.specular).MultScalar(factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
