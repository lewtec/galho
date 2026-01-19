## 2023-10-27 - Remove redundant os.MkdirAll call in scaffold.InstallFS
**Issue:** The `InstallFS` function in `pkg/utils/scaffold/scaffold.go` contained a redundant call to `os.MkdirAll` for each file it processed.
**Root Cause:** The code was likely written without fully relying on the behavior of `fs.WalkDir`, which guarantees that directories are visited before the files inside them. This led to a defensive, but unnecessary, check to ensure the parent directory existed.
**Solution:** I removed the superfluous `os.MkdirAll` call that was being executed for every file. The preceding logic already ensures that all necessary directories are created.
**Pattern:** When using file system walkers like `fs.WalkDir`, trust their contracts. `WalkDir` processes directories first, so manual parent directory creation for files is often redundant and can be omitted for cleaner code.

## 2026-01-19 - Simplify migration listing logic
**Issue:** `listMigrations` in `pkg/entities/database/migration.go` used manual bubble sort and error-prone manual string slicing to parse filenames.
**Root Cause:** Likely implemented before `sort` package was considered or just legacy code.
**Solution:** Replaced manual sort with `sort.Strings` and string parsing with `strings.HasSuffix`/`strings.TrimSuffix`.
**Pattern:** Use standard library functions (`sort`, `strings`) instead of re-implementing basic algorithms and parsing logic.
