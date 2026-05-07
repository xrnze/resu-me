package middleware

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readingHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := io.Copy(io.Discard, r.Body)
		if err != nil {
			var maxBytesErr *http.MaxBytesError
			if errors.As(err, &maxBytesErr) {
				w.WriteHeader(http.StatusRequestEntityTooLarge)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	})
}

func TestBodyLimit_UnderLimit(t *testing.T) {
	middleware := BodyLimit(100 * 1024)
	handler := middleware(readingHandler())

	body := bytes.NewReader([]byte(`{"key":"value"}`))
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBodyLimit_OverLimit(t *testing.T) {
	middleware := BodyLimit(100 * 1024)
	handler := middleware(readingHandler())

	body := bytes.NewReader(make([]byte, 200*1024))
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusRequestEntityTooLarge, w.Code)
}

func TestBodyLimit_ExactLimit(t *testing.T) {
	middleware := BodyLimit(100 * 1024)
	handler := middleware(readingHandler())

	body := bytes.NewReader(make([]byte, 100*1024))
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
