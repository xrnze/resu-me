export interface AnalysisRequest {
  resume_text: string;
  job_description: string;
}

export interface AnalysisResponse {
  score: number;
  missing_keywords: string[];
  section_feedback: {
    summary: string;
    experience: string;
    skills: string;
  };
  rewrite_suggestions: string[];
}
