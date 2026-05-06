package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalysisRequest_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    AnalysisRequest
		wantErr bool
	}{
		{
			name:  "valid request",
			input: `{"resume_text":"experienced Go developer","job_description":"senior backend role"}`,
			want: AnalysisRequest{
				ResumeText:     "experienced Go developer",
				JobDescription: "senior backend role",
			},
		},
		{
			name:  "partial fields",
			input: `{"resume_text":"hello"}`,
			want:  AnalysisRequest{ResumeText: "hello"},
		},
		{
			name:  "empty object",
			input: `{}`,
			want:  AnalysisRequest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req AnalysisRequest
			err := json.Unmarshal([]byte(tt.input), &req)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, req)
		})
	}
}

func TestAnalysisRequest_Validate(t *testing.T) {
	longStr := make([]byte, MaxInputLength+1)
	for i := range longStr {
		longStr[i] = 'a'
	}

	tests := []struct {
		name    string
		req     AnalysisRequest
		wantErr string
	}{
		{name: "valid", req: AnalysisRequest{ResumeText: "dev", JobDescription: "role"}},
		{name: "empty resume", req: AnalysisRequest{ResumeText: "", JobDescription: "role"}, wantErr: "resume_text is required"},
		{name: "empty job", req: AnalysisRequest{ResumeText: "dev", JobDescription: ""}, wantErr: "job_description is required"},
		{name: "whitespace resume", req: AnalysisRequest{ResumeText: "   ", JobDescription: "role"}, wantErr: "resume_text is required"},
		{name: "whitespace job", req: AnalysisRequest{ResumeText: "dev", JobDescription: " \t\n "}, wantErr: "job_description is required"},
		{name: "resume over max length", req: AnalysisRequest{ResumeText: string(longStr), JobDescription: "role"}, wantErr: "resume_text exceeds maximum length"},
		{name: "job over max length", req: AnalysisRequest{ResumeText: "dev", JobDescription: string(longStr)}, wantErr: "job_description exceeds maximum length"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}

func TestAnalysisResponse_Marshal(t *testing.T) {
	resp := AnalysisResponse{
		Score:           78,
		MissingKeywords: []string{"Kubernetes", "CI/CD"},
		SectionFeedback: SectionFeedback{
			Summary:    "Good summary",
			Experience: "Needs metrics",
			Skills:     "Missing keywords",
		},
		RewriteSuggestions: []string{"Add metrics", "Use action verbs"},
	}
	data, err := json.Marshal(resp)
	assert.NoError(t, err)

	var decoded map[string]any
	err = json.Unmarshal(data, &decoded)
	assert.NoError(t, err)

	assert.Equal(t, float64(78), decoded["score"])
	kw := decoded["missing_keywords"].([]any)
	assert.Equal(t, 2, len(kw))
	assert.Equal(t, "Kubernetes", kw[0])
	sf := decoded["section_feedback"].(map[string]any)
	assert.Equal(t, "Good summary", sf["summary"])
	rs := decoded["rewrite_suggestions"].([]any)
	assert.Equal(t, 2, len(rs))
}

func TestErrorResponse_Marshal(t *testing.T) {
	t.Run("with details", func(t *testing.T) {
		apiErr := APIError{
			Error: ErrorResponse{
				Code:    "VALIDATION_ERROR",
				Message: "bad input",
				Details: map[string]string{"field": "resume_text"},
			},
		}
		data, err := json.Marshal(apiErr)
		assert.NoError(t, err)

		var decoded map[string]any
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)

		e := decoded["error"].(map[string]any)
		assert.Equal(t, "VALIDATION_ERROR", e["code"])
		assert.NotNil(t, e["details"])
	})

	t.Run("without details", func(t *testing.T) {
		apiErr := APIError{
			Error: ErrorResponse{
				Code:    "INTERNAL_ERROR",
				Message: "oops",
			},
		}
		data, err := json.Marshal(apiErr)
		assert.NoError(t, err)

		var decoded map[string]any
		err = json.Unmarshal(data, &decoded)
		assert.NoError(t, err)

		assert.NotContains(t, decoded["error"].(map[string]any), "details")
	})
}
