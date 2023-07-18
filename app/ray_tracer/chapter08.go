package ray_tracer

import "math"

func createWorldWithObjects08() World {
	flatScaling := *NewScalingMatrix(10, 0.01, 10)

	w := NewWorld()

	floor := NewDefaultSphere()
	floor.transform = flatScaling
	floor.material = NewDefaultMaterial()
	floor.material.color = NewColor(1, 0.9, 0.9)
	floor.material.specular = 0
	w["floor"] = &floor

	leftWall := NewDefaultSphere()
	// transformations are applied in reverse order
	leftWall.transform = *NewTranslationMatrix(0, 0, 5).MulMat(NewRotationYMatrix(-math.Pi / 4)).MulMat(NewRotationXMatrix(math.Pi / 2)).MulMat(&flatScaling)
	leftWall.material = floor.material
	w["leftWall"] = &leftWall

	rightWall := NewDefaultSphere()
	rightWall.transform = *NewTranslationMatrix(0, 0, 5).MulMat(NewRotationYMatrix(math.Pi / 4)).MulMat(NewRotationXMatrix(math.Pi / 2)).MulMat(&flatScaling)
	rightWall.material = floor.material
	w["rightWall"] = &rightWall

	sphereMaterial := NewDefaultMaterial()
	sphereMaterial.color = GREEN
	sphereMaterial.diffuse = 0.7
	sphereMaterial.specular = 0.3

	middleSphere := NewDefaultSphere()
	middleSphere.transform = *NewTranslationMatrix(-0.5, 1, 0.5)
	middleSphere.material = sphereMaterial
	w["middleSphere"] = &middleSphere

	rightSphere := NewDefaultSphere()
	rightSphere.transform = *NewTranslationMatrix(1.5, 0.5, -0.5).MulMat(NewScalingMatrix(0.5, 0.5, 0.5))
	rightSphere.material = sphereMaterial
	w["rightSphere"] = &rightSphere

	leftSphere := NewDefaultSphere()
	leftSphere.transform = *NewTranslationMatrix(-1.5, 0.33, -0.75).MulMat(NewScalingMatrix(0.33, 0.33, 0.33))
	leftSphere.material = sphereMaterial
	leftSphere.material.color = NewColor(1, 0.8, 0.1)
	w["leftSphere"] = &leftSphere

	w.SetLight(NewPointLight(NewPoint(-10, 10, -10), WHITE))
	return w
}

func Chapter08Shadows(filename string) {
	w := createWorldWithObjects08()

	// Change camera size to get a better resolution
	camera := NewCamera(600, 400, math.Pi/3)
	from, to, up := NewPoint(0, 1.5, -5), NewPoint(0, 1, 0), NewVector(0, 1, 0)
	camera.transform = *NewViewTransformation(from, to, up)

	canvas := camera.Render(w)
	canvas.SavePpm(filename)
}
