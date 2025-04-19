/*
 * Network Services Adapter API
 *
 * This package provides an adapter for interacting with Temporal.io workflows.
 * It handles the initialization and execution of workflows for network service operations.
 *
 * The adapter acts as a bridge between the API layer and Temporal.io, managing
 * workflow lifecycle and result handling.
 */
package adapters

import (
	"context"
	"log"

	"act-core-api/internal/models"
	"act-core-api/internal/workflows"
	"go.temporal.io/sdk/client"
)

// TemporalClient represents a client for interacting with Temporal.io workflows.
// Fields:
//   - Client: The underlying Temporal.io client instance
type TemporalClient struct {
	Client client.Client
}

// NewTemporalClient creates a new TemporalClient instance with the provided
// Temporal.io client.
// Parameters:
//   - temporalClient: The Temporal.io client to use for workflow operations
//
// Returns:
//   - *TemporalClient: A new TemporalClient instance
func NewTemporalClient(temporalClient client.Client) *TemporalClient {
	return &TemporalClient{
		Client: temporalClient,
	}
}

// StartWorkflow initiates a new workflow execution in Temporal.io for processing
// an Automated Communication Task (ACT) request.
// Parameters:
//   - ctx: Context for the operation
//   - request: The ACT request to process
//
// Returns:
//   - models.ActResponse: The response from the workflow execution
//   - error: Any error that occurred during workflow execution
func (t *TemporalClient) StartWorkflow(ctx context.Context, request models.ActRequest) (models.ActResponse, error) {
	log.Printf("\nStartWorkflow Invoked\n *********************\n")

	options := client.StartWorkflowOptions{
		ID:        "act-communication-workflow-" + request.IdentifierId, // Unique workflow ID
		TaskQueue: "act-communication-task-queue",
	}

	we, err := t.Client.ExecuteWorkflow(ctx, options, workflows.GetActRequestWorkflow, request)
	if err != nil {
		log.Printf("Unable to execute workflow: %v", err)
		return models.ActResponse{}, err
	}
	log.Printf("\nStarted workflow WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	// Wait for workflow completion and get the result
	var result models.ActResponse
	err = we.Get(ctx, &result)
	if err != nil {
		log.Printf("Error getting workflow result: %v", err)
		return models.ActResponse{}, err
	}
	status_result := we.GetID() + " : " + we.GetRunID() + " : " + "COMPLETED"

	result.Status = status_result
	log.Printf("\nWorkflow completed successfully with result: %+v", result)
	return result, nil
}
