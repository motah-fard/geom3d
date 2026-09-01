package geom3d

import (
	"math"
	"testing"
)

// These package-level sinks prevent the compiler from optimizing away the
// benchmarked calls as dead code, since their results are otherwise
// discarded.
var (
	vec3Sink  Vec3
	boolSink  bool
	floatSink float64
	mat3Sink  Mat3
	quatSink  Quaternion
)

func BenchmarkVec3Normalize(b *testing.B) {
	v := Vec3{3, 4, 12}
	for i := 0; i < b.N; i++ {
		vec3Sink = v.Normalize()
	}
}

func BenchmarkVec3Cross(b *testing.B) {
	a := Vec3{1, 2, 3}
	v := Vec3{4, 5, 6}
	for i := 0; i < b.N; i++ {
		vec3Sink = a.Cross(v)
	}
}

func BenchmarkClosestPointOnSegment(b *testing.B) {
	seg := Segment3{A: Vec3{0, 0, 0}, B: Vec3{4, 0, 0}}
	p := Vec3{2, 3, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vec3Sink = ClosestPointOnSegment(p, seg)
	}
}

func BenchmarkClosestPointsBetweenSegments(b *testing.B) {
	s1 := Segment3{A: Vec3{0, 0, 0}, B: Vec3{2, 0, 0}}
	s2 := Segment3{A: Vec3{1, 1, 1}, B: Vec3{1, -1, 1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vec3Sink, vec3Sink = ClosestPointsBetweenSegments(s1, s2)
	}
}

func BenchmarkClosestPointOnTriangle(b *testing.B) {
	tri := Triangle{A: Vec3{0, 0, 0}, B: Vec3{4, 0, 0}, C: Vec3{0, 4, 0}}
	p := Vec3{1, 1, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vec3Sink = ClosestPointOnTriangle(p, tri)
	}
}

func BenchmarkIntersectRayTriangle(b *testing.B) {
	tri := Triangle{A: Vec3{0, 0, 0}, B: Vec3{4, 0, 0}, C: Vec3{0, 4, 0}}
	r := Ray3{Origin: Vec3{1, 1, 5}, Dir: Vec3{0, 0, -1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vec3Sink, boolSink = IntersectRayTriangle(r, tri)
	}
}

func BenchmarkIntersectRayAABB(b *testing.B) {
	box := AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}
	r := Ray3{Origin: Vec3{-1, 0.5, 0.5}, Dir: Vec3{1, 0, 0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boolSink, floatSink, floatSink = IntersectRayAABB(r, box)
	}
}

func BenchmarkIntersectRaySphere(b *testing.B) {
	s := Sphere{Center: Vec3{0, 0, 0}, Radius: 2}
	r := Ray3{Origin: Vec3{-5, 0, 0}, Dir: Vec3{1, 0, 0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boolSink, floatSink, floatSink = IntersectRaySphere(r, s)
	}
}

func BenchmarkIntersectRayOBB(b *testing.B) {
	box := OBB{Center: Vec3{0, 0, 0}, HalfExtents: Vec3{2, 1, 1}, Orientation: RotationZ(math.Pi / 4)}
	r := Ray3{Origin: Vec3{-5, 0, 0}, Dir: Vec3{1, 0, 0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boolSink, floatSink, floatSink = IntersectRayOBB(r, box)
	}
}

func BenchmarkIntersectRayCapsule(b *testing.B) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	r := Ray3{Origin: Vec3{-5, 0, 2}, Dir: Vec3{1, 0, 0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boolSink, floatSink, floatSink = IntersectRayCapsule(r, c)
	}
}

func BenchmarkTriangleOverlapsNonCoplanar(b *testing.B) {
	t1 := Triangle{A: Vec3{0, 0, 0}, B: Vec3{4, 0, 0}, C: Vec3{0, 4, 0}}
	t2 := Triangle{A: Vec3{1, 1, -5}, B: Vec3{1, 1, 5}, C: Vec3{3, 1, 5}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boolSink = t1.Overlaps(t2)
	}
}

func BenchmarkTriangleOverlapsCoplanar(b *testing.B) {
	// The coplanar case is the more expensive path (edge-crossing tests
	// all fail first, then falls through to the 2D SAT test).
	t1 := Triangle{A: Vec3{0, 0, 0}, B: Vec3{4, 0, 0}, C: Vec3{0, 4, 0}}
	t2 := Triangle{A: Vec3{1, 1, 0}, B: Vec3{5, 1, 0}, C: Vec3{1, 5, 0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boolSink = t1.Overlaps(t2)
	}
}

func BenchmarkMat3Mul(b *testing.B) {
	m := RotationX(0.4)
	n := RotationY(0.9)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mat3Sink = m.Mul(n)
	}
}

func BenchmarkMat3Inverse(b *testing.B) {
	m := RotationX(0.4).Mul(RotationY(0.9)).Mul(RotationZ(1.3))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mat3Sink, boolSink = m.Inverse()
	}
}

func BenchmarkQuaternionMul(b *testing.B) {
	q1 := QuaternionFromAxisAngle(Vec3{0, 0, 1}, math.Pi/2)
	q2 := QuaternionFromAxisAngle(Vec3{1, 0, 0}, math.Pi/2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quatSink = q1.Mul(q2)
	}
}

func BenchmarkQuaternionRotateVector(b *testing.B) {
	q := QuaternionFromAxisAngle(Vec3{1, 2, 3}, 0.9)
	v := Vec3{4, -1, 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vec3Sink = q.RotateVector(v)
	}
}

func BenchmarkQuaternionSlerp(b *testing.B) {
	q1 := IdentityQuaternion()
	q2 := QuaternionFromAxisAngle(Vec3{0, 1, 0}, math.Pi/2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quatSink = q1.Slerp(q2, 0.5)
	}
}
