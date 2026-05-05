export function TheProcess() {
  return (
    <section id="process" className="px-8 md:px-16 py-xl border-y-4 border-black bg-white overflow-hidden">
      <div className="mb-lg">
        <h2 className="font-headline-lg text-headline-lg uppercase text-black inline-block border-b-8 border-primary-container">
          The Process
        </h2>
      </div>
      <div className="relative">
        <div className="hidden md:block absolute top-1/2 left-0 w-full h-1 bg-black -translate-y-1/2 -z-10"></div>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-xl">
          {/* Step 1: Upload */}
          <div className="flex flex-col items-center text-center space-y-md">
            <div className="w-24 h-24 bg-primary-container border-4 border-black rounded-full flex items-center justify-center font-display-2xl text-black shadow-[4px_4px_0px_0px_rgba(0,0,0,1)]">
              1
            </div>
            <h4 className="font-headline-md text-headline-md uppercase text-black">Upload</h4>
            <p className="font-body-md text-black">
              Drag your dusty PDF or Word doc into the engine. We handle the rest.
            </p>
          </div>

          {/* Step 2: Analyze */}
          <div className="flex flex-col items-center text-center space-y-md">
            <div className="w-24 h-24 bg-tertiary-container border-4 border-black rounded-full flex items-center justify-center font-display-2xl text-black shadow-[4px_4px_0px_0px_rgba(0,0,0,1)]">
              2
            </div>
            <h4 className="font-headline-md text-headline-md uppercase text-black">Analyze</h4>
            <p className="font-body-md text-black">
              Our AI shreds your content and identifies every missed opportunity.
            </p>
          </div>

          {/* Step 3: Conquer */}
          <div className="flex flex-col items-center text-center space-y-md">
            <div className="w-24 h-24 bg-secondary-fixed border-4 border-black rounded-full flex items-center justify-center font-display-2xl text-black shadow-[4px_4px_0px_0px_rgba(0,0,0,1)]">
              3
            </div>
            <h4 className="font-headline-md text-headline-md uppercase text-black">Conquer</h4>
            <p className="font-body-md text-black">
              Download your high-powered resume and go claim your new salary.
            </p>
          </div>
        </div>
      </div>
    </section>
  );
}