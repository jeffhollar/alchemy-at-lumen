#!/bin/bash

# Script to clean up files after model creation
# Preserves essential model files while removing temporary and generated files

# Define the base directory
BASE_DIR="lln_model"

# Files to preserve
PRESERVE_FILES=(
    "saved_model/model.keras"
    "scaler.joblib"
)

# Files to remove
REMOVE_FILES=(
    "training_history.png"
    "anomaly_detection.png"
)

# Function to check if a file exists and remove it
remove_file() {
    local file="$1"
    if [ -f "$BASE_DIR/$file" ]; then
        echo "Removing $file..."
        rm "$BASE_DIR/$file"
    else
        echo "File $file not found, skipping..."
    fi
}

# Main cleanup process
echo "Starting cleanup process..."

# Remove specified files
for file in "${REMOVE_FILES[@]}"; do
    remove_file "$file"
done

echo "Cleanup complete!"
echo "Preserved files:"
for file in "${PRESERVE_FILES[@]}"; do
    if [ -f "$BASE_DIR/$file" ]; then
        echo "- $file"
    fi
done 