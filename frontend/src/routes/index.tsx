import { Header } from '@/components/Header';
import { HeroSection } from '@/components/HeroSection';
import { Features } from '@/components/Features';
import { TheProcess } from '@/components/TheProcess';
import { SuccessStories } from '@/components/SuccessStories';
import { Questions } from '@/components/Questions';
import { Footer } from '@/components/Footer';

export function LandingPage() {
  return (
    <>
      <Header />
      <main className="flex-1">
        <HeroSection />
        <Features />
        <TheProcess />
        <SuccessStories />
        <Questions />
      </main>
      <Footer />
    </>
  );
}
