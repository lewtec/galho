---
name: galho-frontend
description: Generate Galho frontend modules. Use this skill when the user wants to scaffold a new frontend module (React/Vite).
---

# Galho Frontend Skill

This skill provides access to the frontend scaffolding capabilities of the Galho framework.

## CLI Wrapper

The CLI is accessed via the wrapper script: `skills/frontend/scripts/galho`.

## Capabilities

### 1. Generate Frontend Module
Scaffold a new frontend module with React, Vite, TailwindCSS, and DaisyUI.

```bash
skills/frontend/scripts/galho generate frontend <path/to/module>
```

Example:
```bash
skills/frontend/scripts/galho generate frontend pkg/entities/web
```
