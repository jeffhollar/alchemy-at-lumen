package com.lumen.workflow.workflow;

import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;
import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActResponse;

@WorkflowInterface
public interface ActWorkflow {
    
    @WorkflowMethod
    ActResponse processAct(ActRequest request);
} 