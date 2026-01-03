# Valkey Inspector

A web UI for checking what keys are available in a Valkey instance or cluster.

Built by Human and AI (Gemini Code Assist) using the following stack:

- Gemini Code Assist (Free tier)
- VSCode
- Ubuntu
- Golang 1.25
- Valkey
- Docker and Docker compose
- Frontend
    - pnpm
    - Vue and Nuxt
    - Tailwindcss

!Valkey Inspector UI

## Features

*   **Cluster & Single Node Support**: Connect to any Valkey instance.
*   **Real-time Key Browsing**: Efficient `SCAN`-based pagination for browsing millions of keys without blocking the server.
*   **Visual Data Inspector**:
    *   **Strings**: Auto-formatting for JSON values.
    *   **Hashes**: Searchable table view.
    *   **Lists/Sets/ZSets**: Virtualized lists for large collections.
*   **Dark Mode**: Fully supported out of the box.
*   **Performance**: Built with Go for a high-performance backend and Nuxt 3 for a snappy frontend.

## Architecture

*   **Frontend**: Nuxt 3 (Vue 3 + TypeScript), Tailwind CSS.
*   **Backend**: Go (Golang) using `valkey-go` driver.
*   **Infrastructure**: Docker Compose for orchestration.

## Getting Started

### Prerequisites

*   Docker & Docker Compose

### Running the Application

The entire stack (Frontend, Backend, and a sample Valkey instance) can be launched with a single command:

```bash
docker-compose up --build
```

Once running, access the application at:

*   **Web UI**: http://localhost:3000
*   **Backend API**: http://localhost:8080

### Development

To run the frontend in development mode:

```bash
cd frontend
pnpm install
pnpm dev
```
