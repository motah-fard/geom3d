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

// Normalize returns a unit vector in the same direction as a.
// If a is the zero vector, it returns Vec3{}.
func (a Vec3) Normalize() Vec3 {
	n := a.Norm()
	if AlmostZero(n) {
		return Vec3{}
	}
	return a.Scale(1 / n)
}
