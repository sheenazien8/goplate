# Installation

This guide covers different ways to install and set up GoPlate for your development environment.

## System Requirements

### Minimum Requirements

- **Go**: Version 1.22 or higher
- **Operating System**: Linux, macOS, or Windows
- **Memory**: 512MB RAM minimum (2GB recommended)
- **Disk Space**: 100MB for GoPlate + dependencies

### Database Requirements

Choose one of the following databases:

- **MySQL**: Version 5.7 or higher (8.0+ recommended)
- **PostgreSQL**: Version 12 or higher (14+ recommended)

### Optional Tools

- **Make**: For using the provided Makefile commands
- **Docker**: For containerized development (optional)
- **Git**: For version control

## Installation Methods

### Method 1: Automated Installation Script

The easiest way to get started with GoPlate:

```bash
curl -sSL https://raw.githubusercontent.com/sheenazien8/goplate/master/install.sh | bash
```

This script will:
- Download the latest GoPlate release
- Install the CLI tool to your `$GOPATH/bin`
- Set up necessary permissions
- Verify the installation

### Method 2: Go Install

Install directly using Go's package manager:

```bash
go install github.com/sheenazien8/goplate/cmd/goplate@latest
```

Make sure your `$GOPATH/bin` is in your `$PATH`:

```bash
# Add to your shell profile (.bashrc, .zshrc, etc.)
export PATH=$PATH:$(go env GOPATH)/bin
```

### Method 3: Manual Download

1. **Download the latest release:**
   ```bash
   # Replace VERSION with the latest version
   wget https://github.com/sheenazien8/goplate/releases/download/v1.0.0/goplate-linux-amd64.tar.gz
   ```

2. **Extract and install:**
   ```bash
   tar -xzf goplate-linux-amd64.tar.gz
   sudo mv goplate /usr/local/bin/
   chmod +x /usr/local/bin/goplate
   ```

### Method 4: Build from Source

For developers who want the latest features:

```bash
# Clone the repository
git clone https://github.com/sheenazien8/goplate.git
cd goplate

# Build the CLI tool
go build -o goplate cmd/goplate/main.go

# Install to your PATH
sudo mv goplate /usr/local/bin/
```

## Verify Installation

Check that GoPlate is installed correctly:

```bash
goplate --version
```

You should see output similar to:
```
GoPlate v1.0.0
```

## Development Dependencies

Install additional tools for the best development experience:

### Essential Tools

```bash
# Install development dependencies using console commands
go run main.go console db:create
```

This installs:
- **reflex**: For hot reload during development
- **dbmate**: For database migrations (now integrated via console commands)
- **dotenv-cli**: For environment variable management

### Manual Installation

If you prefer to install tools manually:

```bash
# Hot reload tool
go install github.com/cespare/reflex@latest

# Database migration tool (optional - console commands available)
go install github.com/amacneil/dbmate@latest

# Environment CLI (requires Node.js)
npm install -g dotenv-cli
```

### Console Command System

GoPlate now includes a powerful console command system. After installation, you can:

```bash
# View all available console commands
go run main.go console list

# Generate new models, DTOs, jobs, etc.
go run main.go console make:model User
go run main.go console make:dto UserDto
go run main.go console make:job ProcessEmail

# Database operations
go run main.go console db:up
go run main.go console db:status
go run main.go console db:fresh
```

## Database Setup

### MySQL Installation

<!-- tabs:start -->

#### **macOS (Homebrew)**

```bash
# Install MySQL
brew install mysql

# Start MySQL service
brew services start mysql

# Secure installation (optional but recommended)
mysql_secure_installation
```

#### **Ubuntu/Debian**

```bash
# Update package index
sudo apt update

# Install MySQL
sudo apt install mysql-server

# Start MySQL service
sudo systemctl start mysql
sudo systemctl enable mysql

# Secure installation
sudo mysql_secure_installation
```

#### **CentOS/RHEL**

```bash
# Install MySQL repository
sudo yum install mysql-server

# Start MySQL service
sudo systemctl start mysqld
sudo systemctl enable mysqld

# Get temporary root password
sudo grep 'temporary password' /var/log/mysqld.log

# Secure installation
mysql_secure_installation
```

