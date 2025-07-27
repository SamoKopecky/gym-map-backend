# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

**Testing:**
- `make test` - Run all tests with database setup
- `go test ./...` - Run tests directly (requires running database)

**Development:**
- `make run` - Start development server with hot reload using Air
- `make up` - Start all services with Docker Compose
- `make up-d` - Start services in background
- `make down` - Stop and remove all containers

**Database:**
- `make pgcli` - Connect to database via pgcli
- `make db-refresh` - Restart only the database container
- `make reset-db` - Full database reset
- `make add-migration MSG="description"` - Create new migration file

**Build:**
- `go build -v ./...` - Build all packages
- `go mod tidy && go mod download` - Install/update dependencies

## Architecture Overview

**Backend Framework:** Go web API using Echo v4 framework with PostgreSQL database via Bun ORM.

**Authentication:** JWT-based authentication with Keycloak integration. Three user roles: admin, trainer, and regular users with different permission levels.

**Core Entities:**
- **Machines** - Gym equipment with positions on floor map
- **Exercises** - Workouts that can be performed on machines
- **Instructions** - User-created workout instructions with media attachments
- **Media** - File uploads (videos, images) linked to instructions
- **Users** - User management with role-based permissions

**API Structure:**
- REST endpoints organized by entity type
- Role-based middleware (admin, trainer, user permissions)
- JWT authentication middleware for protected routes
- Prometheus metrics integration

**Database:**
- PostgreSQL with Bun ORM
- Migration-based schema management in `/migrations`
- Database connection with debug logging support

**Key Directories:**
- `/api` - HTTP handlers organized by entity
- `/model` - Data models with Bun annotations
- `/crud` - Database operations layer
- `/service` - Business logic layer
- `/store` - Repository interfaces with mock implementations
- `/config` - Environment-based configuration
- `/migrations` - SQL migration files

**Testing:**
- Uses `testify` for assertions and mocks
- Mock implementations in `/store/mock`
- Test utilities in `/testutil`
- GitHub Actions CI/CD with PostgreSQL service

**File Handling:**
- Media file uploads stored in configurable directory
- Floor map SVG file management
- UUID-based file naming

**Configuration:**
- Environment variable based config with `APP_` prefix
- Default development values provided
- Keycloak integration for user management