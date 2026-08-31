---
name: Feature request
about: Suggest a new primitive, query, or capability for geom3d
title: ""
labels: enhancement
assignees: ""
---

**What's the use case?**
Describe the real-world workflow or problem this would solve (engineering,
simulation, robotics, biomechanics, spatial analytics, etc.). Concrete
examples are more useful than abstract ones.

**Proposed API**
A rough sketch of the function/type signature you have in mind, following
the existing conventions in [API_AUDIT.md](../../API_AUDIT.md) (methods for
single-primitive behavior, free functions for relations between objects).

```go
func ExampleNewFunction(...) (...)
```

**Does it fit geom3d's scope?**
`geom3d` intentionally excludes rendering, physics, mesh loading, and CAD
kernel features (see the README's "Non-goals" section). Briefly explain why
this belongs here rather than in a more specialized package.

**Alternatives considered**
Any existing workaround using the current API, or other libraries that
provide this.
