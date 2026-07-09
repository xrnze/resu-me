export const NETWORK_ERROR_MESSAGE =
  "Couldn't reach the analysis server. Check your connection and try again.";

export const UNKNOWN_ERROR_MESSAGE =
  "Something went wrong, please try again later.";

export const ERROR_MESSAGES: Record<string, string> = {
  VALIDATION_ERROR:
    "Please check your resume and job description, make sure both are filled in, and don't contain any code or special formatting.",
  PROMPT_INJECTION_DETECTED:
    "Your input contains content that isn't allowed, please remove any instructions or commands and try again.",
  INTERNAL_ERROR: "Something went wrong on our end, please try again.",
  LLM_ERROR:
    "The analysis service is temporarily unavailable, please try again in a moment.",
  BAD_REQUEST: "Something went wrong with your request, please try again.",
  METHOD_NOT_ALLOWED:
    "Something went wrong with your request, please try again.",
  PAYLOAD_TOO_LARGE:
    "Your input is too large, please try with a shorter resume or job description.",
  RATE_LIMITED:
    "You've made too many requests, please wait a moment and try again.",
};

export const STATUS_MESSAGES: Record<number, string> = {
  429: "You've made too many requests, please wait a moment and try again.",
  502: "The analysis service is temporarily unavailable, please try again in a moment.",
  503: "The analysis service is temporarily unavailable, please try again in a moment.",
  504: "The analysis service is temporarily unavailable, please try again in a moment.",
};
