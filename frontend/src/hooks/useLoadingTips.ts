import { useEffect, useRef, useState } from "react";

const TIPS = [
  "ANALYZING KEYWORDS...",
  "COMPARING EXPERIENCE...",
  "CALCULATING MATCH SCORE...",
];

export function useLoadingTips(intervalMs = 3000) {
  const [index, setIndex] = useState(0);
  const idRef = useRef<ReturnType<typeof setInterval> | null>(null);

  const clear = () => {
    if (idRef.current) {
      clearInterval(idRef.current);
      idRef.current = null;
    }
    setIndex(0);
  };

  const start = () => {
    clear();
    idRef.current = setInterval(() => {
      setIndex((i) => (i + 1 < TIPS.length ? i + 1 : i));
    }, intervalMs);
  };

  useEffect(() => clear, []);

  return { tip: TIPS[index], start, clear };
}
