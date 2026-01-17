# Consistently Ignored Changes

This file lists patterns of changes that have been consistently rejected by human reviewers. All agents MUST consult this file before proposing a new change. If a planned change matches any pattern described below, it MUST be abandoned.

---

## IGNORE: Replacing `GALHO/app` Placeholder

**- Pattern:** Do not replace the `GALHO/app` module name in Go template files.
**- Justification:** This placeholder is intentionally used by the scaffolding logic to generate new modules. Replacing it with the actual project module path (`github.com/lewtec/galho`) breaks the template engine. This change has been attempted and rejected multiple times.
**- Files Affected:** `pkg/entities/core/_template/go.mod.tmpl` (and other `.tmpl` files)

---

## IGNORE: Adding HSTS Security Header

**- Pattern:** Do not add the `Strict-Transport-Security` (HSTS) header.
**- Justification:** This change has been proposed multiple times and rejected. It is considered an operational risk because it can break subdomains that do not support HTTPS. This is a classic example of a security enhancement that, while good in theory, is a breaking change in this specific context.
**- Files Affected:** `src/hooks.server.ts`

---
