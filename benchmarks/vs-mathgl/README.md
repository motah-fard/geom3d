# vs-mathgl

Benchmarks comparing geom3d against
[go-gl/mathgl](https://github.com/go-gl/mathgl) on the small set of
operations the two libraries genuinely share: basic `Vec3` arithmetic,
`Mat3` multiplication, and `Quaternion` multiply/rotate/slerp.

This is its own Go module, not part of the main `geom3d` module, so that
mathgl never appears in `geom3d`'s own `go.mod` — the main library is
deliberately dependency-free (stdlib only), and a benchmark-only
comparison shouldn't compromise that for every consumer of `go get
github.com/motah-fard/geom3d`.

Run it with:

```bash
go test -bench=. -benchmem -run '^$' .
```

See [`../../BENCHMARKS.md`](../../BENCHMARKS.md) for the latest recorded
results and how to read them.
