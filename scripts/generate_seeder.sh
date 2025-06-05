#!/bin/sh
# Usage: ./generate_seeder.sh SeederName

SEEDER_FILENAME=$1

# Ensure a seeder name was provided
if [ -z "$SEEDER_FILENAME" ]; then
    echo "Error: Seeder name is required"
    exit 1
fi

# Function to convert the seeder name to a struct name
to_struct_name() {
    echo "$1" | perl -pe 's/(?:^|-)./uc($&)/ge;s/-//g;'
}

# Convert the seeder name to a struct name
STRUCT_NAME=$(to_struct_name "$SEEDER_FILENAME")

# Define the path and filename
SEEDER_DIR="./db/seeders"
SEEDER_FILE="$SEEDER_DIR/${SEEDER_FILENAME}.go"

# Create the seeders directory if it doesn't exist
mkdir -p $SEEDER_DIR

# Generate the seeder file with a basic structure
cat <<EOL > $SEEDER_FILE
package seeders

import (
	"gorm.io/gorm"
)

type $STRUCT_NAME struct {}

func (s $STRUCT_NAME) Seed(db *gorm.DB) error {
    return nil
}
EOL

echo "Seeder $STRUCT_NAME has been generated at $SEEDER_FILE"
echo "Please register the seeder in ./db/seeders/register.go"

