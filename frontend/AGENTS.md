# PRD: Resume Analyzer Frontend

## 1. Overview

A single-page web application that allows job seekers to upload their resume (PDF/DOCX), paste a target job description, and submit both to a backend analysis service. The frontend displays an overall match score and structured improvement suggestions in a clean, clean neobrutalist layout (see `DESIGN.md`).

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

### 5.1 Theme System (Light-First)

- **Default**: Light mode (`--color-surface: #f9f9f9`, white cards).
- **Toggle**: Manual switch in the header (sun/moon icon) that overrides the default.
- **Persistence**: Selected theme stored in `localStorage` (`theme: 'dark' | 'light'`).
- **Implementation**: A single `ThemeContext` provides the current theme and toggle function. All colors are defined as CSS custom properties in `index.css` via a `@theme` block. No `dark:` variant classes are used — the theme tokens in `@theme` represent the light palette.

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

- **Network error / Backend error**: Inline error via `ErrorAlert` component: "Unable to reach the analysis server. Please try again."
- **File extraction error**: Inline error below dropzone: "Could not read this file. Please try another PDF or DOCX."

### 5.6 Navigation / Reset

- "Analyze another resume" button at the bottom of results → scrolls back to top, clears all form state.

## 6. UI/UX Design Direction

### Aesthetic: Neobrutalist

- **Fonts**: `Space Grotesk` for headlines, `Work Sans` for body. Self-hosted in `public/fonts/`.
- **Background**: Light beige `#f9f9f9` (default), white `#ffffff` (cards).
- **Surfaces**: White cards with `4px` solid black borders and hard offset shadows (`8px 8px 0 0 #000`).
- **Typography**: `Space Grotesk` bold for headings (`letter-spacing: -0.02em`). `Work Sans` for body text. All headers uppercase.
- **Radii**: `0px` — sharp corners everywhere (neobrutalist).
- **Spacing**: Generous. Sections separated by `64px–96px`. Inner padding `32px`.
- **Accent**: Yellow (`#FFDAB9` / `#FFE600`) for primary actions, pink (`#E0BBE4`) for secondary, cyan (`#B2E2D2`) for tertiary.

## 7. Frontend Architecture

### 7.1 Tech Stack (Existing + Additions)

| Layer | Technology | Status |
|-------|-----------|--------|
| Build Tool | Vite 8 | ✅ Existing |
| Framework | React 19 + TypeScript | ✅ Existing |
| Styling | Tailwind CSS v4 | ✅ Existing |
| Routing | TanStack Router (active — `/` + `/analyze`) | ✅ Existing |
| Compiler | babel-plugin-react-compiler | ✅ Existing |
| PDF Parsing | `pdfjs-dist` | ✅ Existing |
| DOCX Parsing | `mammoth` | ✅ Existing |
| Icons | `lucide-react` | ✅ Existing |

### 7.2 State Management

**No global state library. No `useReducer`.**

Analyzer UI state lives in the route component `AnalyzePage` (`src/routes/analyze.tsx`) using plain `useState`:

```typescript
function AnalyzePage() {
  const [resumeText, setResumeText] = useState('');
  const [jobDescription, setJobDescription] = useState('');
  const [isExtracting, setIsExtracting] = useState(false);
  const [isAnalyzing, setIsAnalyzing] = useState(false);
  const [extractError, setExtractError] = useState('');
  const [results, setResults] = useState<AnalysisResponse | null>(null);
  // ...
}
```

State is passed down as props. Callbacks (`handleExtract`, `setJobDescription`, `handleSubmit`) are passed to child components.

**The ONLY exception** is the `ThemeContext`. The root layout uses a minimal context to avoid prop drilling:

```typescript
// contexts/ThemeContext.tsx
const ThemeContext = createContext<ThemeContextType>({ theme: 'light', toggle: () => {} });
```

### 7.3 Component Hierarchy

```
Router
└── RootLayout (ThemeProvider → Outlet)
    ├── LandingPage (/)
    │   ├── Header
    │   ├── HeroSection
    │   ├── Features
    │   ├── TheProcess
    │   ├── SuccessStories
    │   ├── Questions
    │   └── Footer
    └── AnalyzePage (/analyze)
        ├── AnalyzerNavbar
        ├── FileDropzone
        ├── JobDescriptionInput
        ├── SubmitButton
        ├── LoadingOverlay
        └── ResultsSection
            ├── ScoreDisplay
            ├── MissingKeywordsList
            ├── SectionFeedbackCards
            └── RewriteSuggestionsList
```

