package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOpenAIClient(t *testing.T) {
	c := NewOpenAIClient("test-key", "https://example.com", "test-model")
	assert.NotNil(t, c)
	assert.Equal(t, "test-model", c.model)
}

func TestOpenAIClient_Chat_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "unauthorized"}`))
	}))
	defer server.Close()

	client := NewOpenAIClient("bad-key", server.URL, "test-model")
	_, err := client.Chat(context.Background(), "system", "user")
	assert.Error(t, err)
}
