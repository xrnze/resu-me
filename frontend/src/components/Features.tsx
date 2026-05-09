export function Features() {
  return (
    <section id="features" className="px-8 md:px-16 py-xl bg-surface">
      <h2 className="font-headline-lg text-headline-lg uppercase text-black mb-lg">Features</h2>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-md">
        {/* Card 1: Keyword Gap Analysis */}
        <div className="bg-white border-4 border-black shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] flex flex-col">
          <div className="h-16 bg-tertiary-container border-b-4 border-black flex items-center px-4">
            <span className="material-symbols-outlined text-black text-3xl">manage_search</span>
          </div>
          <div className="p-8 space-y-sm flex-1">
            <h3 className="font-headline-md text-headline-md uppercase text-black">Keyword Gap Analysis</h3>
            <p className="font-body-md text-body-md text-black">
              We compare your resume against the job description line by line. Every missing keyword, every overlooked skill &mdash; surfaced instantly.
            </p>
          </div>
        </div>

        {/* Card 2: Section Feedback */}
        <div className="bg-white border-4 border-black shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] flex flex-col">
          <div className="h-16 bg-primary-container border-b-4 border-black flex items-center px-4">
            <span className="material-symbols-outlined text-black text-3xl">rate_review</span>
          </div>
          <div className="p-8 space-y-sm flex-1">
            <h3 className="font-headline-md text-headline-md uppercase text-black">Section Feedback</h3>
            <p className="font-body-md text-body-md text-black">
              Generic summaries. Weak metrics. Vague bullets. We call out exactly what&apos;s wrong in each section so you know where to focus.
            </p>
          </div>
        </div>

        {/* Card 3: Rewrite Suggestions */}
        <div className="bg-white border-4 border-black shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] flex flex-col">
          <div className="h-16 bg-secondary-fixed border-b-4 border-black flex items-center px-4">
            <span className="material-symbols-outlined text-black text-3xl">edit_note</span>
          </div>
          <div className="p-8 space-y-sm flex-1">
            <h3 className="font-headline-md text-headline-md uppercase text-black">Rewrite Suggestions</h3>
            <p className="font-body-md text-body-md text-black">
              No vague advice. We hand you rewritten bullet points you can copy-paste directly into your resume.
            </p>
          </div>
        </div>
      </div>
    </section>
  );
}