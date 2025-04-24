/*
 *  Copyright (c) 2025 Lumen Technologies, Inc. All Rights Reserved
 *
 */

package io.temporal.samples.nexus.caller;

import io.temporal.api.common.v1.WorkflowExecution;
import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.samples.nexus.options.ClientOptions;
import io.temporal.samples.nexus.service.NexusService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class CallerStarter {
  private static final Logger logger = LoggerFactory.getLogger(CallerStarter.class);

  public static void main(String[] args) {
    WorkflowClient client = ClientOptions.getWorkflowClient(args);

    WorkflowOptions workflowOptions =
        WorkflowOptions.newBuilder().setTaskQueue(CallerWorker.DEFAULT_TASK_QUEUE_NAME).build();

    ActCallerWorkflow actWorkflow =
        client.newWorkflowStub(ActCallerWorkflow.class, workflowOptions);

    ActRequest actRequest = new ActRequest();
    actRequest.setActivationTransactionId("1234567890");

    WorkflowExecution execution = WorkflowClient.start(actWorkflow::processAct, actRequest);

    logger.info(
        "Started ActCallerWorkflow workflowId: {} runId: {}",
        execution.getWorkflowId(),
        execution.getRunId());
    logger.info("Workflow result: {}", actWorkflow.processAct(actRequest));
  }
}