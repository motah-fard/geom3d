# Benchmarks

## geom3d's own hot paths

`benchmark_test.go` covers geom3d's own hot paths — closest-point queries,
ray/segment intersections, `Mat3`/`Quaternion` operations — against
themselves, to catch regressions over time. Run them with:

```bash
go test -bench=. -benchmem -run '^$' .
```

## Comparison against go-gl/mathgl

geom3d's actual differentiator — ray/triangle, ray/AABB, closest-point,
and distance queries between primitives — has no equivalent in
[`go-gl/mathgl`](https://github.com/go-gl/mathgl), so those aren't
comparable and aren't benchmarked here. The two libraries only genuinely
overlap on basic vector, matrix, and quaternion arithmetic, so that's all
this comparison covers: `Vec3.Normalize`/`Cross`/`Dot`, `Mat3` multiply,
and `Quaternion` multiply/rotate/slerp, using `mgl64` (mathgl's `float64`
variant, matching geom3d's own numeric type).

This benchmark lives in its own module, at
[`benchmarks/vs-mathgl`](benchmarks/vs-mathgl), specifically so mathgl
never becomes a dependency of the main `geom3d` module — see the module's
own `go.mod` and comment for why. Run it with:

```bash
cd benchmarks/vs-mathgl
go test -bench=. -benchmem -run '^$' .
```

### Results

Measured 2026-09-04, Go 1.26.3, darwin/arm64, Apple M4 Pro, three runs
each (numbers were stable run-to-run; see raw output in the module's
history if you want to reproduce and compare):

| Operation | geom3d | mathgl (mgl64) | 0 allocs/op (both) |
|---|---:|---:|:---:|
| `Vec3.Normalize` | 0.69 ns/op | 1.59 ns/op | ✓ |
| `Vec3.Cross` | 0.33 ns/op | 1.41 ns/op | ✓ |
| `Vec3.Dot` | 0.23 ns/op | 0.98 ns/op | ✓ |
| `Mat3` multiply | 20.8 ns/op | **7.4 ns/op** | ✓ |
| `Quaternion` multiply | 1.18 ns/op | 14.1 ns/op | ✓ |
| `Quaternion.Rotate(Vec3)` | 4.07 ns/op | 14.1 ns/op | ✓ |
| `Quaternion.Slerp` | 23.4 ns/op | 63.9 ns/op | ✓ |

Neither library allocates for any of these operations.

### Reading this honestly

geom3d is faster on every operation measured here **except `Mat3`
multiply, where mathgl is about 2.8x faster**. The reason is visible in
the source, not a fluke of the benchmark: geom3d's `Mat3.Mul` is a
straightforward triple-nested loop over a `[3][3]float64`, while mathgl's
`Mat3` is a flat `[9]float64` with a hand-unrolled `Mul3` — that shape
lets the compiler avoid loop overhead and array-bounds checks that the
nested-loop version doesn't get for free. It's a real, reproducible gap,
not measurement noise (see [ROADMAP.md](ROADMAP.md) — flattening `Mat3`'s
internal storage would be an internal-only change, not a signature
change, so it's a legitimate future optimization, just not one bundled
into this benchmarking pass).

This is not a "geom3d is faster than mathgl" claim — it's a report of
what these specific overlapping operations measured like, on one machine,
on one date, including the one case that didn't go geom3d's way. mathgl
is solving a different problem (see the README's "Compared to other Go
libraries" section) and these micro-benchmarks say nothing about which
library is the right choice for a given project.
