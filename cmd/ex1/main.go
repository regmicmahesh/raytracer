package main

import (
	"fmt"

	"github.com/regmicmahesh/raytracer"
)

// A projectile has a position (a point) and a velocity (a vector).
type Projectile struct {
	position *raytracer.Point
	velocity *raytracer.Vector
}

// An environment has gravity (a vector) and wind (a vector).
type Environment struct {
	gravity *raytracer.Vector
	wind    *raytracer.Vector
}

func tick(p *Projectile, e *Environment) *Projectile {
	return &Projectile{
		position: p.position.AddVector(p.velocity),
		velocity: p.velocity.Add(e.gravity.Add(e.wind)),
	}
}

// projectile(point(0, 1, 0), normalize(vector(1, 1, 0)))

func main() {

	projectilePosition := raytracer.NewPoint(0, 1, 0)
	projectileVelocity := raytracer.NewVector(1, 1, 0).Normalize().MultiplyScalar(11.25)

	// gravity -0.1 unit/tick, and wind is -0.01 unit/tick.
	// e ← environment(vector(0, -0.1, 0), vector(-0.01, 0, 0))
	envGravity := raytracer.NewVector(0, -0.1, 0)
	envWind := raytracer.NewVector(-0.01, 0, 0)

	projectile := &Projectile{
		position: projectilePosition,
		velocity: projectileVelocity,
	}

	env := &Environment{
		gravity: envGravity,
		wind:    envWind,
	}

	tickCount := 0

	for projectile.position.Y > 0 {
		projectile = tick(projectile, env)
		tickCount += 1
		fmt.Printf("Tick %d: [%.2f %.2f %.2f]\n", tickCount, projectile.position.X, projectile.position.Y, projectile.position.Z)
	}

}
