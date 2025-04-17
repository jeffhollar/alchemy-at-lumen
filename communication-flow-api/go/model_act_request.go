/*
 * Communication Flow API
 *
 * RESTful API for Communication Flow operations
 * This file defines the ActRequest type which represents the input
 * parameters for an Automated Communication Task (ACT) request.
 *
 * The type is an alias to the internal models.ActRequest type,
 * providing a clean interface for the API layer.
 */
package swagger

import "communication-flow-api/internal/models"

// ActRequest represents the input parameters for initiating an Automated
// Communication Task. It is an alias to the internal models.ActRequest type,
// which contains the necessary data for task initialization and execution.
type ActRequest = models.ActRequest
