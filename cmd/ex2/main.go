package main

import (
	"os"

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

func (p *Projectile) tick(e *Environment) *Projectile {
	return &Projectile{
		position: p.position.AddVector(p.velocity),
		velocity: p.velocity.Add(e.gravity.Add(e.wind)),
	}
}

func main() {

	canvas := raytracer.NewCanvas(1200, 550)

	for i := 8; i < 9; i++ {
		projectilePosition := raytracer.NewPoint(0, 1, 0)
		projectileVelocity := raytracer.NewVector(2.5, 2.5, 0).Normalize().MultiplyScalar(11.25)

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

		color := raytracer.NewColor(0, 1, 0)

		canvas.WritePixel(int(projectile.position.X), 550-int(projectile.position.Y), color)

		for projectile.position.Y > 0 {
			projectile = projectile.tick(env)
			tickCount += 1
			canvas.WritePixel(int(projectile.position.X), 550-int(projectile.position.Y), color)
			canvas.WritePixel(int(projectile.position.X)+1, 550-int(projectile.position.Y)+1, color)
		}

	}

	file, err := os.Create("data.ppm")
	if err != nil {
		panic(err)
	}
	canvas.ToPPM(file)

}
