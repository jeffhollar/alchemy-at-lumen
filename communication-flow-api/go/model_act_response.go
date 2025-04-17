/*
 * Communication Flow API
 *
 * RESTful API for Communication Flow operations
 * This file defines the ActResponse type which represents the output
 * data returned from an Automated Communication Task (ACT) execution.
 *
 * The type is an alias to the internal models.ActResponse type,
 * providing a clean interface for the API layer.
 */
package swagger

import "communication-flow-api/internal/models"

// ActResponse represents the output data returned after executing an Automated
// Communication Task. It is an alias to the internal models.ActResponse type,
// which contains the results and status information from the completed task.
type ActResponse = models.ActResponse
