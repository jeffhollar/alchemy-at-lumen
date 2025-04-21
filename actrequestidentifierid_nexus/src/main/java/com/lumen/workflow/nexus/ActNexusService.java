package com.lumen.workflow.nexus;

import io.temporal.nexus.NexusService;
import io.temporal.nexus.NexusOperation;
import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActResponse;

@NexusService(name = "actrequestidentifierid")
public interface ActNexusService {
    @NexusOperation
    ActResponse processAct(ActRequest request);
} 

