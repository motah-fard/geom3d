// Package vsmathgl benchmarks the small set of operations that genuinely
// overlap between geom3d and go-gl/mathgl (mgl64): basic Vec3 ops, Mat3
// multiplication, and Quaternion multiply/rotate/slerp. It lives in its own
// module specifically so this comparison — and its mathgl dependency — never
// leaks into geom3d's own go.mod, which is intentionally dependency-free.
//
// geom3d's actual differentiator (ray/triangle, closest-point, distance
// queries, and the rest of internal/../*.go) has no mathgl equivalent, so
// it isn't benchmarked here — see benchmark_test.go in the main module for
// those numbers on their own. This file only compares the overlapping
// foundational math, and isn't a claim that one library is "faster overall."
//
// Run with:
//
//	go test -bench=. -benchmem ./...
package vsmathgl

import (
	"math"
	"testing"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/motah-fard/geom3d"
)

var (
	geom3dVec3Sink geom3d.Vec3
	mathglVec3Sink mgl64.Vec3
	geom3dFloat    float64
	mathglFloat    float64
	geom3dMat3     geom3d.Mat3
	mathglMat3     mgl64.Mat3
	geom3dQuat     geom3d.Quaternion
	mathglQuat     mgl64.Quat
)

func BenchmarkGeom3d_Vec3Normalize(b *testing.B) {
	v := geom3d.Vec3{X: 3, Y: 4, Z: 12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		geom3dVec3Sink = v.Normalize()
	}
}

func BenchmarkMathgl_Vec3Normalize(b *testing.B) {
	v := mgl64.Vec3{3, 4, 12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mathglVec3Sink = v.Normalize()
	}
}

func BenchmarkGeom3d_Vec3Cross(b *testing.B) {
	a := geom3d.Vec3{X: 1, Y: 2, Z: 3}
	v := geom3d.Vec3{X: 4, Y: 5, Z: 6}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		geom3dVec3Sink = a.Cross(v)
	}
}

func BenchmarkMathgl_Vec3Cross(b *testing.B) {
	a := mgl64.Vec3{1, 2, 3}
	v := mgl64.Vec3{4, 5, 6}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mathglVec3Sink = a.Cross(v)
	}
}

func BenchmarkGeom3d_Vec3Dot(b *testing.B) {
	a := geom3d.Vec3{X: 1, Y: 2, Z: 3}
	v := geom3d.Vec3{X: 4, Y: 5, Z: 6}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		geom3dFloat = a.Dot(v)
	}
}

func BenchmarkMathgl_Vec3Dot(b *testing.B) {
	a := mgl64.Vec3{1, 2, 3}
	v := mgl64.Vec3{4, 5, 6}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mathglFloat = a.Dot(v)
	}
}

func BenchmarkGeom3d_Mat3Mul(b *testing.B) {
	m := geom3d.RotationX(0.4)
	n := geom3d.RotationY(0.9)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		geom3dMat3 = m.Mul(n)
	}
}

func BenchmarkMathgl_Mat3Mul(b *testing.B) {
	m := mgl64.Rotate3DX(0.4)
	n := mgl64.Rotate3DY(0.9)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mathglMat3 = m.Mul3(n)
	}
}

func BenchmarkGeom3d_QuaternionMul(b *testing.B) {
	q1 := geom3d.QuaternionFromAxisAngle(geom3d.Vec3{X: 0, Y: 0, Z: 1}, math.Pi/2)
	q2 := geom3d.QuaternionFromAxisAngle(geom3d.Vec3{X: 1, Y: 0, Z: 0}, math.Pi/2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		geom3dQuat = q1.Mul(q2)
	}
}

func BenchmarkMathgl_QuaternionMul(b *testing.B) {
	q1 := mgl64.QuatRotate(math.Pi/2, mgl64.Vec3{0, 0, 1})
	q2 := mgl64.QuatRotate(math.Pi/2, mgl64.Vec3{1, 0, 0})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mathglQuat = q1.Mul(q2)
	}
}

func BenchmarkGeom3d_QuaternionRotateVector(b *testing.B) {
	q := geom3d.QuaternionFromAxisAngle(geom3d.Vec3{X: 1, Y: 2, Z: 3}, 0.9)
	v := geom3d.Vec3{X: 4, Y: -1, Z: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		geom3dVec3Sink = q.RotateVector(v)
	}
}

func BenchmarkMathgl_QuaternionRotateVector(b *testing.B) {
	q := mgl64.QuatRotate(0.9, mgl64.Vec3{1, 2, 3}.Normalize())
	v := mgl64.Vec3{4, -1, 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mathglVec3Sink = q.Rotate(v)
	}
}

func BenchmarkGeom3d_QuaternionSlerp(b *testing.B) {
	q1 := geom3d.IdentityQuaternion()
	q2 := geom3d.QuaternionFromAxisAngle(geom3d.Vec3{X: 0, Y: 1, Z: 0}, math.Pi/2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		geom3dQuat = q1.Slerp(q2, 0.5)
	}
}

func BenchmarkMathgl_QuaternionSlerp(b *testing.B) {
	q1 := mgl64.QuatIdent()
	q2 := mgl64.QuatRotate(math.Pi/2, mgl64.Vec3{0, 1, 0})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mathglQuat = mgl64.QuatSlerp(q1, q2, 0.5)
	}
}
