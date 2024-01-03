package raytracer

import (
	"math"
)

type Matrix4 struct {
	data [4][4]float64
}

func NewIdentityMatrix4() *Matrix4 {
	return &Matrix4{
		data: [4][4]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
}

func NewMatrix4() *Matrix4 {
	return &Matrix4{}
}

func (m *Matrix4) Initialize(data [4][4]float64) {
	m.data = data
}

func (m *Matrix4) Set(row, col int64, val float64) {
	m.data[row][col] = val
}

func (m *Matrix4) Get(row, col int64) float64 {
	return m.data[row][col]
}

func (m *Matrix4) Equals(n *Matrix4) bool {

	for i := 0; i < 4; i++ {

		for j := 0; j < 4; j++ {

			if m.data[i][j] != n.data[i][j] {
				return false
			}

		}

	}
	return true

}

func (m *Matrix4) Multiply(n *Matrix4) *Matrix4 {

	res := &Matrix4{}

	for row := 0; row < 4; row++ {

		for col := 0; col < 4; col++ {

			total := float64(0)
			for k := 0; k < 4; k++ {
				total = total + m.data[row][k]*n.data[k][col]
			}
			res.data[row][col] = total
		}

	}

	return res

}

func (m *Matrix4) Transpose() *Matrix4 {

	n := &Matrix4{}

	for i := 0; i < 4; i++ {

		for j := 0; j < 4; j++ {

			n.data[j][i] = m.data[i][j]

		}

	}

	return n

}

func (m *Matrix4) SubMatrix(col, row int64) *Matrix3 {
	res := &Matrix3{
		data: [3][3]float64{},
	}

	resIndex := 0

	for i := int64(0); i < 4; i++ {
		if i == col {
			continue
		}

		for j := int64(0); j < 4; j++ {
			if j == row {
				continue
			}

			res.data[resIndex/3][resIndex%3] = m.data[i][j]
			resIndex++
		}
	}

	return res
}

func (m *Matrix4) Minor(col, row int64) float64 {

	return m.SubMatrix(col, row).Determinant()

}

func (m *Matrix4) Cofactor(col, row int64) float64 {

	minor := m.Minor(col, row)
	if (row+col)%2 != 0 {
		return -minor
	}
	return minor

}

func (m *Matrix4) Determinant() float64 {

	a1 := m.Cofactor(0, 0)
	a2 := m.Cofactor(0, 1)
	a3 := m.Cofactor(0, 2)
	a4 := m.Cofactor(0, 3)

	return m.data[0][0]*a1 + m.data[0][1]*a2 + m.data[0][2]*a3 + m.data[0][3]*a4

}

func (m *Matrix4) ApplyPoint(p *Point) *Point {

	newPoint := &Point{}

	newPoint.X = m.data[0][0]*p.X + m.data[0][1]*p.Y + m.data[0][2]*p.Z + m.data[0][3]*p.W
	newPoint.Y = m.data[1][0]*p.X + m.data[1][1]*p.Y + m.data[1][2]*p.Z + m.data[1][3]*p.W
	newPoint.Z = m.data[2][0]*p.X + m.data[2][1]*p.Y + m.data[2][2]*p.Z + m.data[2][3]*p.W
	newPoint.W = m.data[3][0]*p.X + m.data[3][1]*p.Y + m.data[3][2]*p.Z + m.data[3][3]*p.W

	return newPoint

}

func (m *Matrix4) ApplyVector(v *Vector) *Vector {
	newVector := &Vector{}

	newVector.X = m.data[0][0]*v.X + m.data[0][1]*v.Y + m.data[0][2]*v.Z + m.data[0][3]*v.W
	newVector.Y = m.data[1][0]*v.X + m.data[1][1]*v.Y + m.data[1][2]*v.Z + m.data[1][3]*v.W
	newVector.Z = m.data[2][0]*v.X + m.data[2][1]*v.Y + m.data[2][2]*v.Z + m.data[2][3]*v.W
	newVector.W = m.data[3][0]*v.X + m.data[3][1]*v.Y + m.data[3][2]*v.Z + m.data[3][3]*v.W

	return newVector

}

func NewTranslationMatrix4(x, y, z float64) *Matrix4 {

	idm := NewIdentityMatrix4()

	idm.data[0][3] = x
	idm.data[1][3] = y
	idm.data[2][3] = z

	return idm

}

func NewScalingMatrix4(x, y, z float64) *Matrix4 {

	idm := NewIdentityMatrix4()

	idm.data[0][0] = x
	idm.data[1][1] = y
	idm.data[2][2] = z

	return idm

}

func NewRotationXMatrix(r float64) *Matrix4 {
	idm := NewIdentityMatrix4()
	idm.data[1][1] = math.Cos(r)
	idm.data[1][2] = -math.Sin(r)
	idm.data[2][1] = math.Sin(r)
	idm.data[2][2] = math.Cos(r)
	return idm
}

func NewRotationYMatrix(r float64) *Matrix4 {
	idm := NewIdentityMatrix4()
	idm.data[0][0] = math.Cos(r)
	idm.data[0][2] = math.Sin(r)
	idm.data[2][0] = -math.Sin(r)
	idm.data[2][2] = math.Cos(r)
	return idm
}

func NewRotationZMatrix(r float64) *Matrix4 {
	idm := NewIdentityMatrix4()
	idm.data[0][0] = math.Cos(r)
	idm.data[0][1] = -math.Sin(r)
	idm.data[1][0] = math.Sin(r)
	idm.data[1][1] = math.Cos(r)
	return idm
}

func NewShearingMatrix(xY, xZ, yX, yZ, zX, zY float64) *Matrix4 {

	idm := NewIdentityMatrix4()

	idm.data[0][1] = xY
	idm.data[0][2] = xZ
	idm.data[1][0] = yX
	idm.data[1][2] = yZ
	idm.data[2][0] = zX
	idm.data[2][1] = zY

	return idm

}

func (m *Matrix4) Translate(x, y, z float64) *Matrix4 {

	tm := NewTranslationMatrix4(x, y, z)
	return tm.Multiply(m)

}

func (m *Matrix4) Scale(x, y, z float64) *Matrix4 {

	sm := NewScalingMatrix4(x, y, z)
	return sm.Multiply(m)

}

func (m *Matrix4) RotateX(r float64) *Matrix4 {

	rm := NewRotationXMatrix(r)
	return rm.Multiply(m)

}

func (m *Matrix4) RotateY(r float64) *Matrix4 {

	rm := NewRotationYMatrix(r)
	return rm.Multiply(m)

}

func (m *Matrix4) RotateZ(r float64) *Matrix4 {

	rm := NewRotationZMatrix(r)
	return rm.Multiply(m)

}

func (m *Matrix4) Shear(xY, xZ, yX, yZ, zX, zY float64) *Matrix4 {

	sm := NewShearingMatrix(xY, xZ, yX, yZ, zX, zY)
	return sm.Multiply(m)

}

func (m *Matrix4) IsInvertible() bool {

	return m.Determinant() != 0

}

func (m *Matrix4) Inverse() *Matrix4 {

	if !m.IsInvertible() {
		panic("cannot invert this matrix")
	}

	res := &Matrix4{}

	determinant := m.Determinant()

	for col := 0; col < 4; col++ {
		for row := 0; row < 4; row++ {

			cofactor := m.Cofactor(int64(col), int64(row))
			// indexing via [row][col] because of transpose
			res.data[row][col] = cofactor / determinant

		}
	}

	return res

}

type Matrix3 struct {
	data [3][3]float64
}

var IdentityMatrix3 = &Matrix3{
	data: [3][3]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	},
}

