package geom3d

import (
	"math"
	"testing"
)

// isFiniteAndBounded rejects NaN, Inf, and values whose magnitude is large
// enough that ordinary float64 precision loss (not a bug) would make
// exact-ish invariant checks unreliable.
func isFiniteAndBounded(vs ...float64) bool {
	const bound = 1e6
	for _, v := range vs {
		if math.IsNaN(v) || math.IsInf(v, 0) || math.Abs(v) > bound {
			return false
		}
	}
	return true
}

// FuzzTriangleOverlaps checks that Triangle.Overlaps — the most
// algorithmically involved function in the package, combining edge-crossing
// tests with a 2D separating-axis fallback for the coplanar case — is
// symmetric for arbitrary finite, bounded triangles.
func FuzzTriangleOverlaps(f *testing.F) {
	f.Add(0.0, 0.0, 0.0, 4.0, 0.0, 0.0, 0.0, 4.0, 0.0, 1.0, 1.0, 0.0, 5.0, 1.0, 0.0, 1.0, 5.0, 0.0)
	f.Add(0.0, 0.0, 0.0, 4.0, 0.0, 0.0, 0.0, 4.0, 0.0, 100.0, 100.0, -5.0, 100.0, 100.0, 5.0, 103.0, 100.0, 5.0)
	f.Add(0.0, 0.0, 0.0, 4.0, 0.0, 0.0, 0.0, 4.0, 0.0, 1.0, 1.0, -5.0, 1.0, 1.0, 5.0, 3.0, 1.0, 5.0)

	f.Fuzz(func(t *testing.T, ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz, ex, ey, ez, fx, fy, fz float64) {
		if !isFiniteAndBounded(ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz, ex, ey, ez, fx, fy, fz) {
			t.Skip()
		}

		t1 := Triangle{A: Vec3{ax, ay, az}, B: Vec3{bx, by, bz}, C: Vec3{cx, cy, cz}}
		t2 := Triangle{A: Vec3{dx, dy, dz}, B: Vec3{ex, ey, ez}, C: Vec3{fx, fy, fz}}

		got := t1.Overlaps(t2)
		symmetric := t2.Overlaps(t1)
		if got != symmetric {
			t.Fatalf("Overlaps is not symmetric: t1.Overlaps(t2)=%v, t2.Overlaps(t1)=%v", got, symmetric)
		}
	})
}

// FuzzMat3Inverse checks that Mat3.Inverse's ok flag is always consistent
// with Mat3.Determinant, and that well-conditioned matrices actually
// invert (m * m.Inverse() close to identity). Near-singular matrices are
// excluded from the accuracy check since they're inherently numerically
// unstable regardless of the inverse formula's correctness.
func FuzzMat3Inverse(f *testing.F) {
	f.Add(1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0)
	f.Add(2.0, 1.0, 0.0, 1.0, 3.0, 1.0, 0.0, 1.0, 2.0)
	f.Add(1.0, 2.0, 3.0, 2.0, 4.0, 6.0, 1.0, 1.0, 1.0) // singular

	f.Fuzz(func(t *testing.T, m00, m01, m02, m10, m11, m12, m20, m21, m22 float64) {
		if !isFiniteAndBounded(m00, m01, m02, m10, m11, m12, m20, m21, m22) {
			t.Skip()
		}

		m := Mat3{M: [3][3]float64{{m00, m01, m02}, {m10, m11, m12}, {m20, m21, m22}}}
		det := m.Determinant()
		inv, ok := m.Inverse()

		if ok == AlmostZero(det) {
			t.Fatalf("Inverse ok=%v inconsistent with Determinant=%v", ok, det)
		}

		if ok && math.Abs(det) > 1 {
			product := m.Mul(inv)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					want := 0.0
					if i == j {
						want = 1.0
					}
					if diff := math.Abs(product.M[i][j] - want); diff > 1e-6 {
						t.Fatalf("m * m.Inverse() not close to identity at (%d,%d): got %v, want %v (det=%v)", i, j, product.M[i][j], want, det)
					}
				}
			}
		}
	})
}

