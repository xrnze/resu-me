import type { AnalysisRequest, AnalysisResponse } from "@/types";
import {
  ERROR_MESSAGES,
  NETWORK_ERROR_MESSAGE,
  STATUS_MESSAGES,
  UNKNOWN_ERROR_MESSAGE,
} from "./error-messages";

const USE_MOCK = false;
const MOCK_DELAY = 1500;

const MOCK_RESPONSE: AnalysisResponse = {
  score: 78,
  missing_keywords: ["Kubernetes", "CI/CD", "TypeScript"],
  section_feedback: {
    summary:
      "Your summary is too generic and lacks quantifiable impact. Focus on specific outcomes rather than responsibilities.",
    experience:
      "Good use of metrics in your experience section. Consider adding more context around team size and project scope.",
    skills:
      "Missing several technical keywords that appear in the job description. Add relevant technologies to increase your ATS match rate.",
  },
  rewrite_suggestions: [
    "Change 'Responsible for' to 'Led a team of 5 engineers to deliver a 30% reduction in deployment time'",
    'Add specific metrics to your second bullet point under Experience (e.g., "Improved API response time by 40%")',
    "Include Kubernetes experience in your Skills section — it appears prominently in the job description",
  ],
};

export async function analyzeResume(
  resumeText: string,
  jobDescription: string,
  signal?: AbortSignal,
): Promise<AnalysisResponse> {
  if (USE_MOCK) {
    await new Promise((r) => setTimeout(r, MOCK_DELAY));
    return MOCK_RESPONSE;
  }

  const body: AnalysisRequest = {
    resume_text: resumeText,
    job_description: jobDescription,
  };

  const backendUrl = import.meta.env.VITE_BACKEND_URL;
  const url = backendUrl ? `http://${backendUrl}/` : "";

  let res: Response;

  try {
    res = await fetch(url + "api/analyze", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body),
      signal,
    });
  } catch (e) {
    if (e instanceof DOMException && e.name === "AbortError") throw e;
    throw new Error(NETWORK_ERROR_MESSAGE, { cause: e });
  }

  if (!res.ok) {
    const body = await res.json().catch(() => ({}));
    const code = body?.error?.code ?? "";
    const msg =
      ERROR_MESSAGES[code] ??
      STATUS_MESSAGES[res.status] ??
      body?.error?.message ??
      UNKNOWN_ERROR_MESSAGE;
    throw new Error(msg);
  }

  return res.json() as Promise<AnalysisResponse>;
}
