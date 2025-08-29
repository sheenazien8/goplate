#!/bin/bash

set -e

REPO_URL="https://github.com/sheenazien8/galaplate"
BINARY_NAME="galaplate"
INSTALL_DIR="/usr/local/bin"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

print_status() {
    echo -e "${BLUE}==>${NC} $1"
}

print_success() {
    echo -e "${GREEN}✅${NC} $1"
}

print_error() {
    echo -e "${RED}❌${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠️${NC} $1"
}

# Check if Go is installed
check_go() {
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed. Please install Go first: https://golang.org/doc/install"
        exit 1
    fi
    print_success "Go is installed: $(go version)"
}

# Install the CLI tool
install_galaplate() {
    print_status "Installing Galaplate CLI..."
    
    # Create temporary directory
    TEMP_DIR=$(mktemp -d)
    cd "$TEMP_DIR"
    
    # Clone the repository
    print_status "Downloading from GitHub..."
    git clone "$REPO_URL" .
    
    # Build the CLI
    print_status "Building CLI tool..."
    cd cmd/galaplate
    go build -o "$BINARY_NAME" .
    
    # Install to system
    if [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
        # Windows
        INSTALL_DIR="$HOME/bin"
        mkdir -p "$INSTALL_DIR"
        mv "$BINARY_NAME.exe" "$INSTALL_DIR/"
        print_warning "Add $INSTALL_DIR to your PATH environment variable"
    else
        # Unix-like systems
        if [[ $(id -u) -eq 0 ]]; then
            # Running as root
            mv "$BINARY_NAME" "$INSTALL_DIR/"
        else
            # Running as regular user
            print_status "Installing to $INSTALL_DIR (requires sudo)..."
            sudo mv "$BINARY_NAME" "$INSTALL_DIR/"
        fi
    fi
    
    # Cleanup
    cd /
    rm -rf "$TEMP_DIR"
    
    print_success "Galaplate CLI installed successfully!"
}

# Usage instructions
show_usage() {
    echo ""
    echo -e "${BLUE}Usage:${NC}"
    echo "  galaplate <project-name>    Create a new Galaplate project"
    echo ""
    echo -e "${BLUE}Example:${NC}"
    echo "  galaplate my-api"
    echo "  cd my-api"
    echo "  go mod tidy"
    echo "  cp .env.example .env"
    echo "  go run main.go"
    echo ""
}

# Main installation process
main() {
    echo -e "${GREEN}"
    echo "  ____       ____  _       _       "
    echo " / ___| ___ |  _ \| | __ _| |_ ___ "
    echo "| |  _ / _ \| |_) | |/ _\` | __/ _ \\"
    echo "| |_| | (_) |  __/| | (_| | ||  __/"
    echo " \____|\___/|_|   |_|\__,_|\__\___|"
    echo ""
    echo -e "${NC}${BLUE}Galaplate Installer${NC}"
    echo ""
    
    check_go
    install_galaplate
    show_usage
    
    print_success "Installation complete! You can now use 'galaplate <project-name>' to create new projects."
}

main "$@"