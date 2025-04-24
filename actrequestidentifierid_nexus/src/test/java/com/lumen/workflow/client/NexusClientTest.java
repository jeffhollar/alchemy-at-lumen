package com.lumen.workflow.client;

import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActRequestMeta;
import com.lumen.workflow.model.ActResponse;
import com.lumen.workflow.service.NexusEndpoint;
import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.testing.TestWorkflowEnvironment;
import io.temporal.worker.Worker;
import com.lumen.workflow.service.NexusEndpointImpl;
import com.lumen.workflow.service.NexusServiceImpl;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

public class NexusClientTest {
    private TestWorkflowEnvironment testEnv;
    private Worker worker;
    private WorkflowClient client;
    private static final String TASK_QUEUE = "test-nexus-task-queue";

    @BeforeEach
    public void setUp() {
        testEnv = TestWorkflowEnvironment.newInstance();
        worker = testEnv.newWorker(TASK_QUEUE);
        worker.registerWorkflowImplementationTypes(NexusEndpointImpl.class);
        worker.registerActivitiesImplementations(new NexusServiceImpl());
        testEnv.start();
        client = testEnv.getWorkflowClient();
    }

    @AfterEach
    public void tearDown() {
        testEnv.close();
    }

    @Test
    public void testProcessNexusRequest() {
        // Create workflow options
        WorkflowOptions options = WorkflowOptions.newBuilder()
                .setTaskQueue(TASK_QUEUE)
                .build();

        // Create workflow stub
        NexusEndpoint workflow = client.newWorkflowStub(NexusEndpoint.class, options);

        // Create test request
        ActRequest request = new ActRequest();
        ActRequestMeta meta = new ActRequestMeta();
        meta.setRequestId("test-" + System.currentTimeMillis());
        meta.setTimestamp(java.time.Instant.now().toString());
        request.setMeta(meta);
        request.setYang("{ \"test\": \"data\" }");
        request.setFeedback("Test feedback");
        request.setActivationTransactionId("test-act-123");

        // Execute workflow
        ActResponse response = workflow.processNexusRequest(request);

        // Verify response
        assertNotNull(response);
        assertEquals("SUCCESS", response.getStatus());
        assertNull(response.getError());

        // Test status query
        String status = workflow.getRequestStatus(request.getMeta().getRequestId());
        assertEquals("COMPLETED", status);

        // Test cancellation
        workflow.cancelRequest(request.getMeta().getRequestId());
        status = workflow.getRequestStatus(request.getMeta().getRequestId());
        assertEquals("CANCELLED", status);
    }

    @Test
    public void testErrorHandling() {
        WorkflowOptions options = WorkflowOptions.newBuilder()
                .setTaskQueue(TASK_QUEUE)
                .build();

        NexusEndpoint workflow = client.newWorkflowStub(NexusEndpoint.class, options);

        // Create invalid request (missing meta)
        ActRequest request = new ActRequest();
        request.setYang("{ \"invalid\": \"data\" }");
        request.setYangError("Test error");

        // Execute workflow and expect error
        ActResponse response = workflow.processNexusRequest(request);

        // Verify error response
        assertNotNull(response);
        assertEquals("ERROR", response.getStatus());
        assertNotNull(response.getError());
        assertNotNull(response.getError().getMessage());
    }
} 