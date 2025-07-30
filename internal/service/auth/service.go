package auth

import (
	"github.com/sirupsen/logrus"
	"lms_system/internal/domain"
	"lms_system/internal/infrastructure/clients/http/keycloak"
)

type Service struct {
	mainRepo       domain.MainRepositoryInterface
	logger         *logrus.Logger
	keycloakClient *keycloak.Client
}

func NewService(mainRepo domain.MainRepositoryInterface, logger *logrus.Logger, keycloakClient *keycloak.Client) *Service {
	return &Service{
		mainRepo:       mainRepo,
		logger:         logger,
		keycloakClient: keycloakClient,
	}
}
