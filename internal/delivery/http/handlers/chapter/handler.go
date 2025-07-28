package chapter

import (
	"lms_system/internal/domain"
)

type Handler struct {
	service domain.ServiceInterface
}

func NewHandler(service domain.ServiceInterface) *Handler {
	return &Handler{
		service: service,
	}
}
