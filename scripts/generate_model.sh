#!/bin/sh

# Usage: ./generate_model.sh ModelName

MODEL_NAME=$1

# Ensure a model name was provided
if [ -z "$MODEL_NAME" ]; then
    echo "Error: Model name is required"
    exit 1
fi

# Function to convert the model name to a struct name
to_struct_name() {
    echo "$1" | perl -pe 's/(?:^|-)./uc($&)/ge;s/-//g;'
}

# Convert the model name to a struct name
STRUCT_NAME=$(to_struct_name "$MODEL_NAME")

# Define the path and filename
MODEL_DIR="./pkg/models"
MODEL_FILE="$MODEL_DIR/${MODEL_NAME}.go"

# Create the models directory if it doesn't exist
mkdir -p $MODEL_DIR

# Generate the model file with a basic structure
cat <<EOL > $MODEL_FILE
package models

import (
	"time"

	"gorm.io/gorm"
)

type $STRUCT_NAME struct {
    ID        uint           \`gorm:"primaryKey" json:"id"\`
	CreatedAt time.Time      \`json:"created_at"\`
	UpdatedAt time.Time      \`json:"updated_at"\`
	DeletedAt gorm.DeletedAt \`gorm:"index" json:"deleted_at"\`
}
EOL

echo "Model $STRUCT_NAME has been generated at $MODEL_FILE"

