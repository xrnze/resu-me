import type { AnalysisResponse } from '@/types';
import { ScoreDisplay } from './ScoreDisplay';
import { MissingKeywordsList } from './MissingKeywordsList';
import { SectionFeedbackCards } from './SectionFeedbackCards';
import { RewriteSuggestionsList } from './RewriteSuggestionsList';

interface ResultsSectionProps {
  results: AnalysisResponse;
  onReset: () => void;
}

export function ResultsSection({ results, onReset }: ResultsSectionProps) {
  return (
    <div className="space-y-10">
      <div className="text-center">
        <h1 className="font-headline-md text-headline-md uppercase tracking-tighter text-black">
          Analysis Results
        </h1>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <ScoreDisplay score={results.score} />
        <MissingKeywordsList keywords={results.missing_keywords} />
      </div>
      <SectionFeedbackCards feedback={results.section_feedback} />
      <RewriteSuggestionsList suggestions={results.rewrite_suggestions} />

      <div className="flex justify-center pt-4">
        <button onClick={onReset} className="btn-neo">
          Analyze Another Resume
        </button>
      </div>
    </div>
  );
}
