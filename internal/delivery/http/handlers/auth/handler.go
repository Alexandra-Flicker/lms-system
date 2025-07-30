package auth

import (
	"lms_system/internal/domain"
)

type Handler struct {
	service domain.AuthServiceInterface
}

func NewHandler(service domain.AuthServiceInterface) *Handler {
	return &Handler{
		service: service,
	}
}