package ray_tracer

import (
	"math"
)

type PointLight struct {
	position  Tuple
	intensity Color
}

func NewPointLight(position Tuple, intensity Color) PointLight {
	return PointLight{position: position, intensity: intensity}
}

func CalcLighting(material Material, light PointLight, position, eyeV, normalV Tuple, isInShadow bool) Color {
	if !position.IsPoint() {
		panic("Position must be a point!")
	}

	if !eyeV.IsVector() || !normalV.IsVector() {
		panic("Eye and Normal must be vectors!")
	}

	effectiveColor := material.color.MultHadamar(light.intensity)
	ligthV := light.position.Sub(position).Normalize()
	ambient := effectiveColor.MultScalar(material.ambient)

	if isInShadow {
		return ambient
	}

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

func ShadeHit(world World, comps *IntersectionComputations) Color {
	// TODO: Add support of multiple lights``

	isShadowed := IsShadowed(world, comps.overPoint)
	return CalcLighting(comps.intersectionObject.material, world.Light(), comps.overPoint,
		comps.eyev, comps.objectNormalv, isShadowed)
}

// Checks if theres smth between point and the light source.
// What to do if there are 1+ light sources?
func IsShadowed(world World, point Tuple) bool {
	point_to_light := world.Light().position.Sub(point)
	distance_to_light := point_to_light.Magnitude()
	point_to_light_ray := NewRay(point, point_to_light.Normalize())

	intersections := world.IntersectWith(&point_to_light_ray)
	i, wasHit := Hit(intersections)
	if !wasHit {
		return false
	}

	return i.time < distance_to_light
}
