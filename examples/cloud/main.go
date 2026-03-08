package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ubiquitousbyte/prefect-go/prefect"
)

func main() {
	// Get credentials from environment variables
	apiKey := os.Getenv("PREFECT_API_KEY")
	if apiKey == "" {
		log.Fatal("PREFECT_API_KEY environment variable is required")
	}

	accountID := os.Getenv("PREFECT_ACCOUNT_ID")
	workspaceID := os.Getenv("PREFECT_WORKSPACE_ID")

	// Create client with authentication
	requestEditor := prefect.WithAPIKey(apiKey)
	if accountID != "" && workspaceID != "" {
		// Use all credentials if available
		requestEditor = prefect.ChainRequestEditors(
			requestEditor,
			prefect.WithAccountID(accountID),
			prefect.WithWorkspaceID(workspaceID),
		)
	}

	client, err := prefect.NewSimpleClient(
		"https://api.prefect.cloud",
		prefect.WithRequestEditorFn(requestEditor),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Health check
	fmt.Println("Checking Prefect Cloud connection...")
	healthy, err := client.HealthCheckHealthGet(ctx)
	if err != nil {
		log.Fatalf("Health check failed: %v", err)
	}

	fmt.Printf("✓ Connected to Prefect Cloud: %v\n\n", healthy)

	// List flows
	fmt.Println("Fetching flows from Prefect Cloud...")
	limit := 10
	offset := 0
	flows, err := client.ReadFlowsFlowsFilterPost(
		ctx,
		&prefect.ReadFlowsFlowsFilterPostParams{},
		prefect.BodyReadFlowsFlowsFilterPost{
			Limit:  &limit,
			Offset: &offset,
		},
	)
	if err != nil {
		log.Fatalf("Failed to list flows: %v", err)
	}

	fmt.Printf("Found %d flows:\n", len(flows))
	for i, flow := range flows {
		tags := ""
		if len(flow.Tags) > 0 {
			tags = fmt.Sprintf(" [tags: %v]", flow.Tags)
		}
		fmt.Printf("%d. %s (ID: %s)%s\n", i+1, flow.Name, flow.ID, tags)
	}

	// List flow runs
	fmt.Println("\nFetching recent flow runs...")
	runs, err := client.ReadFlowRunsFlowRunsFilterPost(
		ctx,
		&prefect.ReadFlowRunsFlowRunsFilterPostParams{},
		prefect.BodyReadFlowRunsFlowRunsFilterPost{
			Limit: &limit,
		},
	)
	if err != nil {
		log.Fatalf("Failed to list flow runs: %v", err)
	}

	fmt.Printf("Found %d flow runs:\n", len(runs))
	for i, run := range runs {
		name := "unnamed"
		if run.Name != nil {
			name = *run.Name
		}
		stateType := "unknown"
		if st, err := run.StateType.Get(); err == nil {
			stateType = string(st)
		}
		stateName := "unknown"
		if sn, err := run.StateName.Get(); err == nil {
			stateName = sn
		}
		fmt.Printf("%d. %s - State: %s/%s\n",
			i+1,
			name,
			stateType,
			stateName,
		)
	}
}
