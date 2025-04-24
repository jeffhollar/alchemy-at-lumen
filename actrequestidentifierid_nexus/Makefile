# Configuration
JAR_NAME=actrequest.jar
JAR_PATH=build/libs/$(JAR_NAME)
PROPERTIES_FILE=config/temporal.properties

# Build the worker jar
.PHONY: build
build:
	./gradlew clean build bootJar

# Build the tests
.PHONY: build-tests
build-tests:
	./gradlew clean test build

# Run the worker
.PHONY: run-worker
run-worker:
	java -jar build/libs/actrequest-1.0.0.jar \
		--spring.profiles.active=dev \
		--server.port=8080

# Run the workflow client service
.PHONY: run-service
run-service:
	java -jar build/libs/actrequest-1.0.0.jar \
		--spring.profiles.active=dev,client \
		--server.port=8081

# Clean build artifacts
.PHONY: clean
clean:
	./gradlew clean

# Run tests
.PHONY: test
test:
	./gradlew test

# Run the Spring Boot application
.PHONY: run
run:
	./gradlew bootRun -Dconfig.file=config.properties

# Help target
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build        - Build the project"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make test         - Run tests"
	@echo "  make run-worker   - Run the worker"
	@echo "  make run-service - Run the service client"

# Default target
.DEFAULT_GOAL := help 