# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

Go URL shortener API. Chi router + GORM/Postgres + Redis, JWT auth. Module path is `github.com/mjrdev` (not the repo name), so internal imports look like `github.com/mjrdev/internal/service`.

## Commands

```bash
docker compose up -d              # Postgres 17 on :5432, Redis 7.4 on :6379
cp .env.example .env              # required — the API log.Fatal's without a .env file

go run ./cmd/migrate              # apply migrations
go run ./cmd/migrate -rollback    # roll back the last one
go run ./cmd/api                  # run API on :3000
air -c .air.toml                  # run with live reload (builds to ./tmp/main)

go build ./...
go vet ./...
go test ./...                     # no test files exist yet
go test -run TestName ./internal/service   # single test
```

### Migrations

`./scripts/create-migration.sh "create something"` copies `templates/migration-template` to `migrations/<unix_ts>_<slug>.go` and stamps the timestamp as the gormigrate ID. Each migration file registers itself into `migrations.AllMigrations` via `init()` — there is no central list to edit, so a new file is picked up just by existing in the package. Edit the generated `Migrate`/`Rollback` bodies to reference the real model (see [migrations/1784647193_create_url.go](migrations/1784647193_create_url.go)).

## Architecture

Request flow: `cmd/api/main.go` → `internal/router` → `internal/handler` → `internal/service` → GORM.

- **Config as singletons.** `config.Db()` and `config.Rdb()` are `sync.Once` globals. There is no dependency injection — services call `config.Db()` directly on each call. Don't thread a `*gorm.DB` through function signatures; follow the existing pattern.
- **Handlers are package-level funcs**, not methods on a struct. They decode+validate with the generic `validator.Validate[T](w, r)`, which writes the 400 response itself and returns `ok=false` — return immediately when it does. Responses go through `pkg/response` (`JSON` / `Error` / `ValidationError`); error messages are in Portuguese.
- **Auth.** `middleware.JwtMiddleware` parses the `Authorization: Bearer` token and puts the `sub` claim (a string user ID) into the request context. Handlers read it with `middleware.GetUserID(ctx)`, which parses it to `uint`. Token generation lives in the same file as the middleware.
- **Routes** ([internal/router/router.go](internal/router/router.go)): `/api/auth` and `/api/user` are public; everything else under `/api` is inside a JWT-protected `chi.Group`. The public redirect endpoint `GET /u/{path}` sits outside `/api` entirely.
- **Short codes** are 12 random alphanumeric chars from `math/rand` (`service.generateRandomString`), with no collision retry. Destinations get `https://` prefixed by `normalizeDestination` if no scheme is present.
- `internal/repository/` exists but is empty — data access currently lives in the service layer.

## Notes

- Redis is connected at startup but not yet used for caching lookups.
- `service.ListUrl` returns all URLs, not just the authenticated user's.
- `middleware.secretKey` is read from `JWT_SECRET` at package-init time, i.e. before `godotenv.Load()` runs in `main` — it only works when the var is set in the real environment.
- CI ([.github/workflows/build.yaml](.github/workflows/build.yaml)) builds and pushes to ECR via AWS OIDC role chaining; the test step is commented out. `terraform/` and `deployments/` are empty placeholders.
- A `postgres` MCP server is configured in `.mcp.json` pointing at the local compose database — useful for inspecting schema and query plans.
