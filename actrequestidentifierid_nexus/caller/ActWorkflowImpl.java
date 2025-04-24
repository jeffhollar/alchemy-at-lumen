/*
 *  Copyright (c) 2025 Lumen Technologies, Inc. All Rights Reserved
 *
 */
package com.lumen.workflow.handler;

import io.temporal.samples.nexus.service.NexusService;
import io.temporal.workflow.NexusOperationOptions;
import io.temporal.workflow.NexusServiceOptions;
import io.temporal.workflow.Workflow;
import io.temporal.activity.ActivityOptions;
import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActResponse;
import com.lumen.workflow.model.ActResponseError;
import com.lumen.workflow.service.ActService;
import org.slf4j.Logger;
import java.time.Duration;

public class ActWorkflowImpl implements ActWorkflow {

    NexusService nexusService =
      Workflow.newNexusServiceStub(
          NexusService.class,
          NexusServiceOptions.newBuilder()
              .setOperationOptions(
                  NexusOperationOptions.newBuilder()
                      .setScheduleToCloseTimeout(Duration.ofSeconds(10))
                      .build())
              .build());

    @Override
    public ActResponse processAct(ActRequest request) {
        return nexusService.processAct(request);
    }
} 