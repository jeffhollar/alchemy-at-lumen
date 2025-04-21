/*
 * Network Services Adapter API
 *
 * This API provides an adapter for managing network services using Temporal.io.
 * It handles workflow execution, processing details, and various callback endpoints
 * for different services (Lava, Rubicon, Vino).
 *
 * The API is designed to be a bridge between HTTP requests and Temporal workflows,
 * managing the lifecycle of network service operations.
 */
package swagger

import (
	"communication-flow-api/internal/adapters"
	"communication-flow-api/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.temporal.io/sdk/client"
	"net/http"
)

// GetActRequest handles HTTP GET requests for retrieving ACT (Automated Communication Task) requests.
// It validates the request, checks authorization, and initiates a Temporal workflow
// for processing the request.
// Parameters:
//   - w: HTTP response writer
//   - r: HTTP request containing the identifierID in the URL path
//
// Returns:
//   - JSON response with workflow execution details or error message
func GetActRequest(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	vars := mux.Vars(r)
	identifierID := vars["identifierID"]

	// Make sure the caller is authorized
	//
	//userToken := r.Header.Get("accessToken")
	//if authorized, err := isAuthorized(userToken); err != nil {
	//	errorJSON(w, err, http.StatusForbidden)
	//	return
	//} else if authorized {
	//	infoMsg("***** Authorized \n( User Token %s)\n", userToken)
	//} else {
	//	errorJSON(w, errors.New("Not Authorized"), http.StatusForbidden)
	//	return
	//}
	//
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	var request models.ActRequest
	request.IdentifierId = identifierID

	// Get Temporal client from context
	temporalClient, ok := r.Context().Value("temporalClient").(client.Client)
	if !ok {
		errorJSON(w, fmt.Errorf("temporal client not found in context"), http.StatusInternalServerError)
		return
	}

	// Create adapter and start workflow
	adapter := adapters.NewTemporalClient(temporalClient)
	infoMsg("\nStartWorkflow Invoked from POST\n *********************\n")

	response, err := adapter.StartWorkflow(r.Context(), request)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	infoMsg("\nStartWorkflow completed", "response", response)
	infoMsg("\n **** **** **** ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GetProcessingDetails is a placeholder endpoint for retrieving the processing details
// of a workflow execution. This currently returns a "NOT IMPLEMENTED" response.
// Parameters:
//   - w: HTTP response writer
//   - r: HTTP request
func GetProcessingDetails(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	defaultResponse := "{\"status\":\"NOT IMPLEMENTED\"}"
	_, _ = fmt.Fprintf(w, string(defaultResponse))
}

// PostActRequest is a placeholder endpoint for submitting new ACT requests.
// This currently returns a "NOT IMPLEMENTED" response.
// Parameters:
//   - w: HTTP response writer
//   - r: HTTP request
func PostActRequest(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	defaultResponse := "{\"status\":\"NOT IMPLEMENTED\"}"
	_, _ = fmt.Fprintf(w, string(defaultResponse))
}

// PostLavaContinueRequest is a placeholder endpoint for handling Lava service
// continuation requests. Currently returns a "NOT IMPLEMENTED" response.
// Parameters:
//   - w: HTTP response writer
//   - r: HTTP request
func PostLavaContinueRequest(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	defaultResponse := "{\"status\":\"NOT IMPLEMENTED\"}"
	_, _ = fmt.Fprintf(w, string(defaultResponse))
}

// PostRubiconCallback is a placeholder endpoint for handling Rubicon service
// callback notifications. Currently returns a "NOT IMPLEMENTED" response.
// Parameters:
//   - w: HTTP response writer
//   - r: HTTP request
func PostRubiconCallback(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	defaultResponse := "{\"status\":\"NOT IMPLEMENTED\"}"
	_, _ = fmt.Fprintf(w, string(defaultResponse))
}

// PostVinoContinueRequest is a placeholder endpoint for handling Vino service
// continuation requests. Currently returns a "NOT IMPLEMENTED" response.
// Parameters:
//   - w: HTTP response writer
//   - r: HTTP request
func PostVinoContinueRequest(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	defaultResponse := "{\"status\":\"NOT IMPLEMENTED\"}"
	_, _ = fmt.Fprintf(w, string(defaultResponse))
}

// errorJSON is a helper function that writes an error response in JSON format.
// Parameters:
//   - w: HTTP response writer
//   - err: Error to be included in the response
//   - statusCode: HTTP status code to be set in the response
func errorJSON(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
