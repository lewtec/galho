## 2026-01-22 - Prevent Zip Slip in Scaffold

**Vulnerability:** The `InstallFS` function in `pkg/utils/scaffold/scaffold.go` was vulnerable to Zip Slip (Path Traversal) attacks. It blindly joined the destination directory with the path from the provided `fs.FS` without verifying if the resulting path was still inside the destination directory. If a malicious `fs.FS` provided a path with `..`, it could write files outside the intended directory.

**Learning:** Even when using `io/fs` abstractions which are generally safe (like `fstest.MapFS` or `embed.FS`), the `InstallFS` function is a public utility that could be used with any implementation of `fs.FS`. Relying on the implicit safety of the input interface implementation is not Defense in Depth. Explicitly validating the destination path prevents this class of vulnerabilities entirely, regardless of the input source.

**Prevention:** Always validate that the resolved absolute path of a file write operation starts with the resolved absolute path of the intended destination directory. Use `filepath.Abs` and `strings.HasPrefix` (with path separator check) to enforce this boundary.
