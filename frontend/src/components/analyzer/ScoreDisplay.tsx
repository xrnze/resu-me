interface ScoreDisplayProps {
  score: number;
}

function getScoreColor(score: number): string {
  if (score >= 80) return 'bg-success';
  if (score >= 50) return 'bg-warning';
  return 'bg-error';
}

export function ScoreDisplay({ score }: ScoreDisplayProps) {
  const colorClass = getScoreColor(score);

  return (
    <div className="card-neo text-center">
      <h2 className="font-headline-md font-bold text-sm uppercase text-on-surface-variant mb-4">
        Overall Match Score
      </h2>
      <div className={`inline-flex items-center justify-center w-32 h-32 border-4 border-black neobrutal-shadow ${colorClass}`}>
        <span className="font-headline-md text-4xl font-black text-black tracking-tighter">
          {score}
        </span>
        <span className="font-headline-md text-xl font-bold text-black/60 mt-2">
          /100
        </span>
      </div>
    </div>
  );
}
