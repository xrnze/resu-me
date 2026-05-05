export function Footer() {
  return (
    <footer className="w-full border-t-4 border-black py-16 bg-white flex flex-col md:flex-row justify-between items-center px-8 md:px-16 gap-8">
      <div className="flex flex-col gap-4">
        <div className="text-xl font-black text-black uppercase">RESUME ENGINE</div>
        <p className="font-['Space_Grotesk'] font-bold uppercase text-sm text-black opacity-60">
          © 2024 RESUME ENGINE. NO MERCY. ALL CAREER GAINS.
        </p>
      </div>
      <div className="flex flex-wrap gap-8 justify-center">
        <a
          href="#"
          className="font-['Space_Grotesk'] font-bold uppercase text-sm text-black hover:text-[#E0BBE4] hover:skew-x-2 transition-transform"
        >
          Terms of Service
        </a>
        <a
          href="#"
          className="font-['Space_Grotesk'] font-bold uppercase text-sm text-black hover:text-[#E0BBE4] hover:skew-x-2 transition-transform"
        >
          Privacy Policy
        </a>
        <a
          href="#"
          className="font-['Space_Grotesk'] font-bold uppercase text-sm text-black hover:text-[#E0BBE4] hover:skew-x-2 transition-transform"
        >
          Contact Support
        </a>
      </div>
      <div className="flex gap-4">
        <a
          href="#"
          className="w-12 h-12 border-4 border-black bg-primary-container flex items-center justify-center shadow-[4px_4px_0px_0px_rgba(0,0,0,1)] hover:translate-x-1 hover:translate-y-1 hover:shadow-none transition-all"
        >
          <span className="material-symbols-outlined text-black">share</span>
        </a>
        <a
          href="#"
          className="w-12 h-12 border-4 border-black bg-tertiary-container flex items-center justify-center shadow-[4px_4px_0px_0px_rgba(0,0,0,1)] hover:translate-x-1 hover:translate-y-1 hover:shadow-none transition-all"
        >
          <span className="material-symbols-outlined text-black">alternate_email</span>
        </a>
      </div>
    </footer>
  );
}