package com.lumen.workflow.service;

import io.temporal.activity.ActivityInterface;
import io.temporal.activity.ActivityMethod;
import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActResponse;

@ActivityInterface
public interface ActService {
    
    @ActivityMethod
    ActResponse processAct(ActRequest request);
} 