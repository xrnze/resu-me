package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"resu-me/model"
	"resu-me/sanitizer"
)

type Analyzer interface {
	Analyze(ctx context.Context, resumeText, jobDescription string) (*model.AnalysisResponse, error)
}

type AnalyzeHandler struct {
	analyzer Analyzer
}

func NewAnalyzeHandler(analyzer Analyzer) *AnalyzeHandler {
	return &AnalyzeHandler{analyzer: analyzer}
}

func (h *AnalyzeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Only POST is accepted")
		return
	}

	var req model.AnalysisRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			writeError(w, http.StatusRequestEntityTooLarge, "PAYLOAD_TOO_LARGE", "Request body exceeds the maximum allowed size")
			return
		}
		writeError(w, http.StatusBadRequest, "BAD_REQUEST", "Request body is not valid JSON")
		return
	}

	if err := req.Validate(); err != nil {
		writeError(w, http.StatusUnprocessableEntity, "VALIDATION_ERROR", err.Error())
		return
	}

	if err := sanitizer.Sanitize(req.ResumeText); err != nil {
		log.Printf("sanitization rejection on resume_text: %v", err)
		writeError(w, http.StatusUnprocessableEntity, "VALIDATION_ERROR", "resume_text contains disallowed patterns")
		return
	}
	if err := sanitizer.Sanitize(req.JobDescription); err != nil {
		log.Printf("sanitization rejection on job_description: %v", err)
		writeError(w, http.StatusUnprocessableEntity, "VALIDATION_ERROR", "job_description contains disallowed patterns")
		return
	}

	resp, err := h.analyzer.Analyze(r.Context(), req.ResumeText, req.JobDescription)
	if err != nil {
		log.Printf("analysis error: %v", err)
		writeError(w, http.StatusBadGateway, "LLM_ERROR", "Failed to communicate with the analysis provider.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(model.APIError{
		Error: model.ErrorResponse{
			Code:    code,
			Message: message,
		},
	})
}
