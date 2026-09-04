# Roadmap

This is a living list of what's likely to come next, and — just as
importantly — what's been deliberately left out and why. It's meant to give
potential contributors a real starting point rather than having to guess
what would be welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) before
picking something up, and please open an issue first for anything larger
than a small fix.

For the detailed, ongoing API-design log (what shipped in which version,
method-vs-free-function conventions, etc.), see [API_AUDIT.md](API_AUDIT.md).
This file is the higher-level "what's next" view.

## Likely next additions

These are scoped, don't require a `v2`, and would be welcome contributions:

- **AABB-OBB and OBB-OBB intersection.** Deliberately not implemented yet —
  an exact test needs the full 15-axis separating-axis theorem (3
  face-normal axes per box, plus 9 pairwise edge-cross-edge axes), and
  those 9 cross-axis cases are the most error-prone part of the standard
  algorithm to get right without a reference implementation to check
  against. A careful, well-tested PR here (ideally cross-checked against
  brute-force sampling, the way `IntersectRayCapsule`'s tests are) would be
  very welcome.
- **Ray-triangle-mesh acceleration** is explicitly out of scope (see
  "Non-goals" in the README) — no BVH, no spatial partitioning. If your use
  case needs that, this library is meant to be the primitive layer
  underneath it, not the acceleration structure itself.
- Additional closest-point/distance pairs that don't exist yet: ray-to-ray,
  ray-to-segment, segment-to-triangle, segment-to-AABB. None of these are
  hard, they just haven't been asked for yet.
- **`Mat3.Mul` is slower than it needs to be.** [`BENCHMARKS.md`](BENCHMARKS.md)
  measured it at ~2.8x slower than mathgl's equivalent: geom3d stores
  `Mat3` as `[3][3]float64` and multiplies with a triple-nested loop,
  while mathgl uses a flat `[9]float64` with a hand-unrolled multiply.
  Flattening `Mat3`'s internal storage and hand-unrolling `Mul` (and
  likely `MulVec`) is purely an internal change — the public `Mat3.M
  [3][3]float64` field would need to become a method-based accessor or
  be dropped in favor of indexed access, which **would** be a breaking
  change requiring a `v2`, so this needs a deliberate decision, not a
  quiet PR.
- JSON marshaling for the basic types (`Vec3`, `Sphere`, etc.), if there's
  real demand for it — not adding speculatively.

## Explicitly not planned

- A rendering engine, OpenGL/graphics helpers, a physics engine, a mesh
  loader, or a CAD kernel (see the README's "Non-goals").
- Making the package generic over `float32`/`float64`. `float64` throughout
  keeps the API simple; a `float32` variant would be a separate module, not
  a generics retrofit of this one.
- Changing the fixed `Epsilon`-based tolerance used internally to something
  scale-aware, without a `v2`. See the README's "Numerical tolerance and
  large coordinates" section — `AlmostEqualRelative`/`AlmostZeroAtScale`
  exist for your own use, but retrofitting the package's own internal
  checks would silently change the observable behavior of already-frozen
  `v1` functions.

## Process notes

- `geom3d` is past `v1.0.0`: existing signatures are frozen (see the
  README's "API stability" section). New functionality lands as additive
  minor releases.
- A breaking change (signature change, removed symbol, or a deliberate
  change to `Epsilon`'s scale-awareness) needs a `v2`, and should start as
  an issue discussing the tradeoff, not a PR.
