package lms

import (
	"lms_system/internal/domain"

	"github.com/sirupsen/logrus"
)

type Service struct {
	repo        domain.MainRepositoryInterface
	logger      *logrus.Logger
	fileService domain.FileServiceInterface
}

func NewService(repo domain.MainRepositoryInterface, logger *logrus.Logger, fileService domain.FileServiceInterface) *Service {
	return &Service{
		repo:        repo,
		logger:      logger,
		fileService: fileService,
	}
}
