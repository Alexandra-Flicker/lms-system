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

func (c *Client) ChangePassword(ctx context.Context, userID string, newPassword string) error {
	// Get admin token
	adminToken, err := c.GetAdminToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get admin token: %w", err)
	}

	endpoint := fmt.Sprintf("%s/admin/realms/%s/users/%s/reset-password", c.host, c.realm, userID)

	credential := model.CredentialRepresentation{
		Type:      "password",
		Value:     newPassword,
		Temporary: false,
	}

	jsonData, err := json.Marshal(credential)
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
		return fmt.Errorf("keycloak change password error: status %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}