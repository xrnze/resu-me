interface JobDescriptionInputProps {
  value: string;
  onChange: (value: string) => void;
}

const MAX_LENGTH = 5000;

export function JobDescriptionInput({ value, onChange }: JobDescriptionInputProps) {
  return (
    <div>
      <label
        htmlFor="job-desc"
        className="block font-['Space_Grotesk'] font-bold text-lg uppercase text-black mb-2"
      >
        Job Description
      </label>
      <textarea
        id="job-desc"
        value={value}
        onChange={(e) => onChange(e.target.value.slice(0, MAX_LENGTH))}
        placeholder="Paste the job description here..."
        className="w-full min-h-[200px] p-4 border-4 border-black bg-card text-black font-body-md text-body-md resize-y focus:outline-none focus:ring-2 focus:ring-black"
      />
      <div className="flex justify-end mt-1">
        <span className="font-['Space_Grotesk'] text-sm font-bold text-on-surface-variant">
          {value.length} / {MAX_LENGTH}
        </span>
      </div>
    </div>
  );
}
