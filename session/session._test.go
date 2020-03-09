package session

import(
	"testing"
)

func TestGetAuthToken(t *testing.T) {
	authResponse := GetAuthToken()
	if authResponse == nil {
		t.Error("Failed to get authorization token")
	}
}