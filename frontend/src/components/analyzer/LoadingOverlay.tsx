import { useEffect, useState } from "react";
import { Loader2 } from "lucide-react";

const TIPS = [
  "Analyzing keywords...",
  "Comparing experience...",
  "Calculating match score...",
];

export function LoadingOverlay() {
  const [tipIndex, setTipIndex] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setTipIndex((i) => (i + 1) % TIPS.length);
    }, 3000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center bg-black/50"
      role="status"
      aria-live="polite"
    >
      <div className="card-neo text-center max-w-2xl md:max-w-2xl w-full ">
        <Loader2 size={40} className="mx-auto mb-4 animate-spin text-black" />
        <p className="font-['Space_Grotesk'] font-bold text-lg uppercase text-black">
          {TIPS[tipIndex]}
        </p>
        <p className="text-on-surface-variant mt-1">
          Please wait while we analyze your resume...
        </p>
      </div>
    </div>
  );
}
