#!/bin/sh

# Usage: ./generate_dto.sh ModelName

DTO_NAME=$1

# Ensure a model name was provided
if [ -z "$DTO_NAME" ]; then
    echo "Error: Model name is required"
    exit 1
fi

# Function to convert the model name to a struct name
to_struct_name() {
    echo "$1" | perl -pe 's/(?:^|-)./uc($&)/ge;s/-//g;'
}

# Convert the model name to a struct name
STRUCT_NAME=$(to_struct_name "$DTO_NAME")

# Define the path and filename
DTO_DIR="./pkg/dto"
DTO_FILE="$DTO_DIR/${DTO_NAME}.go"

# Create the models directory if it doesn't exist
mkdir -p $DTO_DIR

# Generate the model file with a basic structure
cat <<EOL > $DTO_FILE
package dto

import (
    "github.com/gofiber/fiber/v2"
    "github.com/sheenazien8/goplate/pkg/utils"
)

type $STRUCT_NAME struct {
}

func (s *$STRUCT_NAME) Validate(c *fiber.Ctx) (u *$STRUCT_NAME, err error) {
	myValidator := &utils.XValidator{}
	if err := c.BodyParser(s); err != nil {
		return nil, err
	}

	if err := myValidator.Validate(s); err != nil {
		return nil, &fiber.Error{
			Code:    fiber.ErrUnprocessableEntity.Code,
			Message: err.Error(),
		}
	}

	return s, nil
}

EOL

echo "DTO $STRUCT_NAME has been generated at $DTO_FILE"

