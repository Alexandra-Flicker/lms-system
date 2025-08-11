#!/bin/bash

# Test file upload functionality

echo "Testing MinIO file upload..."

# Start the services
echo "Starting docker-compose services..."
docker-compose up -d

# Wait for services to be ready
echo "Waiting for services to be ready..."
sleep 10

# Create a test file
echo "Creating test file..."
echo "This is a test file for MinIO upload" > test.txt

# Test general file upload
echo "Testing general file upload..."
curl -X POST \
  http://localhost:8082/api/v1/admin/files/upload \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -F "file=@test.txt" \
  -F "type=general"

# Clean up
rm test.txt

echo "Test completed!"