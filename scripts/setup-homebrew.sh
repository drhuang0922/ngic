#!/bin/bash

# Setup script for ngic Homebrew installation
# This script helps set up a Homebrew tap for ngic

set -e

REPO_OWNER="drhuang0922"
REPO_NAME="ngic"
TAP_NAME="homebrew-tools"

echo "🍺 Setting up Homebrew tap for ngic..."

# Check if Homebrew is installed
if ! command -v brew &> /dev/null; then
    echo "❌ Homebrew is not installed. Please install Homebrew first:"
    echo "   /bin/bash -c \"\$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\""
    exit 1
fi

echo "✅ Homebrew is installed"

# Instructions for creating a Homebrew tap
echo ""
echo "📋 To make ngic installable via Homebrew, you need to:"
echo ""
echo "1. Create a new repository named 'homebrew-tools' in your GitHub account"
echo "2. Copy the formula from ./homebrew/ngic.rb to the new repository"
echo "3. Create a release of ngic with version tags"
echo ""
echo "Here are the steps:"
echo ""
echo "Step 1: Create the tap repository"
echo "  - Go to https://github.com/new"
echo "  - Repository name: homebrew-tools"
echo "  - Make it public"
echo "  - Initialize with README"
echo ""
echo "Step 2: Add the formula"
echo "  - Clone your homebrew-tools repository"
echo "  - Copy ./homebrew/ngic.rb to Formula/ngic.rb in the tap repo"
echo "  - Update the SHA256 hash after creating a release"
echo ""
echo "Step 3: Create a release"
echo "  - Tag a version: git tag v1.0.0"
echo "  - Push tags: git push origin --tags"
echo "  - GitHub Actions will create the release automatically"
echo ""
echo "Step 4: Update formula with correct SHA256"
echo "  - Download the source tarball from the release"
echo "  - Calculate SHA256: shasum -a 256 v1.0.0.tar.gz"
echo "  - Update the sha256 field in Formula/ngic.rb"
echo ""
echo "After setup, users can install with:"
echo "  brew tap ${REPO_OWNER}/tools"
echo "  brew install ngic"
echo ""
echo "🚀 Current installation methods:"
echo "  1. Go install: go install github.com/${REPO_OWNER}/${REPO_NAME}/cmd/ngic@latest"
echo "  2. Download binary from releases"
echo "  3. Build from source: make install"
echo ""
