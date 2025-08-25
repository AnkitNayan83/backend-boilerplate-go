# Backend Boilerplate with Go

This repository serves as a boilerplate for Go backend applications with database integration, migrations, and SQL query generation.

## Features

- Docker-based PostgreSQL database setup
- Database migration support (up and down)
- SQL query generation with sqlc

## Getting Started

### Prerequisites

- Go (latest version recommended)
- Docker and Docker Compose
- [sqlc](https://github.com/sqlc-dev/sqlc) installed

### Database Setup

The project uses Docker to provide a consistent PostgreSQL database environment. To start the database:

```bash
docker-compose up -d
```

### Database Migrations

The project uses migrations to manage database schema changes:

```bash
# Run migrations up
go run cmd/migrate/main.go up

# Rollback migrations
go run cmd/migrate/main.go down
```

### SQL Query Generation

SQL queries are managed using sqlc which generates type-safe Go code from SQL:

```bash
# Generate SQL queries
sqlc generate
```

## Project Structure

```
.
├── cmd/                # Application entry points
│   └── migrate/        # Migration command tool
├── db/                 # Database related code
│   ├── migrations/     # SQL migration files
│   └── queries/        # SQLC query definitions
├── internal/           # Private application code
│   ├── database/       # Database connection code
│   └── models/         # Generated SQL models
└── docker-compose.yml  # Docker configuration
```

## Development

To set up the development environment:

1. Clone this repository
2. Start the database with `docker-compose up -d`
3. Run migrations with `go run cmd/migrate/main.go up`
4. Generate SQL code with `sqlc generate`
5. Start developing your application
