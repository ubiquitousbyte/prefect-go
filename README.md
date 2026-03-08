# prefect-go

Go client SDK for Prefect, auto-generated from Prefect's OpenAPI specification using [oapi-codegen-exp](https://github.com/oapi-codegen/oapi-codegen-exp).

[![Go Reference](https://pkg.go.dev/badge/github.com/ubiquitousbyte/prefect-go.svg)](https://pkg.go.dev/github.com/ubiquitousbyte/prefect-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/ubiquitousbyte/prefect-go)](https://goreportcard.com/report/github.com/ubiquitousbyte/prefect-go)

## ⚠️ Experimental Status

This project uses **oapi-codegen-exp** to support OpenAPI 3.1 specifications. The experimental repository is under active development and includes this warning from the maintainers:

> "This is an experimental version that is not ready for production use. Do not use for anything important."

**Why experimental?** Prefect's API specification uses OpenAPI 3.1 features (like `exclusiveMinimum`/`exclusiveMaximum` as numbers) that are not supported by the stable oapi-codegen v2.x releases. The experimental version uses a different parser (`libopenapi`) with full OpenAPI 3.1+ support.

**Known Issues:**
- The code generator produces invalid `ApplyDefaults()` functions for certain union types (workaround applied in generated code)
- API is subject to change as the experimental branch evolves

For production applications, consider waiting for oapi-codegen v3.x stable release or use Prefect's Python SDK.

## Features

- **Type-safe** - Full Go type safety from OpenAPI specification
- **Auto-generated** - Reduces manual coding and stays in sync with Prefect API
- **Complete API coverage** - All Prefect REST API endpoints available (46,000+ lines of generated code)
- **OpenAPI 3.1 support** - Compatible with Prefect's modern API specification
- **Dual environment support** - Works with both Prefect Cloud and self-hosted servers
- **Minimal dependencies** - Uses standard `net/http` (Go 1.22+)
- **Easy authentication** - Helper functions for API keys and headers

## Installation

```bash
go get github.com/ubiquitousbyte/prefect-go
```

## Quick Start

### Self-hosted Prefect Server

```go
package main

import (
    "context"
    "log"

    "github.com/ubiquitousbyte/prefect-go"
)

func main() {
    // Create client
    client, err := prefect.NewSimpleClient("http://localhost:4200")
    if err != nil {
        log.Fatal(err)
    }

    // Check health
    resp, err := client.HealthCheckHealthGet(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Server status: %+v", resp)
}
```

### Prefect Cloud

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/ubiquitousbyte/prefect-go"
)

func main() {
    apiKey := os.Getenv("PREFECT_API_KEY")

    // Create authenticated client
    client, err := prefect.NewSimpleClient(
        "https://api.prefect.cloud",
        prefect.WithRequestEditorFn(
            prefect.WithAPIKey(apiKey),
        ),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Get server version info
    version, err := client.ServerVersionVersionGet(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Prefect version: %s", version)
}
```

## Authentication

### API Key (Prefect Cloud)

```go
client, err := prefect.NewSimpleClient(
    "https://api.prefect.cloud",
    prefect.WithRequestEditorFn(
        prefect.WithAPIKey("your-api-key"),
    ),
)
```

### Prefect Cloud with Account/Workspace

```go
client, err := prefect.NewSimpleClient(
    "https://api.prefect.cloud",
    prefect.WithRequestEditorFn(
        prefect.ChainRequestEditors(
            prefect.WithAPIKey("your-api-key"),
            prefect.WithAccountID("your-account-id"),
            prefect.WithWorkspaceID("your-workspace-id"),
        ),
    ),
)
```

### Custom Headers

```go
headers := map[string]string{
    "X-Custom-Header": "value",
}

client, err := prefect.NewSimpleClient(
    "https://api.prefect.cloud",
    prefect.WithRequestEditorFn(
        prefect.ChainRequestEditors(
            prefect.WithAPIKey("your-api-key"),
            prefect.WithCustomHeaders(headers),
        ),
    ),
)
```

## Examples

Complete working examples are available in the [`examples/`](examples/) directory:

- [`examples/cloud/`](examples/cloud/) - Prefect Cloud integration with authentication
- [`examples/selfhosted/`](examples/selfhosted/) - Basic usage with a local Prefect server

## API Coverage

This SDK is generated from Prefect 3.6.21's complete OpenAPI 3.1 specification and includes:

- **46,000+ lines of generated code** covering the full Prefect REST API
- Health checks and version information
- Flow management (create, read, update, delete, filter)
- Flow run operations and state management
- Deployment management
- Work pool and worker operations
- Block type and block document operations
- Task run tracking
- Automation and event management
- And all other Prefect REST API endpoints

The generated client includes 9 different client types for various use cases:
- `SimpleClient` - Basic client with automatic JSON handling (recommended)
- `Client` - Low-level client with manual response handling
- Additional specialized clients for specific needs

## Development

### Prerequisites

- Go 1.22 or later
- Python 3.x (for JSON validation in fetch script)

### Regenerating the Client

If you update the OpenAPI specification, regenerate the client:

```bash
# Install code generation tools
make install-tools

# Regenerate from current spec
make generate

# Or use go generate directly
go generate ./...
```

**Note:** Code generation uses `oapi-codegen-exp` (experimental). A workaround is applied to fix up 3 invalid lines in the generated `ApplyDefaults()` functions due to a known bug in the generator.

### Fetching Latest OpenAPI Spec

To fetch the OpenAPI spec from a running Prefect instance:

```bash
# From local server
make fetch-spec VERSION=3.6.21 API_URL=http://localhost:4200
```

### Building

```bash
# Build the library
make build

# Build examples
make examples

# Run tests
make test

# Run everything
make all
```

## Supported Prefect Versions

This SDK is generated from Prefect's OpenAPI specification. It should work with:

- Prefect 3.x (tested with 3.6.21)
- Both Prefect Cloud and self-hosted instances

To target a specific Prefect version, fetch that version's OpenAPI spec and regenerate the client.

## Why oapi-codegen?

Using `oapi-codegen` provides several benefits:

1. **Automatic updates** - Regenerate when Prefect API changes
2. **Type safety** - Full Go type checking for requests/responses
3. **Less maintenance** - No manual client code to maintain
4. **Complete coverage** - All API endpoints automatically included
5. **Standard patterns** - Uses proven Go HTTP client patterns
6. **OpenAPI 3.1 support** - Via experimental branch, compatible with modern API specs


## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

### How to Contribute

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and examples
5. Submit a pull request

## Resources

- [Prefect Documentation](https://docs.prefect.io/)
- [Prefect REST API Reference](https://docs.prefect.io/api-ref/rest-api/)
- [oapi-codegen Documentation](https://github.com/oapi-codegen/oapi-codegen)

## License

Apache 2.0

## Acknowledgments

- Built with [oapi-codegen-exp](https://github.com/oapi-codegen/oapi-codegen-exp)
- For use with [Prefect](https://github.com/PrefectHQ/prefect)
