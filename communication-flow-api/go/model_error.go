/*
 * Communication Flow API
 *
 * RESTful API for Communication Flow operations
 * This file defines the ModelError type which represents error information
 * returned by the API when an operation fails.
 *
 * The type is an alias to the internal models.ModelError type,
 * providing a clean interface for the API layer.
 */
package swagger

import "communication-flow-api/internal/models"

// ModelError represents error information returned by the API when an operation fails.
// It is an alias to the internal models.ModelError type, which contains
// error details including error codes, messages, and any additional context
// about the failure.
type ModelError = models.ModelError
