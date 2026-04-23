// Package geom3d provides practical 3D geometry and spatial math utilities
// for Go.
//
// It includes vectors, planes, rays, segments, triangles, axis-aligned
// bounding boxes, rotation matrices, rigid transforms, and common operations
// such as distances, projections, closest-point queries, and intersections.
//
// geom3d is designed for engineering, simulation, robotics-adjacent, and
// analytics workflows rather than rendering, physics engines, or CAD kernels.
//
// Core features:
//   - Vec3 operations: dot, cross, norm, normalization, distances
//   - Geometry primitives: Ray3, Segment3, Plane, Triangle, AABB
//   - Practical operations: projection, barycentric coordinates, closest point,
//     signed distance, segment-segment queries, ray queries, box queries,
//     ray-plane intersection, segment-plane intersection, and ray-AABB intersection
//   - Rigid transforms: Mat3 rotations and Transform composition/inversion
//
// Non-goals:
//   - Rendering or OpenGL helpers
//   - Mesh loading or file formats
//   - Physics engine features
//   - General dense linear algebra
package geom3d
