# File Storage System

## Overview

The LMS system uses MinIO as an object storage solution for managing files like course materials, lesson documents, and other media files.

## MinIO Configuration

MinIO is configured with the following default settings:

- **Endpoint**: localhost:9000 (API)
- **Console**: localhost:9001 (Web UI)
- **Access Key**: minioadmin
- **Secret Key**: minioadmin123
- **Bucket**: lms-files

## API Endpoints

### User Endpoints (Authentication Required)

#### Download File
```bash
GET /api/v1/user/files/download?path=<file_path>
```

#### Get File URL
```bash
GET /api/v1/user/files/url?path=<file_path>
```

Returns a presigned URL valid for 1 hour.

### Admin Endpoints (Admin Role Required)

#### Upload General File
```bash
POST /api/v1/admin/files/upload
Content-Type: multipart/form-data

Parameters:
- file: The file to upload
- type: File type/category (optional, defaults to "general")
```

#### Upload Course File
```bash
POST /api/v1/admin/files/courses/{courseId}/upload
Content-Type: multipart/form-data

Parameters:
- file: The file to upload
```

#### Upload Lesson File
```bash
POST /api/v1/admin/files/lessons/{lessonId}/upload
Content-Type: multipart/form-data

Parameters:
- file: The file to upload
```

#### Delete File
```bash
DELETE /api/v1/admin/files
Content-Type: application/json

Body:
{
  "file_path": "path/to/file.pdf"
}
```

## File Organization

Files are organized in the following structure:

```
lms-files/
├── general/          # General files
├── courses/          # Course-specific files
│   ├── 1/           # Course ID 1
│   ├── 2/           # Course ID 2
│   └── ...
└── lessons/          # Lesson-specific files
    ├── 1/           # Lesson ID 1
    ├── 2/           # Lesson ID 2
    └── ...
```

## Usage Examples

### Upload a Course Material
```bash
curl -X POST \
  http://localhost:8082/api/v1/admin/files/courses/1/upload \
  -H "Authorization: Bearer <token>" \
  -F "file=@course_material.pdf"
```

### Get Download URL for a File
```bash
curl -X GET \
  "http://localhost:8082/api/v1/user/files/url?path=courses/1/course_material_123456789.pdf" \
  -H "Authorization: Bearer <token>"
```

### Direct File Download
```bash
curl -X GET \
  "http://localhost:8082/api/v1/user/files/download?path=courses/1/course_material_123456789.pdf" \
  -H "Authorization: Bearer <token>" \
  -o downloaded_file.pdf
```

## MinIO Web Console

You can access the MinIO web console at http://localhost:9001 to:
- Browse files
- Create/delete buckets
- Manage access policies
- Monitor storage usage

Login with:
- Username: minioadmin
- Password: minioadmin123

## Docker Compose Setup

The MinIO service is automatically started with docker-compose. The bucket "lms-files" is created automatically on startup.

## Environment Variables

Configure MinIO in your `.env` file:

```env
MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY_ID=minioadmin
MINIO_SECRET_ACCESS_KEY=minioadmin123
MINIO_USE_SSL=false
MINIO_BUCKET_NAME=lms-files
```

## Security Considerations

1. **Change default credentials** in production
2. Use HTTPS/TLS for MinIO in production
3. Configure proper bucket policies
4. Implement file type validation
5. Set appropriate file size limits
6. Use virus scanning for uploaded files in production