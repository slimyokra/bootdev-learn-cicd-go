package auth

import (
	"net/http"
	"testing"
)

func TestNoAuthorizationHeader(t *testing.T) {
	headers := http.Header{}
	if _, err := GetAPIKey(headers); err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected %v error, got no error", err)
	}
}

func TestApiKeyNotPassedInAuthorizationHeader(t *testing.T) {
	header := http.Header{
		"Authorization": []string{"abc"},
	}

	if _, err := GetAPIKey(header); err == nil {
		t.Errorf("expected %v error, got no error", err)
	}
}
