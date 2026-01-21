## 2023-10-27 - Remove redundant os.MkdirAll call in scaffold.InstallFS
**Issue:** The `InstallFS` function in `pkg/utils/scaffold/scaffold.go` contained a redundant call to `os.MkdirAll` for each file it processed.
**Root Cause:** The code was likely written without fully relying on the behavior of `fs.WalkDir`, which guarantees that directories are visited before the files inside them. This led to a defensive, but unnecessary, check to ensure the parent directory existed.
**Solution:** I removed the superfluous `os.MkdirAll` call that was being executed for every file. The preceding logic already ensures that all necessary directories are created.
**Pattern:** When using file system walkers like `fs.WalkDir`, trust their contracts. `WalkDir` processes directories first, so manual parent directory creation for files is often redundant and can be omitted for cleaner code.

## 2026-01-21 - Simplify module name matching logic
**Issue:** The `moduleMatchesName` function in `pkg/core/resolver.go` contained redundant checks for parent directory and base name before iterating over all path components.
**Root Cause:** Defensive programming or iterative development where specific cases (parent dir, base name) were handled explicitly before a general solution (looping over all parts) was added, leaving the specific checks as dead/redundant code.
**Solution:** I removed the explicit checks for `parentDir` and `dirName`. The loop that checks if `name` matches any part of the split path already covers these cases (and more). I also added a unit test to verify correctness.
**Pattern:** Generalize early. If a general loop covers specific edge cases, remove the specific checks to reduce code volume and cognitive load.
