package com.lumen.workflow.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonInclude(JsonInclude.Include.NON_EMPTY)
public class ActResponse {
    
    @JsonProperty("status")
    private String status;

    @JsonProperty("error")
    private ActResponseError error;

    // Getters and Setters
    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public ActResponseError getError() {
        return error;
    }

    public void setError(ActResponseError error) {
        this.error = error;
    }
} 