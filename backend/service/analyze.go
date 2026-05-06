package service

import (
	"context"
	"encoding/json"
	"fmt"
	"resu-me/model"
)

type LLMClient interface {
	Chat(ctx context.Context, systemPrompt, userPrompt string) (string, error)
}

type AnalyzeService struct {
	client LLMClient
}

func NewAnalyzeService(client LLMClient) *AnalyzeService {
	return &AnalyzeService{client: client}
}

func BuildSystemPrompt() string {
	return `You are a resume analysis assistant. Analyze how well a resume matches a job description and return ONLY valid JSON with this exact schema:
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

Rules:
- Score 0-100 reflects how well the resume matches the job description.
- missing_keywords lists skills/technologies from the job description not found in the resume.
- section_feedback gives specific feedback on the summary, experience, and skills sections.
- rewrite_suggestions provides actionable, specific improvements.
- Return ONLY valid JSON. No markdown, no explanation, no code blocks.`
}

func BuildUserPrompt(resumeText, jobDescription string) string {
	return fmt.Sprintf("Resume:\n%s\n\nJob Description:\n%s", resumeText, jobDescription)
}

func (s *AnalyzeService) Analyze(ctx context.Context, resumeText, jobDescription string) (*model.AnalysisResponse, error) {
	systemPrompt := BuildSystemPrompt()
	userPrompt := BuildUserPrompt(resumeText, jobDescription)

	raw, err := s.client.Chat(ctx, systemPrompt, userPrompt)
	if err != nil {
		return nil, fmt.Errorf("LLM call failed: %w", err)
	}

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
	return nil
}
