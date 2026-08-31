package geom3d

import "math"

// Sphere represents a solid ball in 3D, defined by a center point and a
// radius.
type Sphere struct {
	Center Vec3
	Radius float64
}

// IsValid reports whether the sphere has a non-negative radius.
func (s Sphere) IsValid() bool {
	return s.Radius >= 0
}

// IsDegenerate reports whether the sphere's radius is effectively zero,
// making it equivalent to a single point.
func (s Sphere) IsDegenerate() bool {
	return AlmostZero(s.Radius)
}

// SurfaceArea returns the surface area of the sphere.
//
// If the sphere is invalid, it returns 0.
func (s Sphere) SurfaceArea() float64 {
	if !s.IsValid() {
		return 0
	}
	return 4 * math.Pi * s.Radius * s.Radius
}

// Volume returns the volume of the solid ball.
//
// If the sphere is invalid, it returns 0.
func (s Sphere) Volume() float64 {
	if !s.IsValid() {
		return 0
	}
	return (4.0 / 3.0) * math.Pi * s.Radius * s.Radius * s.Radius
}

// Contains reports whether p lies inside or on the surface of the sphere.
//
// If the sphere is invalid, it returns false.
func (s Sphere) Contains(p Vec3) bool {
	if !s.IsValid() {
		return false
	}
	return p.Distance2(s.Center) <= s.Radius*s.Radius
}
