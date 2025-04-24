package com.lumen.workflow.service;

import io.temporal.activity.ActivityInterface;
import io.temporal.activity.ActivityMethod;

@ActivityInterface
public interface NexusService {
    @ActivityMethod
    String processRequest(String requestId, String requestData);
} 