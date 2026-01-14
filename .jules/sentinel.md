## 2024-08-05 - Fix Path Traversal Vulnerability in `scaffold.InstallFS`

**Vulnerability:** A path traversal vulnerability was identified in the `InstallFS` function in `pkg/utils/scaffold/scaffold.go`. The function did not properly sanitize or validate the file paths of the files being installed, allowing for files to be written outside of the intended destination directory.

**Learning:** The vulnerability existed because the code concatenated a user-provided path with a destination directory without first cleaning and validating the path. This allowed for traversal sequences like `../../` to be used to escape the destination directory.

**Prevention:** To prevent this vulnerability, all user-provided file paths must be cleaned and validated to ensure they are within the intended directory before being used. The fix implemented this by using `filepath.Clean` and checking that the resulting path has the destination directory as a prefix. This ensures that no matter what input is provided, the file will always be written within the intended destination.
