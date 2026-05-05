export function Features() {
  return (
    <section id="features" className="px-8 md:px-16 py-xl bg-surface">
      <div className="grid grid-cols-1 md:grid-cols-3 gap-md">
        {/* Card 1: Keyword Extraction */}
        <div className="bg-white border-4 border-black shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] flex flex-col">
          <div className="h-16 bg-tertiary-container border-b-4 border-black flex items-center px-4">
            <span className="material-symbols-outlined text-black text-3xl">key</span>
          </div>
          <div className="p-8 space-y-sm flex-1">
            <h3 className="font-headline-md text-headline-md uppercase text-black">Keyword Extraction</h3>
            <p className="font-body-md text-body-md text-black">
              We scan thousands of job descriptions to find the exact power words hiring managers are looking for. No fluff, just results.
            </p>
            <div className="flex flex-wrap gap-xs pt-4">
              <span className="bg-secondary-container border-2 border-black px-2 font-label-bold text-black">STRATEGY</span>
              <span className="bg-secondary-container border-2 border-black px-2 font-label-bold text-black">EXECUTION</span>
              <span className="bg-secondary-container border-2 border-black px-2 font-label-bold text-black">ROI</span>
            </div>
          </div>
        </div>

        {/* Card 2: Format Optimization */}
        <div className="bg-white border-4 border-black shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] flex flex-col">
          <div className="h-16 bg-primary-container border-b-4 border-black flex items-center px-4">
            <span className="material-symbols-outlined text-black text-3xl">dashboard_customize</span>
          </div>
          <div className="p-8 space-y-sm flex-1">
            <h3 className="font-headline-md text-headline-md uppercase text-black">Format Optimization</h3>
            <p className="font-body-md text-body-md text-black">
              Ugly resumes get tossed. We restructure your layout for maximum readability and visual impact in under 6 seconds.
            </p>
          </div>
        </div>

        {/* Card 3: ATS Score Prediction */}
        <div className="bg-white border-4 border-black shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] flex flex-col">
          <div className="h-16 bg-secondary-fixed border-b-4 border-black flex items-center px-4">
            <span className="material-symbols-outlined text-black text-3xl">query_stats</span>
          </div>
          <div className="p-8 space-y-sm flex-1">
            <h3 className="font-headline-md text-headline-md uppercase text-black">ATS Score Prediction</h3>
            <p className="font-body-md text-body-md text-black">
              Know exactly how you rank against the bots. Our score gauge tells you if you&apos;re passing or failing before you hit apply.
            </p>
          </div>
        </div>
      </div>
    </section>
  );
}