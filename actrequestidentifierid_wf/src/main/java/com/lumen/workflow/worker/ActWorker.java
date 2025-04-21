package com.lumen.workflow.worker;

import io.temporal.worker.Worker;
import io.temporal.worker.WorkerFactory;
import io.temporal.client.WorkflowClient;
import io.temporal.serviceclient.WorkflowServiceStubs;
import com.lumen.workflow.options.ClientOptions;
import com.lumen.workflow.workflow.ActWorkflowImpl;
import com.lumen.workflow.service.ActServiceImpl;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class ActWorker {
    private static final Logger logger = LoggerFactory.getLogger(ActWorker.class);
    
    public static void main(String[] args) {
        if (args.length != 1) {
            System.err.println("Usage: ActWorker <properties-file>");
            System.exit(1);
        }
        
        try {
            // Load client options from properties file
            ClientOptions options = ClientOptions.fromProperties(args[0]);
            
            // Create a Temporal service client with the loaded options
            WorkflowServiceStubs service = WorkflowServiceStubs.newInstance(
                    options.getServiceStubsOptions());
            WorkflowClient client = WorkflowClient.newInstance(
                    service, options.getClientOptions());
            
            // Create a worker factory
            WorkerFactory factory = WorkerFactory.newInstance(client);
            
            // Create a worker
            Worker worker = factory.newWorker("act-task-queue");
            
            // Register workflow and activities
            worker.registerWorkflowImplementationTypes(ActWorkflowImpl.class);
            worker.registerActivitiesImplementations(new ActServiceImpl());
            
            // Start the worker
            factory.start();
            
            logger.info("Worker started for task queue: act-task-queue");
            logger.info("Connected to Temporal server at: {}", options.getTargetHost());
            logger.info("Using namespace: {}", options.getNamespace());
            
            // Keep the worker running
            Runtime.getRuntime().addShutdownHook(new Thread(() -> {
                logger.info("Shutting down worker...");
                factory.shutdown();
                service.shutdown();
            }));
            
            // Wait for shutdown
            Thread.currentThread().join();
        } catch (Exception e) {
            logger.error("Error starting worker", e);
            System.exit(1);
        }
    }
} 