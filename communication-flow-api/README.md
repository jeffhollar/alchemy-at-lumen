# Act Core Communication Flow

A Go-based adapter service that provides an interface between network services and Temporal.io workflows. This project serves as a significant adapter that bridges the domain's workflow requirements with the underlying execution environment, enabling seamless communication between core business logic and external systems.

## Project Overview

The Communication Flow API is a robust service that leverages Temporal.io's workflow orchestration capabilities to provide:
- Task queuing and execution
- State management
- Fault tolerance
- Reliable and scalable operations
- Modular and adaptable system architecture

The service consists of two main components:
1. **API Service**: A RESTful API that handles incoming requests and manages communication with external systems
2. **Worker Service**: Handles workflow execution and task processing using Temporal.io

## Features

- RESTful API with Swagger documentation
- Secure HTTPS communication with TLS
- Integration with Temporal.io for workflow orchestration
- Health and service information endpoints
- Support for Automated Communication Tasks (ACT)
- Processing status tracking and monitoring

## Prerequisites

- Go 1.23 or later
- Docker (for containerized deployment)
- OpenSSL (for generating SSL certificates)
- Temporal.io server (for workflow execution)

## Development Setup

### Generate SSL Certificates

For development, you'll need to create self-signed certificates:

```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt
```

### Local Development

1. Clone the repository:

```bash
git clone git@github.com:jeffhollar/alchemy-at-lumen.git
cd communication-flow-api
```

2. Build and run the application:

```bash
make clean
make build
make run
```

For debugging:

```bash
make debug
```

3. Access the Swagger UI:

```
https://localhost:8501/swaggerui/
```

### Docker Deployment

1. Build the Docker image:

```bash
make clean
make build
make package
```

2. Run the container:

```bash
docker run -d -p 8501:8501 --name communication-service comapi-1.0.0
```

## Configuration

The service can be configured using environment variables:

- `TEMPORAL_HOST`: Temporal server host and port (default: localhost:7233)
- `TEMPORAL_NAMESPACE`: Temporal namespace (default: lumen-usecases)
- `TEMPORAL_TASK_QUEUE`: Task queue name (default: act-communication-task-queue)
- `SSL_CERT_FILE`: Path to SSL certificate file (default: key.pem)
- `SSL_KEY_FILE`: Path to SSL private key file (default: server.key)
- `SERVER_PORT`: Server port (default: :8501)

## API Documentation

The API documentation is available through the Swagger UI interface at `https://localhost:8501/swaggerui/`. The API provides endpoints for:

- Health checks
- Service information
- ACT request submission and management
- Processing status tracking


## Best Practices

1. **Endpoint Design**
   - Use clear, descriptive names
   - Version all endpoints
   - Document all methods

2. **Deployment**
   - Use rolling updates
   - Implement proper health checks
   - Monitor resource usage

3. **Security**
   - Enable TLS by default
   - Use proper authentication
   - Implement rate limiting

4. **Monitoring**
   - Track endpoint metrics
   - Monitor error rates
   - Set up alerts

## License

Copyright 2025 Lumen Technologies, Inc. All rights reserved.
