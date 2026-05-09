# AGENTS.md — resu-me

## Repo Structure

Monorepo with two packages:

| Package | Path | Stack |
|---------|------|-------|
| Frontend | `frontend/` | React 19 + TypeScript, Vite 8, Tailwind CSS v4, TanStack Router (active — `/` + `/analyze`) |
| Backend | `backend/` | Go 1.25.0, gorilla/mux, openai-go/v3 SDK (OpenRouter-compatible), in-memory rate limiter |

## Frontend (`frontend/`)

**Package manager**: `pnpm` (lockfile: `pnpm-lock.yaml`)

### Commands (run from `frontend/`)
```
pnpm dev          # Vite dev server (HMR)
pnpm build        # tsc -b && vite build
pnpm lint         # ESLint
pnpm preview      # Preview production build
```

### Key Conventions
- **State**: All analyzer UI state in `src/routes/analyze.tsx` via `useState`. No global state library. Only exception: `ThemeContext`.
- **Styling**: Tailwind CSS v4 (CSS-first). Theme uses CSS custom properties in `index.css` via `@theme` block — do NOT use Tailwind's `dark:` classes.
- **PDF.js worker**: Copied via `postinstall` hook to `public/pdf.worker.min.js` (from `pdfjs-dist/build/pdf.worker.min.mjs`). After `pnpm install`, verify the file exists.
- **Design system**: Neobrutalist (see `DESIGN.md`) — sharp corners, hard shadows, bold borders, Space Grotesk + Work Sans fonts.
- **API contract**: `POST /api/analyze` → `{ score, missing_keywords, section_feedback, rewrite_suggestions }`. See `frontend/AGENTS.md` for full spec.

### File Boundaries
- `src/lib/extract-text.ts` — client-side PDF/DOCX parsing (lazy-loaded)
- `src/lib/api.ts` — fetch wrapper to `/api/analyze`
- `src/types/index.ts` — shared TypeScript types
- `src/contexts/ThemeContext.tsx` — only context in the app
- `src/routes/` — TanStack Router (active — `/` + `/analyze`)

## Backend (`backend/`)

**Module**: `resu-me` (Go 1.25.0)

### Commands (run from `backend/`)
```
go run .              # Start server
go build              # Compile
go test ./...         # Run tests
go vet ./...          # Static analysis
go fmt ./...          # Format
```

### Architecture
- **Router**: `github.com/gorilla/mux` — use it for routing, method matching, and middleware.
- **LLM provider**: OpenRouter (OpenAI-compatible). Use official `github.com/openai/openai-go/v3` SDK with custom base URL.
- **Stateless** — no database. Each request is a fresh analysis.
- **Package layout**:
  - `handler/` — HTTP handlers (request parsing, JSON response writing)
  - `service/` — business logic (LLM prompt construction, response parsing)
  - `model/` — request/response structs
  - `config/` — environment variable loading (.env via godotenv)
  - `provider/` — LLM client abstraction (wraps OpenAI SDK for OpenRouter)
  - `middleware/` — CORS, rate limit, body size limit
  - `sanitizer/` — XSS/HTML/SQL injection detection

### API Contract
- `POST /api/analyze` — accepts `{ "resume_text": string, "job_description": string }`
- Returns `{ "score": number, "missing_keywords": string[], "section_feedback": {...}, "rewrite_suggestions": string[] }`
- Content-Type: `application/json`
- CORS: Allow `*` for development (frontend and backend run on different ports during dev)

### LLM Prompt Design
- System prompt should instruct the model to return **valid JSON only** matching the response schema.
- Use structured output / JSON mode if the provider supports it.
- Temperature: low (0.1–0.3) for deterministic scoring.

## Cross-Cutting

### Dev Setup
- Frontend dev server: `http://localhost:5173` (default Vite)
- Backend: pick a port (e.g., `:8080`) and configure frontend Vite proxy or CORS.
- No `.env` files committed. Use environment variables for API keys (`OPENROUTER_API_KEY`).

### What NOT to do
- Don't add a database unless explicitly requested.
- Don't swap gorilla/mux for another framework — it's the chosen router.
- Don't use Tailwind `dark:` classes in the frontend — CSS custom properties handle theming.
- Don't put AI/LLM logic in the frontend — it's purely a presentation layer.
