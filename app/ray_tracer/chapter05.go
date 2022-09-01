package ray_tracer

import "fmt"

// My intuitive approach. Wall is a canvas (z=0), and I use 2 "parameters"
// to make the shadow of the sphere look as big as I'd like:
// 1st "parameter" is a size of the sphere (scale)
// 2n "parameter" is the distance of the "eye" (where all the rays are coming from) to the wall (canvas)
func Chapter05(filename string) {
	const height, width = 100, 100
	canvas := newCanvas(width, height)

	s := newDefaultSphere()
	const scale, eyeCenterXY = 20., 50

	transform := newIdentityMatrix(4).Scale(scale, scale, scale).Translate(eyeCenterXY, eyeCenterXY, 0)
	s.SetTransform(transform)

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
			intersections := ray.Intersect(&s)
			_, ok := Hit(intersections)
			if !ok {
				continue
			}

			hitCount++
			shadowX, shadowY := canvas.ToCanvasCoordinates(x, y)
			canvas.WritePixel(shadowX, shadowY, RED)
		}
	}

	fmt.Println("Total hit count: ", hitCount)
	canvas.SavePpm(filename)
}

// Book solution doesn't scale the sphere, but uses a "virtual" which maps to the canvas
func Chapter05BookSolution(filename string) {
	const canvasSize = 100
	canvas := newCanvas(canvasSize, canvasSize)

	s := newDefaultSphere()

	eyeOrigin := newPoint(0, 0, -5)
	const wallZ, wallSize = 10, 7
	const pixelSize = float64(wallSize) / canvasSize
	const halfWall = wallSize / 2.

	hitCount := 0
	for y := 0.; y < canvasSize; y++ {
		// compute the world y coordinate (top = +half, bottom = -half)
		worldY := halfWall - pixelSize*y

		for x := 0.; x < canvasSize; x++ {
			// compute the world x coordinate (left = -half, right = half)
			worldX := -halfWall + pixelSize*x
			targetPoint := newPoint(worldX, worldY, wallZ)
			direction := targetPoint.Sub(eyeOrigin).Normalize()
			ray := newRay(eyeOrigin, direction)
			xs := ray.Intersect(&s)
			if _, ok := Hit(xs); !ok {
				continue
			}

			hitCount++
			canvas.WritePixel(int(x), int(y), RED)
		}
	}

	fmt.Println("Total hit count: ", hitCount)
	canvas.SavePpm(filename)
}
