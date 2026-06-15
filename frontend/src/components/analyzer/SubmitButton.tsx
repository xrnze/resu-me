import { Loader2 } from "lucide-react";

interface SubmitButtonProps {
  disabled: boolean;
  loading: boolean;
  onClick: () => void;
  tip: string;
}

export function SubmitButton({
  disabled,
  loading,
  onClick,
  tip,
}: SubmitButtonProps) {
  return (
    <button
      onClick={onClick}
      disabled={disabled || loading}
      className="btn-neo text-lg px-12 py-5 min-w-[320px]"
    >
      {loading ? (
        <span
          aria-live="polite"
          className="flex items-center justify-center gap-3 font-['Space_Grotesk'] font-bold uppercase"
        >
          <Loader2 size={20} className="animate-spin" />
          {tip}
        </span>
      ) : (
        "Analyze My Resume"
      )}
    </button>
  );
}
