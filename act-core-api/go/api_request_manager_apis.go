/*
 * ACT Core API
 *
 * API documentation for ACT Core system
 *
 */
package swagger

import (
	"act-core-api/internal/adapters"
	"act-core-api/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.temporal.io/sdk/client"
	"net/http"
)

func ActRequestIdentifierIDGet(w http.ResponseWriter, r *http.Request) {
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

	// THIS LOGIC SHOULD INVOKE THE NEXUS ENDPOINT TO GET THE ACT REQUEST
	//

	var request models.ActRequest
	request.ActivationTransactionId = identifierID

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

	//
	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
}

func ActRequestIdentifierIDYangGet(w http.ResponseWriter, r *http.Request) {
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

func EditDataRequestIDPut(w http.ResponseWriter, r *http.Request) {
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

func LavaIdentifierIdPost(w http.ResponseWriter, r *http.Request) {
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

func RequestGet(w http.ResponseWriter, r *http.Request) {
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

func RequestPost(w http.ResponseWriter, r *http.Request) {
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

func RequestRequestIdGet(w http.ResponseWriter, r *http.Request) {
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

func RequestRequestIdPut(w http.ResponseWriter, r *http.Request) {
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

func RestartActivationRequestIDPut(w http.ResponseWriter, r *http.Request) {
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

func RestartDataGatheringRequestIDPut(w http.ResponseWriter, r *http.Request) {
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

func VinoIdentifierIdPost(w http.ResponseWriter, r *http.Request) {
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
