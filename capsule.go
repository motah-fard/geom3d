package geom3d

import "math"

// Capsule represents a solid capsule in 3D: a cylinder of radius Radius
// with its core axis running from A to B, capped by two hemispheres of the
// same radius. Equivalently, it is the set of points within Radius of the
// segment from A to B.
type Capsule struct {
	A, B   Vec3
	Radius float64
}

// IsValid reports whether the capsule has a non-negative radius.
func (c Capsule) IsValid() bool {
	return c.Radius >= 0
}

// IsDegenerate reports whether the capsule's radius is effectively zero,
// making it equivalent to its core segment.
func (c Capsule) IsDegenerate() bool {
	return AlmostZero(c.Radius)
}

// Segment returns the capsule's core axis as a Segment3.
func (c Capsule) Segment() Segment3 {
	return Segment3{A: c.A, B: c.B}
}

// Volume returns the volume of the solid capsule: a cylinder plus a full
// sphere (the two end caps together form one sphere of radius Radius).
//
// If the capsule is invalid, it returns 0.
func (c Capsule) Volume() float64 {
	if !c.IsValid() {
		return 0
	}
	h := c.Segment().Length()
	return math.Pi*c.Radius*c.Radius*h + (4.0/3.0)*math.Pi*c.Radius*c.Radius*c.Radius
}

// SurfaceArea returns the surface area of the solid capsule: the cylinder's
// lateral surface plus a full sphere's surface (the two end caps together
// form one sphere of radius Radius).
//
// If the capsule is invalid, it returns 0.
func (c Capsule) SurfaceArea() float64 {
	if !c.IsValid() {
		return 0
	}
	h := c.Segment().Length()
	return 2*math.Pi*c.Radius*h + 4*math.Pi*c.Radius*c.Radius
}

// Contains reports whether p lies inside or on the surface of the capsule.
//
// If the capsule is invalid, it returns false.
func (c Capsule) Contains(p Vec3) bool {
	if !c.IsValid() {
		return false
	}
	return DistancePointToSegment(p, c.Segment()) <= c.Radius
}

// Overlaps reports whether c and other intersect or touch.
//
// Two capsules (swept spheres) overlap exactly when the distance between
// their core segments is at most the sum of their radii.
//
// If either capsule is invalid, it returns false.
func (c Capsule) Overlaps(other Capsule) bool {
	if !c.IsValid() || !other.IsValid() {
		return false
	}
	return DistanceBetweenSegments(c.Segment(), other.Segment()) <= c.Radius+other.Radius
}
