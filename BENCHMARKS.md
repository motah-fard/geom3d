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

Measured 2026-09-10, Go 1.26.3, darwin/arm64, Apple M4 Pro, three runs
each (numbers were stable run-to-run; see raw output in the module's
history if you want to reproduce and compare):

| Operation | geom3d | mathgl (mgl64) | 0 allocs/op (both) |
|---|---:|---:|:---:|
| `Vec3.Normalize` | 0.68 ns/op | 1.58 ns/op | ✓ |
| `Vec3.Cross` | 0.32 ns/op | 1.41 ns/op | ✓ |
| `Vec3.Dot` | 0.23 ns/op | 0.98 ns/op | ✓ |
| `Mat3` multiply | 7.6 ns/op | 7.7 ns/op | ✓ |
| `Quaternion` multiply | 1.20 ns/op | 14.3 ns/op | ✓ |
| `Quaternion.Rotate(Vec3)` | 4.11 ns/op | 14.4 ns/op | ✓ |
| `Quaternion.Slerp` | 23.1 ns/op | 64.7 ns/op | ✓ |

Neither library allocates for any of these operations.

### Reading this honestly

geom3d now matches or beats mathgl on every operation measured here.
`Mat3` multiply used to be the one exception — mathgl was ~2.8x faster as
of the previous measurement (2026-09-04) — but that gap is now closed:
`Mat3.Mul` was hand-unrolled into 9 direct index expressions instead of a
triple-nested loop over `[3][3]float64`, which eliminated the loop
overhead and bounds checks the old version paid on every call. That's a
purely internal change — `Mat3`'s public shape didn't move, so this
shipped as a normal, non-breaking release (no `v2` needed; see
`CHANGELOG.md`). The gap wasn't a fundamental storage-layout problem
after all, just an unrolled-vs-looped implementation difference — worth
knowing before assuming a fix requires a breaking API change.

This is not a "geom3d is faster than mathgl" claim in general — it's a
report of what these specific overlapping operations measured like, on
one machine, on one date. mathgl is solving a different problem (see the
README's "Compared to other Go libraries" section) and these
micro-benchmarks say nothing about which library is the right choice for
a given project.
