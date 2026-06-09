# PRD: Resume Analyzer Backend

## 1. Overview

A stateless Go HTTP service that receives a resume and job description, uses an LLM via OpenRouter to analyze the match, and returns structured JSON feedback. The backend is purely an analysis engine — no persistence, no user accounts, no sessions.

## 2. Goals

- Accept resume text + job description, return analysis in under 30 seconds.
- Produce deterministic, well-structured JSON responses matching the frontend contract.
- Be deployable as a single binary with zero external dependencies (no DB, no cache).
- Handle concurrent requests safely with per-IP rate limiting.
- Sanitize inputs against XSS, HTML injection, and SQL injection patterns.

## 3. Non-Goals

- No database or data persistence.
- No user authentication or accounts.
- No file parsing (PDF/DOCX extraction happens client-side).
- No streaming — single JSON response per request.

## 4. Tech Stack

| Layer | Technology |
|-------|-----------|
| Language | Go 1.25.0 |
| Module | `resu-me` |
| Router | `github.com/gorilla/mux` |
| LLM SDK | `github.com/openai/openai-go/v3` (OpenRouter-compatible) |
| Rate Limiter | In-memory hash map with token bucket |
| Logging | Standard library `log` |
| Input Sanitizer | Custom (regex-based HTML/SQL pattern detection) |
| Testing | Go `testing` package, table-driven tests |

## 5. Architecture

### 5.1 Package Layout

```
backend/
├── go.mod
├── main.go                 # Entry point: server setup, router, middleware
├── AGENTS.md               # This file
├── config/
│   ├── config.go           # Environment variable loading (.env via godotenv)
│   └── config_test.go      # Config loading tests (defaults, custom, errors)
├── provider/
│   ├── openai.go           # LLM client abstraction (wraps OpenAI SDK for OpenRouter)
│   └── openai_test.go      # Provider tests (client creation, HTTP errors)
├── model/
│   ├── request.go          # AnalysisRequest, AnalysisResponse, ErrorResponse structs
│   └── request_test.go     # JSON marshaling, validation tests
├── handler/
│   ├── analyze.go          # HTTP handler: parse, validate, sanitize, call service, respond
│   └── analyze_test.go     # Handler tests with mocked service
├── service/
│   ├── analyze.go          # Business logic: prompt construction, LLM call, response parsing
│   ├── analyze_test.go     # Prompt construction, response parsing tests (mock LLM)
│   ├── filter.go           # LLM-based prompt injection filter
│   ├── filter_test.go      # Filter tests: safe, injection, invalid JSON, network error
│   ├── ratelimit.go        # Token bucket rate limiter (per-IP, in-memory, golang.org/x/time/rate)
│   └── ratelimit_test.go   # Token bucket: acquire, refill, burst, concurrency, cleanup
├── middleware/
│   ├── bodylimit.go        # Request body size limit middleware (100KB)
│   ├── bodylimit_test.go   # Body limit: within limit, over limit
│   ├── cors.go             # CORS middleware (Allow *)
│   ├── cors_test.go        # CORS header + preflight tests
│   ├── ratelimit.go        # Rate limit middleware (wraps service/ratelimit.go)
│   └── ratelimit_test.go   # Middleware integration: allowed, blocked, stale cleanup
└── sanitizer/
    ├── input.go            # XSS/HTML/SQL injection detection and sanitization
    └── input_test.go       # All sanitization pattern tests
```

### 5.2 Request Flow

```
POST /api/analyze
  → Body limit middleware (100KB max)
  → CORS middleware (preflight + headers)
  → Rate limit middleware (per-IP token bucket)
  → Analyze handler
    → Parse JSON body → model.AnalysisRequest
    → Validate (non-empty, max length)
    → Sanitize input (strip HTML/scripts, reject SQL injection)
    → Check prompt injection via LLM filter (separate fast model)
      → On injection detected: return 422 PROMPT_INJECTION_DETECTED
    → service.Analyze(resumeText, jobDesc)
      → Build system + user prompt
      → provider.OpenAIClient.Chat() (OpenRouter via OpenAI SDK)
      → Parse LLM response into model.AnalysisResponse
    → Write JSON response (200 OK)
```

## 6. API Contract

### 6.1 POST /api/analyze

