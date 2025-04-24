package com.lumen.workflow.workflow;

import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;

@WorkflowInterface
public interface ActWorkflow {
    @WorkflowMethod
    String processRequest(String requestId, String requestData);
} 