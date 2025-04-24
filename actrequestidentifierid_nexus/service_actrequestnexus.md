# ACT Request Identifier Nexus Service

## Service Overview
The ACT Request Identifier service processes ACT requests with YANG data validation and activation transaction handling. This service provides a Temporal workflow interface for managing ACT requests through the Nexus platform.

## Workflow Interface

### NexusEndpoint
```typescript
workflow NexusEndpoint {
    // Process a new Nexus request
    @workflowMethod
    processNexusRequest(request: ActRequest): ActResponse

    // Query the status of a request
    @queryMethod
    getRequestStatus(requestId: string): string

    // Signal to cancel an ongoing request
    @signalMethod
    cancelRequest(requestId: string): void
}
```

## Data Types

### ActRequest
```typescript
type ActRequest {
    meta: ActRequestMeta
    feedback?: string
    yang?: any
    activationTransactionId?: string
    yangError?: string
}

type ActRequestMeta {
    requestId: string
    timestamp: string
}
```

### ActResponse
```typescript
type ActResponse {
    status: string
    error?: ActResponseError
}

type ActResponseError {
    message: string
}
```

## Configuration

- **Target Namespace**: nexus-lumen-act-ns
- **Task Queue**: nexus-task-queue
- **Service Name**: act-requestidentifier-service

## Service Capabilities

- Processing of ACT requests with YANG data validation
- Request status tracking and querying
- Request cancellation support
- Error handling and reporting
- Activation transaction processing

## Dependencies

- Temporal.io SDK
- Spring Boot Framework
- Jackson for JSON processing

## Usage Examples

### Starting a New Request
```typescript
// Create a workflow client
const workflow = client.newWorkflowStub(NexusEndpoint.class, workflowOptions);

// Create and submit request
const request = {
    meta: {
        requestId: "request-123",
        timestamp: "2024-04-21T18:00:00Z"
    },
    yang: yangData,
    activationTransactionId: "act-123"
};

const response = await workflow.processNexusRequest(request);
```

### Checking Request Status
```typescript
const status = await workflow.getRequestStatus("request-123");
```

### Cancelling a Request
```typescript
await workflow.cancelRequest("request-123");
```

## Error Handling

The service provides detailed error responses through the ActResponse type, including:
- Validation errors
- Processing errors
- YANG data errors
- System errors

## Security

- Requires proper authentication through Temporal
- Supports SSL/TLS for secure communication
- API key authentication available 