**Request:**
```json
{
  "resume_text": "string (required, non-empty, max 50000 chars, sanitized)",
  "job_description": "string (required, non-empty, max 50000 chars, sanitized)"
}
```

**Response (200 OK):**
```json
{
  "score": 78,
  "missing_keywords": ["Kubernetes", "CI/CD", "TypeScript"],
  "section_feedback": {
    "summary": "Your summary is too generic...",
    "experience": "Good use of metrics...",
    "skills": "Missing several technical keywords..."
  },
  "rewrite_suggestions": [
    "Change 'Responsible for' to 'Led a team of 5 engineers...'",
    "Add specific metrics to your second bullet point..."
  ]
}
```

**Error Response (400 Bad Request):**
```json
{
  "error": {
    "code": "BAD_REQUEST",
    "message": "Request body is not valid JSON",
    "details": null
  }
}
```

**Error Response (422 Validation Error):**
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "resume_text contains disallowed patterns (HTML/script tags)",
    "details": { "field": "resume_text" }
  }
}
```

**Error Response (429 Rate Limited):**
```json
{
  "error": {
    "code": "RATE_LIMITED",
    "message": "Too many requests. Please try again later.",
    "details": null
  }
}
```

**Error Response (500 Server Error):**
```json
{
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "Analysis failed. Please try again later.",
    "details": null
  }
}
```

**Error Response (502 LLM Provider Error):**
```json
{
  "error": {
    "code": "LLM_ERROR",
    "message": "Failed to communicate with the analysis provider.",
    "details": null
  }
}
```

### 6.2 GET /health

**Response (200 OK):**
```json
{ "status": "ok" }
```

## 7. Input Validation & Sanitization

### 7.1 Validation Rules

| Rule | Behavior |
|------|----------|
| `resume_text` missing or empty | 422 `VALIDATION_ERROR` |
| `job_description` missing or empty | 422 `VALIDATION_ERROR` |
| Either field > 50,000 characters | 422 `VALIDATION_ERROR` |
| Either field is whitespace-only | 422 `VALIDATION_ERROR` |
| Malformed JSON body | 400 `BAD_REQUEST` |

### 7.2 Sanitization Rules

| Pattern | Action |
|---------|--------|
| HTML tags (`<script>`, `<img>`, `<iframe>`, etc.) | Strip tags, return 422 if found |
| JavaScript URLs (`javascript:`, `data:text/html`) | Reject, return 422 |
| Event handlers (`onerror=`, `onclick=`, `onload=`) | Reject, return 422 |
| SQL injection keywords (`DROP TABLE`, `UNION SELECT`, `INSERT INTO`, `DELETE FROM`, `'; --`) | Reject, return 422 |
| Encoded variants (`%3Cscript%3E`, `&#60;script&#62;`) | Decode then check, reject if found |

**Implementation**: Use compiled regex patterns in `sanitizer/input.go`. Check both fields before passing to the LLM service. Log rejected patterns for debugging (do not log full input).

### 7.3 Prompt Injection Filter

After regex sanitization passes, an LLM-based prompt injection filter runs against both fields in a single call:

| Field | Behavior |
|-------|----------|
| Filter passes (`safe: true`) | Request proceeds to analysis |
| Injection detected (`safe: false`) | Return 422 `PROMPT_INJECTION_DETECTED` |
| Filter system error (timeout, network) | Return 500 `INTERNAL_ERROR` |

**Filter model**: A separate fast/cheap model configured via `FILTER_MODEL` (default: `meta-llama/llama-3.1-8b-instruct`). Uses a separate LLM client instance from the analysis model.

**Filter prompt**: Wraps both inputs in `<RESUME_TEXT>` and `<JOB_DESCRIPTION>` XML-style markers. Instructs the model to reply with `{"safe": true/false, "reason": "<brief explanation>"}`.

**Implementation**: `service/filter.go` — `FilterService.Check(ctx, resumeText, jobDesc)` → calls the LLM → parses JSON verdict → returns `ErrInjectionDetected` if unsafe. Wired via `handler.Injector` interface in the handler.

## 8. Rate Limiting

