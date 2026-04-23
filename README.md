# geom3d

[![Go Reference](https://pkg.go.dev/badge/github.com/motah-fard/geom3d.svg)](https://pkg.go.dev/github.com/motah-fard/geom3d)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/motah-fard/geom3d)](https://github.com/motah-fard/geom3d/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/motah-fard/geom3d)](https://goreportcard.com/report/github.com/motah-fard/geom3d)

`geom3d` is a practical Go library for 3D geometry and spatial calculations.

It provides a focused set of tools for working with 3D vectors, geometric primitives, projections, distances, intersections, and rigid transforms in Go. The package is designed for engineering, simulation, robotics-adjacent workflows, biomechanics, and spatial analytics.

The library is intentionally small, explicit, and easy to use.

## Features

- `Vec3` for 3D vector and point operations
- Geometry primitives:
  - `Ray3`
  - `Segment3`
  - `Plane`
  - `Triangle`
  - `AABB`
- Core operations:
  - dot product
  - cross product
  - norm and normalization
  - distance calculations
  - projections onto planes and lines
  - closest point on a ray, segment, triangle, and AABB
  - closest points between segments
  - ray-plane intersection
  - segment-plane intersection
  - segment-segment intersection at a single point
  - collinear segment overlap detection
  - ray-AABB intersection with hit interval output (`hit`, `tMin`, `tMax`)
  - barycentric coordinates
  - point-to-ray, point-to-segment, point-to-line, point-to-triangle, point-to-AABB, and segment-to-segment distance queries
- 3D rotations with `Mat3`
- Rigid transforms with `Transform`
  - apply to points and vectors
  - compose transforms
  - invert transforms

## Why this library exists

The Go ecosystem already has solid low-level math and graphics-oriented packages, but there is still room for a small, practical geometry library focused on everyday 3D calculations.

`geom3d` targets that middle layer: the kinds of operations developers often need in real applications, without requiring a rendering engine, graphics framework, or heavyweight geometry stack.

Typical use cases include:

- projecting points onto planes or lines
- finding the closest point on a ray, segment, triangle, or bounding box
- computing closest points or minimum distance between segments
- checking ray intersections with planes or bounding boxes
- testing whether two segments intersect at a single point or overlap collinearly
- computing barycentric coordinates for triangle-based workflows
- applying and composing rigid transforms
- working with coordinate frames in engineering or sensor-based applications

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
`Vec3` supports common 3D vector operations such as addition, subtraction, scaling, dot products, cross products, norms, distances, and normalization.

### Primitives
The package includes practical 3D primitives for common geometric workflows:

- `Ray3`
- `Segment3`
- `Plane`
- `Triangle`
- `AABB`

### Matrices and transforms
`Mat3` supports 3D rotation matrices and matrix operations.  
`Transform` supports rigid-body transforms for points and vectors, transform composition, and inversion.

### Geometric helpers
The package includes helpers for:

- point-to-plane distance
- point-to-ray distance
- point-to-segment distance
- point-to-line distance
- point-to-triangle distance
- point-to-AABB distance
- segment-to-segment distance
- point projection to planes and lines
- barycentric coordinates
- closest-point queries on rays, segments, triangles, and AABBs
- closest-point queries between segments
- ray-plane intersection
- segment-plane intersection
- segment-segment intersection
- collinear segment overlap detection
- ray-AABB intersection

## Behavior notes

A few helpers have intentionally specific semantics:

- `IntersectSegments` reports only **single-point** intersections. If two segments overlap over a non-zero interval, it returns `false`.
- `SegmentsOverlap` reports only **collinear overlap over a non-zero interval**. Endpoint-only touching does not count as overlap.
- `IntersectRayAABB` returns `hit, tMin, tMax`. If the ray starts inside the box, `tMin` may be `0`.
- `ClosestPointOnRay` clamps to the ray origin when the orthogonal projection falls behind the origin.
- `ClosestPointOnAABB` returns the input point itself when the point lies inside the box.
- Helpers that receive invalid rays, invalid AABBs, or other degenerate inputs document their fallback behavior in GoDoc.

## Examples

Runnable examples are included under the `examples/` directory, including:

- `basic_vectors`
- `ray_plane`
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
- `triangle_normal`
- `triangle_closest_point`
- `triangle_barycentric`
- `transform_point`

## API stability

`geom3d` is approaching `v1.0.0`.

The current focus is API stability, edge-case confidence, and documentation clarity. Remaining changes before `v1.0.0` should be small, deliberate, and centered on correctness rather than feature growth.

## License

MIT
