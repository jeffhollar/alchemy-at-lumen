package com.lumen.workflow.client;

import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActRequestMeta;
import com.lumen.workflow.model.ActResponse;
import com.lumen.workflow.service.NexusEndpoint;
import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowClientOptions;
import io.temporal.client.WorkflowOptions;
import io.temporal.serviceclient.WorkflowServiceStubs;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.CommandLineRunner;
import org.springframework.context.annotation.Profile;
import org.springframework.stereotype.Component;

@Component
@Profile("integration-test")
public class NexusIntegrationClient implements CommandLineRunner {
    private static final Logger logger = LoggerFactory.getLogger(NexusIntegrationClient.class);
    private static final String TASK_QUEUE = "nexus-task-queue";
    private static final String NAMESPACE = "nexus-lumen-act-ns";

    @Override
    public void run(String... args) {
        try {
            // Connect to Temporal service
            WorkflowServiceStubs service = WorkflowServiceStubs.newInstance();
            WorkflowClient client = WorkflowClient.newInstance(service, 
                WorkflowClientOptions.newBuilder()
                    .setNamespace(NAMESPACE)
                    .build());

            // Create workflow options
            WorkflowOptions options = WorkflowOptions.newBuilder()
                .setTaskQueue(TASK_QUEUE)
                .build();

            // Create workflow stub
            NexusEndpoint workflow = client.newWorkflowStub(NexusEndpoint.class, options);

            // Create test request
            ActRequest request = createTestRequest();
            logger.info("Sending test request: {}", request.getMeta().getRequestId());

            // Execute workflow
            ActResponse response = workflow.processNexusRequest(request);
            logger.info("Received response: {}", response.getStatus());

            // Query status
            String status = workflow.getRequestStatus(request.getMeta().getRequestId());
            logger.info("Request status: {}", status);

            // Test cancellation
            workflow.cancelRequest(request.getMeta().getRequestId());
            status = workflow.getRequestStatus(request.getMeta().getRequestId());
            logger.info("Status after cancellation: {}", status);

        } catch (Exception e) {
            logger.error("Error running integration test", e);
            throw new RuntimeException("Integration test failed", e);
        }
    }

    private ActRequest createTestRequest() {
        ActRequest request = new ActRequest();
        ActRequestMeta meta = new ActRequestMeta();
        meta.setRequestId("integration-test-" + System.currentTimeMillis());
        meta.setTimestamp(java.time.Instant.now().toString());
        request.setMeta(meta);

        // Add test data
        request.setYang("{\n" +
            "  \"service\": {\n" +
            "    \"type\": \"test-service\",\n" +
            "    \"parameters\": {\n" +
            "      \"param1\": \"value1\",\n" +
            "      \"param2\": \"value2\"\n" +
            "    }\n" +
            "  }\n" +
            "}");
        request.setFeedback("Integration test feedback");
        request.setActivationTransactionId("integration-act-" + System.currentTimeMillis());

        return request;
    }

    public static void main(String[] args) {
        // For running directly from IDE or command line
        NexusIntegrationClient client = new NexusIntegrationClient();
        client.run(args);
    }
} 