package geom3d

import "math"

// Vec3 represents a 3D vector or point in Cartesian coordinates.
type Vec3 struct {
	X, Y, Z float64
}

// Add returns the vector sum a + b.
func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

// Sub returns the vector difference a - b.
func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

// Scale returns the vector a multiplied by scalar s.
func (a Vec3) Scale(s float64) Vec3 {
	return Vec3{
		X: a.X * s,
		Y: a.Y * s,
		Z: a.Z * s,
	}
}

// Dot returns the dot product of a and b.
func (a Vec3) Dot(b Vec3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Cross returns the cross product of a and b.
func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

// Norm2 returns the squared Euclidean norm of a.
func (a Vec3) Norm2() float64 {
	return a.Dot(a)
}

// Norm returns the Euclidean norm of a.
func (a Vec3) Norm() float64 {
	return math.Sqrt(a.Norm2())
}

// Distance2 returns the squared Euclidean distance between a and b.
func (a Vec3) Distance2(b Vec3) float64 {
	return a.Sub(b).Norm2()
}

// Distance returns the Euclidean distance between a and b.
func (a Vec3) Distance(b Vec3) float64 {
	return math.Sqrt(a.Distance2(b))
}

// Midpoint returns the point halfway between a and b.
func (a Vec3) Midpoint(b Vec3) Vec3 {
	return a.Add(b).Scale(0.5)
}

// Normalize returns a unit vector in the same direction as a.
// If a is the zero vector, it returns Vec3{}.
func (a Vec3) Normalize() Vec3 {
	n := a.Norm()
	if AlmostZero(n) {
		return Vec3{}
	}
	return a.Scale(1 / n)
}

// Lerp returns the point that is linearly interpolated between a and b by
// fraction t, where t = 0 returns a and t = 1 returns b.
//
// t is not clamped to [0, 1]; values outside that range extrapolate.
func (a Vec3) Lerp(b Vec3, t float64) Vec3 {
	return a.Add(b.Sub(a).Scale(t))
}

// Reflect returns a reflected about the plane through the origin with unit
// normal n, as if n were a surface normal and a an incoming direction.
//
// n is assumed to be normalized; if it is not, the result is scaled
// accordingly.
func (a Vec3) Reflect(n Vec3) Vec3 {
	return a.Sub(n.Scale(2 * a.Dot(n)))
}

// Project returns the vector projection of a onto b: the component of a
// that lies in the direction of b.
//
// If b is the zero vector, it returns Vec3{}.
func (a Vec3) Project(b Vec3) Vec3 {
	denom := b.Norm2()
	if AlmostZero(denom) {
		return Vec3{}
	}
	return b.Scale(a.Dot(b) / denom)
}

// Angle returns the unsigned angle, in radians, between a and b.
//
// If either vector is the zero vector, it returns 0.
func (a Vec3) Angle(b Vec3) float64 {
	denom := a.Norm() * b.Norm()
	if AlmostZero(denom) {
		return 0
	}
	cos := a.Dot(b) / denom
	// Guard against floating-point drift pushing cos slightly outside
	// [-1, 1], which would make math.Acos return NaN.
	cos = clamp(cos, -1, 1)
	return math.Acos(cos)
}

// ClampLength returns a scaled down to have norm at most max.
//
// If a's norm is already less than or equal to max, it returns a unchanged.
func (a Vec3) ClampLength(max float64) Vec3 {
	n := a.Norm()
	if n <= max {
		return a
	}
	return a.Scale(max / n)
}

// Abs returns the component-wise absolute value of a.
func (a Vec3) Abs() Vec3 {
	return Vec3{
		X: math.Abs(a.X),
		Y: math.Abs(a.Y),
		Z: math.Abs(a.Z),
	}
}

// Min returns the component-wise minimum of a and b.
func (a Vec3) Min(b Vec3) Vec3 {
	return Vec3{
		X: math.Min(a.X, b.X),
		Y: math.Min(a.Y, b.Y),
		Z: math.Min(a.Z, b.Z),
	}
}

// Max returns the component-wise maximum of a and b.
func (a Vec3) Max(b Vec3) Vec3 {
	return Vec3{
		X: math.Max(a.X, b.X),
		Y: math.Max(a.Y, b.Y),
		Z: math.Max(a.Z, b.Z),
	}
}
