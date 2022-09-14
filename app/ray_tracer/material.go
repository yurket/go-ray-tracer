package ray_tracer

type Material struct {
	color     Color
	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
}

func NewMaterial(color Color, ambient, diffuse, specular, shininess float64) Material {
	if ambient < 0 || diffuse < 0 || specular < 0 || shininess < 0 {
		panic("All Material's attribues must be nonnegative!")
	}

	return Material{color, ambient, diffuse, specular, shininess}
}

func NewDefaultMaterial() Material {
	return Material{WHITE, 0.1, 0.9, 0.9, 200.}
}
