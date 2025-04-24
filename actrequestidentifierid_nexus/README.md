# ActRequestNexusService
## Technical Documentation


> *A Temporal Nexus service for processing activation requests with YANG data. This service provides reliable, standardized communication between services with built-in idempotency handling.*

## Table of Contents
- [Overview](#overview)
- [Version Information](#version-information)
- [Prerequisites](#prerequisites)
- [Environment Setup](#environment-setup)
- [Setup](#setup)
- [Build and Run](#build-and-run)
- [Configuration](#configuration)
- [Making Requests](#making-requests)
- [Workflow Implementation](#workflow-implementation)
- [Components](#components)
- [Monitoring](#monitoring)
- [Security](#security)
- [Development](#development)
- [Testing](#testing)
- [Troubleshooting](#troubleshooting)

## Overview

The ActRequestNexusService exposes a Nexus endpoint that allows processing of activation requests with the following features:
- Idempotent request handling through unique request IDs
- YANG data processing
- Standardized error responses
- TLS/SSL security
- Authentication via API Key or Client Certificate
- Automatic retries with exponential backoff
- Activity timeouts and error handling

## Version Information

| Component | Version |
|-----------|---------|
| Project Version | 1.0.0 |
| Spring Boot | 3.2.x |
| Temporal SDK | 1.27.1 |
| Spring Boot Starter | 3.2.x |
| SLF4J | 2.0.x |
| Logback | 1.4.x |

## Prerequisites

### 1. Java 17 or higher
```bash
# Check Java version
java -version

# If Java 17 is not installed, install it using:
# For macOS (using Homebrew):
brew install openjdk@17

# For Ubuntu/Debian:
sudo apt-get update
sudo apt-get install openjdk-17-jdk
```

### 2. Gradle 8.13 or higher
```bash
# Check Gradle version
gradle --version

# Install Gradle using SDK Man:
curl -s "https://get.sdkman.io" | bash
source "$HOME/.sdkman/bin/sdkman-init.sh"
sdk install gradle 8.13
```

### 3. Temporal CLI v1.3.0 or higher

Follow the instructions on the [Temporal docs](https://docs.temporal.io/cli) to install Temporal CLI.

### 4. A running Temporal server (local or cloud)

## Environment Setup

### 1. Set required environment variables:
```bash
# Add to your ~/.bashrc or ~/.zshrc
export JAVAHOME=/path/to/java17
export GRADLEHOME=/path/to/gradle
export PATH=$JAVAHOME/bin:$GRADLEHOME/bin:$PATH

# For development, you might want to set:
export SPRINGPROFILESACTIVE=dev
```

### 2. Verify your environment:
```bash
java -version  # Should show Java 17
gradle -version  # Should show Gradle 8.13 or higher
temporal -v  # Should show Temporal CLI version
```

## Setup

### 1. Install Temporal CLI
Follow the instructions on the [Temporal docs](https://docs.temporal.io/cli) to install Temporal CLI.

### 2. Start Temporal Server with Nexus Feature
```bash
temporal server start-dev \
  --ip "0.0.0.0" \
  --http-port 7243 \
  --namespace "network-usecases" \
  --namespace "lumen-usecases" \
  --namespace "nexus-lumen-act-ns" \
  --namespace "act-requestid-caller-ns" \
  --namespace "aiagent-usecases" \
  --dynamic-config-value system.enableNexus=true
```

### 3. Create Required Namespaces
```bash
# Create service namespace
temporal operator namespace create --namespace nexus-lumen-act-ns

# Create caller namespace (if you're also deploying the caller)
temporal operator namespace create --namespace act-requestid-caller-ns
```

### 4. Create Nexus Service and Nexus Endpoint

```bash
temporal operator nexus endpoint create \
  --name act-requestidentifier-service \
  --target-namespace nexus-lumen-act-ns \
  --target-task-queue act-task-queue \
  --description-file ./service_actrequestnexus.md
```

Other Temporal applications can then connect to this service using:

```
Service endpoint: act-requestidentifier-service.nexus-lumen-act-ns.svc.cluster.local
Task queue: nexus-task-queue
Namespace: nexus-lumen-act-ns
```


## Build and Run

The project uses a Makefile for common operations. Here are the available commands:

```bash
# Show all available commands
make help

# Build the project
make build

# Clean all build artifacts
make clean

# Run Worker on port 8080
make run-worker

# Run Service in a separate terminal on port 8081
make run-workflow

# Optional: Run in a separate terminal
make test
```

## Configuration

Create a `config.properties` file in `src/main/resources/` with your service configuration:

### Development Configuration Example
```properties
# Server connection (for local development)
target-host=localhost:7233
namespace=nexus-lumen-act-ns
server.port=8080
api-key=dev-api-key
insecure-skip-verify=true
```

### Production Configuration Example
```properties
# Server connection (for production)
target-host=temporal.production:7233
namespace=nexus-lumen-act-ns
server.port=443
client-cert=/etc/certs/client.crt
client-key=/etc/certs/client.key
server-root-ca-cert=/etc/certs/ca.crt
server-name=temporal.production
insecure-skip-verify=false
```

## Making Requests

The service exposes the following endpoint:
- `POST http://localhost:8080/api/v1/act/process`

### Example Request
```json
{
  "meta": {
    "requestId": "unique-request-id",
    "timestamp": "2024-04-21T18:00:00Z"
  },
  "feedback": "Processing activation request"
}
```

### Example Successful Response
```json
{
  "status": "SUCCESS",
  "requestId": "unique-request-id",
  "timestamp": "2024-04-21T18:00:00Z",
  "data": {
    "activationTraceId": "act-123456"
  }
}
```

### Example Error Response
```json
{
  "status": "ERROR",
  "requestId": "unique-request-id",
  "timestamp": "2024-04-21T18:00:00Z",
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid YANG data format"
  }
}
```

### Response Status Codes

| Status Code | Description |
|-------------|-------------|
| 200 | Success - Request processed successfully |
| 400 | Bad Request - Invalid request format or YANG data |
| 401 | Unauthorized - Invalid or missing authentication |
| 409 | Conflict - Duplicate request ID |
| 422 | Unprocessable Entity - Valid request but processing failed |
| 500 | Internal Server Error - Unexpected server error |

## Workflow Implementation

The service uses Temporal workflows for reliable execution:

### Activity Configuration
- Start to close timeout: 30 seconds
- Retry configuration:
  - Initial interval: 1 second
  - Maximum interval: 10 seconds
  - Backoff coefficient: 2.0
  - Maximum attempts: 3

## Components

- `ActWorkflow` : Interface defining the workflow contract
- `ActWorkflowImpl` : Implementation handling the request processing
- `ActService` : Activity interface for actual request processing
- `ActServiceImpl` : Activity implementation with business logic

## Monitoring

The service logs activity to stdout/stderr using SLF4J with Logback. Monitor the logs for:
- Workflow execution status
- Activity processing
- Request IDs
- Error messages and stack traces

## Security

- TLS/SSL is required for production deployments
- Authentication is required via either:
  - API Key
  - Client Certificate
- Configure security settings in the properties file

## Development

### Project Structure
- `service/` - Service definition and interface
- `src/main/java/com/lumen/workflow/`
  - `client/` - Client implementation
  - `model/` - Request/Response models
  - `service/` - Service implementation
  - `worker/` - Temporal worker implementation
  - `workflow/` - Workflow definitions and implementation

### Build with Tests
```bash
./gradlew clean test build
```

## Testing

### Using the Test Script
A test script is provided in the `scripts` directory to test the Nexus service endpoint. The script supports both API key and certificate-based authentication.

```bash
# Basic usage with default settings
./scripts/test_service.sh

# Using API key authentication
./scripts/test_service.sh --api-key "your-api-key"

# Using certificate-based authentication
./scripts/test_service.sh \
  --cert /path/to/client.crt \
  --key /path/to/client.key \
  --ca-cert /path/to/ca.crt

# Custom host and namespace
./scripts/test_service.sh \
  --host custom.host:7233 \
  --namespace custom-namespace \
  --api-key "your-api-key"

# Show help
./scripts/test_service.sh --help
```

The script will:
1. Generate a unique request ID
2. Create a timestamp in ISO format
3. Construct a test request with sample YANG data
4. Send the request to the service endpoint
5. Display the response

Example successful response:
```json
{ "status": "SUCCESS" }
```

Example error response:
```json
{
  "status": "ERROR",
  "error": {
    "message": "Error description"
  }
}
```

## Troubleshooting

### 1. Connection Issues
- Verify Temporal server is running
- Check target-host configuration
- Validate SSL certificates if using TLS

### 2. Authentication Failures
- Verify API key or certificates are properly configured
- Check namespace permissions

### 3. Request Processing Errors
- Check request format matches ActRequest structure
- Verify unique requestId for each request
- Check logs for detailed error messages

### 4. Workflow Execution Issues
- Check Temporal UI for workflow execution status
- Verify activity timeouts are appropriate
- Review retry configurations if activities are failing
