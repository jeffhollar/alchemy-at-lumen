/*
 *  Copyright (c) 2025 Lumen Technologies, Inc. All Rights Reserved
 *
 */
package com.lumen.workflow.handler;

import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;
import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActResponse;

@WorkflowInterface
public interface ActWorkflow {
    
    @WorkflowMethod
    ActResponse processAct(ActRequest request);
} 