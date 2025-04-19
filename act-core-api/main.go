/*
 * ACT Core API
 *
 * API documentation for ACT Core system
 *
 */
package main

import (
	sw "act-core-api/go"
	"act-core-api/internal/activities"
	"act-core-api/internal/workflows"
	"context"
	"crypto/tls"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Config represents the application configuration settings.
// Fields:
//   - TemporalHost: Host and port of the Temporal server
//   - TemporalNamespace: Namespace for Temporal workflows
//   - TemporalTaskQueue: Queue name for Temporal tasks
//   - CertFile: Path to SSL certificate file
//   - KeyFile: Path to SSL private key file
//   - ServerPort: Port for the HTTP server
type Config struct {
	TemporalHost      string
	TemporalNamespace string
	TemporalTaskQueue string
	CertFile          string
	KeyFile           string
	ServerPort        string
}

var config = Config{
	TemporalHost:      getEnvOrDefault("TEMPORAL_HOST", "192.168.1.233:7233"),
	TemporalNamespace: getEnvOrDefault("TEMPORAL_NAMESPACE", "lumen-usecases"),
	TemporalTaskQueue: getEnvOrDefault("TEMPORAL_TASK_QUEUE", "act-core-api-task-queue"),
	CertFile:          getEnvOrDefault("SSL_CERT_FILE", "key.pem"),
	KeyFile:           getEnvOrDefault("SSL_KEY_FILE", "server.key"),
	ServerPort:        getEnvOrDefault("SERVER_PORT", ":8501"),
}

// temporalClientMiddleware creates a middleware that injects the Temporal client
// into the request context for use by handlers.
// Parameters:
//   - temporalClient: The Temporal.io client instance
//
// Returns:
//   - func(http.Handler) http.Handler: Middleware function
func temporalClientMiddleware(temporalClient client.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "temporalClient", temporalClient)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// startTemporalWorker initializes the Temporal client and worker, registering
// workflows and activities.
// Returns:
//   - client.Client: The Temporal client instance
//   - worker.Worker: The Temporal worker instance
//   - error: Any error that occurred during initialization
func startTemporalWorker() (client.Client, worker.Worker, error) {
	temporalClient, err := client.Dial(client.Options{
		HostPort:  config.TemporalHost,
		Namespace: config.TemporalNamespace,
	})
	if err != nil {
		errorMsg("Failed to initialize Temporal client: %v\n", err)
		return nil, nil, err
	}

	workerClient := worker.New(temporalClient, config.TemporalTaskQueue, worker.Options{})

	// Register workflows and activities
	infoMsg("Registering workflows and activities...")
	workerClient.RegisterWorkflow(workflows.GetActRequestWorkflow)
	workerClient.RegisterActivity(activities.ExecuteActRequestOperation)

	return temporalClient, workerClient, nil
}

func main() {
	log.Printf("Network Service Adapter starting...")

	// Verify SSL certificates exist
	if !fileExists(config.CertFile) || !fileExists(config.KeyFile) {
		errorMsg("SSL certificate files not found. Please ensure %s and %s exist", config.CertFile, config.KeyFile)
		os.Exit(1)
	}

	// Initialize Temporal client and worker
	temporalClient, workerClient, err := startTemporalWorker()
	if err != nil {
		errorMsg("Failed to initialize Temporal client: %v", err)
		os.Exit(1)
	}
	defer temporalClient.Close()

	// Start the worker in a separate goroutine
	workerErr := make(chan error, 1)
	go func() {
		infoMsg("Starting Temporal worker...")
		if err := workerClient.Run(worker.InterruptCh()); err != nil {
			errorMsg("Worker failed to start: %v", err)
			workerErr <- err
		}
	}()

	// Initialize router and middleware
	router := sw.NewRouter()
	temporalRouter := temporalClientMiddleware(temporalClient)(router)

	// Initialize Swagger UI
	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui/")))
	router.PathPrefix("/swaggerui/").Handler(sh)

	// Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		CurvePreferences: []tls.CurveID{
			tls.X25519,
			tls.CurveP256,
			tls.CurveP384,
		},
		InsecureSkipVerify: true, // Skip certificate verification for development
	}

	// Create a HTTPS server
	server := &http.Server{
		Addr:      config.ServerPort,
		Handler:   temporalRouter,
		TLSConfig: tlsConfig,
	}

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start a server in a separate goroutine
	serverErr := make(chan error, 1)
	go func() {
		infoMsg("Starting HTTPS server on port %s", config.ServerPort)
		serverErr <- server.ListenAndServeTLS(config.CertFile, config.KeyFile)
	}()

	// Wait for either server error, worker error, or shutdown signal
	select {
	case err := <-serverErr:
		errorMsg("Server error: %v", err)
	case err := <-workerErr:
		errorMsg("Worker error: %v", err)
	case sig := <-sigChan:
		infoMsg("Received signal %v, shutting down...", sig)
	}

	// Attempt a graceful shutdown
	infoMsg("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		errorMsg("Error during server shutdown: %v", err)
	}

}

// getEnvOrDefault retrieves an environment variable or returns a default value
// if the variable is not set.
// Parameters:
//   - key: The environment variable name
//   - defaultValue: The default value to return if the variable is not set
//
// Returns:
//   - string: The environment variable value or default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// fileExists checks if a file exists at the specified path.
// Parameters:
//   - filename: The path to check
//
// Returns:
//   - bool: True if the file exists, false otherwise
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// debug logs a debug-level message with the actcoreapi prefix.
// Parameters:
//   - template: Message template with format specifiers
//   - values: Values to be formatted into the template
func debug(template string, values ...interface{}) {
	log.Printf("[actcoreapi][debug] "+template+"\n", values...)
}

// infoMsg logs an info-level message with the actcoreapi prefix.
// Parameters:
//   - template: Message template with format specifiers
//   - values: Values to be formatted into the template
func infoMsg(template string, values ...interface{}) {
	log.Printf("[actcoreapi][info] "+template+"\n", values...)
}

// errorMsg logs an error-level message with the actcoreapi prefix.
// Parameters:
//   - template: Message template with format specifiers
//   - values: Values to be formatted into the template
func errorMsg(template string, values ...interface{}) {
	log.Printf("[actcoreapi][error] "+template+"\n", values...)
}
