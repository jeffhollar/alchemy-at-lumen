# ActRequest Identifier Workflow

This project implements a Temporal workflow for processing ActRequest objects. It provides a robust and scalable solution for handling request processing with proper error handling and logging.

## Project Structure

The project is organized into the following packages:

- `options` - Client configuration
- `service` - Activity service implementation
- `workflow` - Workflow implementation
- `worker` - Worker implementation
- `model` - Data models

## Prerequisites

- Java 17 or higher
- Gradle 7.0 or higher
- Access to a Temporal server
- Properly configured Temporal namespace

## Dependencies

The project uses the following main dependencies:
- Temporal SDK (1.22.0)
- Temporal Service Client (1.22.0)
- Jackson Databind (2.15.2)
- SLF4J API (2.0.7)
- Logback Classic (1.4.11)

## Building and Running the Project

### Using Gradle

The project can be built using Gradle with the following commands:

```bash
# Build the project
./gradlew build

# Create a fat JAR (includes all dependencies)
./gradlew fatJar
```

The build artifacts will be created in the `build/libs` directory.

### Using Makefile

The project includes a Makefile that provides convenient commands for common operations:

```bash
# Build the worker jar
make build

# Run the worker
make run-worker

# Run the workflow client
make run-workflow

# Clean build artifacts
make clean

# Show available commands
make help
```

The Makefile commands will automatically handle the proper JAR paths and configuration file locations.

## Running the Project

### Starting the Worker

To start the worker process:

```bash
# Using Java directly
java -jar build/libs/actrequest.jar config/temporal.properties

# Using Makefile
make run-worker
```

### Running the Client

To start a workflow using the client:

```bash
# Using Java directly
java -cp build/libs/actrequest.jar com.lumen.workflow.client.ActClient config/temporal.properties

# Using Makefile
make run-workflow
```

## Configuration

The project requires a `temporal.properties` file in the `config` directory. This file should contain the necessary configuration for connecting to your Temporal server, including:

- Server address
- Namespace
- SSL/TLS configuration (if required)
- API key (if required)

## Features

- Connects to Temporal server with optional SSL/TLS or API key authentication
- Processes ActRequest objects through a workflow
- Returns ActResponse objects
- Implements proper error handling and logging
- Supports scalable worker deployment

## Error Handling

The project includes comprehensive error handling throughout the codebase, ensuring graceful failure handling and proper logging of issues.

## Logging

Logging is implemented using SLF4J with Logback, providing detailed logging of workflow execution, activities, and any errors that occur during processing. 