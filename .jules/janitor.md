## 2023-10-27 - Remove redundant os.MkdirAll call in scaffold.InstallFS
**Issue:** The `InstallFS` function in `pkg/utils/scaffold/scaffold.go` contained a redundant call to `os.MkdirAll` for each file it processed.
**Root Cause:** The code was likely written without fully relying on the behavior of `fs.WalkDir`, which guarantees that directories are visited before the files inside them. This led to a defensive, but unnecessary, check to ensure the parent directory existed.
**Solution:** I removed the superfluous `os.MkdirAll` call that was being executed for every file. The preceding logic already ensures that all necessary directories are created.
**Pattern:** When using file system walkers like `fs.WalkDir`, trust their contracts. `WalkDir` processes directories first, so manual parent directory creation for files is often redundant and can be omitted for cleaner code.

## 2026-01-22 - Simplify Module Name Resolution Logic
**Issue:** `NewDatabaseModule`, `NewFrontendModule`, and `NewGraphQLModule` contained redundant logic where variables were re-assigned to the same values they already held. The `database` module also had dead "fallback" code that was immediately overwritten.
**Root Cause:** The code was likely copied and pasted with modifications that didn't fully account for the initialization logic, or attempted to handle edge cases defensively in a way that became redundant.
**Solution:** Simplified the module name resolution to a single line: `filepath.Base(filepath.Dir(path))`. This covers all standard use cases (`internal/MODULE/db` -> `MODULE`) correctly and is much easier to read. Added tests to verify behavior.
**Pattern:** Always verify if a conditional assignment actually changes the value. If `x = y` is followed by `if condition { x = y }`, the second assignment is redundant.
