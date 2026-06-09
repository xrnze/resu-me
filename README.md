# RESU-ME

AI-powered resume analyzer. Upload your resume, paste a job description, and get instant structured feedback on how well they match.

<img width="1853" height="837" alt="image" src="https://github.com/user-attachments/assets/43dcee74-1142-4d57-914d-37487c661fbb" />
<img width="1876" height="923" alt="image" src="https://github.com/user-attachments/assets/d27f1747-6915-4da7-9874-5829e4ac28a9" />
<img width="1853" height="924" alt="image" src="https://github.com/user-attachments/assets/64065ea5-fad4-4544-89f8-df7a0bdf7af0" />


## What it does

- Upload a PDF or DOCX resume text is extracted client-side in the browser
- Paste a job description
- Submit both to an LLM for analysis
- Receive a match score (0–100), missing keywords, section-by-section feedback, and actionable rewrite suggestions
- All inputs are sanitized against XSS, HTML injection, and SQL injection. An additional LLM-based filter checks for prompt injection before the analysis runs.

## Tech stack

| Layer | Technology |
|-------|-----------|
| Frontend | React 19, TypeScript, Vite 8, Tailwind CSS v4, TanStack Router |
| Backend | Go 1.25.0, gorilla/mux, OpenAI Go SDK (OpenAI-compatible API) |
| Infrastructure | Docker Compose, Traefik (reverse proxy + automatic TLS via Let's Encrypt) |

## Architecture

```
User → Frontend (React) → POST /api/analyze → Backend (Go) → LLM Provider → JSON response
```

- The frontend handles file parsing (PDF/DOCX) client-side and sends extracted text to the backend.
- The backend is stateless no database. Each request is a fresh analysis.
- All AI analysis runs via an OpenAI-compatible API (e.g., [OpenRouter](https://openrouter.ai/)) using a configurable LLM model.

For in-depth architecture and API docs, see [`backend/AGENTS.md`](./backend/AGENTS.md) and [`frontend/AGENTS.md`](./frontend/AGENTS.md).

## Getting started

### Prerequisites

- Node.js 20+ and [pnpm](https://pnpm.io/installation)
- Go 1.25+
- Docker + Docker Compose (for containerized setup)
- An API key for an OpenAI-compatible LLM provider (e.g., [OpenRouter](https://openrouter.ai/))

### Local development

**Backend:**

```bash
cd backend
cp .env.example .env
# Edit .env and set LLM_API_KEY
go run .
```

The backend starts on `localhost:8080`.

**Frontend:**

```bash
cd frontend
cp env.example .env
# .env already contains VITE_BACKEND_URL=localhost:8080
pnpm install
pnpm dev
```

The frontend starts on `localhost:5173`. It points to the backend via the `VITE_BACKEND_URL` environment variable.

### Docker

```bash
cp .env.example .env
# Set ACME_EMAIL for Let's Encrypt TLS (optional)

# Create backend/.env with LLM_API_KEY
cp backend/.env.example backend/.env

docker-compose up -d
```

Traefik handles routing: `/` → frontend, `/api` → backend with automatic TLS via Let's Encrypt.

## Project structure

```
resu-me/
├── frontend/          # React SPA upload, input, results UI
├── backend/           # Go API analysis engine, prompt injection filter, rate limiter
├── traefik/           # Reverse proxy configuration
├── docker-compose.yml
└── Makefile
```

## Environment variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `LLM_API_KEY` | Yes | - | API key for OpenAI-compatible LLM provider |
| `LLM_BASE_URL` | No | `https://openrouter.ai/api/v1` | OpenAI-compatible API base URL |
| `LLM_MODEL` | No | `openai/gpt-oss-120b` | LLM model for resume analysis |
| `FILTER_MODEL` | No | `meta-llama/llama-3.1-8b-instruct` | Fast model for prompt injection detection |
| `ACME_EMAIL` | No | - | Email for Let's Encrypt (Docker only) |

See [`backend/.env.example`](./backend/.env.example) for the full list of configuration options.

## Running tests

```bash
cd backend
go test ./...
go vet ./...
```

The backend uses table-driven Go tests with TDD workflow. Tests cover the analyzer service, filter service, rate limiter, handler, middleware, config loading, input sanitizer, and LLM provider.
