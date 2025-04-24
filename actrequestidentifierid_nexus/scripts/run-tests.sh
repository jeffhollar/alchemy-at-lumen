#!/bin/bash

# Exit on any error
set -e

# Default values
TEST_TYPE="all"  # Can be "unit", "integration", or "all"

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --test-type)
            TEST_TYPE="$2"
            shift 2
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

echo "Running tests: $TEST_TYPE"

case $TEST_TYPE in
    "unit")
        echo "Running unit tests..."
        ./gradlew test
        ;;
    "integration")
        echo "Running integration tests..."
        ./gradlew bootRun --args='--spring.profiles.active=integration-test'
        ;;
    "all")
        echo "Running all tests..."
        echo "1. Unit tests"
        ./gradlew test
        echo "2. Integration tests"
        ./gradlew bootRun --args='--spring.profiles.active=integration-test'
        ;;
    *)
        echo "Invalid test type: $TEST_TYPE"
        echo "Valid options are: unit, integration, all"
        exit 1
        ;;
esac

echo "Tests completed!" 