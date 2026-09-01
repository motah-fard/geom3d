// Package geom3d provides practical 3D geometry and spatial math utilities
// for Go.
//
// It includes vectors, planes, infinite lines, rays, segments, triangles,
// axis-aligned and oriented bounding boxes, spheres, capsules, rotation
// matrices, quaternions, rigid transforms, and common operations such as
// distances, projections, closest-point queries, and intersections.
//
// geom3d is designed for engineering, simulation, robotics-adjacent, and
// analytics workflows rather than rendering, physics engines, or CAD kernels.
//
// Core features:
//   - Vec3 operations: dot, cross, norm, normalization, distances, midpoints,
//     lerp, reflect, projection, angle, length clamping, component-wise
//     abs/min/max
//   - Geometry primitives: Line3, Ray3, Segment3, Plane, Triangle, AABB,
//     Sphere, OBB, Capsule
//   - Practical operations: projection, barycentric coordinates, closest-point
//     queries, signed distance, segment-segment and line-line queries, ray
//     queries, box queries, sphere queries, OBB queries, capsule queries,
//     ray-plane intersection, ray-triangle intersection, ray-sphere
//     intersection, ray-OBB intersection, ray-capsule intersection,
//     line-plane intersection, plane-plane intersection, triangle-triangle
//     overlap (coplanar and non-coplanar), segment-plane intersection,
//     segment-triangle intersection, bounded segment-AABB and
//     segment-sphere raycasts, ray-AABB intersection, AABB-sphere,
//     sphere-sphere, and capsule-capsule intersection, and AABB
//     construction/union/expansion
//   - Rotations: Mat3 (including general Determinant/Inverse) and
//     Quaternion (axis-angle construction, composition, Slerp), with
//     conversion between the two
//   - Rigid transforms: Transform composition and inversion
//
// Non-goals:
//   - Rendering or OpenGL helpers
//   - Mesh loading or file formats
//   - Physics engine features
//   - General dense linear algebra
package geom3d
