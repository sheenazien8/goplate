#!/bin/bash

# Galaplate CLI Installation Script
# This script installs the galaplate CLI tool for easy project creation

set -e

GALAPLATE_VERSION="v0.1.0-dev"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="galaplate"

echo "🚀 Installing Galaplate CLI ${GALAPLATE_VERSION}..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first: https://golang.org/doc/install"
    exit 1
fi

# Check if running from the galaplate directory
if [ ! -f "cli/cmd/galaplate/main.go" ]; then
    echo "❌ Please run this script from the galaplate repository root directory"
    exit 1
fi

echo "📦 Building Galaplate CLI..."

# Build the CLI
cd cli
go build -ldflags "-X main.version=${GALAPLATE_VERSION} -X main.date=$(date -u +%Y-%m-%dT%H:%M:%SZ)" -o ${BINARY_NAME} cmd/galaplate/main.go

echo "📁 Installing to ${INSTALL_DIR}..."

# Install the binary
sudo mv ${BINARY_NAME} ${INSTALL_DIR}/${BINARY_NAME}
sudo chmod +x ${INSTALL_DIR}/${BINARY_NAME}

echo "✅ Galaplate CLI installed successfully!"
echo ""
echo "🎯 Quick Start:"
echo "   galaplate new my-api                    # Create a new API project"
echo "   galaplate templates                     # List available templates"
echo "   galaplate help                          # Show help"
echo ""
echo "🛠️  After creating a project:"
echo "   cd my-api"
echo "   cp .env.example .env                  # Configure environment"
echo "   go mod tidy                           # Install dependencies"
echo "   go run main.go console list           # See all generators"
echo "   go run main.go console make:model User  # Generate models"
echo ""
echo "For more information: https://github.com/sheenazien8/galaplate"