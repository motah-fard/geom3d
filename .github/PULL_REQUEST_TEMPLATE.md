## Summary

<!-- What does this PR change, and why? -->

## Checklist

- [ ] `gofmt -l .` prints nothing
- [ ] `go vet ./...` passes
- [ ] `go test ./... -race -cover` passes
- [ ] New/changed exported symbols have doc comments, including fallback
      behavior for invalid/degenerate input
- [ ] Tests added for new behavior, including edge cases
- [ ] `README.md` updated (feature list / behavior notes) if public API changed
- [ ] `CHANGELOG.md` updated under `[Unreleased]` if public API changed
- [ ] `API_AUDIT.md` updated if a new public symbol was added
- [ ] A runnable example under `examples/` was added, if applicable

## Related issue

<!-- Link the issue this addresses, if any -->
