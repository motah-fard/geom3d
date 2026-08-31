# Public API Audit

## Keep
- Vec3
- Ray3
- Plane
- Segment3
- Triangle
- AABB
- Sphere
- OBB
- Capsule
- Line3
- Mat3
- Quaternion
- Transform
- Norm / Norm2
- Length / Length2
- Distance / Distance2
- IsValid
- IsDegenerate
- ProjectPointToLine
- Current method vs free-function split

## Changed
- IntersectRayAABB now returns `(hit, tMin, tMax)` instead of only `bool`

## Added in v0.3.0
- DistancePointToSegment
- DistancePointToLine
- BarycentricCoordinates
- ClosestPointOnTriangle
- DistancePointToTriangle

## Added in v0.4.0
- ClosestPointsBetweenSegments
- DistanceBetweenSegments
- IntersectSegments

## Added in v0.5.0
- ClosestPointOnRay
- DistancePointToRay
- ClosestPointOnAABB
- DistancePointToAABB
- SegmentsOverlap

## Added after v1.0.0
- IntersectRayTriangle
- Sphere
- ClosestPointOnSphere
- DistancePointToSphere
- IntersectRaySphere
- Vec3.Midpoint
- Vec3.Lerp, Vec3.Reflect, Vec3.Project, Vec3.Angle, Vec3.ClampLength, Vec3.Abs, Vec3.Min, Vec3.Max
- AABBFromPoints, AABB.Union, AABB.ExpandToInclude, AABB.Expand, AABB.Volume, AABB.SurfaceArea
- Sphere.Overlaps, DistanceBetweenSpheres, IntersectAABBSphere
- IntersectSegmentTriangle, IntersectSegmentAABB, IntersectSegmentSphere
- Mat3.Determinant, Mat3.Inverse, Mat3.ToQuaternion
- Quaternion, IdentityQuaternion, QuaternionFromAxisAngle
- OBB, ClosestPointOnOBB, DistancePointToOBB
- Capsule, ClosestPointOnCapsule, DistancePointToCapsule
- Line3, IntersectLinePlane, IntersectPlanePlane, ClosestPointsBetweenLines, DistanceBetweenLines
- Triangle.Overlaps

## Review later
- Whether additional projection helpers should be added to match `ProjectPointToLine`
- Whether future intersection helpers should return richer result types or tuples
- Whether overlapping collinear segment behavior should eventually have a richer relation helper
- Whether invalid-input reporting (currently a documented zero-value fallback, see README's "Error handling" section) should become a `(value, bool)` return uniformly across all queries — this would require a `v2`, since it changes existing signatures
- Ray-OBB and ray-capsule intersection are natural follow-ups to `IntersectRayAABB`/`IntersectRaySphere` but aren't implemented yet
- AABB-OBB, OBB-OBB, and capsule-capsule intersection/overlap tests aren't implemented yet

## Current API direction
- Keep primitive object behavior as methods, including same-type relations
  (e.g. `AABB.Overlaps`, `Sphere.Overlaps`) and derived-value conversions
  (e.g. `Triangle.Normal`, `Plane.UnitNormal`) on the receiver's own type
  - `Vec3.Norm`, `Vec3.Midpoint`, `Vec3.Lerp`, `Vec3.Reflect`, `Vec3.Project`, `Vec3.Angle`, `Vec3.ClampLength`, `Vec3.Abs`, `Vec3.Min`, `Vec3.Max`
  - `Segment3.Length`
  - `Triangle.Area`
  - `Plane.UnitNormal`
  - `AABB.Overlaps`, `AABB.Union`, `AABB.ExpandToInclude`, `AABB.Expand`, `AABB.Volume`, `AABB.SurfaceArea`
  - `Sphere.Overlaps`
  - `OBB.IsValid`, `OBB.Volume`, `OBB.SurfaceArea`, `OBB.Contains`
  - `Capsule.IsValid`, `Capsule.IsDegenerate`, `Capsule.Segment`, `Capsule.Volume`, `Capsule.SurfaceArea`, `Capsule.Contains`
  - `Line3.PointAt`, `Line3.IsValid`
  - `Triangle.Overlaps`
  - `Mat3.Determinant`, `Mat3.Inverse`, `Mat3.ToQuaternion`
  - `Quaternion.Dot`, `Quaternion.Norm`/`Norm2`, `Quaternion.Normalize`, `Quaternion.Conjugate`, `Quaternion.Inverse`, `Quaternion.Mul`, `Quaternion.RotateVector`, `Quaternion.ToMat3`, `Quaternion.Slerp`

- Keep geometric relations between different primitive types, and
  multi-object constructors, as free functions
  - `DistancePointToPlane`
  - `DistancePointToRay`
  - `DistancePointToSegment`
  - `DistancePointToLine`
  - `DistancePointToTriangle`
  - `DistancePointToAABB`
  - `DistancePointToSphere`
  - `DistanceBetweenSegments`
  - `DistanceBetweenSpheres`
  - `ProjectPointToPlane`
  - `ProjectPointToLine`
  - `BarycentricCoordinates`
  - `ClosestPointOnRay`
  - `ClosestPointOnSegment`
  - `ClosestPointOnTriangle`
  - `ClosestPointOnAABB`
  - `ClosestPointOnSphere`
  - `ClosestPointsBetweenSegments`
  - `IntersectRayPlane`
  - `IntersectRayTriangle`
  - `IntersectRaySphere`
  - `IntersectSegmentPlane`
  - `IntersectSegmentTriangle`
  - `IntersectSegmentAABB`
  - `IntersectSegmentSphere`
  - `IntersectSegments`
  - `IntersectAABBSphere`
  - `SegmentsOverlap`
  - `AABBFromPoints`
  - `ClosestPointOnOBB`
  - `DistancePointToOBB`
  - `ClosestPointOnCapsule`
  - `DistancePointToCapsule`
  - `ClosestPointsBetweenLines`
  - `DistanceBetweenLines`
  - `IntersectLinePlane`
  - `IntersectPlanePlane`
  - `QuaternionFromAxisAngle` (constructor from raw scalar/vector input, like `RotationX`/`RotationY`/`RotationZ`)

## Notes toward v1.0.0 (historical)
The public API is now more coherent and practically useful than in early releases.

Before `v1.0.0`, remaining review focused on:
- any missing core geometry queries
- whether current ray, triangle, segment, and AABB support is sufficient
- whether any return shapes should be standardized further
- whether any additional convenience helpers are essential enough to freeze into the public API

`v1.0.0` shipped with that review complete. See "Review later" above for the
current, ongoing list — anything added now is purely additive (see the
README's "API stability" section).
