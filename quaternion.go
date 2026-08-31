package geom3d

import "math"

// Quaternion represents a rotation in 3D as (X, Y, Z, W), where (X, Y, Z) is
// the vector part and W is the scalar part.
//
// Quaternion avoids the gimbal lock of Euler angles and is cheaper to
// compose and interpolate than Mat3. For applying a rotation to many points,
// convert to a Mat3 with ToMat3 first.
type Quaternion struct {
	X, Y, Z, W float64
}

// IdentityQuaternion returns the quaternion representing no rotation.
func IdentityQuaternion() Quaternion {
	return Quaternion{X: 0, Y: 0, Z: 0, W: 1}
}

// QuaternionFromAxisAngle returns the quaternion representing a rotation of
// angle radians about axis.
//
// If axis is the zero vector, it returns IdentityQuaternion().
func QuaternionFromAxisAngle(axis Vec3, angle float64) Quaternion {
	n := axis.Normalize()
	if AlmostZero(n.Norm()) {
		return IdentityQuaternion()
	}

	half := angle / 2
	s := math.Sin(half)

	return Quaternion{
		X: n.X * s,
		Y: n.Y * s,
		Z: n.Z * s,
		W: math.Cos(half),
	}
}

// IsValid reports whether q has non-zero norm.
//
// The zero quaternion does not represent a rotation.
func (q Quaternion) IsValid() bool {
	return !AlmostZero(q.Norm())
}

// Dot returns the dot product of q and other, treating both as 4D vectors.
func (q Quaternion) Dot(other Quaternion) float64 {
	return q.X*other.X + q.Y*other.Y + q.Z*other.Z + q.W*other.W
}

// Norm2 returns the squared Euclidean norm of q.
func (q Quaternion) Norm2() float64 {
	return q.Dot(q)
}

// Norm returns the Euclidean norm of q.
func (q Quaternion) Norm() float64 {
	return math.Sqrt(q.Norm2())
}

// Normalize returns a unit quaternion in the same orientation as q.
//
// If q is the zero quaternion, it returns Quaternion{}.
func (q Quaternion) Normalize() Quaternion {
	n := q.Norm()
	if AlmostZero(n) {
		return Quaternion{}
	}
	inv := 1 / n
	return Quaternion{X: q.X * inv, Y: q.Y * inv, Z: q.Z * inv, W: q.W * inv}
}

// Conjugate returns the conjugate of q: the vector part negated, the scalar
// part unchanged.
//
// For a unit quaternion, the conjugate is also its inverse.
func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{X: -q.X, Y: -q.Y, Z: -q.Z, W: q.W}
}

// Inverse returns the multiplicative inverse of q, such that
// q.Mul(q.Inverse()) is (approximately) IdentityQuaternion().
//
// If q is the zero quaternion, it returns Quaternion{}.
func (q Quaternion) Inverse() Quaternion {
	n2 := q.Norm2()
	if AlmostZero(n2) {
		return Quaternion{}
	}
	c := q.Conjugate()
	inv := 1 / n2
	return Quaternion{X: c.X * inv, Y: c.Y * inv, Z: c.Z * inv, W: c.W * inv}
}

// Mul returns the Hamilton product q * other.
//
// As a rotation, q.Mul(other) applies other first, then q — the same
// composition order as Mat3.Mul and Transform.Compose.
func (q Quaternion) Mul(other Quaternion) Quaternion {
	return Quaternion{
		W: q.W*other.W - q.X*other.X - q.Y*other.Y - q.Z*other.Z,
		X: q.W*other.X + q.X*other.W + q.Y*other.Z - q.Z*other.Y,
		Y: q.W*other.Y - q.X*other.Z + q.Y*other.W + q.Z*other.X,
		Z: q.W*other.Z + q.X*other.Y - q.Y*other.X + q.Z*other.W,
	}
}

// RotateVector applies q's rotation to v.
//
// q does not need to be normalized first.
//
// If q is the zero quaternion, it returns Vec3{}.
func (q Quaternion) RotateVector(v Vec3) Vec3 {
	if !q.IsValid() {
		return Vec3{}
	}

	vq := Quaternion{X: v.X, Y: v.Y, Z: v.Z, W: 0}
	r := q.Mul(vq).Mul(q.Inverse())
	return Vec3{X: r.X, Y: r.Y, Z: r.Z}
}

// ToMat3 returns the rotation matrix equivalent to q.
//
// q does not need to be normalized first.
//
// If q is the zero quaternion, it returns IdentityMat3(), since a zero
// rotation matrix would silently collapse every point it's applied to onto
// the origin.
func (q Quaternion) ToMat3() Mat3 {
	n2 := q.Norm2()
	if AlmostZero(n2) {
		return IdentityMat3()
	}
	s := 2 / n2

	xx := s * q.X * q.X
	yy := s * q.Y * q.Y
	zz := s * q.Z * q.Z
	xy := s * q.X * q.Y
	xz := s * q.X * q.Z
	yz := s * q.Y * q.Z
	wx := s * q.W * q.X
	wy := s * q.W * q.Y
	wz := s * q.W * q.Z

	return Mat3{
		M: [3][3]float64{
			{1 - (yy + zz), xy - wz, xz + wy},
			{xy + wz, 1 - (xx + zz), yz - wx},
			{xz - wy, yz + wx, 1 - (xx + yy)},
		},
	}
}

// Slerp returns the spherical linear interpolation between q and other by
// fraction t, where t = 0 returns q (normalized) and t = 1 returns other
// (normalized, and on the same hemisphere as q).
//
// q and other do not need to be normalized first.
func (q Quaternion) Slerp(other Quaternion, t float64) Quaternion {
	q = q.Normalize()
	other = other.Normalize()

	dot := q.Dot(other)

	// Take the shorter path around the hypersphere.
	if dot < 0 {
		other = Quaternion{X: -other.X, Y: -other.Y, Z: -other.Z, W: -other.W}
		dot = -dot
	}

	const nearParallel = 1 - 1e-6
	if dot > nearParallel {
		// Nearly identical orientations: linear interpolation is a good
		// approximation and avoids dividing by a near-zero sine below.
		return Quaternion{
			X: q.X + (other.X-q.X)*t,
			Y: q.Y + (other.Y-q.Y)*t,
			Z: q.Z + (other.Z-q.Z)*t,
			W: q.W + (other.W-q.W)*t,
		}.Normalize()
	}

	theta := math.Acos(clamp(dot, -1, 1))
	sinTheta := math.Sin(theta)
	a := math.Sin((1-t)*theta) / sinTheta
	b := math.Sin(t*theta) / sinTheta

	return Quaternion{
		X: q.X*a + other.X*b,
		Y: q.Y*a + other.Y*b,
		Z: q.Z*a + other.Z*b,
		W: q.W*a + other.W*b,
	}
}
