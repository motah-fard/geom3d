package geom3d

// AABB represents an axis-aligned bounding box in 3D.
type AABB struct {
	Min Vec3
	Max Vec3
}

// IsValid reports whether Min <= Max in all coordinates.
func (b AABB) IsValid() bool {
	return b.Min.X <= b.Max.X &&
		b.Min.Y <= b.Max.Y &&
		b.Min.Z <= b.Max.Z
}

// Size returns the box dimensions Max - Min.
func (b AABB) Size() Vec3 {
	return b.Max.Sub(b.Min)
}

// Center returns the center point of the box.
func (b AABB) Center() Vec3 {
	return Vec3{
		X: (b.Min.X + b.Max.X) / 2,
		Y: (b.Min.Y + b.Max.Y) / 2,
		Z: (b.Min.Z + b.Max.Z) / 2,
	}
}

// Contains reports whether p lies inside or on the boundary of the box.
func (b AABB) Contains(p Vec3) bool {
	return p.X >= b.Min.X && p.X <= b.Max.X &&
		p.Y >= b.Min.Y && p.Y <= b.Max.Y &&
		p.Z >= b.Min.Z && p.Z <= b.Max.Z
}

// Overlaps reports whether b and other overlap or touch.
func (b AABB) Overlaps(other AABB) bool {
	return b.Min.X <= other.Max.X && b.Max.X >= other.Min.X &&
		b.Min.Y <= other.Max.Y && b.Max.Y >= other.Min.Y &&
		b.Min.Z <= other.Max.Z && b.Max.Z >= other.Min.Z
}
