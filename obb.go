package geom3d

import "math"

// OBB represents an oriented (rotated) bounding box in 3D: a box centered
// at Center, with half-extents HalfExtents along its own local axes, which
// are rotated relative to world space by Orientation.
//
// Orientation is assumed to be a proper (orthonormal) rotation matrix, such
// as one returned by RotationX/Y/Z or Quaternion.ToMat3.
type OBB struct {
	Center      Vec3
	HalfExtents Vec3
	Orientation Mat3
}

// IsValid reports whether all of b's half-extents are non-negative.
func (b OBB) IsValid() bool {
	return b.HalfExtents.X >= 0 && b.HalfExtents.Y >= 0 && b.HalfExtents.Z >= 0
}

// Volume returns the volume of the box.
//
// If the box is invalid, it returns 0.
func (b OBB) Volume() float64 {
	if !b.IsValid() {
		return 0
	}
	return 8 * b.HalfExtents.X * b.HalfExtents.Y * b.HalfExtents.Z
}

// SurfaceArea returns the total surface area of the box.
//
// If the box is invalid, it returns 0.
func (b OBB) SurfaceArea() float64 {
	if !b.IsValid() {
		return 0
	}
	h := b.HalfExtents
	return 8 * (h.X*h.Y + h.Y*h.Z + h.Z*h.X)
}

// Contains reports whether p lies inside or on the boundary of the box.
//
// If the box is invalid, it returns false.
func (b OBB) Contains(p Vec3) bool {
	if !b.IsValid() {
		return false
	}
	local := b.toLocal(p)
	return math.Abs(local.X) <= b.HalfExtents.X &&
		math.Abs(local.Y) <= b.HalfExtents.Y &&
		math.Abs(local.Z) <= b.HalfExtents.Z
}

// toLocal transforms a world-space point into the box's local, axis-aligned
// frame, assuming Orientation is orthonormal (so its transpose is its
// inverse — the same convention Transform.Inverse relies on).
func (b OBB) toLocal(p Vec3) Vec3 {
	return b.Orientation.Transpose().MulVec(p.Sub(b.Center))
}

// toWorld transforms a point from the box's local frame back into world
// space.
func (b OBB) toWorld(local Vec3) Vec3 {
	return b.Orientation.MulVec(local).Add(b.Center)
}

// ClosestPointOnOBB returns the closest point on or in solid box b to
// point p.
//
// If p lies inside the box, it returns p.
//
// If the box is invalid, it returns Vec3{}.
func ClosestPointOnOBB(p Vec3, b OBB) Vec3 {
	if !b.IsValid() {
		return Vec3{}
	}

	local := b.toLocal(p)
	clamped := Vec3{
		X: clamp(local.X, -b.HalfExtents.X, b.HalfExtents.X),
		Y: clamp(local.Y, -b.HalfExtents.Y, b.HalfExtents.Y),
		Z: clamp(local.Z, -b.HalfExtents.Z, b.HalfExtents.Z),
	}

	return b.toWorld(clamped)
}
