import { useState } from 'react';
import { Copy, Check } from 'lucide-react';

interface RewriteSuggestionsListProps {
  suggestions: string[];
}

export function RewriteSuggestionsList({ suggestions }: RewriteSuggestionsListProps) {
  const [copiedIndex, setCopiedIndex] = useState<number | null>(null);

  const handleCopy = async (text: string, index: number) => {
    try {
      await navigator.clipboard.writeText(text);
      setCopiedIndex(index);
      setTimeout(() => setCopiedIndex(null), 1500);
    } catch {
      // silently fail — icon stays as Copy
    }
  };

  if (suggestions.length === 0) {
    return (
      <div className="card-neo">
        <h3 className="font-['Space_Grotesk'] font-bold text-lg uppercase text-black mb-2">
          Rewrite Suggestions
        </h3>
        <p className="text-on-surface-variant font-body-md text-body-md">
          No rewrite suggestions needed.
        </p>
      </div>
    );
  }

  return (
    <div>
      <h2 className="font-['Space_Grotesk'] font-bold text-lg uppercase text-black mb-4">
        Rewrite Suggestions
      </h2>
      <ol className="space-y-4">
        {suggestions.map((suggestion, i) => (
          <li key={i} className="card-neo flex items-start gap-4">
            <span className="font-['Space_Grotesk'] font-black text-xl text-black shrink-0 mt-0.5">
              {String(i + 1).padStart(2, '0')}
            </span>
            <p className="flex-1 font-body-md text-body-md text-on-surface-variant leading-relaxed">
              {suggestion}
            </p>
            <button
              onClick={() => handleCopy(suggestion, i)}
              className="p-2 border-2 border-black shrink-0 hover:bg-black hover:text-white transition-colors"
              aria-label={`Copy suggestion ${i + 1}`}
            >
              {copiedIndex === i ? <Check size={16} /> : <Copy size={16} />}
            </button>
          </li>
        ))}
      </ol>
    </div>
  );
}
