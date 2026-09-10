package geom3d

import (
	"math"
)

// Mat3 represents a 3×3 matrix used for linear transformations in 3D.
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

// MulVec returns the matrix-vector product m * v, applying v first and then m
// when used with column vectors.
func (m Mat3) MulVec(v Vec3) Vec3 {
	return Vec3{
		X: m.M[0][0]*v.X + m.M[0][1]*v.Y + m.M[0][2]*v.Z,
		Y: m.M[1][0]*v.X + m.M[1][1]*v.Y + m.M[1][2]*v.Z,
		Z: m.M[2][0]*v.X + m.M[2][1]*v.Y + m.M[2][2]*v.Z,
	}
}

// Mul returns the matrix product m * n.
//
// Hand-unrolled rather than a triple-nested loop over M — benchmarked at
// roughly 2.8x faster (see BENCHMARKS.md), closing most of the gap to
// go-gl/mathgl's flat-array Mul3 despite Mat3 keeping its [3][3]float64
// shape. Same values in, same values out; only the arithmetic's shape
// changed.
func (m Mat3) Mul(n Mat3) Mat3 {
	return Mat3{M: [3][3]float64{
		{
			m.M[0][0]*n.M[0][0] + m.M[0][1]*n.M[1][0] + m.M[0][2]*n.M[2][0],
			m.M[0][0]*n.M[0][1] + m.M[0][1]*n.M[1][1] + m.M[0][2]*n.M[2][1],
			m.M[0][0]*n.M[0][2] + m.M[0][1]*n.M[1][2] + m.M[0][2]*n.M[2][2],
		},
		{
			m.M[1][0]*n.M[0][0] + m.M[1][1]*n.M[1][0] + m.M[1][2]*n.M[2][0],
			m.M[1][0]*n.M[0][1] + m.M[1][1]*n.M[1][1] + m.M[1][2]*n.M[2][1],
			m.M[1][0]*n.M[0][2] + m.M[1][1]*n.M[1][2] + m.M[1][2]*n.M[2][2],
		},
		{
			m.M[2][0]*n.M[0][0] + m.M[2][1]*n.M[1][0] + m.M[2][2]*n.M[2][0],
			m.M[2][0]*n.M[0][1] + m.M[2][1]*n.M[1][1] + m.M[2][2]*n.M[2][1],
			m.M[2][0]*n.M[0][2] + m.M[2][1]*n.M[1][2] + m.M[2][2]*n.M[2][2],
		},
	}}
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

// Determinant returns the determinant of m.
func (m Mat3) Determinant() float64 {
	return m.M[0][0]*(m.M[1][1]*m.M[2][2]-m.M[1][2]*m.M[2][1]) -
		m.M[0][1]*(m.M[1][0]*m.M[2][2]-m.M[1][2]*m.M[2][0]) +
		m.M[0][2]*(m.M[1][0]*m.M[2][1]-m.M[1][1]*m.M[2][0])
}

// Inverse returns the inverse of m and true, for any m with a non-zero
// determinant.
//
// For an orthonormal rotation matrix (such as one returned by RotationX,
// RotationY, RotationZ, or Quaternion.ToMat3), Transpose is equivalent and
// cheaper to compute; Inverse handles the general case.
//
// If m's determinant is zero (m is singular, and therefore has no inverse),
// it returns Mat3{} and false. This is a property of the specific matrix
// value, not invalid input.
func (m Mat3) Inverse() (Mat3, bool) {
	det := m.Determinant()
	if AlmostZero(det) {
		return Mat3{}, false
	}
	invDet := 1 / det

	return Mat3{
		M: [3][3]float64{
			{
				(m.M[1][1]*m.M[2][2] - m.M[1][2]*m.M[2][1]) * invDet,
				(m.M[0][2]*m.M[2][1] - m.M[0][1]*m.M[2][2]) * invDet,
				(m.M[0][1]*m.M[1][2] - m.M[0][2]*m.M[1][1]) * invDet,
			},
			{
				(m.M[1][2]*m.M[2][0] - m.M[1][0]*m.M[2][2]) * invDet,
				(m.M[0][0]*m.M[2][2] - m.M[0][2]*m.M[2][0]) * invDet,
				(m.M[0][2]*m.M[1][0] - m.M[0][0]*m.M[1][2]) * invDet,
			},
			{
				(m.M[1][0]*m.M[2][1] - m.M[1][1]*m.M[2][0]) * invDet,
				(m.M[0][1]*m.M[2][0] - m.M[0][0]*m.M[2][1]) * invDet,
				(m.M[0][0]*m.M[1][1] - m.M[0][1]*m.M[1][0]) * invDet,
			},
		},
	}, true
}

// ToQuaternion returns the rotation quaternion equivalent to m.
//
// m is assumed to be a proper (orthonormal, right-handed) rotation matrix;
// for any other matrix the result is not meaningful.
func (m Mat3) ToQuaternion() Quaternion {
	trace := m.M[0][0] + m.M[1][1] + m.M[2][2]

	var q Quaternion
	switch {
	case trace > 0:
		s := 0.5 / math.Sqrt(trace+1)
		q.W = 0.25 / s
		q.X = (m.M[2][1] - m.M[1][2]) * s
		q.Y = (m.M[0][2] - m.M[2][0]) * s
		q.Z = (m.M[1][0] - m.M[0][1]) * s
	case m.M[0][0] > m.M[1][1] && m.M[0][0] > m.M[2][2]:
		s := 2 * math.Sqrt(1+m.M[0][0]-m.M[1][1]-m.M[2][2])
		q.W = (m.M[2][1] - m.M[1][2]) / s
		q.X = 0.25 * s
		q.Y = (m.M[0][1] + m.M[1][0]) / s
		q.Z = (m.M[0][2] + m.M[2][0]) / s
	case m.M[1][1] > m.M[2][2]:
		s := 2 * math.Sqrt(1+m.M[1][1]-m.M[0][0]-m.M[2][2])
		q.W = (m.M[0][2] - m.M[2][0]) / s
		q.X = (m.M[0][1] + m.M[1][0]) / s
		q.Y = 0.25 * s
		q.Z = (m.M[1][2] + m.M[2][1]) / s
	default:
		s := 2 * math.Sqrt(1+m.M[2][2]-m.M[0][0]-m.M[1][1])
		q.W = (m.M[1][0] - m.M[0][1]) / s
		q.X = (m.M[0][2] + m.M[2][0]) / s
		q.Y = (m.M[1][2] + m.M[2][1]) / s
		q.Z = 0.25 * s
	}

	return q
}
