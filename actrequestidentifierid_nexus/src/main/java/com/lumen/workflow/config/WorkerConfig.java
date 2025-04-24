package com.lumen.workflow.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import io.temporal.worker.Worker;
import io.temporal.worker.WorkerFactory;
import io.temporal.client.WorkflowClient;
import com.lumen.workflow.workflow.ActWorkflowImpl;
import com.lumen.workflow.service.ActServiceImpl;
import com.lumen.workflow.service.NexusServiceImpl;
import com.lumen.workflow.service.NexusEndpointImpl;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@Configuration
public class WorkerConfig {
    private static final Logger logger = LoggerFactory.getLogger(WorkerConfig.class);
    private static final String ACT_TASK_QUEUE = "act-task-queue";
    private static final String NEXUS_TASK_QUEUE = "nexus-task-queue";

    @Bean(initMethod = "start", destroyMethod = "shutdown")
    public WorkerFactory workerFactory(WorkflowClient workflowClient) {
        WorkerFactory factory = WorkerFactory.newInstance(workflowClient);
        
        // Create and configure ACT worker
        Worker actWorker = factory.newWorker(ACT_TASK_QUEUE);
        actWorker.registerWorkflowImplementationTypes(ActWorkflowImpl.class);
        actWorker.registerActivitiesImplementations(new ActServiceImpl());
        logger.info("Temporal ACT worker configured for task queue: {}", ACT_TASK_QUEUE);
        
        // Create and configure Nexus worker
        Worker nexusWorker = factory.newWorker(NEXUS_TASK_QUEUE);
        nexusWorker.registerWorkflowImplementationTypes(NexusEndpointImpl.class);
        nexusWorker.registerActivitiesImplementations(new NexusServiceImpl());
        logger.info("Temporal Nexus worker configured for task queue: {}", NEXUS_TASK_QUEUE);
        
        return factory;
    }
} 