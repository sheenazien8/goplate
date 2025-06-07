#!/bin/sh
# Usage: ./generate_cron.sh cronName

CRON_FILENAME=$1

# Ensure a cron name was provided
if [ -z "$CRON_FILENAME" ]; then
    echo "Error: cron name is required"
    exit 1
fi

# Function to convert the cron name to a struct name
to_struct_name() {
    echo "$1" | perl -pe 's/(?:^|-)./uc($&)/ge;s/-//g;'
}

# Convert the cron name to a struct name
STRUCT_NAME=$(to_struct_name "$CRON_FILENAME")

# Define the path and filename
CRON_DIR="./pkg/scheduler"
CRON_FILE="$CRON_DIR/${CRON_FILENAME}.go"

# Create the crons directory if it doesn't exist
mkdir -p $CRON_DIR

# Generate the cron file with a basic structure
cat <<EOL > $CRON_FILE
package scheduler

type $STRUCT_NAME struct{}

func ($STRUCT_NAME) Handle() (string, func()) {
	return "@every 5s", func() {
        // This is where you would implement the task logic
	}
}

func init() {
	registerScheduler("$CRON_FILENAME", $STRUCT_NAME{}.Handle)
}
EOL

echo "cron $STRUCT_NAME has been generated at $CRON_FILE"
