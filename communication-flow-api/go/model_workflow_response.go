/*
 * Network Services Adapter API
 *
 * This API provides an adapter for managing network services using Temporal.io.
 * This file defines the WorkflowResponse type which represents the output
 * data returned from a workflow execution.
 *
 * The type is an alias to the internal models.WorkflowResponse type,
 * providing a clean interface for the API layer.
 */
package swagger

import "communication-flow-api/internal/models"

// WorkflowResponse represents the output data returned after workflow execution.
// It is an alias to the internal models.WorkflowResponse type, which contains
// the results and status information from the completed workflow.
type WorkflowResponse = models.WorkflowResponse
