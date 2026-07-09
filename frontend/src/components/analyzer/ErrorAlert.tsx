import { X } from "lucide-react";

interface ErrorAlertProps {
  message: string;
  onDismiss: () => void;
}

export function ErrorAlert({ message, onDismiss }: ErrorAlertProps) {
  return (
    <div
      role="alert"
      aria-live="assertive"
      className="relative bg-error px-4 py-3 border-2 border-black"
    >
      <p className="font-['Space_Grotesk'] font-bold text-sm text-black pr-8">
        {message}
      </p>
      <button
        onClick={onDismiss}
        aria-label="Dismiss error"
        className="absolute top-2 right-2 p-1 border-2 border-black hover:bg-black hover:text-white transition-colors"
      >
        <X size={14} />
      </button>
    </div>
  );
}
