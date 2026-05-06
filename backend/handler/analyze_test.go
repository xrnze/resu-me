package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"resu-me/model"
)

type mockAnalyzer struct {
	resp *model.AnalysisResponse
	err  error
}

func (m *mockAnalyzer) Analyze(ctx context.Context, resumeText, jobDescription string) (*model.AnalysisResponse, error) {
	return m.resp, m.err
}

func TestAnalyzeHandler_ValidRequest(t *testing.T) {
	mock := &mockAnalyzer{
		resp: &model.AnalysisResponse{
			Score:           78,
			MissingKeywords: []string{"Kubernetes"},
			SectionFeedback: model.SectionFeedback{
				Summary:    "Good",
				Experience: "Needs work",
				Skills:     "Missing keywords",
			},
			RewriteSuggestions: []string{"Add metrics"},
		},
	}
	handler := NewAnalyzeHandler(mock)

	body := `{"resume_text":"experienced Go developer","job_description":"senior backend role"}`
	req := httptest.NewRequest("POST", "/api/analyze", bytes.NewReader([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result model.AnalysisResponse
	err := json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, 78, result.Score)
}

func TestAnalyzeHandler_MissingResumeText(t *testing.T) {
	mock := &mockAnalyzer{}
	handler := NewAnalyzeHandler(mock)

	body := `{"job_description":"senior backend role"}`
	req := httptest.NewRequest("POST", "/api/analyze", bytes.NewReader([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)

	var apiErr model.APIError
	json.NewDecoder(resp.Body).Decode(&apiErr)
	assert.Equal(t, "VALIDATION_ERROR", apiErr.Error.Code)
}

func TestAnalyzeHandler_MissingJobDescription(t *testing.T) {
	mock := &mockAnalyzer{}
	handler := NewAnalyzeHandler(mock)

	body := `{"resume_text":"experienced Go developer"}`
	req := httptest.NewRequest("POST", "/api/analyze", bytes.NewReader([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}

func TestAnalyzeHandler_OverMaxLength(t *testing.T) {
	mock := &mockAnalyzer{}
	handler := NewAnalyzeHandler(mock)

	longStr := make([]byte, 50001)
	for i := range longStr {
		longStr[i] = 'a'
	}
	body := `{"resume_text":"` + string(longStr) + `","job_description":"role"}`
	req := httptest.NewRequest("POST", "/api/analyze", bytes.NewReader([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}

func TestAnalyzeHandler_MalformedJSON(t *testing.T) {
	mock := &mockAnalyzer{}
	handler := NewAnalyzeHandler(mock)

	body := `not json`
	req := httptest.NewRequest("POST", "/api/analyze", bytes.NewReader([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var apiErr model.APIError
	json.NewDecoder(resp.Body).Decode(&apiErr)
	assert.Equal(t, "BAD_REQUEST", apiErr.Error.Code)
}

func TestAnalyzeHandler_XSSInInput(t *testing.T) {
	mock := &mockAnalyzer{}
	handler := NewAnalyzeHandler(mock)

	body := `{"resume_text":"<script>alert('x')</script>","job_description":"role"}`
	req := httptest.NewRequest("POST", "/api/analyze", bytes.NewReader([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}

func TestAnalyzeHandler_SQLInjection(t *testing.T) {
	mock := &mockAnalyzer{}
	handler := NewAnalyzeHandler(mock)

	body := `{"resume_text":"experienced Go developer","job_description":"'; DROP TABLE users; --"}`
	req := httptest.NewRequest("POST", "/api/analyze", bytes.NewReader([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}

func TestAnalyzeHandler_LLMError(t *testing.T) {
	mock := &mockAnalyzer{err: errors.New("LLM call failed")}
	handler := NewAnalyzeHandler(mock)

	body := `{"resume_text":"experienced Go developer","job_description":"senior backend role"}`
	req := httptest.NewRequest("POST", "/api/analyze", bytes.NewReader([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusBadGateway, resp.StatusCode)

	var apiErr model.APIError
	json.NewDecoder(resp.Body).Decode(&apiErr)
	assert.Equal(t, "LLM_ERROR", apiErr.Error.Code)
}
