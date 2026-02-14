# Bookstore API — Project Guide

## Overview

A RESTful bookstore API built with Go and [Chi](https://github.com/go-chi/chi). The project uses an in-memory data store (no database) and auto-generates OpenAPI documentation via [swaggo/swag](https://github.com/swaggo/swag).

## Directory Structure

```
cmd/server/main.go          — Entry point, router setup, swagger annotations
internal/handlers/           — HTTP handler functions (one file per resource)
internal/models/             — Data structures (Book, Review, etc.)
internal/store/              — In-memory data store
internal/middleware/          — HTTP middleware (logging, etc.)
docs/                        — Generated swagger spec + Swagger UI shell
```

## Running the Project

```bash
# Run the server
go run cmd/server/main.go

# Run tests
go test ./...

# Run linter
golangci-lint run ./...

# Generate swagger docs
swag init -g cmd/server/main.go
```

## Go Conventions

- **Error handling**: Always check errors. Wrap errors from external packages with `fmt.Errorf("context: %w", err)`.
- **Naming**: Use camelCase for unexported, PascalCase for exported. Acronyms stay uppercase (e.g., `ID`, `URL`).
- **Comments**: All exported types and functions must have doc comments. Comments must end with a period.
- **Blank lines**: Add a blank line before `return` statements (`nlreturn` linter).
- **nolint directives**: Must include a specific linter name and an explanation.

## Testing Requirements

- All new handlers must have table-driven tests.
- Use `net/http/httptest` for handler tests.
- Tests must cover success and error paths.
- Target >80% coverage for new code.

## Swagger Annotations

All endpoints must have swaggo annotations. Example:

```go
// @Summary      Short summary
// @Description  Longer description.
// @Tags         resource-name
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Resource ID"
// @Success      200  {object}  models.Resource
// @Failure      404  {object}  handlers.ErrorResponse
// @Router       /api/resource/{id} [get]
```

After adding or modifying annotations, regenerate docs: `swag init -g cmd/server/main.go`

## PR Conventions

- **Branch naming**: `feature/BOOK-<id>-<short-description>` (e.g., `feature/BOOK-1-add-reviews`)
- **Commit messages**: Conventional commits with Jira ID — `feat(api): add reviews endpoint [BOOK-1]`
- **PR body**: Include a summary of changes, test plan, and link to the Jira ticket.

## CI/CD

- **PR checks** (`.github/workflows/pr.yml`): tests, golangci-lint (strict), build.
- **Build** (`.github/workflows/build.yml`): tests, build binary, generate swagger docs.
- **Deploy** (`.github/workflows/deploy.yml`): generates swagger spec and deploys `docs/` to GitHub Pages.
