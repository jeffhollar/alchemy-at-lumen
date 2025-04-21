package com.lumen.workflow.workflow;

import io.temporal.workflow.Workflow;
import io.temporal.activity.ActivityOptions;
import com.lumen.workflow.model.ActRequest;
import com.lumen.workflow.model.ActResponse;
import com.lumen.workflow.model.ActResponseError;
import com.lumen.workflow.service.ActService;
import org.slf4j.Logger;

import java.time.Duration;

public class ActWorkflowImpl implements ActWorkflow {
    private static final Logger logger = Workflow.getLogger(ActWorkflowImpl.class);
    
    private final ActService actService = Workflow.newActivityStub(
        ActService.class,
        ActivityOptions.newBuilder()
            .setStartToCloseTimeout(Duration.ofSeconds(30))
            .build()
    );

    @Override
    public ActResponse processAct(ActRequest request) {
        logger.info("Starting ActWorkflow with request: {}", request);
        
        try {
            // Call the activity
            ActResponse response = actService.processAct(request);
            logger.info("Received response from ActService: {}", response);
            
            return response;
        } catch (Exception e) {
            logger.error("Error in ActWorkflow", e);
            ActResponse response = new ActResponse();
            response.setStatus("ERROR");
            
            ActResponseError error = new ActResponseError();
            error.setMessage(e.getMessage());
            response.setError(error);
            
            return response;
        }
    }
} 