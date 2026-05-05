# PRD: Resume Analyzer Frontend

## 1. Overview

A single-page web application that allows job seekers to upload their resume (PDF/DOCX), paste a target job description, and submit both to a backend analysis service. The frontend displays an overall match score and structured improvement suggestions in a clean, dark-first UI inspired by Vercel and Linear.

## 2. Goals

- Provide a frictionless upload-to-insight experience in under 60 seconds.
- Display analysis results clearly with an emphasis on actionable feedback.
- Support both dark and light themes, defaulting to dark.
- Remain purely frontend: all file parsing happens client-side before text is sent to the backend.

## 3. Non-Goals

- No AI/LLM logic in the browser. The frontend is strictly a presentation and data-collection layer.
- No user authentication or session persistence (for this version).
- No backend implementation (assumed to exist at `/api/analyze`).

## 4. User Flow

```
┌──────────────┐     ┌──────────────────┐     ┌──────────────────┐
│  Landing     │────▶│  Upload Resume   │────▶│ Paste Job Desc   │
│  (Hero)      │     │  + Extract Text  │     │ + Submit         │
└──────────────┘     └──────────────────┘     └──────────────────┘
                                                        │
                                                        ▼
                                               ┌──────────────────┐
                                               │  Loading State   │
                                               └──────────────────┘
                                                        │
                                                        ▼
                                               ┌──────────────────┐
                                               │  Results View    │
                                               │  Score + List    │
                                               └──────────────────┘
```

## 5. Feature Specification

### 5.1 Theme System (Dark-First)

- **Default**: Dark mode.
- **Toggle**: Manual switch in the header (sun/moon icon) that overrides `prefers-color-scheme`.
- **Persistence**: Selected theme stored in `localStorage` (`theme: 'dark' | 'light'`).
- **Implementation**: A single `ThemeContext` provides the current theme and toggle function. The `<html>` element receives a `.light` class when active. All colors are CSS custom properties in `index.css`.

### 5.2 Landing / Hero Section

- Clean, centered layout.
- Headline: "Optimize your resume for any job."
- Subheadline: "Upload your resume, paste a job description, and get instant AI-powered feedback."
- Primary CTA button: "Analyze my resume" → scrolls to the analyzer section.

### 5.3 Upload & Input Section

**Resume Upload**

- **Supported formats**: PDF, DOCX.
- **Input method**: Drag-and-drop zone + click-to-browse file input.
- **Client-side extraction**:
  - PDF: `pdfjs-dist` (Mozilla's PDF.js) to extract text in the browser.
  - DOCX: `mammoth` to convert .docx to plain text.
- **Validation**:
  - Max file size: 5MB.
  - Reject unsupported types with inline error.
- **State after upload**: Show file name, size, and a "Remove" button. The extracted text is held in memory (not displayed to the user).

**Job Description**

- Large `<textarea>` (min-height: 200px) with a character counter.
- Placeholder: "Paste the job description here..."

**Submit**

- Button state: Disabled until both resume text and job description are present.
- On click: POST to `/api/analyze` with JSON body:

```json
{
  "resume_text": "extracted text...",
  "job_description": "pasted text..."
}
```

### 5.4 Loading State

- Inline spinner in the submit button + overlay on the form.
- Display a rotating set of 3 "tip" messages (e.g., "Analyzing keywords...", "Comparing experience...", "Calculating match score...") to reduce perceived wait time.

### 5.5 Results View

Triggered on successful `200 OK` from backend. Displays the response shape:

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
    "Change 'Responsible for' to 'Led a team of 5 engineers to deliver...'",
    "Add specific metrics to your second bullet point..."
  ]
}
```

**UI Mapping**:

1. **Overall Score**: Large, prominent display (e.g., `78/100`). Circular progress indicator or large numerical display with a color scale:
   - 0-49: Red (`#ef4444`)
   - 50-79: Amber (`#f59e0b`)
   - 80-100: Green (`#22c55e`)
2. **Missing Keywords**: Horizontal scrolling tag list or flex-wrap grid of "pill" badges. Color: subtle border, muted background.
3. **Section Feedback**: Card-per-section layout. Sections: Summary, Experience, Skills. Each card has a title and the feedback paragraph.
4. **Rewrite Suggestions**: Numbered list with light copy-to-clipboard buttons on each item.

