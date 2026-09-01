# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]
### Added
- Added `IntersectRayTriangle` (Möller–Trumbore ray-triangle intersection)
- Added `examples/ray_triangle`
- Added `Sphere` primitive with `IsValid`, `IsDegenerate`, `SurfaceArea`, `Volume`, and `Contains`
- Added `ClosestPointOnSphere`, `DistancePointToSphere`, and `IntersectRaySphere`
- Added `examples/sphere_closest_point` and `examples/sphere_ray`
- Added a GitHub Actions CI workflow running build, vet, gofmt, and race-enabled tests across Linux, macOS, and Windows
- Added `CONTRIBUTING.md`, issue templates, and a pull request template
- Added `Vec3.Midpoint`
- Added a README "Error handling" section documenting the library's invalid-input contract
- Added `Vec3.Lerp`, `Vec3.Reflect`, `Vec3.Project`, `Vec3.Angle`, `Vec3.ClampLength`, `Vec3.Abs`, `Vec3.Min`, and `Vec3.Max`
- Added `AABBFromPoints`, `AABB.Union`, `AABB.ExpandToInclude`, `AABB.Expand`, `AABB.Volume`, and `AABB.SurfaceArea`
- Added `Sphere.Overlaps` and `DistanceBetweenSpheres`
- Added `IntersectAABBSphere`
- Added `IntersectSegmentTriangle`, `IntersectSegmentAABB`, and `IntersectSegmentSphere` (bounded raycasts for finite segments, mirroring the existing `Ray3` intersection tests)
- Added `examples/aabb_from_points`, `examples/segment_bounded_raycast`, and `examples/sphere_sphere`
- Added `Mat3.Determinant` and `Mat3.Inverse` (general 3x3 inverse, for non-rotation matrices; `Transpose` remains the cheaper choice for rotations)
- Added `Quaternion`: `IdentityQuaternion`, `QuaternionFromAxisAngle`, `IsValid`, `Dot`, `Norm`/`Norm2`, `Normalize`, `Conjugate`, `Inverse`, `Mul`, `RotateVector`, `ToMat3`, and `Slerp`
- Added `Mat3.ToQuaternion`
- Added `OBB` (oriented bounding box) with `IsValid`, `Volume`, `SurfaceArea`, and `Contains`, plus `ClosestPointOnOBB` and `DistancePointToOBB`
- Added `examples/quaternion_rotation` and `examples/obb_closest_point`
- Added `Capsule` with `IsValid`, `IsDegenerate`, `Segment`, `Volume`, `SurfaceArea`, and `Contains`, plus `ClosestPointOnCapsule` and `DistancePointToCapsule`
- Added `Line3` (infinite line) with `PointAt` and `IsValid`
- Added `IntersectLinePlane`, `IntersectPlanePlane`, `ClosestPointsBetweenLines`, and `DistanceBetweenLines`
- Added `Triangle.Overlaps`, including a 2D separating-axis test for the coplanar case (partial overlap or full containment), alongside the general non-coplanar edge-crossing test
- Added `examples/capsule_closest_point`, `examples/plane_plane_intersection`, and `examples/triangle_overlap`
- Added `IntersectRayOBB`, reusing `IntersectRayAABB` by transforming the ray into the OBB's local frame
- Added `IntersectRayCapsule` (ray vs. the infinite cylinder along the capsule's axis, clipped to the two hemispherical end caps)
- Added `Capsule.Overlaps` and `DistanceBetweenCapsules`
- Added `examples/ray_obb` and `examples/ray_capsule`
- Added `RelativeEpsilon`, `AlmostEqualRelative`, and `AlmostZeroAtScale` for scale-aware tolerance comparisons at large coordinate magnitudes, where the fixed `Epsilon` breaks down; documented as tools for the caller's own use, not a change to any existing function's internal behavior
- Added `benchmark_test.go` covering the package's hot paths (closest-point queries, ray/segment intersections, `Mat3`/`Quaternion` operations)
- Added `fuzz_test.go` with fuzz targets for `Triangle.Overlaps` (symmetry), `Mat3.Inverse` (consistency with `Determinant`, and numerical accuracy for well-conditioned matrices), `Quaternion.Slerp` (no NaN/Inf for any valid input and any `t`), and `IntersectRayCapsule` (entry/exit points always exactly on the capsule surface, cross-checked against `DistancePointToCapsule`) — each run clean for 10M+ generated cases with zero failures before being committed
- Documented how to run both in `CONTRIBUTING.md`

### Improved
- `IntersectRayTriangle` now shares its Möller–Trumbore implementation with the new `IntersectSegmentTriangle` via a private `intersectLineTriangle` helper, instead of each duplicating the algorithm
- `Segment3.Midpoint` and `AABB.Center` now delegate to `Vec3.Midpoint` instead of duplicating the averaging logic
- Removed a duplicate clamp implementation in `closest.go` (`ClosestPointOnAABB` now shares the same `clamp` helper as `clamp01`)
- `SegmentsOverlap` now uses `math.Abs`/`math.Max`/`math.Min` instead of manual comparisons
- Corrected the README and API_AUDIT.md, which still described the library as "approaching v1.0.0" after v1.0.0 had already shipped
- CI now only runs the race detector on Linux and macOS, since `-race` requires cgo and a C compiler that isn't guaranteed to be preconfigured on the windows-latest runner
- Moved `ClosestPointOnOBB` and `ClosestPointOnCapsule` into `closest.go`, and `DistancePointToCapsule` into `distance.go`, for consistency with every other `ClosestPointOnX`/`DistancePointToX` free function living in those two files rather than the primitive's own file
- `IntersectRaySphere` now shares its quadratic-solving with the new `IntersectRayCapsule` via a private `intersectLineSphere` helper
- Removed two provably-unreachable root-ordering swaps in `intersectLineSphere` and `IntersectRayCapsule`'s cylinder test (the quadratic's leading coefficient is always non-negative in both cases, so the roots are already ordered)

