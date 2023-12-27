package raytracer

import (
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

	// t.Run("inverse check", func(t *testing.T) {
	// 	matrix := NewMatrix4()

	// 	matrix.Set(0, 0, -5)
	// 	matrix.Set(0, 1, 2)
	// 	matrix.Set(0, 2, 6)
	// 	matrix.Set(0, 3, -8)

	// 	matrix.Set(1, 0, 1)
	// 	matrix.Set(1, 1, -5)
	// 	matrix.Set(1, 2, 1)
	// 	matrix.Set(1, 3, 8)

	// 	matrix.Set(2, 0, 7)
	// 	matrix.Set(2, 1, 7)
	// 	matrix.Set(2, 2, -6)
	// 	matrix.Set(2, 3, -7)

	// 	matrix.Set(3, 0, 1)
	// 	matrix.Set(3, 1, -3)
	// 	matrix.Set(3, 2, 7)
	// 	matrix.Set(3, 3, 4)

	// 	inverse := IdentityMatrix4.Inverse()

	// 	fmt.Printf("%#v", inverse.data)

	// })

	t.Run("transformation check", func(t *testing.T) {
		matrix := NewMatrix4()

		matrix.Set(0, 0, 1)
		matrix.Set(0, 1, 0)
		matrix.Set(0, 2, 0)
		matrix.Set(0, 3, 1)

		matrix.Set(1, 0, 0)
		matrix.Set(1, 1, 1)
		matrix.Set(1, 2, 0)
		matrix.Set(1, 3, 2)

		matrix.Set(2, 0, 0)
		matrix.Set(2, 1, 0)
		matrix.Set(2, 2, 1)
		matrix.Set(2, 3, 2)

		matrix.Set(3, 0, 0)
		matrix.Set(3, 1, 0)
		matrix.Set(3, 2, 0)
		matrix.Set(3, 3, 1)

	})

	// 2 7
	// 6 -3

	// t.Run("must be able to set value in matrix", func(t *testing.T) {

	// 	width := 4
	// 	val := float64(12)

	// 	matrix := NewSquareMatrix(int64(width))

	// 	matrix.Set(2, 3, val)

	// 	assert.Equal(t, matrix.data[2][3], val)

	// })

}
