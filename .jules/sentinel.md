## 2023-10-27 - Initial Sentinel Journal Entry

**Vulnerability:** Path Traversal in `pkg/utils/scaffold/scaffold.go`. The `InstallFS` function is vulnerable to path traversal because it doesn't validate that the destination path for file creation is within the intended destination directory. A malicious actor could craft a file path with `../` to write files outside of the intended directory.

**Learning:** This vulnerability likely exists because the developer trusted the contents of the `fs.FS` being walked. The focus was on functionality, and the security implications of joining paths without validation were overlooked.

**Prevention:** Always validate and sanitize user-controllable or otherwise potentially malicious paths. When writing files, ensure that the resolved absolute path is a child of the intended destination directory. Go's `filepath.Clean` and checking for `../` components are good practices, but a prefix check on the resolved path is a robust solution.
