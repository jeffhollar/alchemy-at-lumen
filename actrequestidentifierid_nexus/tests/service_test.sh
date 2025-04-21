#!/bin/bash

# Test the Nexus service endpoint
echo "Testing Act Nexus Service..."

# Example request payload
REQUEST_PAYLOAD='{
  "requestId": "test-123",
  "data": {
    "identifier": "test-identifier"
  }
}'

# Send request to the Nexus service
curl -X POST http://localhost:8080/actrequestidentifierid/processAct \
  -H "Content-Type: application/json" \
  -d "$REQUEST_PAYLOAD"

echo -e "\nTest completed." 