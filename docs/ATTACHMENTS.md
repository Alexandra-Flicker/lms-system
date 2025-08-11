# Lesson Attachments System

## Overview

The LMS system supports file attachments for lessons. Teachers and administrators can upload files that students can download if they have access to the lesson.

## Features

- Upload multiple files per lesson
- Access control based on course enrollment
- Support for various file types (PDF, documents, images, etc.)
- Automatic file storage in MinIO
- Presigned URLs for secure downloads

## Database Schema

```sql
CREATE TABLE attachments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    lesson_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_lesson
        FOREIGN KEY (lesson_id) 
        REFERENCES lessons(id) 
        ON DELETE CASCADE
);
```

## API Endpoints

### Upload Attachment

**Endpoint:** `POST /api/v1/attachments/lessons/{lessonId}/upload`

**Access:** ROLE_ADMIN, ROLE_TEACHER

**Request:**
```bash
curl -X POST \
  http://localhost:8082/api/v1/attachments/lessons/1/upload \
  -H "Authorization: Bearer <token>" \
  -F "file=@document.pdf"
```

**Response:**
```json
{
  "id": 1,
  "name": "document.pdf",
  "url": "lessons/1/document_1234567890.pdf",
  "lesson_id": 1,
  "created_at": "2025-01-11T10:00:00Z",
  "updated_at": "2025-01-11T10:00:00Z"
}
```

### Get Lesson Attachments

**Endpoint:** `GET /api/v1/user/lessons/{lessonId}/attachments`

**Access:** Any authenticated user

**Request:**
```bash
curl -X GET \
  http://localhost:8082/api/v1/user/lessons/1/attachments \
  -H "Authorization: Bearer <token>"
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "document.pdf",
    "url": "lessons/1/document_1234567890.pdf",
    "lesson_id": 1,
    "created_at": "2025-01-11T10:00:00Z",
    "updated_at": "2025-01-11T10:00:00Z"
  }
]
```

### Download Attachment

**Endpoint:** `GET /api/v1/user/attachments/{attachmentId}/download`

**Access:** Users with access to the lesson (enrolled in course)

**Request:**
```bash
curl -X GET \
  http://localhost:8082/api/v1/user/attachments/1/download \
  -H "Authorization: Bearer <token>" \
  -L \
  -o downloaded_file.pdf
```

**Response:** HTTP 307 redirect to presigned MinIO URL

### Delete Attachment

**Endpoint:** `DELETE /api/v1/attachments/{attachmentId}`

**Access:** ROLE_ADMIN only

**Request:**
```bash
curl -X DELETE \
  http://localhost:8082/api/v1/attachments/1 \
  -H "Authorization: Bearer <token>"
```

**Response:**
```json
{
  "message": "Attachment deleted successfully"
}
```

## Access Control

1. **Upload**: Only ROLE_ADMIN and ROLE_TEACHER can upload attachments
2. **Download**: Only users enrolled in the course (with UserCourseAccess.Unlocked = true) can download
3. **Delete**: Only ROLE_ADMIN can delete attachments
4. **List**: Any authenticated user can list attachments for a lesson

## File Storage

Files are stored in MinIO with the following structure:

```
lms-files/
└── lessons/
    ├── 1/
    │   ├── document_1234567890.pdf
    │   └── slides_1234567891.pptx
    ├── 2/
    │   └── homework_1234567892.docx
    └── ...
```

## Security Considerations

1. Files are not publicly accessible
2. Download URLs are presigned and expire after 1 hour
3. Access is checked before generating download URLs
4. File size limit is 50MB per upload
5. Original filenames are preserved in the database

## Implementation Notes

1. When an attachment is deleted, both the database record and the MinIO file are removed
2. If MinIO deletion fails, the operation still succeeds (logged as warning)
3. User ID conversion from Keycloak string ID to numeric ID needs implementation
4. File uploads are processed synchronously

## Future Enhancements

1. Implement proper user ID mapping between Keycloak and database
2. Add file type restrictions
3. Add virus scanning for uploaded files
4. Implement async file processing for large files
5. Add file preview functionality
6. Support for batch uploads
7. File versioning support