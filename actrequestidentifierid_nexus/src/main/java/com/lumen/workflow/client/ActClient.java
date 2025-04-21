package com.lumen.workflow.client;

import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.serviceclient.WorkflowServiceStubs;
import com.lumen.workflow.options.ClientOptions;
import com.lumen.workflow.workflow.ActWorkflow;
import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActRequestMeta;
import com.lumen.workflow.model.ActResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class ActClient {
    private static final Logger logger = LoggerFactory.getLogger(ActClient.class);
    private static final String TASK_QUEUE = "act-task-queue";

    public static void main(String[] args) {
        if (args.length != 1) {
            System.err.println("Usage: ActClient <properties-file>");
            System.exit(1);
        }

        try {
            // Load client options from properties file
            ClientOptions options = ClientOptions.fromProperties(args[0]);

            // Create a Temporal service client
            WorkflowServiceStubs service = WorkflowServiceStubs.newInstance(
                    options.getServiceStubsOptions());
            WorkflowClient client = WorkflowClient.newInstance(
                    service, options.getClientOptions());

            // Create workflow options
            WorkflowOptions workflowOptions = WorkflowOptions.newBuilder()
                    .setTaskQueue(TASK_QUEUE)
                    .build();

            // Create a workflow stub
            ActWorkflow workflow = client.newWorkflowStub(ActWorkflow.class, workflowOptions);

            // Create a sample request
            ActRequest request = new ActRequest();
            ActRequestMeta meta = new ActRequestMeta();
            meta.setRequestId("test-request-1");
            meta.setTimestamp("2024-04-21T18:00:00Z");
            request.setMeta(meta);
            request.setFeedback("Test feedback");
            request.setYang("Test YANG data");

            // Start the workflow
            logger.info("Starting workflow with request: {}", request);
            ActResponse response = workflow.processAct(request);
            logger.info("Workflow completed with response: {}", response);

            // Shutdown the client
            service.shutdown();
        } catch (Exception e) {
            logger.error("Error running workflow", e);
            System.exit(1);
        }
    }
} 