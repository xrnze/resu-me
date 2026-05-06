package model

import (
	"fmt"
	"strings"
)

const MaxInputLength = 50000

type AnalysisRequest struct {
	ResumeText     string `json:"resume_text"`
	JobDescription string `json:"job_description"`
}

func (r *AnalysisRequest) Validate() error {
	if strings.TrimSpace(r.ResumeText) == "" {
		return fmt.Errorf("resume_text is required")
	}
	if strings.TrimSpace(r.JobDescription) == "" {
		return fmt.Errorf("job_description is required")
	}
	if len(r.ResumeText) > MaxInputLength {
		return fmt.Errorf("resume_text exceeds maximum length")
	}
	if len(r.JobDescription) > MaxInputLength {
		return fmt.Errorf("job_description exceeds maximum length")
	}
	return nil
}

type SectionFeedback struct {
	Summary    string `json:"summary"`
	Experience string `json:"experience"`
	Skills     string `json:"skills"`
}

type AnalysisResponse struct {
	Score              int             `json:"score"`
	MissingKeywords    []string        `json:"missing_keywords"`
	SectionFeedback    SectionFeedback `json:"section_feedback"`
	RewriteSuggestions []string        `json:"rewrite_suggestions"`
}

type ErrorResponse struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

type APIError struct {
	Error ErrorResponse `json:"error"`
}
