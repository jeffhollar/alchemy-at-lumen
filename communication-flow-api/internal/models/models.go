/*
 * Network Services Adapter API
 *
 * This package contains the data models used throughout the application.
 * These models define the structure of requests, responses, and internal
 * data representations for network service operations.
 *
 * The models are used for serialization/deserialization of JSON data
 * and provide type safety for the application's data structures.
 */
package models

// ActRequestHeader represents the header information for an ACT request.
// Fields:
//   - Activity: The type of activity to perform
type ActRequestHeader struct {
	Activity string `json:"activity,omitempty"`
}

// ActRequestMeta represents metadata for an ACT request.
// Fields:
//   - GenerateOnly: Flag indicating if the request should only generate data
type ActRequestMeta struct {
	GenerateOnly string `json:"generate-only,omitempty"`
}

// ActRequest represents a request for an Automated Communication Task.
// Fields:
//   - IdentifierId: Unique identifier for the request
//   - Header: Header information for the request
//   - Status: Current status of the request
//   - ErrorMessage: Any error message associated with the request
//   - Meta: Metadata for the request
type ActRequest struct {
	IdentifierId string            `json:"identifier_id,omitempty"`
	Header       *ActRequestHeader `json:"header,omitempty"`
	Status       string            `json:"status,omitempty"`
	ErrorMessage string            `json:"error_message,omitempty"`
	Meta         *ActRequestMeta   `json:"meta,omitempty"`
}

// ActResponseError represents error information in an ACT response.
// Fields:
//   - Message: Description of the error
type ActResponseError struct {
	Message string `json:"message,omitempty"`
}

// ActResponse represents the response from an Automated Communication Task.
// Fields:
//   - Status: Status of the response
//   - Error_: Error information if the request failed
type ActResponse struct {
	Status string            `json:"status,omitempty"`
	Error_ *ActResponseError `json:"error,omitempty"`
}

// ErrorError represents detailed error information.
// Fields:
//   - Message: Description of the error
//   - Code: Error code
type ErrorError struct {
	Message string `json:"message,omitempty"`
	Code    int32  `json:"code,omitempty"`
}

// ModelError represents a standardized error response.
// Fields:
//   - Error_: Detailed error information
type ModelError struct {
	Error_ *ErrorError `json:"error,omitempty"`
}

// ProcessingDetailsProcessingDetailsRequestInstance represents timing and status
// information for a request instance.
// Fields:
//   - RequestStartDateTime: When the request started
//   - RequestEndDateTime: When the request ended
//   - RequestStatus: Current status of the request
//   - TotalRequestDurationInMs: Total duration in milliseconds
//   - TemplateStartDateTime: When the template processing started
//   - TemplateEndDateTime: When the template processing ended
type ProcessingDetailsProcessingDetailsRequestInstance struct {
	RequestStartDateTime     string `json:"requestStartDateTime,omitempty"`
	RequestEndDateTime       string `json:"requestEndDateTime,omitempty"`
	RequestStatus            string `json:"requestStatus,omitempty"`
	TotalRequestDurationInMs int32  `json:"totalRequestDurationInMs,omitempty"`
	TemplateStartDateTime    string `json:"templateStartDateTime,omitempty"`
	TemplateEndDateTime      string `json:"templateEndDateTime,omitempty"`
}

// ProcessingDetailsProcessingDetails represents the processing details structure.
// Fields:
//   - RequestInstance: Information about the request instance
type ProcessingDetailsProcessingDetails struct {
	RequestInstance *ProcessingDetailsProcessingDetailsRequestInstance `json:"requestInstance,omitempty"`
}

// ProcessingDetails represents the overall processing details.
// Fields:
//   - ProcessingDetails: Detailed processing information
type ProcessingDetails struct {
	ProcessingDetails *ProcessingDetailsProcessingDetails `json:"processingDetails,omitempty"`
}

// - - - - - - - - - - - - - - - - - -
// Legacy
//

// NetworkServiceDefinitionCredentials represents authentication credentials
// for a network service.
// Fields:
//   - Username: Username for authentication
//   - Password: Password for authentication
type NetworkServiceDefinitionCredentials struct {
	// Username for authentication
	Username string `json:"username"`
	// Password for authentication
	Password string `json:"password"`
}

// NetworkServiceDefinition represents the configuration for a network service.
// Fields:
//   - ServiceId: Unique identifier for the service
//   - ServiceType: Type of service (e.g., firewall, load_balancer)
//   - Endpoint: API endpoint for the service
//   - Credentials: Authentication credentials for the service
type NetworkServiceDefinition struct {
	// Unique identifier for the service.
	ServiceId string `json:"service_id"`
	// Type of service, e.g., \"firewall\", \"load_balancer\".
	ServiceType string `json:"service_type"`
	// API endpoint for interacting with the service.
	Endpoint string `json:"endpoint"`

	Credentials *NetworkServiceDefinitionCredentials `json:"credentials"`
}

// WorkflowRequest represents a request to start a workflow.
// Fields:
//   - WorkflowId: Unique identifier for the workflow
//   - Operation: Operation to perform on the services
//   - ServiceDefinitions: List of network service definitions
type WorkflowRequest struct {
	// Unique identifier for the workflow
	WorkflowId string `json:"workflow_id"`
	// Operation to perform on the services (e.g., \"create_network\", \"delete_network\").
	Operation string `json:"operation"`

	ServiceDefinitions []NetworkServiceDefinition `json:"service_definitions"`
}

// WorkflowResponse represents the response from starting a workflow.
// Fields:
//   - WorkflowId: Identifier of the workflow
//   - RunId: Identifier of the workflow run
//   - StatusId: Status identifier
//   - Result: Result of the workflow execution
type WorkflowResponse struct {
	WorkflowId string `json:"workflowId"`
	RunId      string `json:"runId"`
	StatusId   string `json:"statusId"`
	Result     string `json:"result"`
}
