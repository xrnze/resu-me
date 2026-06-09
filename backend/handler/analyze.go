package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"resu-me/model"
	"resu-me/sanitizer"
	"resu-me/service"
)

type Analyzer interface {
	Analyze(ctx context.Context, resumeText, jobDescription string) (*model.AnalysisResponse, error)
}

type Injector interface {
	Check(ctx context.Context, resumeText, jobDesc string) error
}

type AnalyzeHandler struct {
	analyzer Analyzer
	injector Injector
}

func NewAnalyzeHandler(analyzer Analyzer, injector Injector) *AnalyzeHandler {
	return &AnalyzeHandler{analyzer: analyzer, injector: injector}
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
		log.Printf("%v", err)
		writeError(w, http.StatusUnprocessableEntity, "VALIDATION_ERROR", "resume_text contains disallowed patterns")
		return
	}
	if err := sanitizer.Sanitize(req.JobDescription); err != nil {
		log.Printf("%v", err)
		writeError(w, http.StatusUnprocessableEntity, "VALIDATION_ERROR", "job_description contains disallowed patterns")
		return
	}

	if err := h.injector.Check(r.Context(), req.ResumeText, req.JobDescription); err != nil {
		if errors.Is(err, service.ErrInjectionDetected) {
			log.Printf("%v", err)
			writeError(w, http.StatusUnprocessableEntity, "PROMPT_INJECTION_DETECTED", err.Error())
			return
		}
		log.Printf("%v", err)
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Input validation failed. Please try again.")
		return
	}

	resp, err := h.analyzer.Analyze(r.Context(), req.ResumeText, req.JobDescription)
	if err != nil {
		log.Printf("%v", err)
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
