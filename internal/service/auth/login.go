package auth

import (
	"context"
	"fmt"
	"lms_system/internal/domain/dto"
)

func (s *Service) Login(ctx context.Context, request *dto.AuthLoginRequest) (response *dto.AuthLoginResponse, err error) {
	// Get token from Keycloak
	tokenResponse, err := s.keycloakClient.GetToken(ctx, request.Username, request.Password)
	if err != nil {
		s.logger.WithError(err).Error("failed to get token from keycloak")
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	return &dto.AuthLoginResponse{
		AccessToken:  tokenResponse.AccessToken,
		RefreshToken: tokenResponse.RefreshToken,
	}, nil
}
