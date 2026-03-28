## 2023-10-27 - Remove redundant os.MkdirAll call in scaffold.InstallFS
**Issue:** The `InstallFS` function in `pkg/utils/scaffold/scaffold.go` contained a redundant call to `os.MkdirAll` for each file it processed.
**Root Cause:** The code was likely written without fully relying on the behavior of `fs.WalkDir`, which guarantees that directories are visited before the files inside them. This led to a defensive, but unnecessary, check to ensure the parent directory existed.
**Solution:** I removed the superfluous `os.MkdirAll` call that was being executed for every file. The preceding logic already ensures that all necessary directories are created.
**Pattern:** When using file system walkers like `fs.WalkDir`, trust their contracts. `WalkDir` processes directories first, so manual parent directory creation for files is often redundant and can be omitted for cleaner code.

## 2024-07-25 - Remove Dead Code from main.go
**Issue:** The application entrypoint at `cmd/galho/main.go` contained an import for a non-existent package (`github.com/lewtec/galho/cmd/galho/entities`) and a call to a function from that package.
**Root Cause:** This was likely a remnant from a previous refactoring. The `entities` command package was probably removed or relocated, but the corresponding import and function call in `main.go` were not cleaned up, creating dead code.
**Solution:** I removed the unused import statement and the associated `entities.AddEntityCommands(Command)` function call to eliminate the "ghost" import and prevent future confusion.
**Pattern:** After refactoring or deleting files, always check for and remove any lingering unused imports or function calls. This keeps the dependency graph clean and the codebase easier to understand.
