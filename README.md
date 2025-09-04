# Top 10

A command-line Go program that displays random David Letterman Top 10 lists from an embedded JSON collection.

## Features

- 1,857 classic Top 10 lists from Late Night with David Letterman
- Embedded JSON data (no external dependencies at runtime)
- Cryptographically secure random selection
- Clean command-line interface

## Installation

### Build from source

```bash
# Clone or download the repository
cd top10

# Build the application
just build

# Or use Go directly
go build -o bin/top10 ./cmd/top10
```

## Usage

### Display a random Top 10 list

```bash
./bin/top10 random
```

### Show help

```bash
./bin/top10 --help
```

## Development

This project uses `just` for task automation. Common commands:

```bash
# Run tests
just test

# Run formatting and linting
just fmt
just lint

# Build the application
just build

# Run the application with arguments
just run random

# Run all checks
just check

# Clean build artifacts
just clean
```
