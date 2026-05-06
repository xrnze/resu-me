interface ScoreDisplayProps {
  score: number;
}

function getScoreColor(score: number): string {
  if (score >= 80) return 'bg-[#22c55e]';
  if (score >= 50) return 'bg-[#f59e0b]';
  return 'bg-[#ef4444]';
}

export function ScoreDisplay({ score }: ScoreDisplayProps) {
  const colorClass = getScoreColor(score);

  return (
    <div className="card-neo text-center">
      <h2 className="font-['Space_Grotesk'] font-bold text-sm uppercase text-on-surface-variant mb-4">
        Overall Match Score
      </h2>
      <div className={`inline-flex items-center justify-center w-32 h-32 border-4 border-black ${colorClass}`}>
        <span className="font-['Space_Grotesk'] text-4xl font-black text-black tracking-tighter">
          {score}
        </span>
        <span className="font-['Space_Grotesk'] text-xl font-bold text-black/60 mt-2">
          /100
        </span>
      </div>
    </div>
  );
}