### 7.4 File Structure

```
src/
├── main.tsx                   # Entry point: TanStack Router setup + render
├── index.css                  # Tailwind import + @theme block + neobrutal utility classes
├── lib/
│   ├── extract-text.ts        # Client-side PDF/DOCX extraction (lazy-loaded)
│   └── api.ts                 # Fetch wrapper for /api/analyze (with mock mode)
├── types/
│   └── index.ts               # AnalysisRequest, AnalysisResponse interfaces
├── contexts/
│   └── ThemeContext.tsx        # Theme provider (light default, toggle function)
├── hooks/
│   └── useTheme.ts            # Convenience hook for ThemeContext
├── routes/
│   ├── __root.tsx             # Root layout: ThemeProvider + Outlet
│   ├── index.tsx              # Landing page (/)
│   └── analyze.tsx            # Analyzer page (/analyze) — all UI state lives here
└── components/
    ├── Header.tsx             # Landing page header (sticky nav)
    ├── HeroSection.tsx        # Hero with CTA → /analyze
    ├── Features.tsx           # Feature cards grid
    ├── TheProcess.tsx         # 3-step process timeline
    ├── SuccessStories.tsx     # Testimonial cards
    ├── Questions.tsx          # FAQ accordion
    ├── Footer.tsx             # Site footer
    └── analyzer/              # Analyzer-specific components (route: /analyze)
        ├── AnalyzerNavbar.tsx
        ├── FileDropzone.tsx
        ├── JobDescriptionInput.tsx
        ├── SubmitButton.tsx
        ├── LoadingOverlay.tsx
        ├── ResultsSection.tsx
        ├── ScoreDisplay.tsx
        ├── MissingKeywordsList.tsx
        ├── SectionFeedbackCards.tsx
        └── RewriteSuggestionsList.tsx
```

### 7.5 Data Flow

1. User navigates to `/analyze` → TanStack Router renders `AnalyzePage`.
2. User drops file → `FileDropzone` calls `extractText(file)` → calls `handleExtract()` in `AnalyzePage` → `setResumeText()`.
3. User types job description → `JobDescriptionInput` calls `onChange` → `setJobDescription()`.
4. User clicks submit → `handleSubmit()` calls `analyzeResume(resumeText, jobDescription)` from `lib/api.ts`.
5. On success: `AnalyzePage` calls `setResults(data)` and renders `ResultsSection`.
6. On error: inline `ErrorAlert` component displays error message.
7. "Analyze Another Resume" clears all state via `handleReset()`.

### 7.6 Theme Implementation Detail

Tailwind v4 is CSS-first. The theme uses a `@theme` block in `index.css` to define design tokens. No Tailwind `dark:` mode classes are used.

```css
/* index.css */
@import "tailwindcss";

@theme {
  --color-primary: #ffdab9;
  --color-primary-container: #ffdab9;
  --color-secondary: #e0bbe4;
  --color-secondary-container: #e0bbe4;
  --color-tertiary: #b2e2d2;
  --color-surface: #f9f9f9;
  --color-background: #f9f9f9;
  --color-on-surface: #1b1b1b;
  --color-on-surface-variant: #4b4731;
  --color-on-primary: #1b1b1b;
  --color-border-brutal: #000000;
  --color-card: #ffffff;
  --color-text: #1b1b1b;
  --color-text-muted: #4b4731;
  --color-success: #22c55e;
  --color-warning: #f59e0b;
  --color-error: #ef4444;
}
```

Components use theme tokens via Tailwind utilities (e.g., `bg-primary-container`, `text-on-surface-variant`) and neobrutal utility classes (`btn-neo`, `card-neo`, `tag-neo`) defined in `@layer components`.

### 7.7 PDF.js Worker Setup

After installing `pdfjs-dist`, copy the worker file to the public directory so the browser can load it at a stable URL:

```bash
cp node_modules/pdfjs-dist/build/pdf.worker.min.mjs public/pdf.worker.min.js
```

In code (`src/lib/extract-text.ts`):

```typescript
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
- Error states use a `role="alert"` element with `aria-live="assertive"` via the `ErrorAlert` component.

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
