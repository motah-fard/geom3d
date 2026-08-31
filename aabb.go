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
	return b.Min.Midpoint(b.Max)
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

// Volume returns the volume of the box.
//
// If the box is invalid, it returns 0.
func (b AABB) Volume() float64 {
	if !b.IsValid() {
		return 0
	}
	size := b.Size()
	return size.X * size.Y * size.Z
}

// SurfaceArea returns the total surface area of the box.
//
// If the box is invalid, it returns 0.
func (b AABB) SurfaceArea() float64 {
	if !b.IsValid() {
		return 0
	}
	size := b.Size()
	return 2 * (size.X*size.Y + size.Y*size.Z + size.Z*size.X)
}

// Union returns the smallest AABB that contains both b and other.
func (b AABB) Union(other AABB) AABB {
	return AABB{
		Min: b.Min.Min(other.Min),
		Max: b.Max.Max(other.Max),
	}
}

// ExpandToInclude returns the smallest AABB that contains both b and p.
func (b AABB) ExpandToInclude(p Vec3) AABB {
	return AABB{
		Min: b.Min.Min(p),
		Max: b.Max.Max(p),
	}
}

// Expand returns a copy of b grown outward by margin on every side.
//
// A negative margin shrinks the box, which can produce an invalid AABB if
// the margin exceeds half the box's size along an axis.
func (b AABB) Expand(margin float64) AABB {
	m := Vec3{X: margin, Y: margin, Z: margin}
	return AABB{
		Min: b.Min.Sub(m),
		Max: b.Max.Add(m),
	}
}

// AABBFromPoints returns the smallest AABB containing every point in points.
//
// If points is empty, it returns AABB{}.
func AABBFromPoints(points []Vec3) AABB {
	if len(points) == 0 {
		return AABB{}
	}

	box := AABB{Min: points[0], Max: points[0]}
	for _, p := range points[1:] {
		box = box.ExpandToInclude(p)
	}
	return box
}
