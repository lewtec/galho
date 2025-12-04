# Galho sample project

This project is the template created by [galho](https://github.com/lewtec/galho) init.

Run `galho --help` for more details.

## Project rename checklist
- [ ] .github/workflows/autorelease.yaml: Look for build/app-$GOOS-$GOARCH, replace app to your binary name
- [ ] cmd/app: Rename the folder to your binary name
- [ ] Dockerfile: `go build -o app ./cmd/app` and the `COPY --from=builder`
- [ ] Look for GALHO and rename to your project URL (ex: github.com/you/project)
