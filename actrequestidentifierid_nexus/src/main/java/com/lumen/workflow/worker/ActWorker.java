package com.lumen.workflow.worker;

import io.temporal.worker.Worker;
import io.temporal.worker.WorkerFactory;
import io.temporal.client.WorkflowClient;
import com.lumen.workflow.workflow.ActWorkflowImpl;
import com.lumen.workflow.service.NexusServiceImpl;
import com.lumen.workflow.config.TemporalProperties;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import jakarta.annotation.PostConstruct;
import jakarta.annotation.PreDestroy;

@Component
public class ActWorker {
    private static final Logger logger = LoggerFactory.getLogger(ActWorker.class);
    
    private final WorkflowClient workflowClient;
    private final WorkerFactory factory;
    private final TemporalProperties temporalProperties;
    private Worker worker;
    
    @Autowired
    public ActWorker(WorkflowClient workflowClient, TemporalProperties temporalProperties) {
        this.workflowClient = workflowClient;
        this.temporalProperties = temporalProperties;
        this.factory = WorkerFactory.newInstance(workflowClient);
    }
    
    @PostConstruct
    public void start() {
        try {
            String taskQueue = temporalProperties.getTaskQueue();
            logger.info("Starting worker with task queue: {}", taskQueue);
            
            // Create a worker
            worker = factory.newWorker(taskQueue);
            
            // Register workflow and activities
            worker.registerWorkflowImplementationTypes(ActWorkflowImpl.class);
            worker.registerActivitiesImplementations(new NexusServiceImpl());
            
            // Start the worker factory
            factory.start();
            
            logger.info("Worker started successfully");
        } catch (Exception e) {
            logger.error("Error starting worker", e);
            throw new RuntimeException("Failed to start worker", e);
        }
    }
    
    @PreDestroy
    public void stop() {
        logger.info("Shutting down worker...");
        if (factory != null) {
            factory.shutdown();
        }
    }
} 