
# Go To-Do API

This is a simple To-Do list API built with Go. It features basic CRUD operations for managing tasks with PostgreSQL as the database backend.

## Features
- Create, read, update, and delete to-dos.
- Basic API documentation (Swagger to be added later).

## Prerequisites
- Go 1.23 or higher
- PostgreSQL

## Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/kostjaog/go-to-do
   cd go-to-do
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Setup Environment Variables
Before starting the application, you need to configure environment variables:

3.1. Copy the `.env.example` file to create your `dev.env` file:
   ```bash
   cp .env.example env/dev.env
   ```

3.2. Modify the values in `env/dev.env` according to your local configuration.

4. Start the application:
   ```bash
   go run cmd/main.go
   ```

5. The API will be available at `http://localhost:8080`.

### Database Setup via Docker
You can set up the PostgreSQL database using Docker by running:

```bash
docker-compose -f deploy/docker-compose.yaml up -d
```

## Environment Variables

| Variable      | Description                        | Example                |
|---------------|------------------------------------|------------------------|
| `DB_USER`     | PostgreSQL database username       | `postgres`             |
| `DB_PASSWORD` | PostgreSQL database password       | `secret`               |
| `DB_NAME`     | Name of the PostgreSQL database    | `todo_app`             |
| `DB_HOST`     | Host for the PostgreSQL database   | `localhost`            |
| `DB_PORT`     | Port for the PostgreSQL database   | `5432`                 |

## Database Structure

| Table    | Column       | Type    | Description                          |
|----------|--------------|---------|--------------------------------------|
| `todos`  | `id`         | `UUID`  | Unique identifier for each task     |
|          | `title`      | `string`| Title of the to-do task             |
|          | `completed`  | `bool`  | Whether the task is completed       |
|          | `created_at` | `timestamp` | Timestamp when the task was created |

## Endpoints

### `GET /todos`
- **Description**: Retrieve all to-dos.
- **Response**: A list of all to-dos in the system.

### `POST /todos`
- **Description**: Create a new to-do.
- **Request body**: A JSON object containing the title and completed status.
  - Example:
    ```json
    {
      "title": "New Task",
      "completed": false
    }
    ```
- **Response**: The created to-do item with all its details.
