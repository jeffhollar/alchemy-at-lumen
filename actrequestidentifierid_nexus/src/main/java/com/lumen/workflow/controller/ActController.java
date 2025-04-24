package com.lumen.workflow.controller;

import org.springframework.web.bind.annotation.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActResponse;
import com.lumen.workflow.model.ActResponseError;
import com.lumen.workflow.workflow.ActWorkflow;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@RestController
@RequestMapping("/api/v1/act")
public class ActController {
    private static final Logger logger = LoggerFactory.getLogger(ActController.class);
    private static final String TASK_QUEUE = "act-task-queue";

    private final WorkflowClient workflowClient;
    private final ObjectMapper objectMapper;

    @Autowired
    public ActController(WorkflowClient workflowClient, ObjectMapper objectMapper) {
        this.workflowClient = workflowClient;
        this.objectMapper = objectMapper;
    }

    @PostMapping("/process")
    public ResponseEntity<ActResponse> processAct(@RequestBody ActRequest request) {
        logger.info("Received request with ID: {}", request.getMeta().getRequestId());

        try {
            // Validate request
            if (request.getMeta() == null || request.getMeta().getRequestId() == null) {
                ActResponse errorResponse = new ActResponse();
                errorResponse.setStatus("ERROR");
                ActResponseError error = new ActResponseError();
                error.setMessage("Request ID is required");
                errorResponse.setError(error);
                return ResponseEntity.badRequest().body(errorResponse);
            }

            // Create workflow options
            WorkflowOptions options = WorkflowOptions.newBuilder()
                    .setTaskQueue(TASK_QUEUE)
                    .build();

            // Convert request to JSON string
            String requestData = objectMapper.writeValueAsString(request);

            // Start the workflow
            ActWorkflow workflow = workflowClient.newWorkflowStub(ActWorkflow.class, options);
            String result = workflow.processRequest(request.getMeta().getRequestId(), requestData);
            
            // Convert result to ActResponse
            ActResponse response = new ActResponse();
            response.setStatus("SUCCESS");
            logger.info("Request {} processed successfully", request.getMeta().getRequestId());
            return ResponseEntity.ok(response);
        } catch (Exception e) {
            logger.error("Error processing request {}", request.getMeta().getRequestId(), e);
            
            ActResponse errorResponse = new ActResponse();
            errorResponse.setStatus("ERROR");
            ActResponseError error = new ActResponseError();
            error.setMessage("Internal server error: " + e.getMessage());
            errorResponse.setError(error);
            
            return ResponseEntity.internalServerError().body(errorResponse);
        }
    }
} 