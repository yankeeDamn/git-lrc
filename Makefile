.PHONY: build build-all build-local run bump release clean test

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=lrc

# Build lrc for the current platform
build:
	$(GOBUILD) -o $(BINARY_NAME) .

# Build lrc for all platforms (linux/darwin/windows × amd64/arm64)
# Output: dist/<platform>/lrc[.exe] + SHA256SUMS
# Version is extracted from appVersion constant in main.go
build-all:
	@echo "🔨 Building lrc CLI for all platforms..."
	@python scripts/lrc_build.py -v build

# Build lrc locally for the current platform and install
build-local:
	@echo "🔨 Building lrc CLI locally (dirty tree allowed)..."
	@go build -o /tmp/lrc .
	@sudo rm -f /usr/local/bin/lrc || true
	@sudo install -m 0755 /tmp/lrc /usr/local/bin/lrc
	@sudo cp /usr/local/bin/lrc /usr/bin/git-lrc
	@echo "✅ Installed lrc to /usr/local/bin and git-lrc to /usr/bin"

# Run the locally built lrc CLI (pass args via ARGS="--flag value")
run: build-local
	@echo "▶️ Running lrc CLI locally..."
	@lrc $(ARGS)

# Bump lrc version by editing appVersion in main.go
# Prompts for version bump type (patch/minor/major)
bump:
	@echo "📝 Bumping lrc version..."
	@python scripts/lrc_build.py bump

# Build and upload lrc to Backblaze B2
release:
	@echo "🚀 Building and releasing lrc..."
	@python scripts/lrc_build.py -v release

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf dist/ $(BINARY_NAME)
	@echo "✅ Clean complete"

# Run tests
test:
	$(GOTEST) -count=1 ./...
