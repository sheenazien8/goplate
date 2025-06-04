#!/bin/bash

# Migration script for GoPlate
# Handles all database migration operations using dbmate

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Print functions
print_success() {
    echo -e "${GREEN}✅${NC} $1"
}

print_error() {
    echo -e "${RED}❌${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠️${NC} $1"
}

print_info() {
    echo -e "${BLUE}ℹ️${NC} $1"
}

# Configuration
MIGRATION_DIR="./db/migrations"
ENV_FILE=".env"

# Check if .env file exists
check_env_file() {
    if [[ ! -f "$ENV_FILE" ]]; then
        print_error ".env file not found. Please create one from .env.example"
        exit 1
    fi
}

# Load environment variables
load_env() {
    if [[ -f "$ENV_FILE" ]]; then
        export $(grep -v '^#' "$ENV_FILE" | xargs)
    fi
}

# Build database URL
build_db_url() {
    if [[ -z "$DB_CONNECTION" || -z "$DB_HOST" || -z "$DB_PORT" || -z "$DB_DATABASE" ]]; then
        print_error "Database configuration missing in .env file"
        print_info "Required variables: DB_DRIVER, DB_HOST, DB_PORT, DB_DATABASE, DB_USERNAME, DB_PASSWORD"
        exit 1
    fi

    case "$DB_CONNECTION" in
        "mysql")
            DB_URL="mysql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?parseTime=true"
            ;;
        "postgres"|"postgresql")
            DB_URL="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable"
            ;;
        *)
            print_error "Unsupported database driver: $DB_CONNECTION"
            print_info "Supported drivers: mysql, postgres"
            exit 1
            ;;
    esac
}

# Check if dbmate is installed
check_dbmate() {
    if ! command -v dbmate &> /dev/null; then
        print_error "dbmate is not installed"
        print_info "Install with: go install github.com/amacneil/dbmate@latest"
        print_info "Or run: make install-deps"
        exit 1
    fi
}

# Create migration directory if it doesn't exist
ensure_migration_dir() {
    if [[ ! -d "$MIGRATION_DIR" ]]; then
        mkdir -p "$MIGRATION_DIR"
        print_success "Created migration directory: $MIGRATION_DIR"
    fi
}

# Show help
show_help() {
    echo -e "${BLUE}GoPlate Migration Tool${NC}"
    echo ""
    echo "Usage: $0 <command> [arguments]"
    echo ""
    echo "Commands:"
    echo "  create <name>      Create a new migration file"
    echo "  up                 Run pending migrations"
    echo "  down               Rollback last migration"
    echo "  status             Show migration status"
    echo "  reset              Drop and recreate database"
    echo "  fresh              Drop, recreate and run all migrations"
    echo "  dump               Dump database schema"
    echo "  load               Load database schema"
    echo "  version            Show current migration version"
    echo "  connect            Connect to database interactive shell"
    echo "  help               Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 create create_users_table"
    echo "  $0 up"
    echo "  $0 status"
    echo "  $0 down"
    echo "  $0 connect"
}

# Create a new migration
create_migration() {
    if [[ -z "$1" ]]; then
        print_error "Migration name is required"
        echo "Usage: $0 create <migration_name>"
        exit 1
    fi

    local migration_name="$1"
    print_info "Creating migration: $migration_name"

    dbmate -d "$MIGRATION_DIR" --url "$DB_URL" new "$migration_name"
    print_success "Migration created successfully"
}

# Run pending migrations
migrate_up() {
    print_info "Running pending migrations..."
    dbmate -d "$MIGRATION_DIR" --url "$DB_URL" up
    print_success "Migrations completed"
}

# Rollback last migration
migrate_down() {
    print_warning "Rolling back last migration..."
    read -p "Are you sure you want to rollback? (y/N): " confirm
    if [[ $confirm =~ ^[Yy]$ ]]; then
        dbmate -d "$MIGRATION_DIR" --url "$DB_URL" down
        print_success "Rollback completed"
    else
        print_info "Rollback cancelled"
    fi
}

