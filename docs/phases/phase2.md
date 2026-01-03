# Phase 2: Frontend Scaffolding & Core Functionality

## Goal: Build the Nuxt frontend to connect to the backend and display keys and their values.

Task 2.1: Create a ./frontend directory and initialize a new Nuxt.js project.
Task 2.2: Install bootstrap.
Task 2.3: Create a frontend/Dockerfile.dev for the development environment.
Task 2.4: Define the frontend service in docker-compose.yml to use this Dockerfile and map the source code for hot-reloading.
Task 2.5: Structure the main UI using Nuxt layouts (layouts/default.vue) and create core Vue components (components/TheHeader.vue, components/KeyList.vue, components/KeyViewer.vue).
Task 2.6: Implement the connection form and state management.
Task 2.7: Use Nuxt's useFetch composable to call the backend /api/keys endpoint and display the results in the KeyList component.
Task 2.8: When a key is clicked, use useFetch to get its details from /api/keys/{key_name} and pass the data to the KeyViewer component.
