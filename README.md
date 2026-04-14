# geom3d

`geom3d` is a practical Go library for 3D geometry and spatial calculations.

It provides a focused set of tools for working with 3D vectors, geometric primitives, projections, distances, intersections, and rigid transforms in Go. The package is designed for engineering, simulation, robotics-adjacent, biomechanics, and spatial analytics workflows.

It is intentionally small, explicit, and easy to use.

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
  - closest point on a segment
  - ray-plane intersection
  - segment-plane intersection
  - ray-AABB intersection
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
- finding the closest point on a segment
- checking ray intersections with planes or bounding boxes
- applying and composing rigid transforms
- working with coordinate frames in engineering or sensor-based applications

## Installation

```bash
go get github.com/motah-fard/geom3d

## Non-goals

geom3d is not:
- a rendering engine
- an OpenGL helper library
- a physics engine
- a mesh loader
- a CAD kernel

