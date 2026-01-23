# Code Style & Architecture

## Overview
This document outlines the established code style, architectural patterns, and technology stack for the repository. All new code should adhere to these conventions to maintain consistency.

## Key Files
- `main.go` - The entry point for dependency injection and application startup.
- `cmd/` - Application entry points (e.g., `cmd/discord`, `cmd/frontend`).
- `internal/` - Private application logic, strictly separated from `cmd`.
- `internal/database` - Database models and access layer using `go-jet`.
- `Taskfile.yaml` - Task definitions for building, testing, and running the project.

## How It Works

### Architecture
The project follows a standard Go monorepo layout:
1.  **Entry Points (`cmd/`)**: Each directory here is a main application or tool. They wire together dependencies but contain minimal business logic.
2.  **Internal Logic (`internal/`)**: Contains the core business logic. Code here is library-agnostic where possible.
3.  **Dependency Injection**: DI is performed manually in `main.go`. Interfaces (like `core.Client`, `database.Client`) are defined where they are used or in a common package, and implementations are injected at startup.

### Frontend
The frontend uses a modern server-side rendering approach:
-   **Templ**: Type-safe HTML generation (`.templ` files).
-   **HTMX**: Used for interactivity and dynamic updates without full page reloads.
-   **TailwindCSS + DaisyUI**: Utility-first CSS framework for styling.
-   **Assets**: Static assets are embedded or served via `internal/assets`.

### Database
-   **PostgreSQL**: The primary data store.
-   **go-jet**: Used for type-safe SQL query generation.
-   **Atlas**: Used for schema migrations (`atlas migrate ...`).

## Patterns Used

### Error Handling
-   **Wrapping**: Use `github.com/pkg/errors` to wrap errors with context.
    -   *Good*: `return errors.Wrap(err, "failed to fetch user")`
    -   *Bad*: `return err` (unless it's the top level)

### Logging
-   **Zerolog**: Use `github.com/rs/zerolog` for structured JSON logging.
    -   Pattern: `log.Info().Str("key", "value").Msg("message")`

### Concurrency
-   **Synchronization**: Heavy use of `sync.WaitGroup` for parallel tasks (e.g., fetching stats for multiple players).
-   **Context**: `context.Context` is passed down the call stack for cancellation and timeouts.
-   **Safety**: Use `sync.Mutex` to protect shared maps or resources in concurrent operations.

### Configuration
-   **Environment Variables**: Config is loaded from `.env` files into `internal/constants` or directly in `main.go`.

## Gotchas
-   **Manual wiring**: Adding a new service requires manually updating `main.go` to initialize and inject it.
-   **Templ Generation**: You must run `go generate` or the appropriate Taskfile command to regenerate `.templ` files after changes.
-   **Database Models**: `go-jet` models are generated. Do not edit them manually; use `task db-generate`.

## Related
-   [Documentation README](../README.md)
-   `AGENTS.md` - Context for AI agents working in this repo.
