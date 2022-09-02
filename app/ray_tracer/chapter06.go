package ray_tracer

import "fmt"

func Chapter06LightAndShading(filename string) {
	const height, width = 75, 75
	canvas := newCanvas(width, height)

	m := newDefaultMaterial()
	m.color = GREEN
	sphere := newSphere("sphere_id", m)

	const scale, eyeCenterXY = 25., 37
	transform := newIdentityMatrix(4).Scale(scale, scale, scale).Translate(eyeCenterXY, eyeCenterXY, 0)
	sphere.SetTransform(transform)

	lightPosition := newPoint(eyeCenterXY+30, eyeCenterXY+30, 80)
	light := newPointLight(lightPosition, WHITE)

	eyeOrigin := newPoint(eyeCenterXY, eyeCenterXY, 100)
	hitCount := 0
	for x := 0.; x < width; x++ {
		for y := 0.; y < height; y++ {
			targetCanvasPoint := newPoint(x, y, 0)
			direction := targetCanvasPoint.Sub(eyeOrigin).Normalize()
			if !direction.IsVector() {
				panic("Directions should definitely be a vector!")
			}

			ray := newRay(eyeOrigin, direction)
			intersections := ray.Intersect(&sphere)
			hit, ok := Hit(intersections)
			if !ok {
				continue
			}
			hitCount++

			intersectionPoint := ray.CalcPosition(hit.time)
			normal := sphere.NormalAt(intersectionPoint)
			eye := ray.direction.Mul(-1)
			colorOnTheSphere := CalcLighting(m, light, intersectionPoint, eye, normal)

			shadowX, shadowY := canvas.ToCanvasCoordinates(x, y)
			canvas.WritePixel(shadowX, shadowY, colorOnTheSphere)
		}
	}

	fmt.Println("Total hit count: ", hitCount)
	canvas.SavePpm(filename)
}
