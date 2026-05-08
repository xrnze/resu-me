package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"resu-me/model"
)

type LLMClient interface {
	Chat(ctx context.Context, prompt string) (string, error)
}

type AnalyzeService struct {
	client  LLMClient
	timeout time.Duration
}

func NewAnalyzeService(client LLMClient, timeout time.Duration) *AnalyzeService {
	return &AnalyzeService{client: client, timeout: timeout}
}

func BuildPrompt(resumeText, jobDescription string) string {
	prompt := `You are a resume analysis assistant. Analyze how well a resume matches a job description and return ONLY valid JSON with this exact schema:
{
  "score": <number 0-100>,
  "missing_keywords": ["<keyword>", ...],
  "section_feedback": {
    "summary": "<feedback>",
    "experience": "<feedback>",
    "skills": "<feedback>"
  },
  "rewrite_suggestions": ["<suggestion>", ...]
}

Resume:
{{ .RESUME_TEXT }}

Job Description:
{{ .JOB_DESCRIPTION }}

Scoring criteria:
- 0-40: Poor match, major gaps
- 41-70: Partial match, improvable
- 71-90: Strong match, minor gaps
- 91-100: Excellent match

Rules:
- Score 0-100 reflects how well the resume matches the job description.
- missing_keywords lists skills/technologies from the job description not found in the resume.
- section_feedback gives specific feedback on the summary, experience, and skills sections.
- rewrite_suggestions provides actionable, specific improvements.
- Return ONLY valid JSON. No markdown, no explanation, no code blocks.`

	prompt = strings.ReplaceAll(prompt, "{{ .RESUME_TEXT }}", resumeText)
	prompt = strings.ReplaceAll(prompt, "{{ .JOB_DESCRIPTION }}", jobDescription)
	return prompt
}

func (s *AnalyzeService) Analyze(ctx context.Context, resumeText, jobDescription string) (*model.AnalysisResponse, error) {
	callCtx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	prompt := BuildPrompt(resumeText, jobDescription)

	start := time.Now()
	raw, err := s.client.Chat(callCtx, prompt)
	if err != nil {
		return nil, fmt.Errorf("LLM call failed: %w", err)
	}
	log.Printf("LLM request completed in %v", time.Since(start))

	var resp model.AnalysisResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		return nil, fmt.Errorf("failed to parse LLM response: %w", err)
	}

	if err := validateResponse(&resp); err != nil {
		return nil, fmt.Errorf("invalid LLM response: %w", err)
	}

	return &resp, nil
}

func validateResponse(resp *model.AnalysisResponse) error {
	if resp.Score < 0 || resp.Score > 100 {
		return fmt.Errorf("score %d out of range (0-100)", resp.Score)
	}
	if resp.SectionFeedback.Summary == "" || resp.SectionFeedback.Experience == "" || resp.SectionFeedback.Skills == "" {
		return fmt.Errorf("section_feedback fields are required")
	}
	if resp.MissingKeywords == nil {
		resp.MissingKeywords = []string{}
	}
	if resp.RewriteSuggestions == nil {
		resp.RewriteSuggestions = []string{}
	}
	return nil
}
