package lms

import (
	"lms_system/internal/domain"

	"github.com/sirupsen/logrus"
)

type Service struct {
	mainRepo domain.MainRepositoryInterface
	logger   *logrus.Logger
}

func NewService(mainRepo domain.MainRepositoryInterface, logger *logrus.Logger) *Service {
	return &Service{
		mainRepo: mainRepo,
		logger:   logger,
	}
}
