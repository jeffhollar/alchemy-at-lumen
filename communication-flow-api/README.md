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

```bash
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

## Java Integration

This project can be extended to include Java-based Temporal.io workflows while maintaining the existing Go-based implementation. Here's the recommended approach:

### Architecture Overview

The integration involves running a separate Java service alongside the existing Go service, both connecting to the same Temporal.io server:

```
communication-flow-api/
├── go/                    # Existing Go service
├── java/                  # New Java service
│   ├── src/
│   │   ├── main/
│   │   │   ├── java/     # Java workflow implementations
│   │   │   └── resources/
│   │   └── test/
│   ├── pom.xml           # Maven configuration
│   └── Dockerfile
├── docker-compose.yml    # Orchestrate both services
└── README.md
```

### Implementation Guidelines

1. **Service Separation**
   - Maintain the existing Go service unchanged
   - Create a new Java service for Java-based workflows
   - Use different task queues for Go and Java workflows
   - Implement proper service discovery between services

2. **Communication Methods**
   - REST API endpoints
   - gRPC (recommended for better performance)
   - Direct Temporal workflow orchestration

3. **Best Practices**
   - Use consistent naming conventions across services
   - Share common workflow definitions (protobufs)
   - Implement proper error handling and logging
   - Use environment variables for configuration
   - Implement health checks for both services

4. **Deployment Strategy**
   - Deploy services independently
   - Use Kubernetes for orchestration
   - Implement proper monitoring and logging
   - Scale services independently as needed

### Example Java Service Setup

1. Create a new Maven project with Temporal Java SDK dependencies
2. Implement Java workflows in a separate package
3. Use a different task queue (e.g., `java-communication-task-queue`)
4. Add the Java service to Docker Compose configuration

### Integration Points

- Go service can start Java workflows using Temporal client
- Java workflows can call Go activities and vice versa
- Use shared data structures (protobufs) for communication
- Implement proper error handling and retry mechanisms

### Development Setup for Java Integration

1. Install Java Development Kit (JDK) 11 or later
2. Install Maven for dependency management
3. Set up the Java service project structure
4. Configure the Temporal.io connection
5. Implement sample Java workflows
6. Test the integration with the Go service

### Configuration for Java Service

The Java service can be configured using environment variables:

- `TEMPORAL_HOST`: Temporal server host and port (default: localhost:7233)
- `TEMPORAL_NAMESPACE`: Temporal namespace (default: lumen-usecases)
- `JAVA_TASK_QUEUE`: Java task queue name (default: java-communication-task-queue)
- `JAVA_SERVICE_PORT`: Java service port (default: 8502)

## Nexus Integration

Temporal.io's Nexus feature provides a powerful way to integrate Java workflows with the existing Go service. Nexus allows for seamless cross-language workflow execution and communication through a unified API.

### Nexus Overview

Nexus provides:
- Unified API for cross-language workflow execution
- Type-safe communication between services
- Simplified service discovery and routing
- Built-in support for Java and Go workflows
- Automatic serialization/deserialization of workflow data

### Implementation with Nexus

1. **Service Architecture**
```
communication-flow-api/
├── go/                    # Existing Go service
├── java/                  # Java service with Nexus
│   ├── src/
│   │   ├── main/
│   │   │   ├── java/
│   │   │   │   ├── nexus/  # Nexus interface definitions
│   │   │   │   └── workflows/
│   │   │   └── resources/
│   │   └── test/
│   ├── pom.xml
│   └── Dockerfile
├── proto/                 # Shared protocol definitions
└── docker-compose.yml
```

2. **Nexus Interface Definition**
```java
@NexusInterface
public interface CommunicationWorkflow {
    @NexusMethod
    void processCommunication(CommunicationRequest request);
}
```

3. **Workflow Implementation**
```java
@WorkflowInterface
public class JavaCommunicationWorkflowImpl implements CommunicationWorkflow {
    @Override
    public void processCommunication(CommunicationRequest request) {
        // Workflow implementation
    }
}
```

### Benefits of Using Nexus

1. **Simplified Integration**
   - No need for custom serialization/deserialization
   - Automatic type conversion between languages
   - Unified error handling across services

2. **Type Safety**
   - Compile-time type checking
   - Automatic validation of workflow parameters
   - Consistent data structures across services

3. **Performance**
   - Optimized communication between services
   - Reduced overhead in cross-language calls
   - Efficient serialization of workflow data

### Configuration for Nexus

Add the following to your Java service configuration:

```properties
# Nexus Configuration
temporal.nexus.enabled=true
temporal.nexus.port=7234
temporal.nexus.namespace=lumen-usecases
```

### Example Go Service Integration

```go
// Go service using Nexus client
client, err := nexus.NewClient(nexus.ClientOptions{
    Service: "java-communication-service",
    Address: "localhost:7234",
})

