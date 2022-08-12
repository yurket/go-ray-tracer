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
	new_proj.position = proj.position.Add(proj.velocity)
	new_proj.velocity = (proj.velocity.Add(env.gravity)).Add(env.wind)

	return new_proj
}

func Chapter01Projectile() {
	env := environment{gravity: newVector(0, -0.1, 0), wind: newVector(-0.01, 0, 0)}
	proj := projectile{position: newPoint(0, 1, 0), velocity: newVector(1, 1, 0).Normalize()}
	i := 0
	for ; proj.position.y >= 0; i++ {
		fmt.Printf("iter %d, pos: %3.3f, vel: %3.2f\n", i, proj.position, proj.velocity)
		proj = tick(env, proj)
	}
	fmt.Printf("Projectile hit the ground on tick %d\n", i)
}
