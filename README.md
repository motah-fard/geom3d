# geom3d

`geom3d` is a practical Go library for 3D geometry and spatial calculations.

It provides vectors, planes, rays, segments, triangles, axis-aligned bounding
boxes, rotation matrices, rigid transforms, and common operations such as
distances, projections, closest-point queries, and intersections.

`geom3d` is designed for engineering, simulation, robotics-adjacent, and
analytics workflows rather than rendering or game engines.

## Why this library exists

The Go ecosystem already has strong low-level math and graphics-oriented
packages. `geom3d` focuses on the practical middle layer: the geometry
operations developers often need in real applications.

Examples include:
- projecting points onto planes or lines
- finding the closest point on a segment
- intersecting rays with planes or boxes
- applying and composing rigid transforms

## Install

```bash
go get github.com/motah-fard/geom3d

## Non-goals

geom3d is not:
- a rendering engine
- an OpenGL helper library
- a physics engine
- a mesh loader
- a CAD kernel

