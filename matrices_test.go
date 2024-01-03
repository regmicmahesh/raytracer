package raytracer

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrix(t *testing.T) {

	t.Run("must be able to create new matrix", func(t *testing.T) {

		{

			matrix := NewMatrix4()
			assert.IsType(t, *matrix, Matrix4{})
		}

		{
			matrix := NewMatrix3()
			assert.IsType(t, *matrix, Matrix3{})

		}
		{

			matrix := NewMatrix2()
			assert.IsType(t, *matrix, Matrix2{})
		}

	})

	t.Run("must be able to set and get values in the matrix", func(t *testing.T) {
		{
			val := rand.Float64() * 10
			col, row := rand.Int63n(4), rand.Int63n(4)
			matrix := NewMatrix4()
			matrix.Set(col, row, val)
			assert.Equal(t, matrix.Get(col, row), val)
		}
	})

	t.Run("set and get must leave other values unchanged", func(t *testing.T) {
		{
			val := rand.Float64() * 10
			col, row := rand.Int63n(4), rand.Int63n(4)
			matrix := NewMatrix4()
			matrix.Set(col, row, val)

			for i := 0; i < 4; i++ {
				for j := 0; j < 4; j++ {

					if int64(i) == col && int64(j) == row {
						assert.Equal(t, matrix.Get(col, row), val)
						continue
					}

					assert.Zero(t, matrix.Get(int64(i), int64(j)))

				}
			}
		}
	})

	t.Run("must be able to initialize matrix with data", func(t *testing.T) {

		m := NewMatrix4()
		m.Initialize([4][4]float64{
			{1, 2, 3, 4},
			{4, 3, 2, 1},
			{5, 6, -7, 8},
			{3, 4, 1, 0},
		})
		assert.Equal(t, m.Get(0, 0), float64(1))
		assert.Equal(t, m.Get(0, 1), float64(2))
		assert.Equal(t, m.Get(0, 2), float64(3))
		assert.Equal(t, m.Get(0, 3), float64(4))

		assert.Equal(t, m.Get(1, 0), float64(4))
		assert.Equal(t, m.Get(1, 1), float64(3))
		assert.Equal(t, m.Get(1, 2), float64(2))
		assert.Equal(t, m.Get(1, 3), float64(1))

		assert.Equal(t, m.Get(2, 0), float64(5))
		assert.Equal(t, m.Get(2, 1), float64(6))
		assert.Equal(t, m.Get(2, 2), float64(-7))
		assert.Equal(t, m.Get(2, 3), float64(8))

		assert.Equal(t, m.Get(3, 0), float64(3))
		assert.Equal(t, m.Get(3, 1), float64(4))
		assert.Equal(t, m.Get(3, 2), float64(1))
		assert.Equal(t, m.Get(3, 3), float64(0))
	})

	t.Run("must be able to compute equality of matrix", func(t *testing.T) {

		m := NewMatrix4()
		m.Initialize([4][4]float64{
			{1, 2, 3, 4},
			{4, 3, 2, 1},
			{5, 6, -7, 8},
			{3, 4, 1, 0},
		})

		m2 := NewMatrix4()
		m2.Initialize([4][4]float64{
			{1, 2, 3, 4},
			{4, 3, 2, 1},
			{5, 6, -7, 8},
			{3, 4, 1, 0},
		})

		assert.True(t, m.Equals(m2))

		m2.Set(0, 0, 100)
		assert.False(t, m.Equals(m2))

	})

	t.Run("must be able to multiply matrix", func(t *testing.T) {

		m1 := NewMatrix4()
		m1.Initialize([4][4]float64{
			{1, 3, 4, 6},
			{2, 3, 1, 32},
			{3, 20, 31, 32},
			{2, 0, 2, 0},
		})

		m2 := NewMatrix4()
		m2.Initialize([4][4]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{-1, 3, 4, 5},
			{1, 43, 3, 5},
		})

		mul := m1.Multiply(m2)

		assert.Equal(t, mul.Get(0, 0), float64(18))
		assert.Equal(t, mul.Get(0, 1), float64(290))
		assert.Equal(t, mul.Get(0, 2), float64(58))
		assert.Equal(t, mul.Get(0, 3), float64(78))

		assert.Equal(t, mul.Get(1, 0), float64(48))
		assert.Equal(t, mul.Get(1, 1), float64(1401))
		assert.Equal(t, mul.Get(1, 2), float64(127))
		assert.Equal(t, mul.Get(1, 3), float64(197))

		assert.Equal(t, mul.Get(2, 0), float64(104))
		assert.Equal(t, mul.Get(2, 1), float64(1595))
		assert.Equal(t, mul.Get(2, 2), float64(369))
		assert.Equal(t, mul.Get(2, 3), float64(487))

		assert.Equal(t, mul.Get(3, 0), float64(0))
		assert.Equal(t, mul.Get(3, 1), float64(10))
		assert.Equal(t, mul.Get(3, 2), float64(14))
		assert.Equal(t, mul.Get(3, 3), float64(18))

	})

	t.Run("must be able to calculate sub matrix", func(t *testing.T) {

		matrix := NewMatrix3()

		matrix.Set(0, 0, 1)
		matrix.Set(0, 1, 5)
		matrix.Set(0, 2, 0)

		matrix.Set(1, 0, -3)
		matrix.Set(1, 1, 2)
		matrix.Set(1, 2, 7)

		matrix.Set(2, 0, 0)
		matrix.Set(2, 1, 6)
		matrix.Set(2, 2, -3)

		matrix.SubMatrix(0, 0)

	})

}
