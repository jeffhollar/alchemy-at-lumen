package com.lumen.workflow.service;

import io.temporal.activity.ActivityOptions;
import io.temporal.workflow.Workflow;
import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.HashMap;
import java.util.Map;

public class NexusEndpointImpl implements NexusEndpoint {
    private static final Logger logger = LoggerFactory.getLogger(NexusEndpointImpl.class);
    private final NexusService nexusService;
    private final Map<String, String> requestStatuses = new HashMap<>();

    public NexusEndpointImpl() {
        this.nexusService = Workflow.newActivityStub(
            NexusService.class,
            ActivityOptions.newBuilder()
                .setStartToCloseTimeout(java.time.Duration.ofMinutes(5))
                .setRetryOptions(io.temporal.common.RetryOptions.newBuilder()
                    .setInitialInterval(java.time.Duration.ofSeconds(1))
                    .setMaximumInterval(java.time.Duration.ofSeconds(10))
                    .setBackoffCoefficient(2.0)
                    .setMaximumAttempts(3)
                    .build())
                .build()
        );
    }

    @Override
    public ActResponse processNexusRequest(ActRequest request) {
        try {
            String requestId = request.getMeta().getRequestId();
            logger.info("Processing Nexus request ID: {}", requestId);
            requestStatuses.put(requestId, "PROCESSING");

            // Process the request directly without data conversion
            String result = nexusService.processRequest(requestId, request.toString());

            // Update status and create response
            requestStatuses.put(requestId, "COMPLETED");
            ActResponse response = new ActResponse();
            response.setStatus("SUCCESS");
            return response;
        } catch (Exception e) {
            String errorMsg = "Error processing Nexus request: " + e.getMessage();
            logger.error(errorMsg);
            requestStatuses.put(request.getMeta().getRequestId(), "FAILED");
            
            ActResponse response = new ActResponse();
            response.setStatus("ERROR");
            com.lumen.workflow.model.ActResponseError error = new com.lumen.workflow.model.ActResponseError();
            error.setMessage(errorMsg);
            response.setError(error);
            return response;
        }
    }

    @Override
    public String getRequestStatus(String requestId) {
        return requestStatuses.getOrDefault(requestId, "NOT_FOUND");
    }

    @Override
    public void cancelRequest(String requestId) {
        if (requestStatuses.containsKey(requestId)) {
            requestStatuses.put(requestId, "CANCELLED");
            logger.info("Request {} has been cancelled", requestId);
        } else {
            logger.warn("Attempt to cancel non-existent request: {}", requestId);
        }
    }
} 