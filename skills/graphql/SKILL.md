---
name: galho-graphql
description: Generate Galho GraphQL modules. Use this skill when the user wants to scaffold a new GraphQL module (gqlgen).
---

# Galho GraphQL Skill

This skill provides access to the GraphQL scaffolding capabilities of the Galho framework.

## CLI Wrapper

The CLI is accessed via the wrapper script: `skills/graphql/scripts/galho`.

## Capabilities

### 1. Generate GraphQL Module
Scaffold a new GraphQL module using gqlgen.

```bash
skills/graphql/scripts/galho generate graphql <path/to/module>
```

Example:
```bash
skills/graphql/scripts/galho generate graphql pkg/entities/api
```
