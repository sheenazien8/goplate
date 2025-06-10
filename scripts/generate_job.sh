#!/bin/sh
# Usage: ./generate_job.sh jobName

JOB_FILENAME=$1

# Ensure a job name was provided
if [ -z "$JOB_FILENAME" ]; then
    echo "Error: job name is required"
    exit 1
fi

# Function to convert the job name to a struct name
to_struct_name() {
    echo "$1" | perl -pe 's/(?:^|-)./uc($&)/ge;s/-//g;'
}

# Convert the job name to a struct name
STRUCT_NAME=$(to_struct_name "$JOB_FILENAME")

# Determine module name
GO_MOD_FILE="./go.mod"
MODULE_NAME=$(grep -o '^module .*' "$GO_MOD_FILE" | cut -d ' ' -f 2)

# Load DB_CONNECTION from .env if present
if [ -f .env ]; then
    DB_CONNECTION=$(grep -m1 '^DB_CONNECTION=' .env | cut -d '=' -f2-)
fi
if [ -z "$DB_CONNECTION" ]; then
    echo "Error: DB_CONNECTION is not set. Please define it in .env"
    exit 1
fi

# Map DB_CONNECTION to stub suffix
case "$DB_CONNECTION" in
    mysql)
        DB_SUFFIX="mysql"
        ;;
    postgres|postgresql)
        DB_SUFFIX="pgsql"
        ;;
    *)
        echo "Error: Unsupported DB_CONNECTION '$DB_CONNECTION'"
        exit 1
        ;;
esac

# Define the path and filename for job handler
JOB_DIR="./pkg/queue/jobs"
JOB_FILE="$JOB_DIR/${JOB_FILENAME}.go"

# Create the jobs directory if it doesn't exist
mkdir -p "$JOB_DIR"

# Generate the job handler file with a basic structure
cat <<EOL > "$JOB_FILE"
package jobs

import (
	"time"
	"encoding/json"

	"$MODULE_NAME/pkg/queue"
)

type $STRUCT_NAME struct {
}

func (e $STRUCT_NAME) MaxAttempts() int {
	return 3
}

func (e $STRUCT_NAME) RetryAfter() time.Duration {
	return 2 * time.Minute
}

func ($STRUCT_NAME) Type() string {
	return "$JOB_FILENAME"
}

func ($STRUCT_NAME) Handle(payload json.RawMessage) error {
	return nil
}

func init() {
	queue.RegisterJob($STRUCT_NAME{})
}
EOL

echo "Job $STRUCT_NAME has been generated at $JOB_FILE"
# Generate migration for jobs table if not exists
MIGRATIONS_DIR="./db/migrations"
STUB_DIR="./scripts/stubs/migrations"

mkdir -p "$MIGRATIONS_DIR"

# Only create migration if none exists
if ! ls "$MIGRATIONS_DIR"/*_create_jobs_table.*.sql 1> /dev/null 2>&1; then
    STUB_FILE=$(ls "$STUB_DIR"/*_create_jobs_table."$DB_SUFFIX".sql.stub 2>/dev/null | head -n1)
    if [ -z "$STUB_FILE" ]; then
        echo "Error: Migration stub for $DB_CONNECTION not found in $STUB_DIR"
        exit 1
    fi
    TARGET_NAME=$(basename "$STUB_FILE" | sed -E "s/\\.$DB_SUFFIX\\.sql\\.stub$/.sql/")
    TARGET_FILE="$MIGRATIONS_DIR/$TARGET_NAME"
    cp "$STUB_FILE" "$TARGET_FILE"
    echo "Migration has been generated at $TARGET_FILE"
fi
