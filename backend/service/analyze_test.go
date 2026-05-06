package service

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"resu-me/model"
)

type mockLLMClient struct {
	response string
	err      error
}

func (m *mockLLMClient) Chat(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	return m.response, m.err
}

func TestBuildSystemPrompt(t *testing.T) {
	prompt := BuildSystemPrompt()
	assert.True(t, strings.Contains(prompt, "JSON"), "system prompt should mention JSON")
	assert.True(t, strings.Contains(prompt, "score"), "system prompt should mention score")
	assert.True(t, strings.Contains(prompt, "missing_keywords"), "system prompt should mention missing_keywords")
	assert.True(t, strings.Contains(prompt, "section_feedback"), "system prompt should mention section_feedback")
	assert.True(t, strings.Contains(prompt, "rewrite_suggestions"), "system prompt should mention rewrite_suggestions")
}

func TestBuildUserPrompt(t *testing.T) {
	resume := "Experienced Go developer"
	jd := "Senior backend role"
	prompt := BuildUserPrompt(resume, jd)
	assert.True(t, strings.Contains(prompt, resume), "user prompt should include resume text")
	assert.True(t, strings.Contains(prompt, jd), "user prompt should include job description")
}

func TestAnalyze_ValidResponse(t *testing.T) {
	mock := &mockLLMClient{
		response: `{
			"score": 78,
			"missing_keywords": ["Kubernetes", "CI/CD"],
			"section_feedback": {
				"summary": "Good summary",
				"experience": "Needs work",
				"skills": "Missing keywords"
			},
			"rewrite_suggestions": ["Add metrics", "Use action verbs"]
		}`,
	}
	svc := &AnalyzeService{client: mock}
	resp, err := svc.Analyze(context.Background(), "resume", "job desc")
	assert.NoError(t, err)
	assert.Equal(t, 78, resp.Score)
	assert.Equal(t, 2, len(resp.MissingKeywords))
	assert.Equal(t, "Kubernetes", resp.MissingKeywords[0])
	assert.Equal(t, "Good summary", resp.SectionFeedback.Summary)
	assert.Equal(t, 2, len(resp.RewriteSuggestions))
}

func TestAnalyze_InvalidJSON(t *testing.T) {
	mock := &mockLLMClient{response: "not json at all"}
	svc := &AnalyzeService{client: mock}
	_, err := svc.Analyze(context.Background(), "resume", "job desc")
	assert.Error(t, err, "expected error for invalid JSON")
}

func TestAnalyze_MissingFields(t *testing.T) {
	mock := &mockLLMClient{
		response: `{"score": 50}`,
	}
	svc := &AnalyzeService{client: mock}
	_, err := svc.Analyze(context.Background(), "resume", "job desc")
	assert.Error(t, err, "expected error for missing fields")
}

func TestAnalyze_ScoreOutOfRange(t *testing.T) {
	tests := []struct {
		name  string
		score int
	}{
		{"negative score", -1},
		{"over 100", 101},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := model.AnalysisResponse{Score: tt.score}
			assert.Error(t, validateResponse(&resp), "expected error for score %d", tt.score)
		})
	}
}

func TestAnalyze_ScoreInRange(t *testing.T) {
	tests := []int{0, 50, 100}
	for _, score := range tests {
		resp := model.AnalysisResponse{
			Score:               score,
			MissingKeywords:     []string{},
			SectionFeedback:     model.SectionFeedback{Summary: "a", Experience: "b", Skills: "c"},
			RewriteSuggestions:  []string{},
		}
		assert.NoError(t, validateResponse(&resp), "unexpected error for score %d", score)
	}
}

func TestAnalyze_LLMError(t *testing.T) {
	mock := &mockLLMClient{err: assertError("network error")}
	svc := &AnalyzeService{client: mock}
	_, err := svc.Analyze(context.Background(), "resume", "job desc")
	assert.Error(t, err, "expected error from LLM")
}

type assertError string

func (e assertError) Error() string { return string(e) }
