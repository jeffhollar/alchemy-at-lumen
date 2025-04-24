/*
 *  Copyright (c) 2025 Lumen Technologies, Inc. All Rights Reserved
 *
 */

package io.temporal.samples.nexus.caller;

import io.temporal.client.WorkflowClient;
import io.temporal.samples.nexus.options.ClientOptions;
import io.temporal.worker.Worker;
import io.temporal.worker.WorkerFactory;
import io.temporal.worker.WorkflowImplementationOptions;
import io.temporal.workflow.NexusServiceOptions;
import java.util.Collections;

public class CallerWorker {
  public static final String DEFAULT_TASK_QUEUE_NAME = "act-task-queue";

  public static void main(String[] args) {
    WorkflowClient client = ClientOptions.getWorkflowClient(args);

    WorkerFactory factory = WorkerFactory.newInstance(client);

    Worker worker = factory.newWorker(DEFAULT_TASK_QUEUE_NAME);
    worker.registerWorkflowImplementationTypes(
        WorkflowImplementationOptions.newBuilder()
            .setNexusServiceOptions(
                Collections.singletonMap(
                    "NexusService",
                    NexusServiceOptions.newBuilder().setEndpoint("act-requestidentifier-service").build()))
            .build(),
        ActWorkflowImpl.class);

    factory.start();
  }
}