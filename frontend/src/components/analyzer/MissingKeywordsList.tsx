interface MissingKeywordsListProps {
  keywords: string[];
}

export function MissingKeywordsList({ keywords }: MissingKeywordsListProps) {
  if (keywords.length === 0) {
    return (
      <div className="card-neo">
        <h3 className="font-['Space_Grotesk'] font-bold text-lg uppercase text-black mb-2">
          Missing Keywords
        </h3>
        <p className="text-on-surface-variant font-body-md text-body-md">
          No missing keywords detected. Great job!
        </p>
      </div>
    );
  }

  return (
    <div className="card-neo">
      <h3 className="font-['Space_Grotesk'] font-bold text-lg uppercase text-black mb-4">
        Missing Keywords
      </h3>
      <div className="flex flex-wrap gap-3">
        {keywords.map((kw) => (
          <span key={kw} className="tag-neo bg-surface-container">
            {kw}
          </span>
        ))}
      </div>
    </div>
  );
}
