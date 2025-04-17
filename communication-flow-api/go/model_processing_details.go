/*
 * Communication Flow API
 *
 * RESTful API for Communication Flow operations
 * This file defines the ProcessingDetails type which represents the status
 * and progress information for an ongoing task or workflow.
 *
 * The type is an alias to the internal models.ProcessingDetails type,
 * providing a clean interface for the API layer.
 */
package swagger

import "communication-flow-api/internal/models"

// ProcessingDetails represents the current status and progress information
// for an ongoing task or workflow. It is an alias to the internal
// models.ProcessingDetails type, which contains detailed information about
// the processing state, progress, and any relevant metadata.
type ProcessingDetails = models.ProcessingDetails
