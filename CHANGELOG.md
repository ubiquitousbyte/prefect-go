# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2026-03-08

### ⚠️ Experimental Release

This is the initial experimental release of prefect-go. This SDK uses **oapi-codegen-exp** (experimental) to support OpenAPI 3.1 specifications. It is **not ready for production use**.

**Important warnings:**
- Built with experimental oapi-codegen branch (not stable)
- Code generator has known bugs requiring manual workarounds
- API may change as the experimental branch evolves
- For production, use Prefect's official Python SDK

### Added

- Initial Go client SDK for Prefect REST API
- Full API coverage auto-generated from Prefect 3.6.21 OpenAPI 3.1 specification
- Support for both Prefect Cloud and self-hosted servers
- Type-safe client with full Go type checking
- Authentication helpers for API keys and custom headers
- Examples for common use cases (cloud and self-hosted)
- Comprehensive documentation and README

### Technical Details

- Generated using [oapi-codegen-exp](https://github.com/oapi-codegen/oapi-codegen-exp) for OpenAPI 3.1 support
- 46,000+ lines of auto-generated client code
- Requires Go 1.22+
- Minimal dependencies (standard `net/http`)

### Known Issues

- Code generator produces invalid `ApplyDefaults()` functions for 3 union types (workaround applied in generated code)
- Cannot exclude these invalid functions via config (inline `anyOf` types, not separate schemas)
- Manual fix required after each code generation: Fix up problematic lines in `prefect/client.gen.go`

### Supported Prefect Versions

- Prefect 3.x (tested with 3.6.21)
- Compatible with Prefect Cloud and self-hosted instances

[Unreleased]: https://github.com/ubiquitousbyte/prefect-go/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/ubiquitousbyte/prefect-go/releases/tag/v0.1.0