- **Strategy**: Token bucket per client IP, stored in a concurrent-safe `map[string]*TokenBucket`.
- **Default limits**: 10 requests per minute per IP, burst of 3.
- **Behavior**: When limit exceeded, return 429 `RATE_LIMITED`.
- **Cleanup**: Periodic goroutine to evict stale entries (older than 10 minutes).
- **Thread safety**: Use `sync.Mutex` for map access.

## 9. LLM Integration

### 9.1 Provider Configuration

- **Provider**: OpenRouter (OpenAI-compatible API).
- **Base URL**: Configurable via `OPENROUTER_BASE_URL` env var. Default: `https://openrouter.ai/api/v1`.
- **API Key**: Required via `OPENROUTER_API_KEY` env var. Fail fast on startup if missing.
- **Model**: Configurable via `OPENROUTER_MODEL` env var. Default: `openai/gpt-oss-120b`.
- **Temperature**: 0.1–0.3 for deterministic output.
- **Response format**: JSON mode / structured output.

### 9.2 SDK Usage

Use `github.com/openai/openai-go/v3` with a custom `WithBaseURL` option pointing to the OpenRouter endpoint. The OpenAI-compatible API means the same SDK works without modification. The SDK is wrapped in the `provider/` package (`provider.OpenAIClient`) which exposes a `Chat(ctx, prompt)` method and is injected into the service layer via the `LLMClient` interface.

### 9.3 System Prompt

The system prompt must instruct the LLM to return **valid JSON only**, matching this exact schema:

```json
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
```

Prompt should instruct:
- Score reflects how well the resume matches the job description (0-100).
- Missing keywords are skills/technologies from the JD not found in the resume.
- Section feedback addresses summary, experience, and skills sections specifically.
- Rewrite suggestions are actionable, specific improvements.
- Return ONLY valid JSON. No markdown, no explanation, no code blocks.

### 9.4 Error Handling

- **API key missing**: Fail on startup with clear error message.
- **Network timeout**: Return 502 `LLM_ERROR`.
- **Invalid LLM response** (non-JSON or schema mismatch): Return 502 `LLM_ERROR`.
- **Rate limit from OpenRouter**: Return 502 `LLM_ERROR` with retry-after header if available.

## 10. CORS

- **Origins**: `*` (all origins) for development.
- **Methods**: `GET`, `POST`, `OPTIONS`.
- **Headers**: `Content-Type`.
- **Preflight**: Handle `OPTIONS` requests with 204 No Content.

## 11. Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `PORT` | No | `8080` | HTTP server port |
| `OPENROUTER_API_KEY` | Yes | — | OpenRouter API key |
| `OPENROUTER_BASE_URL` | No | `https://openrouter.ai/api/v1` | OpenRouter API base URL |
| `OPENROUTER_MODEL` | No | `openai/gpt-oss-120b` | LLM model to use |
| `LLM_TIMEOUT` | No | `60s` | LLM call timeout (must be less than `SERVER_WRITE_TIMEOUT`) |
| `FILTER_MODEL` | No | `meta-llama/llama-3.1-8b-instruct` | Fast model for prompt injection detection |
| `FILTER_TIMEOUT` | No | `10s` | Filter LLM call timeout |
| `SERVER_READ_TIMEOUT` | No | `30s` | HTTP server read timeout |
| `SERVER_WRITE_TIMEOUT` | No | `60s` | HTTP server write timeout |
| `SERVER_IDLE_TIMEOUT` | No | `120s` | HTTP server idle timeout |
| `RATE_LIMIT_RATE` | No | `10` | Requests per second per IP |
| `RATE_LIMIT_BURST` | No | `3` | Burst capacity per IP |
| `RATE_LIMIT_TTL` | No | `10m` | Rate limiter stale entry TTL |

## 12. Commands

```bash
go run .              # Start server
go build              # Compile binary
go test ./...         # Run all tests
go vet ./...          # Static analysis
go fmt ./...          # Format code
```

## 13. Error Codes Reference

| Code | HTTP Status | Meaning |
|------|-------------|---------|
| `BAD_REQUEST` | 400 | Malformed request body |
| `VALIDATION_ERROR` | 422 | Input validation or sanitization failed |
| `PROMPT_INJECTION_DETECTED` | 422 | LLM filter flagged prompt injection |
| `RATE_LIMITED` | 429 | Too many requests |
| `INTERNAL_ERROR` | 500 | Unexpected server error |
| `LLM_ERROR` | 502 | OpenRouter API call failed |

