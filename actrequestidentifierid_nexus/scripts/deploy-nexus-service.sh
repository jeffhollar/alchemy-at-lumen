#!/bin/bash

# Exit on any error
set -e

# Default values
NAMESPACE="nexus-lumen-act-ns"
SERVICE_NAME="act-requestidentifier-service"
TASK_QUEUE="nexus-task-queue"
DESCRIPTION_FILE="./service_actrequestnexus.md"

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --namespace)
            NAMESPACE="$2"
            shift 2
            ;;
        --service-name)
            SERVICE_NAME="$2"
            shift 2
            ;;
        --task-queue)
            TASK_QUEUE="$2"
            shift 2
            ;;
        --description-file)
            DESCRIPTION_FILE="$2"
            shift 2
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

# Verify required files exist
if [ ! -f "$DESCRIPTION_FILE" ]; then
    echo "Error: Service description file not found: $DESCRIPTION_FILE"
    exit 1
fi

if [ ! -f "nexus-service-config.yaml" ]; then
    echo "Error: Service configuration file not found: nexus-service-config.yaml"
    exit 1
fi

# Create namespace if it doesn't exist
echo "Ensuring namespace exists: $NAMESPACE"
temporal operator namespace ensure --name "$NAMESPACE"

# Deploy the Nexus service
echo "Deploying Nexus service: $SERVICE_NAME"
temporal operator nexus endpoint create \
    --name "$SERVICE_NAME" \
    --target-namespace "$NAMESPACE" \
    --target-task-queue "$TASK_QUEUE" \
    --description-file "$DESCRIPTION_FILE"

# Apply the service configuration
echo "Applying service configuration"
kubectl apply -f nexus-service-config.yaml

# Wait for deployment to complete
echo "Waiting for deployment to complete..."
kubectl wait --for=condition=available --timeout=300s deployment/$SERVICE_NAME -n $NAMESPACE

echo "Deployment completed successfully!"
echo "Service endpoint: $SERVICE_NAME.$NAMESPACE.svc.cluster.local"
echo "Task queue: $TASK_QUEUE"

# Display service status
echo "Service status:"
temporal operator nexus endpoint describe --name "$SERVICE_NAME" --namespace "$NAMESPACE" 