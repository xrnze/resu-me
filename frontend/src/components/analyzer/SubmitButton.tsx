import { Loader2 } from "lucide-react";

interface SubmitButtonProps {
  disabled: boolean;
  loading: boolean;
  onClick: () => void;
}

export function SubmitButton({
  disabled,
  loading,
  onClick,
}: SubmitButtonProps) {
  return (
    <button
      onClick={onClick}
      disabled={disabled || loading}
      className={"btn-neo text-lg px-12 py-5"}
    >
      {loading ? (
        <span className="flex items-center gap-3">
          <Loader2 size={20} className="animate-spin" />
          Analyzing...
        </span>
      ) : (
        "Analyze My Resume"
      )}
    </button>
  );
}
