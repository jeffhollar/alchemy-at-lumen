/*
 * Network Services Adapter API
 *
 * This API provides an adapter for managing network services using Temporal.io.
 * This file defines the WorkflowRequest type which represents the input
 * parameters for initiating a workflow execution.
 *
 * The type is an alias to the internal models.WorkflowRequest type,
 * providing a clean interface for the API layer.
 */
package swagger

import "communication-flow-api/internal/models"

// WorkflowRequest represents the input parameters required to start a workflow.
// It is an alias to the internal models.WorkflowRequest type, which contains
// the necessary data for workflow initialization and execution.
type WorkflowRequest = models.WorkflowRequest
