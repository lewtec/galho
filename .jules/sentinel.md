## 2026-01-17 - Prevent Zip Slip Vulnerability in Scaffold
**Vulnerability:** The `InstallFS` function in `pkg/utils/scaffold/scaffold.go` did not verify that the destination path for extracted files was contained within the target directory. This could potentially allow path traversal (Zip Slip) if a malicious filesystem was provided.
**Learning:** Even when using trusted embedded filesystems, utility functions should be robust against misuse or malicious inputs. Checking that the resolved path starts with the intended destination path is a standard mitigation.
**Prevention:** Always validate that the final path of a file operation is within the expected directory boundaries using `filepath.Abs` and `strings.HasPrefix`.
