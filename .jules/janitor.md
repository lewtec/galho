## 2023-10-27 - Remove redundant os.MkdirAll call in scaffold.InstallFS
**Issue:** The `InstallFS` function in `pkg/utils/scaffold/scaffold.go` contained a redundant call to `os.MkdirAll` for each file it processed.
**Root Cause:** The code was likely written without fully relying on the behavior of `fs.WalkDir`, which guarantees that directories are visited before the files inside them. This led to a defensive, but unnecessary, check to ensure the parent directory existed.
**Solution:** I removed the superfluous `os.MkdirAll` call that was being executed for every file. The preceding logic already ensures that all necessary directories are created.
**Pattern:** When using file system walkers like `fs.WalkDir`, trust their contracts. `WalkDir` processes directories first, so manual parent directory creation for files is often redundant and can be omitted for cleaner code.

## 2026-01-23 - Simplify module matching logic in core
**Issue:** The `moduleMatchesName` function in `pkg/core/resolver.go` contained redundant checks for parent directory and base name, followed by a loop that checked all path components.
**Root Cause:** Likely defensive coding or iterative development where specific checks were added before a general loop, or simply a misunderstanding that the loop covers the specific cases.
**Solution:** Removed the specific checks for `parentDir` and `dirName`, relying solely on splitting the path and iterating over components. Added a test suite to verify correctness.
**Pattern:** When implementing search logic over hierarchical data (like paths), a general iteration often covers specific "head" or "tail" checks. Verify redundant logic with tests before removing.
