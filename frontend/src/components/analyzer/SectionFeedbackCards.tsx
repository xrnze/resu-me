interface SectionFeedbackCardsProps {
  feedback: {
    summary: string;
    experience: string;
    skills: string;
  };
}

const SECTIONS: { key: keyof SectionFeedbackCardsProps['feedback']; label: string }[] = [
  { key: 'summary', label: 'Summary' },
  { key: 'experience', label: 'Experience' },
  { key: 'skills', label: 'Skills' },
];

export function SectionFeedbackCards({ feedback }: SectionFeedbackCardsProps) {
  return (
    <div>
      <h2 className="font-['Space_Grotesk'] font-bold text-lg uppercase text-black mb-4">
        Section Feedback
      </h2>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        {SECTIONS.map(({ key, label }) => (
          <div key={key} className="card-neo">
            <h3 className="font-['Space_Grotesk'] font-bold text-base uppercase text-black mb-2">
              {label}
            </h3>
            <p className="font-body-md text-body-md text-on-surface-variant leading-relaxed">
              {feedback[key]}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
}
