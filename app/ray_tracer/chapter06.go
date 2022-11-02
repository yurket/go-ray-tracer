package ray_tracer

import "fmt"

func Chapter06LightAndShading(filename string) {
	const height, width = 75, 75
	canvas := NewCanvas(width, height)

	m := NewDefaultMaterial()
	m.color = GREEN
	sphere := NewSphere("sphere_id", m)

	const scale, eyeCenterXY = 25., 37
	transform := NewIdentityMatrix(4).Scale(scale, scale, scale).Translate(eyeCenterXY, eyeCenterXY, 0)
	sphere.SetTransform(transform)

	lightPosition := NewPoint(eyeCenterXY+30, eyeCenterXY+30, 80)
	light := NewPointLight(lightPosition, WHITE)

	eyeOrigin := NewPoint(eyeCenterXY, eyeCenterXY, 100)
	hitCount := 0
	for x := 0.; x < width; x++ {
		for y := 0.; y < height; y++ {
			targetCanvasPoint := NewPoint(x, y, 0)
			direction := targetCanvasPoint.Sub(eyeOrigin).Normalize()
			if !direction.IsVector() {
				panic("Directions should definitely be a vector!")
			}

			ray := NewRay(eyeOrigin, direction)
			intersections := sphere.IntersectWith(&ray)
			hit, ok := Hit(intersections)
			if !ok {
				continue
			}
			hitCount++

			intersectionPoint := ray.CalcPosition(hit.time)
			normal := sphere.NormalAt(intersectionPoint)
			eye := ray.direction.Mul(-1)
			isInShadow := false
			colorOnTheSphere := CalcLighting(m, light, intersectionPoint, eye, normal, isInShadow)

			shadowX, shadowY := canvas.ToCanvasCoordinates(x, y)
			canvas.WritePixel(shadowX, shadowY, colorOnTheSphere)
		}
	}

	fmt.Println("Total hit count: ", hitCount)
	canvas.SavePpm(filename)
}
