# ACT Core API

The ACT Core API is a Go-based service that provides a communication layer for the ACT Core system. It integrates with Temporal.io for workflow orchestration and provides a secure HTTPS interface with Swagger UI documentation.

## Features

- RESTful API with Swagger UI documentation
- Integration with Temporal.io for workflow orchestration
- Secure HTTPS communication with TLS 1.2+
- Docker containerization support
- Graceful shutdown handling
- Environment-based configuration

## Prerequisites

- Go 1.23.6 or later
- Docker (for containerized deployment)
- Temporal.io server (for workflow orchestration)
- SSL certificates (key.pem and server.key)

## Dependencies

The project uses the following major dependencies:

- github.com/gorilla/mux v1.8.1 (HTTP router)
- go.temporal.io/sdk v1.34.0 (Temporal.io client)
- go.temporal.io/api v1.46.0 (Temporal.io API)

For a complete list of dependencies, see `go.mod` and `go.sum` files.

## Configuration

The service can be configured using environment variables:

- `TEMPORAL_HOST`: Temporal server host and port (default: "192.168.1.233:7233")
- `TEMPORAL_NAMESPACE`: Temporal namespace (default: "lumen-usecases")
- `TEMPORAL_TASK_QUEUE`: Temporal task queue name (default: "act-core-api-task-queue")
- `SSL_CERT_FILE`: Path to SSL certificate file (default: "key.pem")
- `SSL_KEY_FILE`: Path to SSL private key file (default: "server.key")
- `SERVER_PORT`: Server port (default: ":8501")

## Building the Project

### Local Build

1. Ensure Go is installed and properly configured
2. Clone the repository
3. Run the following commands:

```bash
export GO111MODULE=on
go mod download
go build -o act-core-service
```

### Docker Build

The project includes a Dockerfile for containerized deployment. To build the Docker image:

```bash
docker build -t apicore-1.0.0 .
```

## Running the Service

### Local Development

1. Ensure SSL certificates (key.pem and server.key) are present in the project directory
2. Run the compiled binary:

```bash
./communication-service
```

The service will start and be available at `https://localhost:8501`

### Docker Deployment

1. Build the Docker image (as shown above)
2. Run the container:

```bash
docker run -d -p 8501:8501 --name act-core-service comapi-1.0.0
```

For debugging purposes, you can run the container in interactive mode:

```bash
docker run -it -p 8501:8501 --name act-core-service apicore-1.0.0 /bin/bash
```

## API Documentation

The API documentation is available through Swagger UI at:
```
https://localhost:8501/swaggerui/
```

## Development

### Project Structure

- `api/`: API definitions and specifications
- `go/`: Generated Go code
- `internal/`: Internal packages including activities and workflows
- `swaggerui/`: Swagger UI static files
- `main.go`: Main application entry point

### Makefile Commands

The project includes a Makefile with the following commands:

- `make build`: Build the Docker image
- `make package`: Build and package the Docker image
- `make run`: Run the Docker container
- `make debug`: Run the Docker container in debug mode
- `make clean`: Remove Docker container and image

## Security

The service uses TLS 1.2+ with the following security features:
- Strong cipher suites
- Modern elliptic curves
- SSL certificate verification

Note: Certificate verification is currently disabled for development purposes (`InsecureSkipVerify: true`). This should be changed in production environments.

## License

Copyright 2025 Lunmen Technologies, Inc. All rights reserved. 