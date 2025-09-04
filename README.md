# Top Ten

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
cd top-ten

# Build the application
just build

# Or use Go directly
go build -o bin/top-ten ./cmd/top-ten
```

## Usage

### Display a random Top 10 list

```bash
./bin/top-ten random
```

### Show help

```bash
./bin/top-ten --help
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
