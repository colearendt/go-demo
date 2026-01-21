# Start hot-reload development server
dev:
    air

# Build without embedding (reads from disk)
build:
    go build -o go-demo

# Build with embedded files
build-embed:
    VERSION=$(git describe --tags --always --dirty) && \
    go build -tags embed -ldflags "-X main.version=${VERSION}" -o go-demo

# Run the built binary
run:
    ./go-demo

# Clean build artifacts
clean:
    rm -f go-demo
    rm -rf tmp/
    rm -f build-errors.log

# Install development dependencies
install-deps:
    go install github.com/cosmtrek/air@latest
    go mod download
