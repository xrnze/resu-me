import { Link } from "@tanstack/react-router";

export function HeroSection() {
  return (
    <section
      id="hero"
      className="bg-white px-8 md:px-16 py-xl border-b-4 border-black flex flex-col md:flex-row items-center gap-md"
    >
      <div className="flex-1 space-y-md">
        <h1 className="font-display-2xl text-display-2xl uppercase italic tracking-tighter text-black">
          REWRITE YOUR FUTURE.{" "}
          <span className="bg-black text-primary-container px-4">
            NO MERCY.
          </span>
        </h1>
        <p className="font-body-lg text-body-lg max-w-2xl text-black">
          Stop begging for interviews. Our AI-powered engine dismantles your
          weak resume and rebuilds it into a high-conversion career weapon that
          bypasses ATS filters instantly.
        </p>
        <div className="flex gap-sm">
          <Link
            to="/analyze"
            className="btn-neo font-headline-md text-headline-md px-10 py-6"
          >
            ANALYZE NOW
          </Link>
        </div>
      </div>
      <div className="flex-1 w-full md:flex justify-end hidden">
        <div className="w-full max-w-120 aspect-square bg-white border-4 border-black shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] overflow-hidden relative">
          <img
            alt="A high-contrast, neobrutalist digital illustration of a stylized paper resume being scanned by glowing laser lines."
            className="w-full h-full object-cover grayscale contrast-125"
            src="https://lh3.googleusercontent.com/aida-public/AB6AXuDBNcgzdhOLNyuoQXVwo7gCuYDTUMUYtD9JC6P8H7bMMSNJE4rQsgzlIIaetSk7Kh4haGCumWYm8oHphOe9Qgy84zsh9LyJt_VR4nQfrnasfR8KeF9eygsM5SMj18fC0GGdZQBrfPte1_E3IyIR3-3yVyEqMH5N1tG0XLY-ymzNcGHV05W9HQdrWWutXxCecmImP4g-NZziHG3vJhpRKkO3cePhx85G1938iU3uHoNsiRa8u3H1ghpuNhlEIV6CEx5_1jS1ga7eCGzy"
          />
          <div className="absolute inset-0 border-4 border-black pointer-events-none"></div>
        </div>
      </div>
    </section>
  );
}
