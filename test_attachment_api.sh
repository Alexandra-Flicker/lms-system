#!/bin/bash

# Test attachment API functionality

API_URL="http://localhost:8082/api/v1"
ADMIN_TOKEN="YOUR_ADMIN_TOKEN_HERE"
USER_TOKEN="YOUR_USER_TOKEN_HERE"

echo "=== Testing Attachment API ==="

# 1. Create a test file
echo "Creating test file..."
echo "This is a test PDF content for lesson attachment" > test_lesson.pdf

# 2. Upload attachment to lesson (as admin/teacher)
echo -e "\n1. Testing attachment upload..."
UPLOAD_RESPONSE=$(curl -X POST \
  "${API_URL}/attachments/lessons/1/upload" \
  -H "Authorization: Bearer ${ADMIN_TOKEN}" \
  -F "file=@test_lesson.pdf" \
  -s)

echo "Upload response: ${UPLOAD_RESPONSE}"

# Extract attachment ID from response (assuming JSON response with "id" field)
ATTACHMENT_ID=$(echo "${UPLOAD_RESPONSE}" | grep -o '"id":[0-9]*' | cut -d':' -f2)

# 3. Get lesson attachments
echo -e "\n2. Getting lesson attachments..."
curl -X GET \
  "${API_URL}/user/lessons/1/attachments" \
  -H "Authorization: Bearer ${USER_TOKEN}" \
  -s | jq .

# 4. Download attachment (as user with access)
echo -e "\n3. Testing attachment download..."
curl -X GET \
  "${API_URL}/user/attachments/${ATTACHMENT_ID}/download" \
  -H "Authorization: Bearer ${USER_TOKEN}" \
  -L \
  -o downloaded_lesson.pdf

if [ -f "downloaded_lesson.pdf" ]; then
  echo "File downloaded successfully!"
else
  echo "Failed to download file"
fi

# 5. Try to download without access (should fail)
echo -e "\n4. Testing download without access (should fail)..."
curl -X GET \
  "${API_URL}/user/attachments/999/download" \
  -H "Authorization: Bearer ${USER_TOKEN}" \
  -v

# 6. Delete attachment (as admin)
echo -e "\n5. Testing attachment deletion..."
curl -X DELETE \
  "${API_URL}/attachments/${ATTACHMENT_ID}" \
  -H "Authorization: Bearer ${ADMIN_TOKEN}" \
  -s | jq .

# Clean up
rm -f test_lesson.pdf downloaded_lesson.pdf

echo -e "\n=== Test completed ===

API Endpoints tested:
- POST   /api/v1/attachments/lessons/{lessonId}/upload   (Admin/Teacher only)
- GET    /api/v1/user/lessons/{lessonId}/attachments     (Any authenticated user)
- GET    /api/v1/user/attachments/{attachmentId}/download (Users with lesson access)
- DELETE /api/v1/attachments/{attachmentId}              (Admin only)

Notes:
1. Replace YOUR_ADMIN_TOKEN_HERE with a valid admin token
2. Replace YOUR_USER_TOKEN_HERE with a valid user token
3. Make sure lesson ID 1 exists and user has access to it
4. File upload limit is 50MB
"