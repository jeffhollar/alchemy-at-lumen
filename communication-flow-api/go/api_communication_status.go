/*
 * Network Services Adapter API
 *
 * This API provides an adapter for managing network services using Temporal.io.
 *
 */
package swagger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ServiceInformation Service Information
type ServiceInformation struct {
	App struct {
		Name        string `json:"name"`
		Version     string `json:"version"`
		Description string `json:"description"`
		Attributes  struct {
			DisplayName string `json:"displayName"`
			Parent      string `json:"parent"`
			Type        string `json:"type"`
		} `json:"attributes"`
	} `json:"app"`
	Build struct {
		Version  string  `json:"version"`
		Number   string  `json:"number"`
		Artifact string  `json:"artifact"`
		Name     string  `json:"name"`
		Time     float64 `json:"time"`
		Group    string  `json:"group"`
	} `json:"build"`
}

func GetNetworkServiceHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == "OPTIONS" {
		return
	}

	// This is a simplistic example of checking the health of this service.
	// We should add more logic to validate the access points and health of the container.
	//
	defaultHealth := "{\"status\":\"UP\"}"
	_, _ = fmt.Fprintf(w, string(defaultHealth))
}

func GetNetworkServiceInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == "OPTIONS" {
		return
	}

	now := time.Now()

	serviceInformation := ServiceInformation{}
	serviceInformation.App.Name = "act-core-communication-flow"
	serviceInformation.App.Version = "1.0.0=000"
	serviceInformation.App.Description = "A Go-based service that provides an ACT core interface for communication services and Temporal.io workflows."
	serviceInformation.App.Attributes.DisplayName = "Act Core Communication Flow"
	serviceInformation.App.Attributes.Parent = "platform"
	serviceInformation.App.Attributes.Type = "platform"
	serviceInformation.Build.Version = "1.0.0-000"
	serviceInformation.Build.Number = "001"
	serviceInformation.Build.Artifact = "act-core-communication-flow"
	serviceInformation.Build.Name = "Act Core Communication Flow"
	serviceInformation.Build.Time = float64(now.UnixNano())
	serviceInformation.Build.Group = "com.lumen"

	outputJSON, err := json.Marshal(serviceInformation)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, _ = fmt.Fprintf(w, string(outputJSON))
}
