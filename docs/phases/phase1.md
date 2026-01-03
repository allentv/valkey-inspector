# Phase 1: Environment & Backend API

## Goal: Set up the complete Dockerized environment and build a functional Go backend API.

- Task 1.1: Create the root project directory and a docker-compose.yml file.
- Task 1.2: Define the valkey service in docker-compose.yml using a public Valkey image.
- Task 1.3: Create an init-valkey.sh script containing valkey-cli commands to SET, HSET, LPUSH sample data into the database.
- Task 1.4: Configure the valkey service to run this script on startup by mounting it as a volume.
- Task 1.5: Create a ./backend directory and initialize a Go module (go mod init).
- Task 1.6: Install the valkey-go dependency.
- Task 1.7: Implement the core net/http server and a Valkey connection service.
- Task 1.8: Implement the API endpoints:
    - POST /api/connect: Validate connection to Valkey.
    - GET /api/keys: List keys using SCAN.
    - GET /api/keys/{key_name}: Get type and value for a specific key.
- Task 1.9: Create a backend/Dockerfile to build a minimal, multi-stage Go container image.
- Task 1.10: Define the backend service in docker-compose.yml to use the new Dockerfile.
