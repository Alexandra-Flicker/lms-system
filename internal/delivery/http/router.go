package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "lms_system/docs"
	"lms_system/internal/delivery/http/handlers/attachment"
	"lms_system/internal/delivery/http/handlers/auth"
	"lms_system/internal/delivery/http/handlers/chapter"
	"lms_system/internal/delivery/http/handlers/course"
	"lms_system/internal/delivery/http/handlers/file"
	"lms_system/internal/delivery/http/handlers/lesson"
	"lms_system/internal/delivery/http/middleware"
	"lms_system/internal/domain"
	"lms_system/internal/domain/common"
)

func NewRouter(service domain.ServiceInterface, authService domain.AuthServiceInterface, fileService domain.FileServiceInterface) *chi.Mux {
	router := chi.NewRouter()

	// Global middleware
	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)
	router.Use(chimiddleware.RequestID)
	router.Use(chimiddleware.RealIP)

	// Create handlers
	courseHandler := course.NewHandler(service)
	chapterHandler := chapter.NewHandler(service)
	lessonHandler := lesson.NewHandler(service)
	authHandler := auth.NewHandler(authService)
	fileHandler := file.NewHandler(fileService)
	attachmentHandler := attachment.NewHandler(service)

	router.Route("/api/v1", func(r chi.Router) {
		// Auth routes
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", authHandler.Login)
			r.Post("/refresh", authHandler.Refresh)
		})

		// Public routes (no authentication required)
		r.Route("/public", func(r chi.Router) {
			r.Get("/courses", courseHandler.GetAllCourses)
			r.Get("/courses/{id}", courseHandler.GetCourseById)
			r.Get("/courses/{id}/chapters", courseHandler.GetChaptersInfoByCourseId)
		})

		// User routes (authentication required)
		r.Route("/user", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)

			// Profile management
			r.Put("/profile", authHandler.UpdateProfile)
			r.Put("/change-password", authHandler.ChangePassword)

			r.Route("/courses", func(r chi.Router) {
				r.Get("/", courseHandler.GetAllCourses)
				r.Get("/{id}", courseHandler.GetCourseById)
				r.Post("/buy", courseHandler.BuyCourse)
				r.Get("/{id}/chapters", courseHandler.GetChaptersInfoByCourseId)
			})

			r.Route("/lessons", func(r chi.Router) {
				r.Get("/{id}", lessonHandler.GetLessonById)
				r.Get("/{lessonId}/attachments", attachmentHandler.GetLessonAttachments)
			})

			// File routes for users
			r.Route("/files", func(r chi.Router) {
				r.Get("/download", fileHandler.DownloadFile)
				r.Get("/url", fileHandler.GetFileURL)
			})

			// Attachment download (requires lesson access)
			r.Get("/attachments/{attachmentId}/download", attachmentHandler.DownloadAttachment)
		})

		// Course management (admin only)
		r.Route("/course", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Use(middleware.OnlyRoles(common.RoleAdmin))

			r.Post("/", courseHandler.CreateCourse)
			r.Put("/{id}", courseHandler.UpdateCourseById)
			r.Delete("/{id}", courseHandler.DeleteCourseById)
		})

		// Chapter management (admin only)
		r.Route("/chapter", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Use(middleware.OnlyRoles(common.RoleAdmin))

			r.Post("/", chapterHandler.CreateChapterStandalone)
			r.Put("/{id}", chapterHandler.UpdateChapterById)
			r.Delete("/{id}", chapterHandler.DeleteChapterById)
		})

		// Lesson management (admin only)
		r.Route("/lesson", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Use(middleware.OnlyRoles(common.RoleAdmin))

			r.Post("/", lessonHandler.CreateLessonStandalone)
			r.Put("/{id}", lessonHandler.UpdateLessonById)
			r.Delete("/{id}", lessonHandler.DeleteLessonById)
		})

		// Attachment management (admin and teacher)
		r.Route("/attachments", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			
			// Upload (admin and teacher only)
			r.Group(func(r chi.Router) {
				r.Use(middleware.OnlyRoles(common.RoleAdmin, common.RoleTeacher))
				r.Post("/lessons/{lessonId}/upload", attachmentHandler.UploadAttachment)
			})

			// Delete (admin only)
			r.Group(func(r chi.Router) {
				r.Use(middleware.OnlyRoles(common.RoleAdmin))
				r.Delete("/{attachmentId}", attachmentHandler.DeleteAttachment)
			})
		})

		// Admin routes (authentication + admin role required)
		r.Route("/admin", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Use(middleware.OnlyRoles(common.RoleAdmin))

			// User management routes
			r.Route("/users", func(r chi.Router) {
				r.Post("/register", authHandler.RegisterUser)
			})

			r.Route("/courses", func(r chi.Router) {
				r.Post("/", courseHandler.CreateCourse)
				r.Put("/{id}", courseHandler.UpdateCourseById)
				r.Delete("/{id}", courseHandler.DeleteCourseById)

				r.Route("/{courseId}/chapters", func(r chi.Router) {
					r.Post("/", chapterHandler.CreateChapter)
				})
			})

			r.Route("/chapters", func(r chi.Router) {
				r.Put("/{id}", chapterHandler.UpdateChapterById)
				r.Delete("/{id}", chapterHandler.DeleteChapterById)

				r.Route("/{chapterId}/lessons", func(r chi.Router) {
					r.Post("/", lessonHandler.CreateLesson)
				})
			})

			r.Route("/lessons", func(r chi.Router) {
				r.Put("/{id}", lessonHandler.UpdateLessonById)
				r.Delete("/{id}", lessonHandler.DeleteLessonById)
			})

			// File management routes
			r.Route("/files", func(r chi.Router) {
				r.Post("/upload", fileHandler.UploadFile)
				r.Delete("/", fileHandler.DeleteFile)
				r.Post("/courses/{courseId}/upload", fileHandler.UploadCourseFile)
				r.Post("/lessons/{lessonId}/upload", fileHandler.UploadLessonFile)
			})
		})
	})

	// Swagger documentation
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	// Serve static swagger files
	router.Handle("/swagger/swagger.yaml", http.FileServer(http.Dir("./docs/")))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("LMS system is running"))
		if err != nil {
			return
		}
	})

	return router
}
