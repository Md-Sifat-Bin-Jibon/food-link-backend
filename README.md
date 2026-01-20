# Foodlink Backend

A Go backend API server for Foodlink application.

## Prerequisites

- Go 1.21 or higher
- Git

## Getting Started

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd Foodlink_backend
```

2. Install dependencies:
```bash
go mod download
```

### Running the Server

1. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080` by default.

### Environment Variables

You can configure the server using environment variables:

- `PORT` - Server port (default: 8080)
- `ENVIRONMENT` - Environment mode (default: development)
- `DATABASE_URL` - Database connection string (optional)

Example:
```bash
PORT=3000 go run main.go
```

## API Documentation

### Swagger UI

Interactive API documentation is available via Swagger UI:

- **Swagger UI**: `http://localhost:8080/swagger/index.html`

The Swagger documentation is automatically generated from code annotations. To regenerate the docs after making changes:

```bash
swag init
```

## API Endpoints

### Health Check
- `GET /health` - Check server health status

### API v1
- `GET /api/v1/` - API v1 welcome message

## Project Structure

```
.
├── main.go                      # Application entry point
├── config/                      # Configuration management
│   └── config.go
├── database/                    # Database layer (to be created)
│   ├── connection.go
│   ├── migrations/
│   └── schema.go
├── middleware/                  # HTTP middleware (to be created)
├── utils/                       # Utility functions (to be created)
├── errors/                      # Custom error types (to be created)
├── features/                    # Feature modules (to be created)
│   ├── auth/
│   ├── inventory/
│   ├── consumption/
│   └── ... (see IMPLEMENTATION_PLAN.md)
├── handlers/                   # HTTP request handlers (legacy)
│   └── handlers.go
├── routes/                     # Route definitions
│   └── routes.go
├── docs/                       # Swagger documentation (auto-generated)
├── schema.sql                  # Database schema
├── IMPLEMENTATION_PLAN.md      # Detailed implementation plan
├── FEATURES.md                 # Features overview
├── go.mod
└── README.md
```

## Implementation Plan

This project follows a feature-based architecture. See the detailed implementation plan:

- **[IMPLEMENTATION_PLAN.md](./IMPLEMENTATION_PLAN.md)** - Complete implementation plan with phases, features, and guidelines
- **[FEATURES.md](./FEATURES.md)** - Quick reference guide for all features

### Implementation Phases

1. **Phase 0**: Foundation & Infrastructure (Database, Middleware)
2. **Phase 1**: Authentication & User Management
3. **Phase 2**: Core Family Features (Inventory, Consumption, Shopping Lists, Meal Plans)
4. **Phase 3**: User Preferences & Nutrition
5. **Phase 4**: Gamification (Badges, XP)
6. **Phase 5**: Community Features
7. **Phase 6**: Restaurant Module
8. **Phase 7**: NGO Module
9. **Phase 8**: Shop Module
10. **Phase 9**: Supporting Features (Uploads, Resources, Notifications)

## Development

### Building

Build the application:
```bash
go build -o bin/server main.go
```

### Running Tests

Run tests:
```bash
go test ./...
```

### Generating Swagger Documentation

After adding or modifying API endpoints with Swagger annotations, regenerate the documentation:

```bash
swag init
```

Make sure `swag` is installed:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## License

MIT
