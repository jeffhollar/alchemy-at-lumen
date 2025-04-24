package com.lumen.workflow.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.lumen.workflow.model.ActRequest;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

@Service
public class NexusServiceImpl implements NexusService {
    private static final Logger logger = LoggerFactory.getLogger(NexusServiceImpl.class);
    private final ObjectMapper objectMapper = new ObjectMapper();

    @Override
    public String processRequest(String requestId, String requestData) {
        try {
            logger.info("Processing Nexus request ID: {}", requestId);
            
            // Deserialize the request data
            ActRequest request = objectMapper.readValue(requestData, ActRequest.class);
            
            // Process YANG data if available
            if (request.getYang() != null) {
                logger.info("Processing YANG data for request ID: {}", requestId);
                // Add your YANG data processing logic here
            }
            
            // Process feedback if available
            if (request.getFeedback() != null) {
                logger.info("Processing feedback for request ID: {}", requestId);
                // Add your feedback processing logic here
            }
            
            // Handle activation transaction if available
            if (request.getActivationTransactionId() != null) {
                logger.info("Processing activation transaction ID: {}", request.getActivationTransactionId());
                // Add your activation transaction processing logic here
            }
            
            // Check for YANG errors
            if (request.getYangError() != null) {
                logger.warn("YANG error detected for request ID {}: {}", requestId, request.getYangError());
                // Add your YANG error handling logic here
            }
            
            logger.info("Successfully processed Nexus request ID: {}", requestId);
            return "Request " + requestId + " processed successfully";
        } catch (Exception e) {
            String errorMsg = "Failed to process Nexus request: " + e.getMessage();
            logger.error(errorMsg, e);
            throw new RuntimeException(errorMsg, e);
        }
    }
} 