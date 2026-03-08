package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ubiquitousbyte/prefect-go/prefect"
)

func main() {
	// Get server URL from environment or use default
	serverURL := os.Getenv("PREFECT_API_URL")
	if serverURL == "" {
		serverURL = "http://localhost:4200"
	}

	fmt.Printf("Connecting to self-hosted Prefect server at %s\n\n", serverURL)

	// Create client (no authentication for basic self-hosted setup)
	client, err := prefect.NewSimpleClient(serverURL)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Health check
	fmt.Println("Checking server health...")
	healthy, err := client.HealthCheckHealthGet(ctx)
	if err != nil {
		log.Fatalf("Health check failed: %v\n", err)
	}

	fmt.Printf("✓ Server is healthy: %v\n\n", healthy)

	// Get server version
	fmt.Println("Getting server version...")
	version, err := client.ReadVersionAdminVersionGet(ctx, &prefect.ReadVersionAdminVersionGetParams{})
	if err != nil {
		log.Printf("Failed to get version: %v\n", err)
	} else {
		fmt.Printf("✓ Server version: %s\n\n", version)
	}

	// List all flows
	fmt.Println("Listing all flows...")
	limit := 20
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
	fmt.Println("\nListing flow runs...")
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