// Start Java workflow through Nexus
workflow := client.NewWorkflowStub(CommunicationWorkflow{})
err = workflow.ProcessCommunication(ctx, request)
```

### Best Practices for Nexus Integration

1. **Interface Design**
   - Define clear, versioned interfaces
   - Use consistent naming conventions
   - Document interface contracts

2. **Error Handling**
   - Implement proper error propagation
   - Use Nexus-specific error types
   - Handle cross-language exceptions

3. **Monitoring**
   - Track Nexus-specific metrics
   - Monitor cross-service communication
   - Implement proper logging

4. **Security**
   - Configure proper authentication
   - Use TLS for Nexus communication
   - Implement proper access controls

### Migration Strategy

1. **Phase 1: Setup**
   - Add Nexus dependencies to Java service
   - Define initial interfaces
   - Configure Nexus server

2. **Phase 2: Implementation**
   - Implement Java workflows
   - Update Go service to use Nexus client
   - Test cross-language communication

3. **Phase 3: Production**
   - Gradual rollout of Nexus-enabled workflows
   - Monitor performance and stability
   - Optimize based on usage patterns

## Deploying Java Workflows as Nexus Endpoints

### Nexus Endpoint Deployment Overview

Nexus endpoints allow Java workflows to be exposed as services that can be called from other languages (like Go) through a unified API. This section details how to deploy and manage Java workflows as Nexus endpoints.

### 1. Nexus Endpoint Configuration

```java
// Example Nexus endpoint configuration
@NexusEndpoint(
    name = "communication-workflow",
    version = "1.0",
    description = "Handles communication workflow processing"
)
public class CommunicationWorkflowEndpoint {
    
    @NexusMethod(
        name = "processCommunication",
        description = "Processes a communication request"
    )
    public void processCommunication(CommunicationRequest request) {
        // Workflow implementation
    }
}
```

### 2. Deployment Steps

1. **Build Configuration**
```xml
<!-- pom.xml Nexus dependencies -->
<dependencies>
    <dependency>
        <groupId>io.temporal</groupId>
        <artifactId>temporal-nexus</artifactId>
        <version>1.0.0</version>
    </dependency>
    <dependency>
        <groupId>io.temporal</groupId>
        <artifactId>temporal-nexus-server</artifactId>
        <version>1.0.0</version>
    </dependency>
</dependencies>
```

2. **Nexus Server Configuration**
```yaml
# application.yml
temporal:
  nexus:
    server:
      port: 7234
      namespace: lumen-usecases
      endpoints:
        - name: communication-workflow
          version: 1.0
          package: com.lumen.workflows
```

3. **Docker Configuration**
```dockerfile
# Dockerfile for Nexus endpoint
FROM openjdk:11-jdk

# Copy application
COPY target/communication-workflow.jar /app/
COPY src/main/resources/application.yml /app/config/

# Expose Nexus port
EXPOSE 7234

# Run application
CMD ["java", "-jar", "/app/communication-workflow.jar"]
```

### 3. Kubernetes Deployment

```yaml
# nexus-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: communication-workflow-nexus
spec:
  replicas: 3
  selector:
    matchLabels:
      app: communication-workflow-nexus
  template:
    metadata:
      labels:
        app: communication-workflow-nexus
    spec:
      containers:
      - name: communication-workflow-nexus
        image: communication-workflow-nexus:1.0.0
        ports:
        - containerPort: 7234
        env:
        - name: TEMPORAL_NEXUS_SERVER_PORT
          value: "7234"
        - name: TEMPORAL_NAMESPACE
          value: "lumen-usecases"
```

### 4. Service Discovery

```yaml
# nexus-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: communication-workflow-nexus
spec:
  selector:
    app: communication-workflow-nexus
  ports:
  - port: 7234
    targetPort: 7234
  type: ClusterIP
```

### 5. Health Checks

```java
@NexusEndpoint
public class CommunicationWorkflowEndpoint {
    
    @NexusMethod
    public HealthCheckResponse healthCheck() {
        return HealthCheckResponse.builder()
            .status(HealthStatus.UP)
            .build();
    }
}
```

### 6. Monitoring and Metrics

```java
@NexusEndpoint
public class CommunicationWorkflowEndpoint {
    
    @NexusMethod
    @Timed(value = "process_communication_duration", 
           description = "Time taken to process communication")
    public void processCommunication(CommunicationRequest request) {
        // Workflow implementation
    }
}
```

### 7. Security Configuration

```yaml
# security-config.yml
temporal:
  nexus:
    security:
      enabled: true
      tls:
        enabled: true
        keyStore: /path/to/keystore.jks
        keyStorePassword: ${KEYSTORE_PASSWORD}
      authentication:
        type: JWT
        issuer: lumen-auth
```

### 8. Deployment Checklist

1. **Pre-deployment**
   - [ ] Verify Nexus endpoint annotations
   - [ ] Test local deployment
   - [ ] Validate security configuration
   - [ ] Check health endpoints

2. **Deployment**
   - [ ] Build and package application
   - [ ] Deploy to Kubernetes cluster
   - [ ] Verify service discovery
   - [ ] Test endpoint connectivity

3. **Post-deployment**
   - [ ] Monitor metrics and logs
   - [ ] Verify health checks
   - [ ] Test failover scenarios
   - [ ] Document endpoint details

### 9. Troubleshooting

Common issues and solutions:

1. **Connection Issues**
   - Verify Nexus server port accessibility
   - Check network policies
   - Validate TLS configuration

2. **Performance Problems**
   - Monitor endpoint metrics
   - Check resource utilization
   - Review workflow execution times

3. **Security Issues**
   - Verify authentication tokens
   - Check TLS certificates
   - Review access logs

### 10. Best Practices

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