#### **Windows**

1. Download MySQL installer from [mysql.com](https://dev.mysql.com/downloads/installer/)
2. Run the installer and follow the setup wizard
3. Choose "Developer Default" for a complete installation
4. Set a root password during installation

<!-- tabs:end -->

### PostgreSQL Installation

<!-- tabs:start -->

#### **macOS (Homebrew)**

```bash
# Install PostgreSQL
brew install postgresql

# Start PostgreSQL service
brew services start postgresql

# Create a database user (optional)
createuser --interactive
```

#### **Ubuntu/Debian**

```bash
# Update package index
sudo apt update

# Install PostgreSQL
sudo apt install postgresql postgresql-contrib

# Start PostgreSQL service
sudo systemctl start postgresql
sudo systemctl enable postgresql

# Switch to postgres user
sudo -u postgres psql
```

#### **CentOS/RHEL**

```bash
# Install PostgreSQL
sudo yum install postgresql-server postgresql-contrib

# Initialize database
sudo postgresql-setup initdb

# Start PostgreSQL service
sudo systemctl start postgresql
sudo systemctl enable postgresql
```

#### **Windows**

1. Download PostgreSQL installer from [postgresql.org](https://www.postgresql.org/download/windows/)
2. Run the installer and follow the setup wizard
3. Remember the password you set for the postgres user

<!-- tabs:end -->

## IDE and Editor Setup

### Visual Studio Code

Recommended extensions for Go development:

```bash
# Install VS Code extensions
code --install-extension golang.go
code --install-extension ms-vscode.vscode-json
code --install-extension bradlc.vscode-tailwindcss
```

### GoLand

GoLand by JetBrains provides excellent Go support out of the box. No additional setup required.

### Vim/Neovim

For Vim users, consider these plugins:
- **vim-go**: Comprehensive Go support
- **coc.nvim**: Language server support
- **nerdtree**: File explorer

## Environment Setup

### Git Configuration

Set up Git hooks for better development workflow:

```bash
# In your GoPlate project directory
git config core.hooksPath .githooks
chmod +x .githooks/*
```

## Troubleshooting

### Common Installation Issues

#### Go Not Found

```bash
# Check if Go is installed
go version

# If not installed, download from https://golang.org/dl/
# Or use package manager:
# macOS: brew install go
# Ubuntu: sudo apt install golang-go
```

#### Permission Denied

```bash
# Fix permissions for global installation
sudo chown -R $(whoami) $(go env GOPATH)
sudo chown -R $(whoami) $(go env GOROOT)
```

#### PATH Issues

```bash
# Add Go binary path to your shell profile
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

### Database Connection Issues

#### MySQL Connection Refused

```bash
# Check if MySQL is running
sudo systemctl status mysql

# Start MySQL if not running
sudo systemctl start mysql

# Check MySQL port
netstat -tlnp | grep :3306
```

#### PostgreSQL Connection Issues

```bash
# Check PostgreSQL status
sudo systemctl status postgresql

# Check PostgreSQL port
netstat -tlnp | grep :5432

# Connect to PostgreSQL
sudo -u postgres psql
```

### Performance Optimization

#### Go Module Proxy

Speed up dependency downloads:

```bash
# Set Go module proxy
export GOPROXY=https://proxy.golang.org,direct
export GOSUMDB=sum.golang.org
```

#### Build Cache

Enable Go build cache for faster builds:

```bash
# Check build cache location
go env GOCACHE

# Clean build cache if needed
go clean -cache
```

## Next Steps

After successful installation:

1. **[Quick Start](/quick-start)** - Create your first project
2. **[Configuration](/configuration)** - Set up your environment
3. **[Project Structure](/project-structure)** - Understand the codebase

## Getting Help

If you encounter issues during installation:

- **[GitHub Issues](https://github.com/sheenazien8/goplate/issues)** - Report bugs
- **[Discussions](https://github.com/sheenazien8/goplate/discussions)** - Ask questions
- **[Documentation](/)** - Browse the full documentation

---

**Installation complete!** 🎉 You're ready to start building with GoPlate!
