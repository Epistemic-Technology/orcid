# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go client library for interacting with the ORCID Public API v3.0. The library provides a comprehensive client with retry logic, rate limiting, and support for both JSON and XML response formats.

## Commands

### Testing
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./orcid/...
```

### Building
```bash
# Build the orcid-search CLI tool
go build -o orcid-search ./cmd/orcid-search

# Build all commands
go build ./...
```

### Module Management
```bash
# Download dependencies
go mod download

# Tidy dependencies
go mod tidy

# Verify dependencies
go mod verify
```

## Architecture

### Package Structure
- **orcid/** - Main client library package containing:
  - `client.go` - Core HTTP client with retry logic, rate limiting, and configuration options
  - `models.go` - ORCID API response models and data structures
  - `search.go` - Search functionality including query builder and iterator pattern
  - Tests use standard Go testing with httptest for mocking API responses

- **cmd/orcid-search/** - CLI tool for searching and retrieving ORCID records

### Key Design Patterns
1. **Functional Options Pattern**: Client configuration uses `WithXxx()` option functions for clean initialization
2. **Iterator Pattern**: Search results implement an iterator for efficient pagination handling
3. **Builder Pattern**: Search queries use a fluent builder interface for constructing complex queries
4. **Context-Aware**: All API methods accept context.Context for cancellation and timeout support

### API Endpoints
The client implements methods for all ORCID v3.0 public API endpoints:
- Record retrieval (GetRecord, GetPerson, GetWorks, etc.)
- Search functionality with query builder
- Support for both production (pub.orcid.org) and sandbox (pub.sandbox.orcid.org) environments

### Testing Approach
- Unit tests use httptest.Server to mock ORCID API responses
- Tests are located alongside implementation files (*_test.go)
- Example tests demonstrate client usage patterns