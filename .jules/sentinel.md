# Sentinel's Journal

## 2026-01-21 - Prevent Zip Slip in Scaffold
**Vulnerability:** The `InstallFS` function in `pkg/utils/scaffold` blindly joins the destination path with the file path from the source FS. While standard `fs.FS` implementations enforce safe paths, a compromised or non-standard FS implementation could supply paths containing `..`, leading to writing files outside the intended destination directory (Zip Slip vulnerability).
**Learning:** Always validate that the final destination path resides within the intended directory boundary when performing file operations based on external inputs, even when using abstractions like `fs.FS`.
**Prevention:** Use `filepath.Abs` to resolve paths and check that the destination path starts with the intended root directory path.
