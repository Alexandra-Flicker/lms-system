package keycloak

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"lms_system/internal/infrastructure/clients/http/keycloak/model"
	"net/http"
)

func (c *Client) GetUserByID(ctx context.Context, userID string) (*model.UserRepresentation, error) {
	// Get admin token
	adminToken, err := c.GetAdminToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get admin token: %w", err)
	}

	endpoint := fmt.Sprintf("%s/admin/realms/%s/users/%s", c.host, c.realm, userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+adminToken.AccessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("keycloak get user error: status %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var user model.UserRepresentation
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *Client) UpdateUser(ctx context.Context, userID string, user *model.UserRepresentation) error {
	// Get admin token
	adminToken, err := c.GetAdminToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get admin token: %w", err)
	}

	endpoint := fmt.Sprintf("%s/admin/realms/%s/users/%s", c.host, c.realm, userID)

	// Ensure we don't try to update the ID
	user.ID = ""

	jsonData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken.AccessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("keycloak update user error: status %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}