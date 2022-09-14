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

func CalcLighting(material Material, light PointLight, position, eyeV, normalV Tuple) Color {
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

func ShadeHit(world World, comps *IntersectionComputations) Color {
	// TODO: Add support of multiple lights
	return CalcLighting(comps.intersectionObject.material, world.Light(), comps.intersectionPoint, comps.eyev, comps.objectNormalv)
}