func NewMatrix3() *Matrix3 {
	return &Matrix3{}
}

func (m *Matrix3) Set(row, col int64, val float64) {
	m.data[row][col] = val
}

func (m *Matrix3) Get(row, col int64) float64 {
	return m.data[row][col]
}

func (m Matrix3) Equals(n *Matrix3) bool {

	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {

			if m.data[i][j] != n.data[i][j] {
				return false
			}

		}

	}
	return true

}

func (m *Matrix3) Transpose() *Matrix3 {

	n := &Matrix3{}

	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {

			n.data[j][i] = m.data[i][j]

		}

	}

	return n

}

func (m *Matrix3) SubMatrix(col, row int64) *Matrix2 {
	res := &Matrix2{}

	resIndex := 0

	for i := int64(0); i < 3; i++ {
		if i == col {
			continue
		}

		for j := int64(0); j < 3; j++ {
			if j == row {
				continue
			}

			res.data[resIndex/2][resIndex%2] = m.data[i][j]
			resIndex++
		}
	}

	return res
}

func (m *Matrix3) Minor(col, row int64) float64 {

	return m.SubMatrix(col, row).Determinant()

}

func (m *Matrix3) Cofactor(col, row int64) float64 {

	minor := m.Minor(col, row)
	if (row+col)%2 != 0 {
		return -minor
	}
	return minor

}

func (m *Matrix3) Determinant() float64 {

	a1 := m.Cofactor(0, 0)
	a2 := m.Cofactor(0, 1)
	a3 := m.Cofactor(0, 2)

	return m.data[0][0]*a1 + m.data[0][1]*a2 + m.data[0][2]*a3

}

type Matrix2 struct {
	data [2][2]float64
}

var IdentityMatrix2 = &Matrix2{
	data: [2][2]float64{
		{1, 0},
		{0, 1},
	},
}

func NewMatrix2() *Matrix2 {
	return &Matrix2{}
}

func (m *Matrix2) Set(col, row int64, val float64) {
	m.data[col][row] = val
}

func (m *Matrix2) Get(col, row int64) float64 {
	return m.data[col][row]
}

func (m Matrix2) Equals(n *Matrix2) bool {

	for i := 0; i < 2; i++ {

		for j := 0; j < 2; j++ {

			if m.data[i][j] != n.data[i][j] {
				return false
			}

		}

	}
	return true

}

func (m *Matrix2) Transpose() *Matrix2 {

	n := &Matrix2{}

	for i := 0; i < 2; i++ {

		for j := 0; j < 2; j++ {

			n.data[j][i] = m.data[i][j]

		}

	}

	return n

}

func (m *Matrix2) Determinant() float64 {
	return m.data[0][0]*m.data[1][1] - m.data[0][1]*m.data[1][0]
}
