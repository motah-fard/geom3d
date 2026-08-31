# Contributing to geom3d

Thanks for considering a contribution. `geom3d` is intentionally small and
explicit, so the bar for changes is about correctness and consistency more
than volume of code.

## Before you start

For anything beyond a small fix (a new primitive, a new query, a signature
change), please open an issue first to discuss the approach. This avoids
wasted work on a PR that doesn't fit the project's scope — see the
[README's "Non-goals" section](README.md#non-goals) and
[API_AUDIT.md](API_AUDIT.md) for the current design direction.

## Development setup

```bash
git clone https://github.com/motah-fard/geom3d.git
cd geom3d
go build ./...
go test ./...
```

No external dependencies are required; the module only uses the standard
library.

## Before opening a PR

Run these locally and make sure they all pass:

```bash
gofmt -l .        # should print nothing
go vet ./...
go build ./...
go test ./... -race -cover
```

CI runs the same checks (build, vet, gofmt, race-enabled tests) on Linux,
macOS, and Windows for every PR.

## Code conventions

- Follow standard Go style; run `gofmt` before committing.
- Every exported type, function, and method needs a doc comment. Document
  fallback behavior for invalid or degenerate inputs explicitly (see
  existing functions for the pattern, e.g. `ClosestPointOnAABB`).
- Match the existing API shape described in [API_AUDIT.md](API_AUDIT.md):
  - Behavior intrinsic to a single primitive is a **method**
    (`Vec3.Norm`, `Triangle.Area`, `AABB.Contains`).
  - Relations between multiple objects are **free functions**
    (`DistancePointToPlane`, `ClosestPointOnTriangle`, `IntersectRayPlane`).
- Prefer returning zero values and a `bool`/fallback over panicking or
  returning an `error`. Degenerate or invalid input should degrade
  predictably, not crash — and the doc comment should say what happens.
- Use `AlmostZero` / `AlmostEqual` from `epsilon.go` for floating-point
  comparisons instead of exact equality.

## Tests

- Add table-driven or individual `Test...` functions for new behavior,
  including edge cases (degenerate/invalid input, parallel/collinear
  cases, boundary conditions).
- Add a `godoc` `Example...` function for new exported functions where a
  short usage example is illustrative (see `intersection_test.go` for the
  pattern). These are checked by `go test` via their `// Output:` comments.
- If you add a new top-level query or primitive, consider adding a runnable
  program under `examples/`, matching the style of the existing examples.

## Commit messages

This repo loosely follows [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add sphere-ray intersection
fix: correct AABB overlap check for touching boxes
docs: update README for sphere queries
refactor: simplify closest-point clamping
```

## Updating docs alongside code

A PR that adds or changes public API should typically touch:

- `README.md` (feature list, and "Behavior notes" if the function has
  non-obvious fallback semantics)
- `CHANGELOG.md` (add an entry under `[Unreleased]`)
- `API_AUDIT.md` (note the addition under the relevant section)
- `doc.go` if it changes the package-level feature summary

## Reporting bugs / requesting features

Please use the issue templates under `.github/ISSUE_TEMPLATE/`. For bugs,
a minimal reproducing example (ideally as a Go playground link or a short
snippet) is the most useful thing you can include.
