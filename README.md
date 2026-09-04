# geom3d

[![CI](https://github.com/motah-fard/geom3d/actions/workflows/ci.yml/badge.svg)](https://github.com/motah-fard/geom3d/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/motah-fard/geom3d.svg)](https://pkg.go.dev/github.com/motah-fard/geom3d)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/motah-fard/geom3d)](https://github.com/motah-fard/geom3d/releases)

`geom3d` is a lightweight, engine-independent **computational geometry**
library for Go: 3D primitives, intersection tests, distance queries,
closest-point operations, projections, and rigid transforms for
simulations, CAD, robotics, games, visualization, and spatial
applications — without requiring a rendering or physics engine.

The library is intentionally small, explicit, and easy to use.

**[Try it live →](https://geom-3d-playground.netlify.app)** — geom3d
Playground is an interactive companion site where several of these exact
queries (point-to-plane projection, ray-plane intersection, closest point
on a segment, segment-to-segment distance, ray-AABB intersection, closest
point on an AABB) run against this library through a small Go API, with
the request and response shown alongside the live 3D scene. [Source →](https://github.com/motah-fard/geom3d-Playground)

## Features

- `Vec3` for 3D vector and point operations
- Geometry primitives:
  - `Ray3`
  - `Segment3`
  - `Plane`
  - `Triangle`
  - `AABB`
  - `Sphere`
  - `OBB` (oriented bounding box)
  - `Capsule`
  - `Line3` (infinite line)
- Core operations:
  - dot product
  - cross product
  - norm and normalization
  - distance calculations
  - lerp, reflect, vector projection, angle between vectors, length clamping
  - component-wise abs/min/max
  - projections onto planes and lines
  - closest point on a ray, segment, triangle, AABB, and sphere
  - closest points between segments
  - ray-plane intersection
  - ray-triangle intersection
  - ray-sphere intersection with hit interval output (`hit`, `tMin`, `tMax`)
  - segment-plane intersection
  - segment-triangle intersection
  - segment-AABB and segment-sphere intersection with hit interval output (bounded raycasts)
  - segment-segment intersection at a single point
  - collinear segment overlap detection
  - ray-AABB intersection with hit interval output (`hit`, `tMin`, `tMax`)
  - ray-OBB and ray-capsule intersection with hit interval output
  - AABB-sphere, sphere-sphere, and capsule-capsule intersection
  - AABB construction from a point set, union, margin expansion, volume, and surface area
  - triangle-triangle overlap, including the coplanar case (partial overlap and full containment)
  - plane-plane intersection (returns the line of intersection)
  - line-plane intersection, and closest points/distance between two infinite lines
  - barycentric coordinates
  - point-to-ray, point-to-segment, point-to-line, point-to-triangle, point-to-AABB, point-to-sphere, point-to-OBB, point-to-capsule, sphere-to-sphere, capsule-to-capsule, segment-to-segment, and line-to-line distance queries
  - closest point on, and distance to, an oriented bounding box (`OBB`) or a `Capsule`
- 3D rotations with `Mat3`
  - matrix inverse and determinant, for the general (non-rotation) case
  - conversion to and from `Quaternion`
- `Quaternion` for rotation without gimbal lock
  - construction from axis-angle
  - composition, conjugate, inverse
  - spherical linear interpolation (`Slerp`)
  - rotating a `Vec3` directly, or converting to `Mat3` to rotate many points cheaply
- Rigid transforms with `Transform`
  - apply to points and vectors
  - compose transforms
  - invert transforms

## Why this library exists

Go already has solid low-level vector/matrix math packages aimed at
graphics pipelines (see "Compared to other Go libraries" below). `geom3d`
isn't another one of those — it answers a different question: given two
geometric objects, what's their relationship? Do they intersect? What's
the closest point? How far apart are they?

That's the computational-geometry layer: the queries developers actually
reach for once they already have points and vectors, without needing to
assemble the algorithms themselves or pull in an entire game/physics
engine just to get a ray-triangle test.

Typical use cases include:

- projecting points onto planes or lines
- finding the closest point on a ray, segment, triangle, bounding box, or sphere
- computing closest points or minimum distance between segments
- checking ray intersections with planes, triangles, bounding boxes, or spheres
- bounded raycasts against a triangle, box, or sphere using a finite `Segment3` instead of an infinite `Ray3`
- building an AABB around a point cloud and testing it against other boxes or spheres
- testing whether two segments intersect at a single point or overlap collinearly
- computing barycentric coordinates for triangle-based workflows
- representing and composing rotations with `Quaternion` for robotics/biomechanics workflows where gimbal lock or interpolation quality matters
- testing a rotated bounding volume (`OBB`) against points, for tighter collision bounds than an `AABB`
- modeling swept-sphere volumes (character capsules, cylindrical links) with `Capsule`
- raycasting against a rotated bounding volume (`OBB`) or a capsule, for hit-testing and picking
- finding where two planes meet, or testing overlap between two triangles that may or may not be coplanar
- applying and composing rigid transforms
- working with coordinate frames in engineering or sensor-based applications

### Compared to other Go libraries

[`go-gl/mathgl`](https://github.com/go-gl/mathgl) is the other well-known
pure-Go 3D math package, and it's a good library — but it solves a
different problem. It's GLM-inspired vector/matrix math aimed at graphics
pipelines: statically-sized vectors and matrices, projection/perspective
matrix construction, and generating basic primitives for rendering. It's
the right choice if you're driving a graphics API and need that
GLM-shaped API surface.

`geom3d` doesn't compete with that; it targets a different, narrower
question: given two geometric objects, what's their relationship
(distance, closest point, do they intersect)? That's a smaller, more
specialized problem area — this is a spatial-query and collision-geometry
library, not a general vector-math or rendering-adjacent one. If you need
`glClipPlane`-style graphics pipeline math, reach for `mathgl`. If you need
"does this ray hit this capsule" or "closest point on this triangle,"
that's what `geom3d` is for. The two are complementary rather than
competing — nothing stops a project from using both.

## Non-goals

`geom3d` is not:

- a rendering engine
- an OpenGL helper library
- a physics engine
- a mesh loader
- a CAD kernel

## Installation

```bash
go get github.com/motah-fard/geom3d
```

## Quick example

```go
package main

import (
    "fmt"

    "github.com/motah-fard/geom3d"
)

func main() {
    box := geom3d.AABB{
        Min: geom3d.Vec3{X: 0, Y: 0, Z: 0},
        Max: geom3d.Vec3{X: 2, Y: 2, Z: 2},
    }

    ray := geom3d.Ray3{
        Origin: geom3d.Vec3{X: -1, Y: 1, Z: 1},
        Dir:    geom3d.Vec3{X: 1, Y: 0, Z: 0},
    }

    hit, tMin, tMax := geom3d.IntersectRayAABB(ray, box)

    fmt.Println("hit:", hit)
    fmt.Println("tMin:", tMin)
    fmt.Println("tMax:", tMax)
}
```

## Package overview

### Vectors
`Vec3` supports common 3D vector operations such as addition, subtraction, scaling, dot products, cross products, norms, distances, midpoints, normalization, linear interpolation (`Lerp`), reflection, vector-onto-vector projection, angle between vectors, length clamping, and component-wise `Abs`/`Min`/`Max`.

### Primitives
The package includes practical 3D primitives for common geometric workflows:

- `Ray3`
- `Segment3`
- `Plane`
- `Triangle`
- `AABB`
- `Sphere`
- `OBB`
- `Capsule`
- `Line3`

### Matrices, quaternions, and transforms
`Mat3` supports 3D rotation matrices and matrix operations, including a general `Determinant`/`Inverse` (for the non-rotation case; use `Transpose` for rotation matrices) and conversion to/from `Quaternion`.  
`Quaternion` supports rotation without gimbal lock: construction from axis-angle, composition (`Mul`), `Conjugate`/`Inverse`, spherical interpolation (`Slerp`), and applying the rotation directly to a `Vec3` or converting to `Mat3`.  
`Transform` supports rigid-body transforms for points and vectors, transform composition, and inversion.

### Geometric helpers
The package includes helpers for:

- point-to-plane distance
- point-to-ray distance
- point-to-segment distance
- point-to-line distance
- point-to-triangle distance
- point-to-AABB distance
- point-to-sphere distance
- point-to-OBB distance
- point-to-capsule distance
- sphere-to-sphere distance
- capsule-to-capsule distance
- segment-to-segment distance
- line-to-line distance
- point projection to planes and lines
- barycentric coordinates
- closest-point queries on rays, segments, triangles, AABBs, spheres, OBBs, and capsules
- closest-point queries between segments, and between infinite lines
- ray-plane intersection
- ray-triangle intersection
- ray-sphere intersection
- ray-OBB intersection
- ray-capsule intersection
- line-plane intersection
- plane-plane intersection
- triangle-triangle overlap (coplanar and non-coplanar)
- segment-plane intersection
- segment-triangle intersection
- segment-AABB intersection (bounded raycast)
- segment-sphere intersection (bounded raycast)
- segment-segment intersection
- collinear segment overlap detection
- ray-AABB intersection
- AABB-sphere intersection
- sphere-sphere intersection
- capsule-capsule overlap
- `AABBFromPoints`, `AABB.Union`, `AABB.ExpandToInclude`, `AABB.Expand`, `AABB.Volume`, `AABB.SurfaceArea`

## Error handling

`geom3d` does not use panics or the `error` type for invalid geometric input
(an invalid `AABB`, a zero-length `Ray3` direction, a degenerate `Triangle`,
a negative-radius `Sphere`, and so on). This is a deliberate choice, not an
oversight:

- Every primitive that can be invalid or degenerate exposes an `IsValid()`
  and/or `IsDegenerate()` method (e.g. `Ray3.IsValid`, `AABB.IsValid`,
  `Sphere.IsValid`, `Triangle.IsDegenerate`, `Segment3.IsDegenerate`).
  Call these at your program's boundary if you need to reject bad input
  explicitly, the same way you'd validate any external data.
- Queries that can legitimately have **no answer** for valid input — "does
  this ray hit this plane," "do these segments intersect" — report that
  with a `bool` return, following Go's own comma-ok idiom
  (`p, ok := IntersectRayPlane(r, pl)`). This is not error handling in the
  `error`-type sense; a `false` here is an expected, meaningful result, not
  a failure.
- Queries that receive **invalid** input (rather than valid input with no
  answer) return a documented zero-value fallback instead of panicking.
  Every such function's GoDoc comment states exactly what it returns for
  invalid or degenerate input — that comment is the authoritative contract,
  not this README.
- All types are plain value structs (no pointers, no `nil` in the public
  API), so there is no possibility of a `nil` dereference from this package.

Because `geom3d` is past `v1.0.0`, these return shapes are frozen: adding a
`bool`/`error` to an existing function's signature would be a breaking
change and won't happen within `v1`. If your use case needs to distinguish
"invalid input" from "no result" more strictly than the zero-value fallback
allows, check `IsValid()`/`IsDegenerate()` before calling.

## Numerical tolerance and large coordinates

Every `IsValid`/`IsDegenerate` check, and every internal "are these two
values equal" comparison, in this package is built on `AlmostZero`/
`AlmostEqual`, which use a fixed absolute tolerance (`Epsilon = 1e-9`). That
works well at "ordinary" coordinate magnitudes — roughly the range a scene
described in meters, or a part described in millimeters, would use — but it
breaks down at large ones: two values can agree to 12 significant figures
and still differ by far more than `1e-9` in absolute terms once they're in
the millions. If you're working at large coordinate magnitudes (e.g. a
world-scale simulation, or GPS-derived coordinates), be aware that
`geom3d`'s own degeneracy checks (a triangle that's technically non-zero
area but numerically flat at that scale, for instance) use this fixed
tolerance and won't automatically adapt.

`AlmostEqualRelative` and `AlmostZeroAtScale` are provided for **your own**
comparisons in this situation — for interpreting `geom3d`'s output, or for
your own tolerance decisions — but they do not change how `geom3d`'s
internal functions behave; that would be a silent behavior change to a
frozen `v1` API. There is currently no scale-aware variant of the library's
own internal checks; see `API_AUDIT.md`'s "Review later" section.

## Behavior notes

The general invalid-input contract is described above; a few helpers also
have intentionally specific semantics worth calling out:

- `IntersectSegments` reports only **single-point** intersections. If two segments overlap over a non-zero interval, it returns `false`.
- `SegmentsOverlap` reports only **collinear overlap over a non-zero interval**. Endpoint-only touching does not count as overlap.
- `IntersectRayAABB` returns `hit, tMin, tMax`. If the ray starts inside the box, `tMin` may be `0`.
- `IntersectRayTriangle` does not perform back-face culling; a hit is reported regardless of which side of the triangle the ray approaches from.
- `IntersectRaySphere` returns `hit, tMin, tMax`. If the ray starts inside the sphere, `tMin` is clamped to `0`.
- `ClosestPointOnSphere` and `DistancePointToSphere` treat `Sphere` as a solid ball: a point inside the sphere returns itself (distance `0`), matching `ClosestPointOnAABB`'s behavior for points inside a box.
- `ClosestPointOnRay` clamps to the ray origin when the orthogonal projection falls behind the origin.
- `ClosestPointOnAABB` returns the input point itself when the point lies inside the box.
- `IntersectSegmentAABB`, `IntersectSegmentSphere`, and `IntersectSegmentTriangle` mirror their `Ray3` counterparts but bound the hit interval to the segment's own length (`t` in `[0, 1]`, where `0` is `s.A` and `1` is `s.B`); a shape the infinite ray would hit is correctly reported as a miss if it lies beyond the segment's endpoint.
- `AABBFromPoints` returns `AABB{}` for an empty slice, rather than a sentinel error — check `len(points) == 0` yourself first if that distinction matters to you.
- `AABB.Union`, `AABB.ExpandToInclude`, and `AABB.Expand` compute their result component-wise without checking `IsValid()` first, matching `AABB.Overlaps`'s existing convention; feeding them an invalid box propagates that invalidity into the result rather than silently discarding it.
- `Sphere.Overlaps` and `IntersectAABBSphere` treat touching shapes (distance exactly equal to the combined radius) as overlapping, matching `AABB.Overlaps`'s inclusive-boundary convention.
- `Mat3.Inverse` returns `(Mat3{}, false)` when the matrix is singular (zero determinant). This is different from the rest of the library's "invalid input" convention: a singular matrix is a perfectly valid `Mat3` value that simply has no inverse, the same way `IntersectRayPlane` returning `false` means "no intersection for this valid input," not "your input was invalid."
- `Mat3.ToQuaternion` and `Quaternion.ToMat3` assume the value they're converting is a proper rotation (an orthonormal `Mat3`, or a `Quaternion` that is at least non-zero); passing an arbitrary matrix or a zero quaternion in does not produce a meaningful rotation out.
- `Quaternion.ToMat3` returns `IdentityMat3()` (not `Mat3{}`) for the zero quaternion, because a zero matrix would silently collapse every point it's applied to onto the origin — a worse failure mode than "no rotation."
- `OBB.Contains`, `ClosestPointOnOBB`, and `DistancePointToOBB` treat `OBB` as a solid box, mirroring `AABB`'s and `Sphere`'s conventions exactly (a point inside returns itself, at distance `0`).
- `Capsule.Contains`, `ClosestPointOnCapsule`, and `DistancePointToCapsule` treat `Capsule` as a solid volume, the same way `Sphere` and `OBB` do.
- `IntersectLinePlane` and `ClosestPointsBetweenLines`/`DistanceBetweenLines` are the unbounded-both-directions counterparts of `IntersectRayPlane` and `ClosestPointsBetweenSegments`/`DistanceBetweenSegments`; they exist because clamping to `t >= 0` or `t` in `[0, 1]` isn't always what you want. For a line given as two points rather than a `Line3`, `DistancePointToLine` and `ProjectPointToLine` remain the simpler choice.
- `Triangle.Overlaps` handles coplanar triangles (partial overlap or full containment) via a 2D separating-axis test on the shared plane, since edges lying within a triangle's own plane are always reported as "parallel" (not intersecting) by `IntersectSegmentTriangle`. Non-coplanar overlap is detected by edge crossings instead.
- `IntersectPlanePlane` returns `false` for **coincident** planes (the same plane specified two different ways), the same way `IntersectSegments` returns `false` for collinear overlap: the "intersection" isn't a single well-defined line, it's the entire plane.
- `IntersectRayOBB` and `IntersectRayCapsule` return `hit, tMin, tMax` exactly like `IntersectRayAABB`/`IntersectRaySphere`, including the same "`tMin` clamped to `0` if the ray starts inside" convention.
- `AlmostEqualRelative` and `AlmostZeroAtScale` compare against a tolerance that scales with the magnitude of the values involved, unlike every other function in the package (which use the fixed `Epsilon`) — see "Numerical tolerance and large coordinates" above.
- `Capsule.Overlaps` and `DistanceBetweenCapsules` compare the distance between the two capsules' **core segments** against the sum of their radii — the same reasoning as `Sphere.Overlaps`, just with a segment instead of a point at the center.

## Examples

Runnable examples are included under the `examples/` directory, including:

- `basic_vectors`
- `ray_plane`
- `ray_triangle`
- `ray_closest_point`
- `ray_distance`
- `plane_projection`
- `segment_closest_point`
- `segment_overlap`
- `segment_segment_closest`
- `segment_segment_distance`
- `aabb_ray`
- `aabb_closest_point`
- `aabb_distance`
- `sphere_closest_point`
- `sphere_ray`
- `sphere_sphere`
- `aabb_from_points`
- `segment_bounded_raycast`
- `triangle_normal`
- `triangle_closest_point`
- `triangle_barycentric`
- `transform_point`
- `quaternion_rotation`
- `obb_closest_point`
- `capsule_closest_point`
- `plane_plane_intersection`
- `triangle_overlap`
- `ray_obb`
- `ray_capsule`

## API stability

`geom3d` has reached `v1.0.0`. Existing exported function and method
signatures are frozen: they will not change in a breaking way within the
`v1` line.

New functionality (new primitives, new queries) is still added as minor
releases (`v1.1.0`, `v1.2.0`, ...) under standard [semantic
versioning](https://semver.org/), and is purely additive. A breaking change
to an existing signature would require a `v2`.

## Contributing

Contributions are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) for
development setup, coding conventions, and PR expectations; see
[ROADMAP.md](ROADMAP.md) for what's likely to be worked on next (and what's
deliberately out of scope). This project follows the
[Contributor Covenant](CODE_OF_CONDUCT.md).

## License

MIT