// FuzzQuaternionSlerp checks that Slerp never produces NaN or Inf for any
// pair of valid (non-zero) quaternions and any t, including t outside
// [0, 1] and quaternions on opposite hemispheres (dot < 0), which exercise
// Slerp's shortest-path and near-parallel branches.
func FuzzQuaternionSlerp(f *testing.F) {
	f.Add(0.0, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0, 0.0, 0.5)
	f.Add(1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.3)
	f.Add(0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, -1.0, 0.7) // opposite hemispheres

	f.Fuzz(func(t *testing.T, x, y, z, w, ox, oy, oz, ow, tt float64) {
		if !isFiniteAndBounded(x, y, z, w, ox, oy, oz, ow, tt) {
			t.Skip()
		}

		q := Quaternion{X: x, Y: y, Z: z, W: w}
		other := Quaternion{X: ox, Y: oy, Z: oz, W: ow}
		if !q.IsValid() || !other.IsValid() {
			t.Skip()
		}

		result := q.Slerp(other, tt)
		if math.IsNaN(result.X) || math.IsNaN(result.Y) || math.IsNaN(result.Z) || math.IsNaN(result.W) {
			t.Fatalf("Slerp produced NaN: q=%+v other=%+v t=%v result=%+v", q, other, tt, result)
		}
		if math.IsInf(result.X, 0) || math.IsInf(result.Y, 0) || math.IsInf(result.Z, 0) || math.IsInf(result.W, 0) {
			t.Fatalf("Slerp produced Inf: q=%+v other=%+v t=%v result=%+v", q, other, tt, result)
		}
	})
}

// FuzzIntersectRayCapsule cross-checks IntersectRayCapsule — the most
// intricate intersection function in the package — against the
// independently implemented DistancePointToCapsule: whenever it reports a
// hit, the entry and exit points it returns must actually lie on the
// capsule's surface.
func FuzzIntersectRayCapsule(f *testing.F) {
	f.Add(0.0, 0.0, 0.0, 0.0, 0.0, 4.0, 1.0, -5.0, 0.0, 2.0, 1.0, 0.0, 0.0)
	f.Add(0.0, 0.0, 0.0, 0.0, 0.0, 4.0, 1.0, 0.0, 0.0, 10.0, 0.0, 0.0, -1.0)
	f.Add(1.0, -2.0, 0.0, 3.0, 2.0, 5.0, 0.7, -10.0, -2.0, 1.0, 1.0, 0.15, 0.08)

	f.Fuzz(func(t *testing.T, ax, ay, az, bx, by, bz, radius, ox, oy, oz, dx, dy, dz float64) {
		if !isFiniteAndBounded(ax, ay, az, bx, by, bz, radius, ox, oy, oz, dx, dy, dz) {
			t.Skip()
		}
		if radius <= 0 || radius > 1e3 {
			t.Skip()
		}

		c := Capsule{A: Vec3{ax, ay, az}, B: Vec3{bx, by, bz}, Radius: radius}
		r := Ray3{Origin: Vec3{ox, oy, oz}, Dir: Vec3{dx, dy, dz}}
		if !r.IsValid() {
			t.Skip()
		}

		hit, tMin, tMax := IntersectRayCapsule(r, c)
		if !hit {
			return
		}

		if math.IsNaN(tMin) || math.IsNaN(tMax) {
			t.Fatalf("IntersectRayCapsule produced NaN t values: tMin=%v tMax=%v", tMin, tMax)
		}
		if tMin > tMax {
			t.Fatalf("tMin > tMax: %v > %v", tMin, tMax)
		}

		const tol = 1e-4
		if d := DistancePointToCapsule(r.PointAt(tMin), c); d > tol {
			t.Fatalf("entry point not on capsule surface: distance=%v (t=%v)", d, tMin)
		}
		if d := DistancePointToCapsule(r.PointAt(tMax), c); d > tol {
			t.Fatalf("exit point not on capsule surface: distance=%v (t=%v)", d, tMax)
		}
	})
}
