export function Header() {
  return (
    <header className="sticky top-0 w-full border-b-4 border-black z-50 bg-white shadow-[4px_4px_0px_0px_rgba(0,0,0,1)] flex justify-between items-center h-24 px-8 md:px-16">
      <a
        href="#hero"
        className="text-link-neo text-2xl font-black italic tracking-tighter uppercase text-black"
      >
        RESU_ME
      </a>
      <nav className="hidden md:flex gap-8 items-center">
        <a
          href="#features"
          className="text-link-neo font-['Space_Grotesk'] uppercase font-bold tracking-tighter text-black"
        >
          Features
        </a>
        <a
          href="#process"
          className="text-link-neo font-['Space_Grotesk'] uppercase font-bold tracking-tighter text-black"
        >
          How It Works
        </a>
        <a
          href="#reviews"
          className="text-link-neo font-['Space_Grotesk'] uppercase font-bold tracking-tighter text-black"
        >
          Reviews
        </a>
        <a
          href="#faq"
          className="text-link-neo font-['Space_Grotesk'] uppercase font-bold tracking-tighter text-black"
        >
          FAQ
        </a>
      </nav>
    </header>
  );
}
