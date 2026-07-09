import { Link } from "@tanstack/react-router";
import { ArrowLeft } from "lucide-react";

export function AnalyzerNavbar() {
  return (
    <header className="sticky top-0 w-full border-b-4 border-black z-50 bg-white neobrutal-shadow flex justify-between items-center h-20 px-6 md:px-12">
      <Link
        to="/"
        className="flex items-center gap-2 font-['Space_Grotesk'] font-bold text-black"
      >
        <ArrowLeft size={20} />
        <span className="text-link-neo hidden sm:inline">Back</span>
      </Link>
      <div className="text-xl font-black italic tracking-tighter uppercase text-black">
        RESU_ME
      </div>
      <div className="w-20" />
    </header>
  );
}
