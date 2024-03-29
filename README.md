
# Go Application with Gin, GORM, and PostgreSQL

This application is a simple REST API built with Go, using Gin as the web framework, GORM as the ORM, and PostgreSQL as the database. It demonstrates basic CRUD (Create, Read, Update, Delete) operations for a `Resource` entity.

## Features

- REST API endpoints for managing `Resource` entities.
- Use of GORM for database operations.
- PostgreSQL as the database backend.
- Gin for routing and handling HTTP requests.

## Prerequisites

- Go (version 1.13 or higher)
- PostgreSQL database server running locally on port 5432

## Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/nik-nenkov/skeleton-go-gin-viper-gorm
   cd skeleton-go-gin-viper-gorm
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

3. **Database Configuration**

   Ensure PostgreSQL is running on `localhost:5432`. Use the following credentials:
   
   - Username: `postgres`
   - Password: `postgres`
   - Database Name: `mydatabase`

   These can be modified in the `main.go` file as needed.

4. **Run the Application**

   ```bash
   go run main.go
   ```

   The application will start running on `http://localhost:8080`.

## API Endpoints

- `GET /resources/:name`
  - Retrieves a resource by its name.
- `POST /resources`
  - Creates a new resource. Requires a JSON body with `name`, `description`, and `value`.
- `PUT /resources/:name`
  - Updates an existing resource by its name. Requires a JSON body with `name`, `description`, and `value`.
- `DELETE /resources/:name`
  - Deletes a resource by its name.
