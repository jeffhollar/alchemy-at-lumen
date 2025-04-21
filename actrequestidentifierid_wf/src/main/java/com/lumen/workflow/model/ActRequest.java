package com.lumen.workflow.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class ActRequest {
    
    @JsonProperty("meta")
    private ActRequestMeta meta;
    
    @JsonProperty("feedback")
    private String feedback;
    
    @JsonProperty("yang")
    private Object yang;
    
    @JsonProperty("activationTransactionId")
    private String activationTransactionId;
    
    @JsonProperty("yang.error")
    private String yangError;

    // Getters and Setters
    public ActRequestMeta getMeta() {
        return meta;
    }

    public void setMeta(ActRequestMeta meta) {
        this.meta = meta;
    }

    public String getFeedback() {
        return feedback;
    }

    public void setFeedback(String feedback) {
        this.feedback = feedback;
    }

    public Object getYang() {
        return yang;
    }

    public void setYang(Object yang) {
        this.yang = yang;
    }

    public String getActivationTransactionId() {
        return activationTransactionId;
    }

    public void setActivationTransactionId(String activationTransactionId) {
        this.activationTransactionId = activationTransactionId;
    }

    public String getYangError() {
        return yangError;
    }

    public void setYangError(String yangError) {
        this.yangError = yangError;
    }
} 