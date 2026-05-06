import { Link } from '@tanstack/react-router';
import { ArrowLeft } from 'lucide-react';

export function AnalyzerNavbar() {
  return (
    <header className="sticky top-0 w-full border-b-4 border-black z-50 bg-[#FFDAB9] shadow-[4px_4px_0px_0px_rgba(0,0,0,1)] flex justify-between items-center h-20 px-6 md:px-12">
      <Link
        to="/"
        className="flex items-center gap-2 font-['Space_Grotesk'] font-bold text-black hover:opacity-70 transition-opacity"
      >
        <ArrowLeft size={20} />
        <span className="hidden sm:inline">Back</span>
      </Link>
      <div className="text-xl font-black italic tracking-tighter uppercase text-black">
        RESUME_ENGINE
      </div>
      <div className="w-20" />
    </header>
  );
}
