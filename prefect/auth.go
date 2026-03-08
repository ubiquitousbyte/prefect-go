package prefect

import (
	"context"
	"net/http"
)

// WithAPIKey adds Prefect Cloud API key authentication to requests.
// This should be used when connecting to Prefect Cloud.
//
// Example:
//
//	client, err := prefect.NewSimpleClient(
//	    "https://api.prefect.cloud",
//	    prefect.WithRequestEditorFn(prefect.WithAPIKey("your-api-key")),
//	)
func WithAPIKey(apiKey string) RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+apiKey)
		return nil
	}
}

// WithAccountID sets the Prefect Cloud account ID header.
// This may be required for certain Prefect Cloud operations.
//
// Example:
//
//	client, err := prefect.NewSimpleClient(
//	    "https://api.prefect.cloud",
//	    prefect.WithRequestEditorFn(
//	        prefect.ChainRequestEditors(
//	            prefect.WithAPIKey("your-api-key"),
//	            prefect.WithAccountID("your-account-id"),
//	        ),
//	    ),
//	)
func WithAccountID(accountID string) RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set("X-PREFECT-ACCOUNT-ID", accountID)
		return nil
	}
}

// WithWorkspaceID sets the Prefect Cloud workspace ID header.
// This may be required for certain Prefect Cloud operations.
func WithWorkspaceID(workspaceID string) RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Set("X-PREFECT-WORKSPACE-ID", workspaceID)
		return nil
	}
}

// ChainRequestEditors combines multiple request editors into a single editor.
// This is useful when you need to apply multiple headers or modifications to requests.
//
// Example:
//
//	editor := prefect.ChainRequestEditors(
//	    prefect.WithAPIKey("your-api-key"),
//	    prefect.WithAccountID("your-account-id"),
//	    prefect.WithWorkspaceID("your-workspace-id"),
//	)
//	client, err := prefect.NewSimpleClient(
//	    "https://api.prefect.cloud",
//	    prefect.WithRequestEditorFn(editor),
//	)
func ChainRequestEditors(editors ...RequestEditorFn) RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		for _, editor := range editors {
			if err := editor(ctx, req); err != nil {
				return err
			}
		}
		return nil
	}
}

// WithCustomHeaders adds custom headers to all requests.
// This is useful for adding custom authentication or tracing headers.
func WithCustomHeaders(headers map[string]string) RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
		return nil
	}
}
