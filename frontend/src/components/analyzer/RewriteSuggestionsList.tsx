import { Copy } from 'lucide-react';
import { toast } from 'sonner';

interface RewriteSuggestionsListProps {
  suggestions: string[];
}

export function RewriteSuggestionsList({ suggestions }: RewriteSuggestionsListProps) {
  const handleCopy = async (text: string) => {
    try {
      await navigator.clipboard.writeText(text);
      toast.success('Copied to clipboard');
    } catch {
      toast.error('Failed to copy');
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
              onClick={() => handleCopy(suggestion)}
              className="p-2 border-2 border-black shrink-0 hover:bg-black hover:text-white transition-colors"
              aria-label={`Copy suggestion ${i + 1}`}
            >
              <Copy size={16} />
            </button>
          </li>
        ))}
      </ol>
    </div>
  );
}
