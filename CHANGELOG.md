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

### Improved
- `IntersectRayTriangle` now shares its Möller–Trumbore implementation with the new `IntersectSegmentTriangle` via a private `intersectLineTriangle` helper, instead of each duplicating the algorithm
- `Segment3.Midpoint` and `AABB.Center` now delegate to `Vec3.Midpoint` instead of duplicating the averaging logic
- Removed a duplicate clamp implementation in `closest.go` (`ClosestPointOnAABB` now shares the same `clamp` helper as `clamp01`)
- `SegmentsOverlap` now uses `math.Abs`/`math.Max`/`math.Min` instead of manual comparisons
- Corrected the README and API_AUDIT.md, which still described the library as "approaching v1.0.0" after v1.0.0 had already shipped
- CI now only runs the race detector on Linux and macOS, since `-race` requires cgo and a C compiler that isn't guaranteed to be preconfigured on the windows-latest runner

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
