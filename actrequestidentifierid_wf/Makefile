# Configuration
JAR_NAME=actrequest.jar
JAR_PATH=build/libs/$(JAR_NAME)
PROPERTIES_FILE=config/temporal.properties

# Build the worker jar
.PHONY: build
build:
	./gradlew fatJar

# Run the worker
.PHONY: run-worker
run-worker: build
	java -jar $(JAR_PATH) $(PROPERTIES_FILE)

# Run the workflow client
.PHONY: run-workflow
run-workflow: build
	java -cp $(JAR_PATH) com.lumen.workflow.client.ActClient $(PROPERTIES_FILE)

# Clean build artifacts
.PHONY: clean
clean:
	./gradlew clean

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build        - Build the worker jar"
	@echo "  run-worker   - Run the worker"
	@echo "  run-workflow - Run the workflow client"
	@echo "  clean        - Clean build artifacts"
	@echo "  help         - Show this help message"

# Default target
.DEFAULT_GOAL := help 