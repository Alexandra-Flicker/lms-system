package keycloak

import "net/http"

type Client struct {
	host         string
	realm        string
	clientID     string
	clientSecret string
	adminUser    string
	adminPass    string
	httpClient   *http.Client
}

func NewClient(host, realm, clientID, clientSecret, adminUser, adminPass string) *Client {
	return &Client{
		host:         host,
		realm:        realm,
		clientID:     clientID,
		clientSecret: clientSecret,
		adminUser:    adminUser,
		adminPass:    adminPass,
		httpClient:   &http.Client{},
	}
}
