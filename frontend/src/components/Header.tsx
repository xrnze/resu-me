export function Header() {
  return (
    <header className="sticky top-0 w-full border-b-4 border-black z-50 bg-[#FFDAB9] shadow-[4px_4px_0px_0px_rgba(0,0,0,1)] flex justify-between items-center h-24 px-8 md:px-16">
      <div className="text-2xl font-black italic tracking-tighter uppercase text-black">
        RESUME_ENGINE
      </div>
      <nav className="hidden md:flex gap-8 items-center">
        <a
          href="#features"
          className="font-['Space_Grotesk'] uppercase font-bold tracking-tighter text-black hover:bg-[#E0BBE4] hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all px-2"
        >
          Features
        </a>
        <a
          href="#process"
          className="font-['Space_Grotesk'] uppercase font-bold tracking-tighter text-black hover:bg-[#E0BBE4] hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all px-2"
        >
          How It Works
        </a>
        <a
          href="#reviews"
          className="font-['Space_Grotesk'] uppercase font-bold tracking-tighter text-black hover:bg-[#E0BBE4] hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all px-2"
        >
          Reviews
        </a>
        <a
          href="#faq"
          className="font-['Space_Grotesk'] uppercase font-bold tracking-tighter text-black hover:bg-[#E0BBE4] hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all px-2"
        >
          FAQ
        </a>
      </nav>
      <button className="bg-black text-white font-['Space_Grotesk'] font-black uppercase px-6 py-3 border-2 border-black shadow-[4px_4px_0px_0px_rgba(0,0,0,1)] hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none active:translate-x-[4px] active:translate-y-[4px] transition-all">
        ANALYZE NOW
      </button>
    </header>
  );
}