## [v0.5.0]
### Added
- Added `ClosestPointOnRay`
- Added `DistancePointToRay`
- Added `ClosestPointOnAABB`
- Added `DistancePointToAABB`
- Added `SegmentsOverlap`

### Improved
- Expanded ray query support
- Expanded AABB point-query support
- Added a dedicated helper for collinear segment overlap detection
- Updated documentation, examples, and API audit notes for the new ray, AABB, and segment helpers
- Strengthened practical pre-v1 geometry coverage across rays, segments, triangles, and boxes

## [v0.4.0]
### Added
- Added `ClosestPointsBetweenSegments`
- Added `DistanceBetweenSegments`
- Added `IntersectSegments`

### Improved
- Expanded segment-segment query support
- Added edge-case coverage for overlapping collinear segments, endpoint-touching segments, skew disjoint segments, and degenerate segment cases
- Strengthened practical support for segment-based geometry workflows

## [v0.3.0]
### Added
- Added `DistancePointToSegment`
- Added `DistancePointToLine`
- Added `BarycentricCoordinates`
- Added `ClosestPointOnTriangle`
- Added `DistancePointToTriangle`

### Improved
- Expanded triangle support and point-query capabilities
- Updated documentation and package overview to reflect the new geometry helpers
- Strengthened practical coverage for line, segment, and triangle workflows

## [v0.2.0]
### Changed
- Refined `IntersectRayAABB` to return `(hit, tMin, tMax)` instead of only a boolean
- Updated tests and examples to match the improved AABB intersection API
- Continued public API consistency review ahead of `v1.0.0`

### Improved
- Polished documentation and example coverage
- Cleaned up test naming and general package organization

## [v0.1.2]
### Improved
- Expanded example coverage for core geometry types and helpers
- Added pkg.go.dev example functions for key APIs
- Improved exported API documentation
- Expanded unit test coverage across vectors, primitives, helpers, matrices, and transforms
- Polished README and overall package presentation

### Fixed
- Minor cleanup and typo fixes across examples and tests

## [v0.1.1]
### Improved
- Updated the README for better clarity, structure, and presentation

## [v0.1.0]
### Added
- Initial public release of `geom3d`
- `Vec3` operations
- 3D primitives: `Ray3`, `Segment3`, `Plane`, `Triangle`, `AABB`
- Core geometry helpers for distances, projections, closest-point queries, and intersections
- `Mat3` rotation helpers
- `Transform` for rigid 3D transforms
