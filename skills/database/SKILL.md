---
name: galho-database
description: Manage Galho database modules (generate, migrate). Use this skill when the user wants to work with databases, create migrations, or scaffold database modules.
---

# Galho Database Skill

This skill provides access to the database capabilities of the Galho framework.

## CLI Wrapper

The CLI is accessed via the wrapper script: `skills/database/scripts/galho`.

## Capabilities

### 1. Generate Database Module
Scaffold a new database module with default configuration.

```bash
skills/database/scripts/galho generate database <path/to/module>
```

Example:
```bash
skills/database/scripts/galho generate database pkg/entities/users
```

### 2. Manage Migrations
Manage database migrations using the `db migration` command.

- **Create Migration:**
  ```bash
  skills/database/scripts/galho db migration create <name> --module <path/to/module>
  ```

- **List Migrations:**
  *(Check help for other commands)*
  ```bash
  skills/database/scripts/galho db migration --help
  ```
