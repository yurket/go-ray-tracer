package ray_tracer

import (
	"math"
)

func Chapter04DrawAnalogClock(filename string) {
	canvas := NewCanvas(200, 200)

	center := NewPoint(100, 100, 0)
	zeroHand := NewVector(0, 50, 0)
	for angle := 0.; angle < 2*math.Pi; angle += math.Pi / 6 {
		// X axis goes to the right, Y up, Z forward (into the screen).
		// So positive rotation around Z in Left-handed coordinate system (like ours)
		// goes counter clockwise and negative rotation around Z goes clockwise.
		rotatedHand := NewIdentityMatrix(4).RotateZ(-angle).MulTuple(zeroHand)
		rotatedHand = center.Add(rotatedHand)
		x, y := canvas.ToCanvasCoordinates(rotatedHand.x, rotatedHand.y)
		canvas.WritePixel(x, y, WHITE)
	}

	canvas.SavePpm(filename)
}
