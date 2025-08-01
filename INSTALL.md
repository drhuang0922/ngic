# NGIC Installation and Release Guide

This document provides instructions for setting up ngic for distribution via `go install` and Homebrew.

## Project Structure

The project is now structured for easy installation:

```
ngic/
├── cmd/ngic/              # Main application entry point
│   └── main.go           # CLI interface with flag parsing
├── pkg/converter/         # Core conversion logic
│   ├── converter.go      # Image conversion functionality  
│   └── converter_test.go # Unit tests
├── homebrew/             # Homebrew formula
│   └── ngic.rb          # Formula for brew installation
├── scripts/              # Utility scripts
│   └── setup-homebrew.sh # Homebrew setup helper
├── .github/workflows/    # GitHub Actions
│   ├── ci.yml           # Continuous integration
│   └── release.yml      # Automated releases
├── build/               # Build artifacts (gitignored)
├── Makefile            # Build automation
├── go.mod              # Go module definition
└── README.md           # User documentation
```

## Installation Methods

### 1. Go Install (Recommended for Go users)

Users with Go installed can use:

```bash
go install github.com/drhuang0922/ngic/cmd/ngic@latest
```

This installs the latest version directly from the repository.

### 2. Homebrew Installation

For Homebrew installation, you need to:

1. **Create a Homebrew tap repository:**
   - Create a new repository named `homebrew-tools` in your GitHub account
   - Make it public

2. **Set up the formula:**
   - Copy `homebrew/ngic.rb` to `Formula/ngic.rb` in your tap repository
   - Create a release (see Release Process below)
   - Update the SHA256 hash in the formula

3. **Users can then install with:**
   ```bash
   brew tap drhuang0922/tools
   brew install ngic
   ```

### 3. Binary Releases

GitHub Actions automatically creates binary releases for multiple platforms:
- macOS (Intel and Apple Silicon)
- Linux (AMD64 and ARM64)
- Windows (AMD64)

### 4. Build from Source

```bash
git clone https://github.com/drhuang0922/ngic.git
cd ngic
make build
# Binary will be in ./build/ngic
```

## Release Process

### Creating a Release

1. **Tag a version:**
   ```bash
   git tag v1.0.0
   git push origin --tags
   ```

2. **GitHub Actions automatically:**
   - Builds binaries for all platforms
   - Creates release archives
   - Publishes the release with assets

3. **For Homebrew (manual step):**
   - Download the source tarball from the release
   - Calculate SHA256: `shasum -a 256 v1.0.0.tar.gz`
   - Update the `sha256` field in your tap's Formula/ngic.rb

### Build Commands

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Create release archives
make release

# Install locally
make install

# Run tests
make test

# Clean build artifacts
make clean
```

## Testing Installation

### Test Go Install
```bash
# Install
go install github.com/drhuang0922/ngic/cmd/ngic@latest

# Test
ngic -version
ngic -h
```

### Test Homebrew (after tap setup)
```bash
# Install
brew tap drhuang0922/tools
brew install ngic

# Test
ngic -version
ngic -h
```

### Test Binary Download
```bash
# Download appropriate binary from releases
wget https://github.com/drhuang0922/ngic/releases/download/v1.0.0/ngic-darwin-amd64.tar.gz

# Extract and test
tar -xzf ngic-darwin-amd64.tar.gz
./ngic-darwin-amd64 -version
```

## Module Configuration

The `go.mod` file is configured with:
- Module name: `github.com/drhuang0922/ngic`
- Go version: 1.21 (for broad compatibility)
- Zero external dependencies (uses only standard library)

## GitHub Actions

- **CI Workflow:** Runs tests on Go 1.21 and 1.22, builds for verification
- **Release Workflow:** Triggered on version tags, creates cross-platform binaries

## Key Features for Distribution

1. **Zero Dependencies:** Uses only Go standard library
2. **Cross-Platform:** Builds for macOS, Linux, Windows
3. **Small Binary Size:** Optimized builds with trimpath and ldflags
4. **Proper Versioning:** Version information embedded at build time
5. **CLI Standards:** Follows Unix CLI conventions with proper help and flags

## Usage Examples

```bash
# Single file conversion
ngic input.jpg output.webp webp

# With quality setting
ngic -q 75 input.png output.jpg jpeg

# Batch conversion
ngic -batch images/ webp

# Batch with custom output directory
ngic -batch -o converted/ images/ webp
```

## Next Steps

1. Push the code to GitHub
2. Create a v1.0.0 release
3. Set up Homebrew tap repository
4. Update documentation with actual release links
5. Announce the tool to relevant communities
