package com.lumen.workflow.service;

import io.temporal.workflow.QueryMethod;
import io.temporal.workflow.SignalMethod;
import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;
import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActResponse;

@WorkflowInterface
public interface NexusEndpoint {
    @WorkflowMethod
    ActResponse processNexusRequest(ActRequest request);

    @QueryMethod
    String getRequestStatus(String requestId);

    @SignalMethod
    void cancelRequest(String requestId);
} 