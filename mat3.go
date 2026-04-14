package geom3d

import "math"

// Mat3 represents a 3x3 matrix.
type Mat3 struct {
	M [3][3]float64
}

// IdentityMat3 returns the 3x3 identity matrix.
func IdentityMat3() Mat3 {
	return Mat3{
		M: [3][3]float64{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	}
}

// RotationX returns a rotation matrix for a rotation of theta radians about the X axis.
func RotationX(theta float64) Mat3 {
	c := math.Cos(theta)
	s := math.Sin(theta)

	return Mat3{
		M: [3][3]float64{
			{1, 0, 0},
			{0, c, -s},
			{0, s, c},
		},
	}
}

// RotationY returns a rotation matrix for a rotation of theta radians about the Y axis.
func RotationY(theta float64) Mat3 {
	c := math.Cos(theta)
	s := math.Sin(theta)

	return Mat3{
		M: [3][3]float64{
			{c, 0, s},
			{0, 1, 0},
			{-s, 0, c},
		},
	}
}

// RotationZ returns a rotation matrix for a rotation of theta radians about the Z axis.
func RotationZ(theta float64) Mat3 {
	c := math.Cos(theta)
	s := math.Sin(theta)

	return Mat3{
		M: [3][3]float64{
			{c, -s, 0},
			{s, c, 0},
			{0, 0, 1},
		},
	}
}

// MulVec returns the matrix-vector product m * v.
func (m Mat3) MulVec(v Vec3) Vec3 {
	return Vec3{
		X: m.M[0][0]*v.X + m.M[0][1]*v.Y + m.M[0][2]*v.Z,
		Y: m.M[1][0]*v.X + m.M[1][1]*v.Y + m.M[1][2]*v.Z,
		Z: m.M[2][0]*v.X + m.M[2][1]*v.Y + m.M[2][2]*v.Z,
	}
}

// Mul returns the matrix product m * n.
func (m Mat3) Mul(n Mat3) Mat3 {
	var out Mat3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				out.M[i][j] += m.M[i][k] * n.M[k][j]
			}
		}
	}

	return out
}

// Transpose returns the transpose of m.
func (m Mat3) Transpose() Mat3 {
	return Mat3{
		M: [3][3]float64{
			{m.M[0][0], m.M[1][0], m.M[2][0]},
			{m.M[0][1], m.M[1][1], m.M[2][1]},
			{m.M[0][2], m.M[1][2], m.M[2][2]},
		},
	}
}
