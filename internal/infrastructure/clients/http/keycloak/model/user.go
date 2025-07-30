package model

type UserRepresentation struct {
	ID            string                 `json:"id,omitempty"`
	Username      string                 `json:"username"`
	Email         string                 `json:"email"`
	FirstName     string                 `json:"firstName,omitempty"`
	LastName      string                 `json:"lastName,omitempty"`
	Enabled       bool                   `json:"enabled"`
	EmailVerified bool                   `json:"emailVerified"`
	Attributes    map[string][]string    `json:"attributes,omitempty"`
	Credentials   []CredentialRepresentation `json:"credentials,omitempty"`
	RealmRoles    []string               `json:"realmRoles,omitempty"`
}

type CredentialRepresentation struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Temporary bool   `json:"temporary"`
}