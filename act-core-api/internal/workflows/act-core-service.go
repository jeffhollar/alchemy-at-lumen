/*
 * Network Services Adapter API
 *
 * This package contains the workflow definitions for network service operations.
 * Workflows orchestrate the execution of activities and manage the overall
 * process flow of network service operations.
 *
 * The workflows in this package define the high-level business logic and
 * error handling for network service operations.
 */
package workflows

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"act-core-api/internal/activities"
	"act-core-api/internal/models"
)

// GetActRequestWorkflow orchestrates the execution of an Automated Communication
// Task (ACT) request. It sets up the workflow context, executes the necessary
// activities, and handles the response.
// Parameters:
//   - ctx: Workflow context containing execution information
//   - request: The ACT request to process
//
// Returns:
//   - *models.ActResponse: The response from the workflow execution
//   - error: Any error that occurred during workflow execution
func GetActRequestWorkflow(ctx workflow.Context, request models.ActRequest) (*models.ActResponse, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("\nStarting GetActRequestWorkflow\n *********************\n", "request", request)

	// Set workflow options
	options := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Hour,
		HeartbeatTimeout:    time.Second * 30,
	})

	// Execute the act request
	var result models.ActResponse
	err := workflow.ExecuteActivity(options, activities.ExecuteActRequestOperation, request).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed", "error", err)
		return nil, err
	}

	logger.Info("\nNetworkServiceWorkflow completed", "result", result)
	return &result, nil
}
