package com.lumen.workflow.nexus;

import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActResponse;
import com.lumen.workflow.workflow.ActWorkflow;

public class ActNexusServiceImpl implements ActNexusService {
    private final WorkflowClient client;
    private final String taskQueue;

    public ActNexusServiceImpl(WorkflowClient client, String taskQueue) {
        this.client = client;
        this.taskQueue = taskQueue;
    }

    @Override
    public ActResponse processAct(ActRequest request) {
        WorkflowOptions options = WorkflowOptions.newBuilder()
            .setTaskQueue(taskQueue)
            .build();
            
        ActWorkflow workflow = client.newWorkflowStub(ActWorkflow.class, options);
        return workflow.processAct(request);
    }
} 