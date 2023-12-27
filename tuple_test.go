package raytracer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTuple(t *testing.T) {

	t.Run("NewPoint() creates a new point", func(t *testing.T) {
		NewPoint(1.0, 2.0, 3.0)

	})

	t.Run("NewVector() creates a new vector", func(t *testing.T) {
		NewVector(1.0, 2.0, 3.0)
	})

	t.Run("adding a vector to vector must produce vector", func(t *testing.T) {

		vecA := NewVector(1.0, 2.0, 5.0)
		vecB := NewVector(1.0, 2.0, 5.0)

		vecA.Add(vecB)

	})

	t.Run("adding a vector to point must give a point", func(t *testing.T) {

		point := NewPoint(1.0, 2.0, 5.0)
		vec := NewVector(1.0, 2.0, 5.0)

		newPoint := point.AddVector(vec)

		assert.IsType(t, *newPoint, Point{})

	})

	t.Run("subtracting two point must give a vector", func(t *testing.T) {

		pointA := NewPoint(1.0, 2.0, 5.0)
		pointB := NewPoint(1.0, 2.0, 5.0)

		vec := pointA.Sub(pointB)

		assert.IsType(t, *vec, Vector{})

	})

	t.Run("subtracting a vector from point must give point", func(t *testing.T) {

		point := NewPoint(1.0, 2.0, 5.0)
		vec := NewVector(1.0, 2.0, 5.0)

		newPoint := point.SubVector(vec)

		assert.IsType(t, *newPoint, Point{})

	})

	t.Run("Negate() must subtract the vector from zero vector", func(t *testing.T) {

		vec := NewVector(1.0, -2.0, 3.0)
		negatedVec := vec.Negate()

		assert.True(t, Equals(vec.X, -negatedVec.X))
		assert.True(t, Equals(vec.Y, -negatedVec.Y))
		assert.True(t, Equals(vec.Z, -negatedVec.Z))

	})

}

func TestColors(t *testing.T) {

	t.Run("NewColor() must create a new color", func(t *testing.T) {

		color := NewColor(12, 13, 14)

		assert.True(t, Equals(color.R, 12))
		assert.True(t, Equals(color.G, 13))
		assert.True(t, Equals(color.B, 14))

	})

	t.Run("colors must be able to add each other", func(t *testing.T) {

		colorOne := NewColor(12, 13, 14)
		colorTwo := NewColor(12, 13, 14)

		colorAdded := colorOne.Add(colorTwo)

		assert.True(t, Equals(colorOne.R+colorTwo.R, colorAdded.R))
		assert.True(t, Equals(colorOne.G+colorTwo.G, colorAdded.G))
		assert.True(t, Equals(colorOne.B+colorTwo.B, colorAdded.B))

	})

	t.Run("colors must be subtract from each other", func(t *testing.T) {

		colorOne := NewColor(12, 13, 14)
		colorTwo := NewColor(12, 13, 14)

		colorSubbed := colorOne.Sub(colorTwo)

		assert.True(t, Equals(colorOne.R-colorTwo.R, colorSubbed.R))
		assert.True(t, Equals(colorOne.G-colorTwo.G, colorSubbed.G))
		assert.True(t, Equals(colorOne.B-colorTwo.B, colorSubbed.B))

	})

	t.Run("colors must be able to multiply a scalar", func(t *testing.T) {

		scalar := float64(12)

		colorOne := NewColor(12, 13, 14)

		colorMul := colorOne.MultiplyScalar(float64(scalar))

		assert.True(t, Equals(colorMul.R, colorOne.R*scalar))
		assert.True(t, Equals(colorMul.G, colorOne.G*scalar))
		assert.True(t, Equals(colorMul.B, colorOne.B*scalar))

	})

	t.Run("colors must be able to calculate hamadard product", func(t *testing.T) {

		colorOne := NewColor(12, 13, 14)
		colorTwo := NewColor(12, 13, 14)

		colorMul := colorOne.Multiply(colorTwo)

		assert.True(t, Equals(colorMul.R, colorOne.R*colorTwo.R))
		assert.True(t, Equals(colorMul.G, colorOne.G*colorTwo.G))
		assert.True(t, Equals(colorMul.B, colorOne.B*colorTwo.B))

	})

}