## 14. TDD Workflow

**Tests are mandatory for every function and service.** Follow this workflow:

1. **Write tests first** — define expected inputs, outputs, and edge cases.
2. **Run tests** — they must fail (red).
3. **Implement the minimum code** to make tests pass (green).
4. **Refactor** — clean up while keeping tests green.
5. **Run `go test ./...`** — all tests must pass before moving to the next component.

### 14.1 Test Coverage Requirements

#### `model/request_test.go`
- JSON unmarshal of valid `AnalysisRequest`
- JSON unmarshal of missing fields
- JSON marshal of `AnalysisResponse` matches expected shape
- JSON marshal of `ErrorResponse` with and without details
- `Validate()`: valid input, empty fields, whitespace-only, over max length

#### `sanitizer/input_test.go`
- Table-driven tests for each pattern category
- Clean input passes through
- `<script>alert('x')</script>` → rejected
- `<img src=x onerror=alert(1)>` → rejected
- `javascript:void(0)` → rejected
- `data:text/html,<script>` → rejected
- `'; DROP TABLE users; --` → rejected
- `UNION SELECT * FROM passwords` → rejected
- `%3Cscript%3E` (URL-encoded) → rejected
- `&#60;script&#62;` (HTML-encoded) → rejected
- Mixed case: `<ScRiPt>` → rejected
- Legitimate text with angle brackets: "C++ > Java, 3 < 5" → passes

#### `config/config_test.go`
- Load with defaults: all config values match expected defaults
- Load with custom env: all config values match custom env vars
- Missing API key: returns error containing "OPENROUTER_API_KEY"
- Invalid duration: returns error for malformed TTL
- Invalid int: returns error for malformed burst value
- Invalid float: returns error for malformed rate value

#### `provider/openai_test.go`
- NewOpenAIClient: client created with correct model string
- Chat with HTTP error: returns error on non-2xx response

#### `service/ratelimit_test.go`
- Token bucket: initial state allows burst
- Token bucket: exhausts tokens, then refills over time
- Token bucket: concurrent access is safe (race detector)
- Rate limiter map: new IP gets fresh bucket
- Rate limiter map: stale entry cleanup works

#### `service/analyze_test.go`
- Prompt construction: system prompt contains correct instructions
- Prompt construction: user prompt includes both resume and JD
- Response parsing: valid JSON → `AnalysisResponse`
- Response parsing: invalid JSON → error
- Response parsing: missing fields → error
- Response parsing: score out of range → error

#### `service/filter_test.go`
- Filter returns `safe: true` → passes through
- Filter returns `safe: false` → error wrapping `ErrInjectionDetected`
- Filter returns invalid JSON → error
- Filter returns missing `safe` field → error
- LLM client network error → error
- LLM client context timeout → error
- `buildFilterPrompt` contains both inputs and XML-style delimiters

#### `handler/analyze_test.go`
- Valid request → 200 + correct response body
- Missing `resume_text` → 422 `VALIDATION_ERROR`
- Missing `job_description` → 422 `VALIDATION_ERROR`
- Over max length → 422 `VALIDATION_ERROR`
- Malformed JSON → 400 `BAD_REQUEST`
- XSS in input → 422 `VALIDATION_ERROR`
- SQL injection in input → 422 `VALIDATION_ERROR`
- Prompt injection detected → 422 `PROMPT_INJECTION_DETECTED`
- Filter system error → 500 `INTERNAL_ERROR`
- Service error → 502 `LLM_ERROR`

#### `middleware/cors_test.go`
- OPTIONS preflight → 204 with correct headers
- POST request → CORS headers present
- All origins allowed (`*`)

#### `middleware/ratelimit_test.go`
- Request within limit → passes through
- Request over limit → 429 `RATE_LIMITED`

## 15. Future Considerations

- Add request logging with request IDs for debugging.
- Add metrics (request count, latency, error rate) via Prometheus.
- Support multiple LLM providers via interface abstraction.
- Add analysis caching for identical inputs.
- Switch to production-grade LLM (GPT-4o, Claude) when moving beyond dev.
