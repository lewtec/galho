# Consistently Ignored Changes

This file lists patterns of changes that have been consistently rejected by human reviewers. All agents MUST consult this file before proposing a new change. If a planned change matches any pattern described below, it MUST be abandoned.

---

## IGNORE: Replacing `GALHO/app` Placeholder

**- Pattern:** Do not replace the `GALHO/app` module name in Go template files.
**- Justification:** This placeholder is intentionally used by the scaffolding logic to generate new modules. Replacing it with the actual project module path (`github.com/lewtec/galho`) breaks the template engine. This change has been attempted and rejected multiple times.
**- Files Affected:** `pkg/entities/core/_template/go.mod.tmpl` (and other `.tmpl` files)

---

## IGNORE: Adding Zip Slip/Path Traversal Checks in Scaffold

**- Pattern:** Do not add `filepath.Abs` and `strings.HasPrefix` validation checks to `InstallFS` to prevent path traversal (Zip Slip).
**- Justification:** These changes are consistently rejected. The `fs.FS` input to `InstallFS` is considered trusted (embedded or internal), and the project maintainers do not want this specific validation logic added here.
**- Files Affected:** `pkg/utils/scaffold/scaffold.go`

---

## IGNORE: Unsolicited Logic Simplification

**- Pattern:** Do not refactor core logic (like `moduleMatchesName` or `NewDatabaseModule`) purely to simplify code that is already working (e.g., removing explicit parent directory checks).
**- Justification:** The maintainers prefer the existing explicit logic over simplified versions. "If it ain't broke, don't fix it."
**- Files Affected:** `pkg/core/resolver.go`, `pkg/entities/*/module.go`

---

## IGNORE: Global Linter/Formatter Reconfiguration

**- Pattern:** Do not submit PRs that purely reconfigure global linters (ESLint, Prettier), add a root `package.json`, or overhaul `mise.toml` tasks, especially if they involve large-scale reformatting of YAML or Go files.
**- Justification:** These changes are consistently rejected as noise. The project likely has its own conventions or manages these dependencies differently.
**- Files Affected:** `mise.toml`, `package.json`, `eslint.config.mjs`, `.github/workflows/*.yml`

---
