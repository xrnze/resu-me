package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}

func TestCORS_HeadersOnPOST(t *testing.T) {
	handler := CORS(testHandler())
	req := httptest.NewRequest("POST", "/api/analyze", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, "*", resp.Header.Get("Access-Control-Allow-Origin"))
	assert.NotEmpty(t, resp.Header.Get("Access-Control-Allow-Methods"))
	assert.NotEmpty(t, resp.Header.Get("Access-Control-Allow-Headers"))
}

func TestCORS_Preflight(t *testing.T) {
	handler := CORS(testHandler())
	req := httptest.NewRequest("OPTIONS", "/api/analyze", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	assert.Equal(t, "*", resp.Header.Get("Access-Control-Allow-Origin"))
	assert.NotEmpty(t, resp.Header.Get("Access-Control-Allow-Methods"))
	assert.NotEmpty(t, resp.Header.Get("Access-Control-Allow-Headers"))
}

func TestCORS_MethodsHeader(t *testing.T) {
	handler := CORS(testHandler())
	req := httptest.NewRequest("OPTIONS", "/api/analyze", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	methods := resp.Header.Get("Access-Control-Allow-Methods")
	assert.Equal(t, "GET, POST, OPTIONS", methods)
}

func TestCORS_HeadersHeader(t *testing.T) {
	handler := CORS(testHandler())
	req := httptest.NewRequest("OPTIONS", "/api/analyze", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	headers := resp.Header.Get("Access-Control-Allow-Headers")
	assert.Equal(t, "Content-Type", headers)
}
