// Package prefect provides a Go client for the Prefect REST API.
//
// This client is auto-generated from Prefect's OpenAPI 3.1 specification using
// oapi-codegen-exp (experimental). It provides type-safe access to all Prefect
// API endpoints for both Prefect Cloud and self-hosted Prefect servers.
//
// ⚠️ EXPERIMENTAL: This package uses oapi-codegen-exp which is under active
// development. The maintainers note: "Do not use for anything important."
// See README.md for details on known issues and limitations.
//
// # Quick Start
//
// For self-hosted Prefect server:
//
//	client, err := prefect.NewSimpleClient("http://localhost:4200")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Check server health
//	resp, err := client.HealthCheckHealthGet(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	log.Printf("Server status: %+v", resp)
//
// For Prefect Cloud with authentication:
//
//	apiKey := os.Getenv("PREFECT_API_KEY")
//	client, err := prefect.NewSimpleClient(
//	    "https://api.prefect.cloud",
//	    prefect.WithRequestEditorFn(
//	        prefect.WithAPIKey(apiKey),
//	    ),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Get server version
//	version, err := client.ReadVersionAdminVersionGet(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	log.Printf("Prefect version: %s", version)
//
// # Client Types
//
// This package provides two main client types:
//
//   - SimpleClient: High-level client with automatic JSON marshaling (recommended)
//   - Client: Low-level client that returns *http.Response for manual handling
//
// Most users should use NewSimpleClient() which provides a cleaner API with
// typed responses and automatic error handling.
//
// # Authentication
//
// The package provides several helper functions for authentication:
//
//   - WithAPIKey: For Prefect Cloud API key authentication (Bearer token)
//   - WithAccountID: For setting Prefect Cloud account ID header
//   - WithWorkspaceID: For setting Prefect Cloud workspace ID header
//   - WithCustomHeaders: For adding custom headers to all requests
//   - ChainRequestEditors: For combining multiple request modifications
//
// Example with multiple headers:
//
//	client, err := prefect.NewSimpleClient(
//	    "https://api.prefect.cloud",
//	    prefect.WithRequestEditorFn(
//	        prefect.ChainRequestEditors(
//	            prefect.WithAPIKey("your-api-key"),
//	            prefect.WithAccountID("your-account-id"),
//	            prefect.WithWorkspaceID("your-workspace-id"),
//	        ),
//	    ),
//	)
//
// # Working with Nullable Types
//
// The generated code uses Nullable[T] for fields that can be null, unspecified,
// or have a value. This is more precise than Go's pointer types.
//
// To work with Nullable fields:
//
//	// Check if a value is set
//	if stateType, err := run.StateType.Get(); err == nil {
//	    fmt.Printf("State: %s\n", stateType)
//	}
//
//	// Create a Nullable with a value
//	nullable := prefect.NewNullableWithValue("my-value")
//
//	// Create an explicit null
//	nullValue := prefect.NewNullNullable[string]()
//
// See the generated code for more Nullable helper methods.
//
// # Generated Code
//
// Most of this package is auto-generated from Prefect's
// OpenAPI 3.1 specification using oapi-codegen-exp. The generated code includes:
//
//   - Type definitions for all API models
//   - Client methods for all API endpoints
//   - Request and response types
//   - Helper types like Nullable[T] for null-safe handling
//
// To regenerate the client after updating the OpenAPI spec:
//
//	make generate
//	# or directly:
//	go generate ./...
//
// # Known Issues
//
// Due to bugs in oapi-codegen-exp, three ApplyDefaults() functions are
// commented out in the generated code after each generation:
//
//   - FlowsSettings.DefaultRetryDelaySeconds
//   - ServerServicesDBVacuumSettings.Enabled
//   - TasksSettings.DefaultRetryDelaySeconds
//
// These fields remain accessible but their default values won't be applied
// automatically. This is a known limitation of the experimental code generator.
// The workaround is applied manually after generation - see the generated
// client.gen.go file for details.
//
// # Examples
//
// See the examples/ directory for complete working examples:
//
//   - examples/cloud/ - Prefect Cloud integration with authentication
//   - examples/selfhosted/ - Self-hosted server with filtering and operations
//
// Run examples with:
//
//	cd examples/cloud && go run main.go
//
// # API Reference
//
// For detailed Prefect REST API documentation:
// https://docs.prefect.io/api-ref/rest-api/
//
// For this package's Go documentation:
// https://pkg.go.dev/github.com/ubiquitousbyte/prefect-go
//
// For the OpenAPI specification used:
// Generated from Prefect 3.6.21 OpenAPI 3.1 specification
package prefect
