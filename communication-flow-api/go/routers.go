/*
 * Network Services Adapter API
 *
 * This API provides an adapter for managing network services using Temporal.io.
 * This file defines the routing configuration for the API endpoints and includes
 * utility functions for logging and authorization.
 *
 * The router configuration maps HTTP endpoints to their corresponding handler functions
 * and includes support for CORS through OPTIONS requests.
 */
package swagger

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Route represents a single HTTP route configuration
// Fields:
//   - Name: The name of the route
//   - Method: HTTP method (GET, POST, etc.)
//   - Pattern: URL pattern with path parameters
//   - HandlerFunc: The function that handles the request
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of Route configurations
type Routes []Route

// NewRouter creates and configures a new Gorilla Mux router with all defined routes
// Returns:
//   - *mux.Router: Configured router instance with all routes and middleware
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// Index is the root endpoint handler that returns a simple welcome message
// Parameters:
//   - w: HTTP response writer
//   - r: HTTP request
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Act Core Communication Flow")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"GetNetworkServiceHealth",
		strings.ToUpper("Options"),
		"/Network/v1/Provisioning/health",
		GetNetworkServiceHealth,
	},
	Route{
		"GetNetworkServiceHealth",
		strings.ToUpper("Get"),
		"/Network/v1/Provisioning/health",
		GetNetworkServiceHealth,
	},

	Route{
		"GetNetworkServiceInfo",
		strings.ToUpper("Options"),
		"/Network/v1/Provisioning/info",
		GetNetworkServiceInfo,
	},
	Route{
		"GetNetworkServiceInfo",
		strings.ToUpper("Get"),
		"/Network/v1/Provisioning/info",
		GetNetworkServiceInfo,
	},

	Route{
		"GetActRequest",
		strings.ToUpper("Options"),
		"/Network/v1/Provisioning/actRequest/{identifierID}",
		GetActRequest,
	},
	Route{
		"GetActRequest",
		strings.ToUpper("Get"),
		"/Network/v1/Provisioning/actRequest/{identifierID}",
		GetActRequest,
	},

	Route{
		"GetProcessingDetails",
		strings.ToUpper("Options"),
		"/Network/v1/Provisioning/actRequest/{identifierID}/processingDetails",
		GetProcessingDetails,
	},
	Route{
		"GetProcessingDetails",
		strings.ToUpper("Get"),
		"/Network/v1/Provisioning/actRequest/{identifierID}/processingDetails",
		GetProcessingDetails,
	},

	Route{
		"PostActRequest",
		strings.ToUpper("Options"),
		"/Network/v1/Provisioning/actRequest",
		PostActRequest,
	},
	Route{
		"PostActRequest",
		strings.ToUpper("Post"),
		"/Network/v1/Provisioning/actRequest",
		PostActRequest,
	},

	Route{
		"PostLavaContinueRequest",
		strings.ToUpper("Options"),
		"/Network/v1/Provisioning/lava/{identifier_id}",
		PostLavaContinueRequest,
	},
	Route{
		"PostLavaContinueRequest",
		strings.ToUpper("Post"),
		"/Network/v1/Provisioning/lava/{identifier_id}",
		PostLavaContinueRequest,
	},

	Route{
		"PostRubiconCallback",
		strings.ToUpper("Options"),
		"/Network/v1/Provisioning/rubicon/callback/{requestId}",
		PostRubiconCallback,
	},
	Route{
		"PostRubiconCallback",
		strings.ToUpper("Post"),
		"/Network/v1/Provisioning/rubicon/callback/{requestId}",
		PostRubiconCallback,
	},

	Route{
		"PostVinoContinueRequest",
		strings.ToUpper("Options"),
		"/Network/v1/Provisioning/vino/{identifier_id}",
		PostVinoContinueRequest,
	},
	Route{
		"PostVinoContinueRequest",
		strings.ToUpper("Post"),
		"/Network/v1/Provisioning/vino/{identifier_id}",
		PostVinoContinueRequest,
	},

	//Route{
	//	"WorkflowsPost",
	//	strings.ToUpper("Options"),
	//	"/netadapter/api/v1/workflows",
	//	WorkflowsPost,
	//},
	//Route{
	//	"WorkflowsPost",
	//	strings.ToUpper("Post"),
	//	"/netadapter/api/v1/workflows",
	//	WorkflowsPost,
	//},
}

// debug logs debug-level messages with the act-core prefix
// Parameters:
//   - template: Message template with format specifiers
//   - values: Values to be formatted into the template
func debug(template string, values ...interface{}) {
	log.Printf("[act-core][info] "+template+"\n", values...)
}

// infoMsg logs info-level messages with the act-core prefix
// Parameters:
//   - template: Message template with format specifiers
//   - values: Values to be formatted into the template
func infoMsg(template string, values ...interface{}) {
	log.Printf("[act-core][info] "+template+"\n", values...)
}

// errorMsg logs error-level messages with the act-core prefix
// Parameters:
//   - template: Message template with format specifiers
//   - values: Values to be formatted into the template
func errorMsg(template string, values ...interface{}) {
	log.Printf("[act-core][failure] "+template+"\n", values...)
}

// isAuthorized checks if a user token is valid and authorized
// Parameters:
//   - user_token: The user's authentication token
// Returns:
//   - bool: Whether the user is authorized
//   - error: Any error that occurred during authorization
func isAuthorized(user_token string) (bool, error) {

	// Validate User Token to determine if authorized
	// TBA
	// infoMsg("isAuthorized: ", ollama_token)

	// Add logic to validate token
	return true, nil
}
