# Public API Audit

## Keep
- Vec3
- Ray3
- Plane
- Segment3
- Triangle
- AABB
- Mat3
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

## Review later
- Whether additional projection helpers should be added to match `ProjectPointToLine`
- Whether more closest-point helpers should be added for rays or line-line / segment-segment queries
- Whether future intersection helpers should return richer result types or tuples
- Whether triangle-related helpers should expand further before `v1.0.0`
- Whether overlapping collinear segment intersection should eventually have a richer helper

## Current API direction
- Keep primitive object behavior as methods
  - `Vec3.Norm`
  - `Segment3.Length`
  - `Triangle.Area`
  - `Plane.UnitNormal`

- Keep geometric relations between multiple objects as free functions
  - `DistancePointToPlane`
  - `DistancePointToSegment`
  - `DistancePointToLine`
  - `DistancePointToTriangle`
  - `DistanceBetweenSegments`
  - `ProjectPointToPlane`
  - `ProjectPointToLine`
  - `BarycentricCoordinates`
  - `ClosestPointOnSegment`
  - `ClosestPointOnTriangle`
  - `ClosestPointsBetweenSegments`
  - `IntersectRayPlane`
  - `IntersectSegmentPlane`
  - `IntersectSegments`

## Notes toward v1.0.0
The public API is now more coherent and practically useful than in early releases.

Before `v1.0.0`, remaining review should focus on:
- any missing core geometry queries
- whether current triangle support is sufficient
- whether any return shapes should be standardized further
- whether any additional convenience helpers are essential enough to freeze into the public API
