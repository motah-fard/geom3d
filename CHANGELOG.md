# Changelog

All notable changes to this project will be documented in this file.

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
