package keycloak

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"lms_system/internal/infrastructure/clients/http/keycloak/model"
	"net/http"
	"strings"
)

func (c *Client) CreateUser(ctx context.Context, user *model.UserRepresentation) (string, error) {
	// Get admin token
	adminToken, err := c.GetAdminToken(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get admin token: %w", err)
	}

	endpoint := fmt.Sprintf("%s/admin/realms/%s/users", c.host, c.realm)

	jsonData, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken.AccessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("keycloak create user error: status %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Extract user ID from Location header
	location := resp.Header.Get("Location")
	parts := strings.Split(location, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1], nil
	}

	return "", fmt.Errorf("failed to extract user ID from response")
}

// AssignRoleToUser assigns a realm role to a user
func (c *Client) AssignRoleToUser(ctx context.Context, userID string, roleName string) error {
	// Get admin token
	adminToken, err := c.GetAdminToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get admin token: %w", err)
	}

	// First, get the role representation
	roleEndpoint := fmt.Sprintf("%s/admin/realms/%s/roles/%s", c.host, c.realm, roleName)
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, roleEndpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+adminToken.AccessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to get role: status %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var role map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&role); err != nil {
		return err
	}

	// Assign the role to the user
	assignEndpoint := fmt.Sprintf("%s/admin/realms/%s/users/%s/role-mappings/realm", c.host, c.realm, userID)
	
	roles := []map[string]interface{}{role}
	jsonData, err := json.Marshal(roles)
	if err != nil {
		return err
	}

	req, err = http.NewRequestWithContext(ctx, http.MethodPost, assignEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken.AccessToken)

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to assign role: status %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}