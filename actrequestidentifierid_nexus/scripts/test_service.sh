#!/bin/bash

# Default values
TARGET_HOST="192.168.1.233:7233"
NAMESPACE="nexus-lumen-act-ns"
TASK_QUEUE="act-task-queue"
API_KEY=""
CLIENT_CERT=""
CLIENT_KEY=""
CA_CERT=""

# Help function
show_help() {
    echo "Usage: $0 [options]"
    echo "Options:"
    echo "  -h, --help         Show this help message"
    echo "  -t, --target-host  Target host (default: 192.168.1.233:7233)"
    echo "  -n, --namespace    Namespace (default: nexus-lumen-act-ns)"
    echo "  -q, --task-queue   Task queue (default: act-task-queue)"
    echo "  -a, --api-key      API Key for authentication"
    echo "  --cert             Client certificate file"
    echo "  --key              Client key file"
    echo "  --ca-cert          CA certificate file"
}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_help
            exit 0
            ;;
        -t|--target-host)
            TARGET_HOST="$2"
            shift 2
            ;;
        -n|--namespace)
            NAMESPACE="$2"
            shift 2
            ;;
        -q|--task-queue)
            TASK_QUEUE="$2"
            shift 2
            ;;
        -a|--api-key)
            API_KEY="$2"
            shift 2
            ;;
        --cert)
            CLIENT_CERT="$2"
            shift 2
            ;;
        --key)
            CLIENT_KEY="$2"
            shift 2
            ;;
        --ca-cert)
            CA_CERT="$2"
            shift 2
            ;;
        *)
            echo "Unknown option: $1"
            show_help
            exit 1
            ;;
    esac
done

# Create temporary properties file with a clean name
TEMP_DIR=$(mktemp -d)
TEMP_PROPS="${TEMP_DIR}/application.properties"

cat > "$TEMP_PROPS" << EOF
# Spring Boot Configuration
spring.application.name=act-request-nexus
server.port=8081

# Temporal Server Configuration
temporal.target-host=${TARGET_HOST}
temporal.namespace=${NAMESPACE}
temporal.task-queue=${TASK_QUEUE}

# Logging Configuration
logging.level.root=INFO
logging.level.com.lumen.workflow=DEBUG

# Security Configuration
EOF

if [ -n "$API_KEY" ]; then
    echo "security.api-key=$API_KEY" >> "$TEMP_PROPS"
fi

if [ -n "$CLIENT_CERT" ]; then
    echo "security.ssl.client-cert=$CLIENT_CERT" >> "$TEMP_PROPS"
fi

if [ -n "$CLIENT_KEY" ]; then
    echo "security.ssl.client-key=$CLIENT_KEY" >> "$TEMP_PROPS"
fi

if [ -n "$CA_CERT" ]; then
    echo "security.ssl.server-root-ca-cert=$CA_CERT" >> "$TEMP_PROPS"
fi

# Run the test with the temporary properties file
java -jar build/libs/actrequest-1.0.0.jar --spring.config.location=file:${TEMP_PROPS}

# Clean up
rm -rf "$TEMP_DIR" 