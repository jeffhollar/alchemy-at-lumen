/*
 * Network Services Adapter API
 *
 * This API provides an adapter for managing network services using Temporal.io.
 * This file defines the NetworkServiceDefinition model which represents
 * the configuration and credentials for a network service.
 *
 * The model is used to store and manage service-specific information
 * required for interacting with various network services.
 */
package swagger

// NetworkServiceDefinition represents the configuration for a network service
// including its identification, type, endpoint, and authentication credentials.
// Fields:
//   - ServiceId: Unique identifier for the service instance
//   - ServiceType: Category of the service (e.g., firewall, load_balancer)
//   - Endpoint: Base URL for the service's API
//   - Credentials: Authentication details for the service
type NetworkServiceDefinition struct {
	// Unique identifier for the service.
	ServiceId string `json:"service_id"`
	// Type of service, e.g., \"firewall\", \"load_balancer\".
	ServiceType string `json:"service_type"`
	// API endpoint for interacting with the service.
	Endpoint string `json:"endpoint"`

	Credentials *NetworkServiceDefinitionCredentials `json:"credentials"`
}
