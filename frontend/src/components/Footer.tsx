import { Share2, AtSign } from "lucide-react";

export function Footer() {
  return (
    <footer className="w-full border-t-4 border-black py-16 bg-white flex flex-col md:flex-row justify-between items-center px-8 md:px-16 gap-8">
      <div className="flex flex-col gap-4">
        <div className="text-xl font-black text-black uppercase">
          RESUME ENGINE
        </div>
        <p className="font-['Space_Grotesk'] font-bold uppercase text-sm text-black opacity-60">
          © 2024 RESUME ENGINE. NO MERCY. ALL CAREER GAINS.
        </p>
      </div>
      <div className="flex flex-wrap gap-8 justify-center">
        <a
          href="#"
          className="text-link-neo font-['Space_Grotesk'] font-bold uppercase text-sm text-black"
        >
          Terms of Service
        </a>
        <a
          href="#"
          className="text-link-neo font-['Space_Grotesk'] font-bold uppercase text-sm text-black"
        >
          Privacy Policy
        </a>
        <a
          href="#"
          className="text-link-neo font-['Space_Grotesk'] font-bold uppercase text-sm text-black"
        >
          Contact Support
        </a>
      </div>
      <div className="flex gap-4">
        <a
          href="#"
          className="btn-neo w-12 h-12 px-0 py-0 bg-primary-container flex items-center justify-center"
        >
          <Share2 size={20} className="text-black" />
        </a>
        <a
          href="#"
          className="btn-neo w-12 h-12 px-0 py-0 bg-tertiary-container flex items-center justify-center"
        >
          <AtSign size={20} className="text-black" />
        </a>
      </div>
    </footer>
  );
}