**Error Handling**:

- **Network error / Backend error**: Global toast via `sonner`: "Unable to reach the analysis server. Please try again."
- **File extraction error**: Inline error below dropzone: "Could not read this file. Please try another PDF or DOCX."

### 5.6 Navigation / Reset

- "Analyze another resume" button at the bottom of results → scrolls back to top, clears all form state.

## 6. UI/UX Design Direction

### Aesthetic: Vercel / Linear

- **Font**: `Inter` via Google Fonts CDN.
- **Background**: Deep dark `#0a0a0a` (dark), clean white `#ffffff` (light).
- **Surfaces**: Slightly elevated cards with `1px` borders (`#27272a` dark, `#e4e4e7` light) and subtle shadows.
- **Typography**: `Inter` font family. Headings are tight (`letter-spacing: -0.02em`).
- **Radii**: Small and sharp. `8px` for buttons, `12px` for cards. No excessive rounded corners.
- **Spacing**: Generous. Sections separated by `80px–120px`. Inner padding `24px–32px`.
- **Accent**: Purple (`#7c3aed` / `#8b5cf6`) for primary actions and score indicators.

## 7. Frontend Architecture

### 7.1 Tech Stack (Existing + Additions)

| Layer | Technology | Status |
|-------|-----------|--------|
| Build Tool | Vite 8 | ✅ Existing |
| Framework | React 19 + TypeScript | ✅ Existing |
| Styling | Tailwind CSS v4 | ✅ Existing |
| Routing | TanStack Router | ✅ Existing (unused for now) |
| Compiler | babel-plugin-react-compiler | ✅ Existing |
| PDF Parsing | `pdfjs-dist` | ⬜ Add |
| DOCX Parsing | `mammoth` | ⬜ Add |
| Icons | `lucide-react` | ⬜ Add |
| Toasts | `sonner` | ⬜ Add |

### 7.2 State Management

**No global state library. No `useReducer`.**

All UI state lives in `App.tsx` using plain `useState`:

```typescript
function App() {
  const [phase, setPhase] = useState<'idle' | 'ready' | 'analyzing' | 'results' | 'error'>('idle');
  const [resumeText, setResumeText] = useState('');
  const [jobDescription, setJobDescription] = useState('');
  const [results, setResults] = useState<AnalysisResponse | null>(null);
  const [error, setError] = useState('');
  // ...
}
```

State is passed down as props. Callbacks (`setResumeText`, `setJobDescription`, `handleSubmit`) are passed to child components.

**The ONLY exception** is the `ThemeContext`. Since `Header` (toggle button) and the root `<html>` element both need theme awareness, a minimal context prevents prop drilling:

```typescript
// contexts/ThemeContext.tsx
const ThemeContext = createContext({ theme: 'dark', toggle: () => {} });
```

### 7.3 Component Hierarchy

```
App (holds all state)
├── ThemeProvider
├── Header
│   └── ThemeToggle (consumes ThemeContext)
├── HeroSection
├── AnalyzerSection
│   ├── FileDropzone (receives setResumeText)
│   ├── JobDescriptionInput (receives jobDescription, setJobDescription)
│   └── SubmitButton (receives onClick, disabled state)
├── LoadingOverlay
└── ResultsSection (receives results data)
    ├── ScoreDisplay
    ├── MissingKeywordsList
    ├── SectionFeedbackCards
    └── RewriteSuggestionsList
```

### 7.4 File Structure

```
src/
├── main.tsx
├── App.tsx
├── index.css                  # Tailwind import + CSS custom properties for theming
├── lib/
│   ├── extract-text.ts        # PDF/DOCX extraction logic
│   └── api.ts                 # Fetch wrapper for /api/analyze
├── types/
│   └── index.ts               # AnalysisRequest, AnalysisResponse
├── contexts/
│   └── ThemeContext.tsx       # Theme provider + hook
├── components/
│   ├── Header.tsx
│   ├── ThemeToggle.tsx
│   ├── HeroSection.tsx
│   ├── AnalyzerSection.tsx
│   ├── FileDropzone.tsx
│   ├── JobDescriptionInput.tsx
│   ├── LoadingOverlay.tsx
│   ├── ResultsSection.tsx
│   ├── ScoreDisplay.tsx
│   ├── MissingKeywordsList.tsx
│   ├── SectionFeedbackCards.tsx
│   └── RewriteSuggestionsList.tsx
└── hooks/
    └── useTheme.ts            # Convenience hook for ThemeContext
```

