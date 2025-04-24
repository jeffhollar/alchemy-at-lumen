package com.lumen.workflow.workflow;

import io.temporal.activity.ActivityOptions;
import io.temporal.workflow.Workflow;
import com.lumen.workflow.service.NexusService;
import java.time.Duration;

public class ActWorkflowImpl implements ActWorkflow {
    private final NexusService nexusService;
    private final ActivityOptions options;

    public ActWorkflowImpl() {
        this.options = ActivityOptions.newBuilder()
                .setStartToCloseTimeout(Duration.ofSeconds(30))
                .setRetryOptions(
                    io.temporal.common.RetryOptions.newBuilder()
                        .setInitialInterval(Duration.ofSeconds(1))
                        .setMaximumInterval(Duration.ofSeconds(10))
                        .setBackoffCoefficient(2.0)
                        .setMaximumAttempts(3)
                        .build()
                )
                .build();
        
        this.nexusService = Workflow.newActivityStub(NexusService.class, options);
    }

    @Override
    public String processRequest(String requestId, String requestData) {
        try {
            return nexusService.processRequest(requestId, requestData);
        } catch (Exception e) {
            Workflow.getLogger(ActWorkflowImpl.class).error("Error processing request: " + e.getMessage());
            throw e;
        }
    }
} 