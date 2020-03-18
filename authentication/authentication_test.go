package authentication

import (
	"testing"
)

func TestGetAuthToken(t *testing.T) {
	authResponse := GetAuthToken()
	if authResponse == nil {
		t.Error("Failed to get authorization token")
	}
}

func TestGetClient(t *testing.T) {
	client := GetClient()
	if client == nil {
		t.Error("Failed to initialize Twitter client")
	}
}
