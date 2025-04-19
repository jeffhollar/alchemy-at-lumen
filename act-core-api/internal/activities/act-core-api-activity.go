/*
 * Network Services Adapter API
 *
 * This package contains the activities that perform the actual network service operations.
 * Activities are the building blocks of workflows and represent the actual work being done.
 *
 * The activities in this package handle the execution of network service operations
 * and manage the interaction with external services.
 */
package activities

import (
	"context"
	// "log"
	"time"

	"act-core-api/internal/models"
	"go.temporal.io/sdk/activity"
)

// ExecuteActRequestOperation performs the actual network service operation for an
// Automated Communication Task (ACT) request. It simulates the network service
// operation and returns a response with the operation status.
// Parameters:
//   - ctx: Context containing workflow execution information
//   - request: The ACT request to process
//
// Returns:
//   - *models.ActResponse: The response containing operation status and results
//   - error: Any error that occurred during the operation
func ExecuteActRequestOperation(ctx context.Context, request models.ActRequest) (*models.ActResponse, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("\nExecuting activity operation\n *********************\n", "request", request)

	// Get workflow info from context
	info := activity.GetInfo(ctx)
	workflowID := info.WorkflowExecution.ID
	runID := info.WorkflowExecution.RunID

	// Simulate network service operation
	time.Sleep(2 * time.Second)

	status_result := workflowID + " : " + runID + " : " + "COMPLETED"

	// Create response
	response := &models.ActResponse{
		Status: status_result,
		Error_: nil,
	}

	logger.Info("\nExecuteNetworkServiceOperation completed", "response", response)
	return response, nil
}
