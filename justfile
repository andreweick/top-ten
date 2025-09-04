# Run tests
test:
    go test -v ./...

# Run linting
lint:
    golangci-lint run

# Format code
fmt:
    go fmt ./...

# Build for current platform
build:
    go build -o bin/top-ten ./cmd/top-ten

# Run the application (depends on build)
run *args: build
    ./bin/top-ten {{args}}

# Clean build artifacts
clean:
    rm -rf bin/

# Run all checks (test, lint, format)
check: fmt lint test

# Install golangci-lint if not present
install-tools:
    @which golangci-lint > /dev/null || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Decrypt the age file to JSON (for local development/testing)
decrypt:
    #!/usr/bin/env bash
    if [ -z "$AGE_ENCRYPTION_PASSWORD" ]; then
        echo "AGE_ENCRYPTION_PASSWORD not set. Please enter password:"
        read -s password
        export AGE_ENCRYPTION_PASSWORD="$password"
    fi
    if [ ! -f "data/all-top-10.json.age" ]; then
        echo "Error: data/all-top-10.json.age not found"
        exit 1
    fi
    echo "$AGE_ENCRYPTION_PASSWORD" | age -d -i /dev/stdin data/all-top-10.json.age > data/all-top-10.json
    echo "Successfully decrypted data/all-top-10.json"

# Encrypt the JSON file to age (for updating data)
encrypt:
    #!/usr/bin/env bash
    if [ -z "$AGE_ENCRYPTION_PASSWORD" ]; then
        echo "AGE_ENCRYPTION_PASSWORD not set. Please enter password:"
        read -s password
        export AGE_ENCRYPTION_PASSWORD="$password"
    fi
    if [ ! -f "data/all-top-10.json" ]; then
        echo "Error: data/all-top-10.json not found"
        exit 1
    fi
    echo "$AGE_ENCRYPTION_PASSWORD" | age -p -a -i /dev/stdin data/all-top-10.json > data/all-top-10.json.age
    echo "Successfully encrypted data/all-top-10.json.age"

# Run the application with password prompt if needed
run-with-password: build
    #!/usr/bin/env bash
    if [ -z "$AGE_ENCRYPTION_PASSWORD" ]; then
        echo "AGE_ENCRYPTION_PASSWORD not set. Please enter password:"
        read -s password
        export AGE_ENCRYPTION_PASSWORD="$password"
    fi
    ./bin/top-ten random