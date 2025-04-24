package com.lumen.workflow.client;

import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import com.lumen.workflow.workflow.ActWorkflow;
import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActRequestMeta;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.CommandLineRunner;
import org.springframework.context.annotation.Profile;
import org.springframework.stereotype.Component;

@Component
@Profile("client")
public class ActClient implements CommandLineRunner {
    private static final Logger logger = LoggerFactory.getLogger(ActClient.class);
    
    @Value("${temporal.task-queue}")
    private String taskQueue;
    
    private final WorkflowClient workflowClient;
    private final ObjectMapper objectMapper;

    public ActClient(WorkflowClient workflowClient, ObjectMapper objectMapper) {
        this.workflowClient = workflowClient;
        this.objectMapper = objectMapper;
    }

    @Override
    public void run(String... args) {
        try {
            // Create workflow options
            WorkflowOptions workflowOptions = WorkflowOptions.newBuilder()
                    .setTaskQueue(taskQueue)
                    .build();

            // Create a workflow stub
            ActWorkflow workflow = workflowClient.newWorkflowStub(ActWorkflow.class, workflowOptions);

            // Create a sample request
            ActRequest request = new ActRequest();
            ActRequestMeta meta = new ActRequestMeta();
            meta.setRequestId("test-request-1");
            meta.setTimestamp("2024-04-21T18:00:00Z");
            request.setMeta(meta);
            request.setFeedback("Test feedback");
            request.setYang("Test YANG data");

            // Convert request to JSON string
            String requestData = objectMapper.writeValueAsString(request);

            // Start the workflow
            logger.info("Starting workflow with request: {}", request);
            String result = workflow.processRequest(request.getMeta().getRequestId(), requestData);
            logger.info("Workflow completed with result: {}", result);
        } catch (Exception e) {
            logger.error("Error running workflow", e);
            throw new RuntimeException("Failed to run workflow", e);
        }
    }
} 