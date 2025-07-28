package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "lms_system/docs"
	"lms_system/internal/delivery/http/handlers/chapter"
	"lms_system/internal/delivery/http/handlers/course"
	"lms_system/internal/delivery/http/handlers/lesson"
	"lms_system/internal/delivery/http/middleware"
	"lms_system/internal/domain"
	"lms_system/internal/domain/common"
)

func NewRouter(service domain.ServiceInterface) *chi.Mux {
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

	router.Route("/api/v1", func(r chi.Router) {
		// Public routes (no authentication required)
		r.Route("/public", func(r chi.Router) {
			r.Get("/courses", courseHandler.GetAllCourses)
			r.Get("/courses/{id}", courseHandler.GetCourseById)
			r.Get("/courses/{id}/chapters", courseHandler.GetChaptersInfoByCourseId)
		})

		// User routes (authentication required)
		r.Route("/user", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)

			r.Route("/courses", func(r chi.Router) {
				r.Get("/", courseHandler.GetAllCourses)
				r.Get("/{id}", courseHandler.GetCourseById)
				r.Post("/buy", courseHandler.BuyCourse)
				r.Get("/{id}/chapters", courseHandler.GetChaptersInfoByCourseId)
			})

			r.Route("/lessons", func(r chi.Router) {
				r.Get("/{id}", lessonHandler.GetLessonById)
			})
		})

		// Admin routes (authentication + admin role required)
		r.Route("/admin", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Use(middleware.OnlyRoles(common.RoleAdmin))

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
		})
	})

	// Swagger documentation
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	// Serve static swagger files
	router.Handle("/swagger/swagger.yaml", http.FileServer(http.Dir("./docs/")))

	return router
}
