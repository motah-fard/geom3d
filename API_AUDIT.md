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

## Review
- IntersectRayAABB returns only bool
- ProjectPointToLine naming vs future projection naming
- Whether more methods should become type methods vs free functions