# Show migration status
migration_status() {
    print_info "Migration status:"
    dbmate -d "$MIGRATION_DIR" --url "$DB_URL" status
}

# Reset database (drop and recreate)
reset_database() {
    print_warning "This will drop and recreate the database!"
    read -p "Are you sure? Type 'yes' to continue: " confirm
    if [[ "$confirm" == "yes" ]]; then
        print_info "Dropping database..."
        dbmate -d "$MIGRATION_DIR" --url "$DB_URL" drop
        print_info "Creating database..."
        dbmate -d "$MIGRATION_DIR" --url "$DB_URL" create
        print_success "Database reset completed"
    else
        print_info "Reset cancelled"
    fi
}

# Fresh migration (reset + migrate)
fresh_migration() {
    print_warning "This will drop the database and run all migrations!"
    read -p "Are you sure? Type 'yes' to continue: " confirm
    if [[ "$confirm" == "yes" ]]; then
        reset_database
        migrate_up
        print_success "Fresh migration completed"
    else
        print_info "Fresh migration cancelled"
    fi
}

# Dump database schema
dump_schema() {
    print_info "Dumping database schema..."
    dbmate -d "$MIGRATION_DIR" --url "$DB_URL" dump
    print_success "Schema dumped successfully"
}

# Load database schema
load_schema() {
    print_info "Loading database schema..."
    dbmate -d "$MIGRATION_DIR" --url "$DB_URL" load
    print_success "Schema loaded successfully"
}

# Show current migration version
show_version() {
    print_info "Current migration version:"
    dbmate -d "$MIGRATION_DIR" --url "$DB_URL" status | grep "Applied" | tail -1 || echo "No migrations applied"
}

# Connect to database interactive shell
connect_database() {
    print_info "Connecting to database..."
    
    case "$DB_CONNECTION" in
        "mysql")
            if command -v mysql &> /dev/null; then
                print_info "Opening MySQL shell..."
                if [[ -n "$DB_PASSWORD" ]]; then
                    mysql -h "$DB_HOST" -P "$DB_PORT" -u "$DB_USERNAME" -p"$DB_PASSWORD" "$DB_DATABASE"
                else
                    mysql -h "$DB_HOST" -P "$DB_PORT" -u "$DB_USERNAME" "$DB_DATABASE"
                fi
            else
                print_error "mysql client not found. Please install MySQL client."
                print_info "On macOS: brew install mysql-client"
                print_info "On Ubuntu: sudo apt-get install mysql-client"
                exit 1
            fi
            ;;
        "postgres"|"postgresql")
            if command -v psql &> /dev/null; then
                print_info "Opening PostgreSQL shell..."
                export PGPASSWORD="$DB_PASSWORD"
                psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USERNAME" -d "$DB_DATABASE"
                unset PGPASSWORD
            else
                print_error "psql client not found. Please install PostgreSQL client."
                print_info "On macOS: brew install postgresql"
                print_info "On Ubuntu: sudo apt-get install postgresql-client"
                exit 1
            fi
            ;;
        *)
            print_error "Database connection not supported for driver: $DB_CONNECTION"
            exit 1
            ;;
    esac
}

# Main function
main() {
    # Check prerequisites
    check_env_file
    load_env
    check_dbmate
    build_db_url
    ensure_migration_dir

    # Handle commands
    case "$1" in
        "create")
            create_migration "$2"
            ;;
        "up"|"migrate")
            migrate_up
            ;;
        "down"|"rollback")
            migrate_down
            ;;
        "status")
            migration_status
            ;;
        "reset")
            reset_database
            ;;
        "fresh")
            fresh_migration
            ;;
        "dump")
            dump_schema
            ;;
        "load")
            load_schema
            ;;
        "version")
            show_version
            ;;
        "connect")
            connect_database
            ;;
        "help"|"")
            show_help
            ;;
        *)
            print_error "Unknown command: $1"
            show_help
            exit 1
            ;;
    esac
}

main "$@"
