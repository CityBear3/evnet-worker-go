# Event Worker

## Overview

This repository is a Proof of Concept (POC) project for exploring and evaluating implementation approaches for an Event Worker system. The code and architecture presented here are intended for learning and verification purposes only and are not meant to be used directly in production environments.

## Disclaimer

**This is a Proof of Concept (POC). It is not recommended for use in production environments.**

This project is experimental and may not implement best practices, comprehensive error handling, or security measures. For actual application development, please use this code as a reference while making appropriate improvements suitable for production environments.

## Project Structure

- `cmd/`: Application entry points
- `internal/`: Internal packages
- `terraform/`: Infrastructure configuration

## Development

### Prerequisites

- Go 1.24 or later
- Docker and Docker Compose

### Setup

1. Start the PubSub emulator:

```bash
# Set the project ID for the PubSub emulator
export PUBSUB_PROJECT_ID=local-project

# Start the PubSub emulator using Docker Compose
docker-compose up -d
```

2. Set up the PubSub topics and subscriptions:

```bash
# Run the setup-pubsub target from the Makefile
make setup-pubsub
```

### Build

To build the applications:

```bash
# Build the event worker
go build -o bin/event_worker cmd/event_worker/main.go

# Build the event publisher
go build -o bin/event_publish cmd/event_publish/main.go

# Build all applications
go build ./...
```

### Run

To run the applications:

```bash
# Run the event worker
go run cmd/event_worker/main.go

# Run the event publisher
go run cmd/event_publish/main.go
```

### Test

To run the tests:

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test -v ./internal/core/env
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
