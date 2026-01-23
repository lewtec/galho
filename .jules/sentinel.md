## 2026-01-23 - Path Traversal in Database Migration Creation
**Vulnerability:** The `db migration create <name>` command was vulnerable to path traversal. A malicious name like `../../evil` could write files outside the migrations directory.
**Learning:** CLI tools that accept filenames as arguments must sanitize them, even if the user is trusted. `filepath.Join` does not prevent path traversal if the input itself contains parent directory references that resolve to a valid path.
**Prevention:** Always validate user-provided filenames to ensure they do not contain path separators or parent directory references (`..`). Use `filepath.Base` or explicit character validation.
