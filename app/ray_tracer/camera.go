package ray_tracer

import "math"

// Primary responsibility of the Camera is to map 3D scene onto a 2D canvas
type Camera struct {
	hSize       int
	vSize       int
	fieldOfView float64
	transform   Matrix
	halfWidth   float64
	halfHeight  float64
	pixelSize   float64
}

func calcCameraParameters(hsize, vsize int, fieldOfView float64) (halfWidth, halfHeight, pixelSize float64) {
	halfView := math.Tan(fieldOfView / 2.)
	aspectRatio := float64(hsize) / float64(vsize)

	if aspectRatio >= 1 {
		halfWidth = halfView
		halfHeight = halfView / aspectRatio
	} else {
		halfWidth = halfView * aspectRatio
		halfHeight = halfView
	}

	pixelSize = (halfWidth * 2) / float64(hsize)
	return
}

func NewCamera(hsize, vsize int, fieldOfView float64) Camera {
	halfWidth, halfHeight, pixelSize := calcCameraParameters(hsize, vsize, fieldOfView)
	return Camera{
		hSize:       hsize,
		vSize:       vsize,
		fieldOfView: fieldOfView,
		transform:   *NewIdentityMatrix(4),
		halfWidth:   halfWidth,
		halfHeight:  halfHeight,
		pixelSize:   pixelSize,
	}
}

func (c *Camera) CastRayIntoPixel(px, py int) Ray {
	xOffset := (float64(px) + 0.5) * c.pixelSize
	yOffset := (float64(py) + 0.5) * c.pixelSize

	// the untransformed coordinates of the pixel in world space.
	worldX := c.halfWidth - xOffset
	worldY := c.halfHeight - yOffset

	// remember that canvas is at z=-1
	worldPixel := NewPoint(worldX, worldY, -1)

	// pixel in Camera space (?)
	pixel := c.transform.Inverse().MulTuple(worldPixel)
	origin := c.transform.Inverse().MulTuple(NewPoint(0, 0, 0))
	direction := pixel.Sub(origin).Normalize()

	return NewRay(origin, direction)
}

func (c *Camera) Render(w World) Canvas {
	canvas := NewCanvas(c.hSize, c.vSize)
	for y := 0; y < c.vSize; y++ {
		for x := 0; x < c.hSize; x++ {
			r := c.CastRayIntoPixel(x, y)
			color := w.ColorAtIntersection(r)
			canvas.WritePixel(x, y, color)
		}
	}
	return canvas
}
