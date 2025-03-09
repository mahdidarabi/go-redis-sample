# Go + Redis + gin Sample Project

A sample API demonstrating integration of Go, Gin, and Redis for user profile management.

## Installation and Make Commands

1. Clone the repository:

```bash
git clone https://github.com/your-repo/go-redis-sample.git
cd go-redis-sample
```

2. Install dependencies:

```bash
make get
```

3. Available make commands:

```bash
make run       # Run the application locally
make build     # Build the application
make start     # Build and start the application
make clean     # Clean build artifacts
make swag-init # Initialize/update Swagger documentation
```

## Sample Usage

The API provides CRUD operations for user profiles. You can use the `api.http` file for testing:

1. Start the application:

```bash
make start
```

2. Use the following sample requests (can be executed directly in VS Code with REST Client extension):

```http
### Create User Profile
POST http://localhost:3000/user-profile/12345
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "phone": "+1234567890",
  "address": "123 Main St",
  "city": "Anytown",
  "state": "CA",
  "zip": "90210"
}

### Get User Profile
GET http://localhost:3000/user-profile/12345

### Update User Profile
PUT http://localhost:3000/user-profile/12345
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "phone": "+1234567890",
  "address": "456 Elm St",
  "city": "Othertown",
  "state": "NY",
  "zip": "10001"
}

### Delete User Profile
DELETE http://localhost:3000/user-profile/12345
```

## Running with Docker Compose

1. Start the application and Redis:

```bash
docker-compose up
```

2. The application will be available at `http://localhost:3000`

3. Redis will be available at `localhost:6379`

4. To stop and remove containers:

```bash
docker-compose down
```

## Auto Build Pipeline

The project includes a GitHub Actions workflow for automated Docker image building and pushing:

1. On every push to the `main` branch, a new Docker image is built and pushed to GitHub Container Registry (GHCR)

2. When creating a Git tag in the format `v*.*.*` (e.g., `v1.2.3`):

   - A new Docker image is built
   - The image is tagged with the version number
   - The image is pushed to GHCR

3. The workflow uses the following naming convention for images:

   - `ghcr.io/<owner>/<repo>:<version>` for tagged releases
   - `ghcr.io/<owner>/<repo>:latest` for main branch builds

4. The pipeline includes:
   - Automatic dependency caching
   - Metadata extraction for proper tagging
