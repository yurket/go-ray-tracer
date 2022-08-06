package main

import "fmt"

type projectile struct {
	position Tuple
	velocity Tuple
}

type environment struct {
	gravity Tuple
	wind    Tuple
}

func tick(env environment, proj projectile) projectile {
	var new_proj projectile
	new_proj.position = Add(proj.position, proj.velocity)
	new_proj.velocity = Add(Add(proj.velocity, env.gravity), env.wind)

	return new_proj
}

func Chapter01Projectile() {
	env := environment{gravity: Vector(0, -0.1, 0), wind: Vector(-0.01, 0, 0)}
	proj := projectile{position: Point(0, 1, 0), velocity: Normalize(Vector(1, 1, 0))}
	i := 0
	for ; proj.position.y >= 0; i++ {
		fmt.Printf("iter %d, pos: %3.3f, vel: %3.2f\n", i, proj.position, proj.velocity)
		proj = tick(env, proj)
	}
	fmt.Printf("Projectile hit the ground on tick %d\n", i)
}
