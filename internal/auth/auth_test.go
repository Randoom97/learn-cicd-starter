package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-api-key")
	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fail()
		return
	}
	if apiKey != "test-api-key" {
		t.Fail()
		return
	}
}
