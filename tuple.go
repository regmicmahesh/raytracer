package raytracer

import "math"

type Tuple struct {
	X float64
	Y float64
	Z float64
}

func NewTuple(x, y, z float64) *Tuple {

	return &Tuple{
		X: x,
		Y: y,
		Z: z,
	}

}

type Point Tuple

func NewPoint(x, y, z float64) *Point {
	return &Point{
		X: x,
		Y: y,
		Z: z,
	}
}

func (p *Point) AddVector(n *Vector) *Point {
	return &Point{
		X: p.X + n.X,
		Y: p.Y + n.Y,
		Z: p.Z + n.Z,
	}
}

func (p *Point) SubVector(n *Vector) *Point {
	return &Point{
		X: p.X - n.X,
		Y: p.Y - n.Y,
		Z: p.Z - n.Z,
	}
}

func (p *Point) Sub(n *Point) *Vector {
	return &Vector{
		X: p.X - n.X,
		Y: p.Y - n.Y,
		Z: p.Z - n.Z,
	}
}

type Vector Tuple

var ZeroVector = &Vector{}

func NewVector(x, y, z float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v *Vector) Add(n *Vector) *Vector {
	return &Vector{
		X: v.X + n.X,
		Y: v.Y + n.Y,
		Z: v.Z + n.Z,
	}
}

func (v *Vector) Sub(n *Vector) *Vector {
	return &Vector{
		X: v.X - n.X,
		Y: v.Y - n.Y,
		Z: v.Z - n.Z,
	}
}

func (t *Vector) Negate() *Vector {
	return ZeroVector.Sub(t)
}

func (t *Vector) MultiplyScalar(s float64) *Vector {
	return &Vector{
		X: t.X * s,
		Y: t.Y * s,
		Z: t.Z * s,
	}
}

func (t *Vector) Divide(s float64) *Vector {
	return t.MultiplyScalar(1 / s)
}

func (t *Vector) Magnitude() float64 {
	return math.Sqrt(
		t.X*t.X + t.Y*t.Y + t.Z*t.Z,
	)
}

func (t *Vector) Normalize() *Vector {
	m := t.Magnitude()
	return t.Divide(m)
}

func (t *Vector) Dot(n *Vector) float64 {
	return t.X*n.X + t.Y*n.Y + t.Z*n.Z
}

type Color struct {
	R float64
	G float64
	B float64
}

var BlackColor = &Color{
	R: 0,
	G: 0,
	B: 0,
}

func NewColor(r, g, b float64) *Color {
	return &Color{
		R: r,
		G: g,
		B: b,
	}
}

func (c *Color) Add(n *Color) *Color {
	return &Color{
		R: c.R + n.R,
		G: c.G + n.G,
		B: c.B + n.B,
	}
}

func (c *Color) Sub(n *Color) *Color {
	return &Color{
		R: c.R - n.R,
		G: c.G - n.G,
		B: c.B - n.B,
	}
}

func (c *Color) MultiplyScalar(s float64) *Color {
	return &Color{
		R: c.R * s,
		G: c.G * s,
		B: c.B * s,
	}
}

func (c *Color) Multiply(n *Color) *Color {
	return &Color{
		R: c.R * n.R,
		G: c.G * n.G,
		B: c.B * n.B,
	}
}

func (c *Color) Clamp(base float64) *Color {

	return &Color{
		R: math.Min(math.Max(c.R*base, 0), base),
		G: math.Min(math.Max(c.G*base, 0), base),
		B: math.Min(math.Max(c.B*base, 0), base),
	}

}
