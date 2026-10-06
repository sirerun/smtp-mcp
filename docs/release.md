# Releases

The Go module is `github.com/sirerun/smtp-mcp`. Release tags use semantic versions
such as `v0.1.0`; `main.version` is set by GoReleaser and exposed by `--version`
and the MCP initialization response.

Before tagging an independently reviewed revision, run:

```sh
go mod tidy
git diff --exit-code -- go.mod go.sum
go test -race ./...
go vet ./...
golangci-lint run ./...
govulncheck ./...
goreleaser check
```

Use golangci-lint v2.13.2 and govulncheck v1.8.0 or a qualified newer release.
A successful vulnerability scan means no reachable advisories were found at scan
time; it does not guarantee that every dependency is vulnerability-free.

GoReleaser builds six archives: Linux, macOS and Windows, each on amd64 and arm64.
The archives include the executable, README and Apache 2.0 license. Checksums are
SHA-256. Cross-compilation is not runtime testing on every target platform.

After merging and verifying the landed revision, tag it and publish:

```sh
git tag -a v0.1.0 -m 'smtp-mcp v0.1.0'
git push origin v0.1.0
goreleaser release
```

Provide `GITHUB_TOKEN` securely in the publishing environment; never commit it.
Alternatively, generate archives locally with `goreleaser release --skip=publish`
and upload the reviewed archives and checksums with the GitHub CLI. Download the
published assets and verify their checksums and the native binary's version and
stdio protocol before recording the release as complete.

When hosted CI is unavailable, record the exact revision, local commands and
results in the release notes. Do not describe local checks as hosted CI success.
