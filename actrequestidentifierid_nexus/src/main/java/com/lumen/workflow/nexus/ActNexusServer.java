package com.lumen.workflow.nexus;

import io.temporal.nexus.NexusServer;
import io.temporal.client.WorkflowClient;
import io.temporal.serviceclient.WorkflowServiceStubs;
import io.temporal.client.WorkflowClientOptions;

public class ActNexusServer {
    public static void main(String[] args) {
        // Create the client
        WorkflowServiceStubs service = WorkflowServiceStubs.newInstance();
        WorkflowClientOptions clientOptions = WorkflowClientOptions.newBuilder()
            .setNamespace("nexus-lumen-act-ns")
            .build();
        WorkflowClient client = WorkflowClient.newInstance(service, clientOptions);
        
        // Create the Nexus service implementation
        ActNexusService serviceImpl = new ActNexusServiceImpl(client, "act-task-queue");
        
        // Create and start the Nexus server
        NexusServer server = NexusServer.newBuilder()
            .setPort(8080)
            .addService(ActNexusService.class, serviceImpl)
            .enableOpenApi()
            .build();
            
        server.start();
    }
} 