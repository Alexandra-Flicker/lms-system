package utils

import (
	"crypto/md5"
	"encoding/binary"
)

// ConvertKeycloakIDToUint converts Keycloak string ID to uint for internal use
// This is a simple hash-based approach for demo purposes
// In production, you'd want to maintain a mapping table in the database
func ConvertKeycloakIDToUint(keycloakID string) uint {
	if keycloakID == "" {
		return 1 // Default fallback for testing
	}
	
	// Create MD5 hash of the Keycloak ID
	hash := md5.Sum([]byte(keycloakID))
	
	// Convert first 4 bytes to uint32, then to uint
	// This ensures we get a consistent mapping for the same ID
	id := binary.BigEndian.Uint32(hash[:4])
	
	// Ensure we don't return 0
	if id == 0 {
		id = 1
	}
	
	return uint(id)
}

// For testing purposes, we can also provide a simple mapping
var testUserMappings = map[string]uint{
	"admin":                                    1,
	"32bfb3d7-5b2c-4502-b08a-92ae81984f57":   1, // Known admin ID from Keycloak
	"test-user":                               2,
	"teacher":                                 3,
}