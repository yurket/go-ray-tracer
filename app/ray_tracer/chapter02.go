package ray_tracer

import (
	"fmt"
)

func Chapter02DrawProjectilePpm(ppmFilename string) {
	WIDTH := 1000
	HEIGHT := 500
	canvas := newCanvas(WIDTH, HEIGHT)

	env := environment{gravity: newVector(0, -0.1, 0), wind: newVector(-0.01, 0, 0)}
	proj := projectile{position: newPoint(0, 10, 0), velocity: newVector(6, 9.5, 0)}
	i := 0
	for ; proj.position.y >= 0; i++ {
		pos := proj.position
		canvas_x, canvas_y := canvas.ToCanvasCoordinates(proj.position.x, proj.position.y)
		canvas.WritePixel(canvas_x, canvas_y, RED)
		fmt.Printf("iter %d, pos: %3.3f, vel: %3.2f\n", i, pos, proj.velocity)

		proj = tick(env, proj)
	}
	fmt.Printf("Projectile hit the ground on tick %d\n", i)

	canvas.SavePpm(ppmFilename)
}
