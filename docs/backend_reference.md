# Backend Source Code Documentation

This document provides a reference for the backend codebase of the Valkey Inspector application.

## Project Structure

The backend is organized following standard Go project layout conventions:

- **`cmd/valkey-inspector/`**: Contains the `main.go` entry point.
- **`pkg/`**: Contains the library code.
    - **`api/`**: HTTP handlers and router configuration.
    - **`models/`**: Data structures for API requests and responses.
    - **`server/`**: HTTP server setup and lifecycle management.
    - **`valkey/`**: Valkey connection management logic.

## Key Packages & Components

### 1. `pkg/valkey` (Connection Manager)
- **File**: `manager.go`
- **Purpose**: Manages dynamic connections to Valkey instances.
- **Key Structs**:
    - `Manager`: Holds a map of active clients protected by a mutex.
    - `ConnectionConfig`: Struct for connection details (Host, Port, Auth).
- **Key Methods**:
    - `GetClient(cfg ConnectionConfig)`: Returns an existing client or creates a new one based on the config.

### 2. `pkg/api` (HTTP Layer)
- **Files**: `handlers.go`, `router.go`
- **Purpose**: Defines the API endpoints and handles HTTP requests.
- **Key Structs**:
    - `Handler`: Holds a reference to `valkey.Manager`.
- **Endpoints**:
    - `POST /api/connect`: Validates connection details.
    - `GET /api/keys`: Lists keys using `SCAN` with cursor-based pagination. Pipelines `TYPE` and `TTL` commands for efficiency.
    - `GET /api/keys/{key}`: Retrieves details for a specific key (String, Hash, List, Set, ZSet).
- **Router**: Uses `chi` router and `cors` middleware.

### 3. `pkg/models` (Data Models)
- **File**: `models.go`
- **Purpose**: Defines JSON structures for API responses.
- **Key Structs**:
    - `KeyItem`: Summary of a key (Name, Type, TTL).
    - `ScanResponse`: Response format for key listing (Cursor, Keys).
    - `KeyDetail`: Detailed view of a key (Value, Type, TTL).

### 4. `pkg/server` (Server)
- **File**: `server.go`
- **Purpose**: Wires up the dependencies and starts the HTTP server.
- **Responsibilities**:
    - Initializing the `valkey.Manager`.
    - Setting up the `api` router.
    - Starting the `net/http` server on port 8080.

## API Contract

The backend is stateless regarding connections. The frontend must send the following headers with every request to identify the target Valkey instance:
- `X-Valkey-Host`
- `X-Valkey-Port`
- `X-Valkey-User` (Optional)
- `X-Valkey-Pass` (Optional)

## Dependencies
- `github.com/go-chi/chi/v5`: HTTP Router.
- `github.com/go-chi/cors`: CORS middleware.
- `github.com/valkey-io/valkey-go`: Valkey client driver.