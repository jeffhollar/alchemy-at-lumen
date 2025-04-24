package com.lumen.workflow.controller;

import org.springframework.web.bind.annotation.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.client.WorkflowStub;
import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActResponse;
import com.lumen.workflow.service.NexusEndpoint;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@RestController
@RequestMapping("/api/v1/nexus")
public class NexusController {
    private static final Logger logger = LoggerFactory.getLogger(NexusController.class);
    private static final String TASK_QUEUE = "nexus-task-queue";

    private final WorkflowClient workflowClient;

    @Autowired
    public NexusController(WorkflowClient workflowClient) {
        this.workflowClient = workflowClient;
    }

    @PostMapping("/process")
    public ResponseEntity<ActResponse> processNexusRequest(@RequestBody ActRequest request) {
        logger.info("Received Nexus request with ID: {}", request.getMeta().getRequestId());

        try {
            // Validate request
            if (request.getMeta() == null || request.getMeta().getRequestId() == null) {
                ActResponse errorResponse = new ActResponse();
                errorResponse.setStatus("ERROR");
                com.lumen.workflow.model.ActResponseError error = new com.lumen.workflow.model.ActResponseError();
                error.setMessage("Request ID is required");
                errorResponse.setError(error);
                return ResponseEntity.badRequest().body(errorResponse);
            }

            // Create workflow options
            WorkflowOptions options = WorkflowOptions.newBuilder()
                    .setTaskQueue(TASK_QUEUE)
                    .setWorkflowId("nexus-" + request.getMeta().getRequestId())
                    .build();

            // Start the workflow
            NexusEndpoint workflow = workflowClient.newWorkflowStub(NexusEndpoint.class, options);
            ActResponse response = workflow.processNexusRequest(request);
            
            logger.info("Request {} processed successfully", request.getMeta().getRequestId());
            return ResponseEntity.ok(response);
        } catch (Exception e) {
            logger.error("Error processing Nexus request {}", request.getMeta().getRequestId(), e);
            
            ActResponse errorResponse = new ActResponse();
            errorResponse.setStatus("ERROR");
            com.lumen.workflow.model.ActResponseError error = new com.lumen.workflow.model.ActResponseError();
            error.setMessage("Internal server error: " + e.getMessage());
            errorResponse.setError(error);
            
            return ResponseEntity.internalServerError().body(errorResponse);
        }
    }

    @GetMapping("/status/{requestId}")
    public ResponseEntity<String> getRequestStatus(@PathVariable String requestId) {
        try {
            WorkflowStub existingWorkflow = workflowClient.newUntypedWorkflowStub("nexus-" + requestId);
            String status = existingWorkflow.query("getRequestStatus", String.class, requestId);
            return ResponseEntity.ok(status);
        } catch (Exception e) {
            logger.error("Error getting status for request {}", requestId, e);
            return ResponseEntity.internalServerError().body("Error getting status: " + e.getMessage());
        }
    }

    @PostMapping("/cancel/{requestId}")
    public ResponseEntity<String> cancelRequest(@PathVariable String requestId) {
        try {
            WorkflowStub existingWorkflow = workflowClient.newUntypedWorkflowStub("nexus-" + requestId);
            existingWorkflow.signal("cancelRequest", requestId);
            return ResponseEntity.ok("Cancel signal sent for request: " + requestId);
        } catch (Exception e) {
            logger.error("Error cancelling request {}", requestId, e);
            return ResponseEntity.internalServerError().body("Error cancelling request: " + e.getMessage());
        }
    }
} 