### 7.5 Data Flow

1. User drops file → `FileDropzone` calls `extractText(file)` → calls `setResumeText()` passed from `App`.
2. User types job description → `JobDescriptionInput` calls `setJobDescription()` passed from `App`.
3. User clicks submit → `App` calls `analyzeResume(resumeText, jobDescription)` from `lib/api.ts`.
4. On success: `App` calls `setResults(data)` and `setPhase('results')`.
5. `App` conditionally renders `ResultsSection`.

### 7.6 Theme Implementation Detail

Tailwind v4 is CSS-first. We will **not** use Tailwind's dark mode classes. Instead:

```css
/* index.css */
@import "tailwindcss";

:root {
  --bg: #0a0a0a;
  --bg-elevated: #16171d;
  --text: #9ca3af;
  --text-heading: #f3f4f6;
  --border: #27272a;
  --accent: #7c3aed;
  --accent-text: #ffffff;
  --success: #22c55e;
  --warning: #f59e0b;
  --error: #ef4444;
}

:root.light {
  --bg: #ffffff;
  --bg-elevated: #f4f4f5;
  --text: #52525b;
  --text-heading: #18181b;
  --border: #e4e4e7;
  --accent: #7c3aed;
  --accent-text: #ffffff;
  --success: #16a34a;
  --warning: #d97706;
  --error: #dc2626;
}
```

Components reference these variables directly or use Tailwind utilities mapped via `@theme` if needed.

### 7.7 PDF.js Worker Setup

After installing `pdfjs-dist`, copy the worker file to the public directory so the browser can load it at a stable URL:

```bash
cp node_modules/pdfjs-dist/build/pdf.worker.min.js public/pdf.worker.min.js
```

In code:

```typescript
import * as pdfjsLib from 'pdfjs-dist';
pdfjsLib.GlobalWorkerOptions.workerSrc = '/pdf.worker.min.js';
```

## 8. API Contract (Frontend Assumption)

```typescript
// POST /api/analyze
interface AnalysisRequest {
  resume_text: string;
  job_description: string;
}

interface AnalysisResponse {
  score: number; // 0-100
  missing_keywords: string[];
  section_feedback: {
    summary: string;
    experience: string;
    skills: string;
  };
  rewrite_suggestions: string[];
}
```

## 9. Accessibility

- All interactive elements keyboard-navigable.
- Drag-and-drop zone has a hidden native `<input type="file">` accessible via keyboard.
- Color contrast meets WCAG AA.
- Loading states announced via `aria-live="polite"`.
- Results sections use semantic headings (`<h2>`, `<h3>`).
- Toasts from `sonner` are screen-reader friendly by default.

## 10. Performance

- **Code splitting**: `pdfjs-dist` and `mammoth` are heavy. Lazy-load them inside `extract-text.ts` so they are only downloaded when a user interacts with the dropzone.
- **Bundle**: Target <200KB initial JS (excluding lazy-loaded parsers).

## 11. Edge Cases & Error States

| Scenario | Behavior |
|----------|----------|
| User uploads image/txt | Reject immediately with inline error: "Please upload a PDF or DOCX file." |
| File > 5MB | Reject with inline error: "File size must be under 5MB." |
| Extraction fails (corrupted PDF) | Inline error: "Could not read this file." |
| Backend returns 500 | Toast: "Analysis failed. Please try again later." |
| Backend returns 422 | Toast with backend message, or fallback: "Invalid input." |
| User navigates away during loading | Abort fetch via `AbortController`. |
| Job description is empty on submit | Disable submit button. |

## 12. Future Considerations

- Add route for `/history` using TanStack Router (already installed) if user accounts are added later.
- Add a "Compare Mode" to analyze against multiple job descriptions.
- Export results as a PDF report.
- Add a side-by-side diff view for rewrite suggestions.
