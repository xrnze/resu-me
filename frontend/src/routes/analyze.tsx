import { useState } from "react";
import { Toaster, toast } from "sonner";
import { extractText } from "@/lib/extract-text";
import { analyzeResume } from "@/lib/api";
import type { AnalysisResponse } from "@/types";
import { AnalyzerNavbar } from "@/components/analyzer/AnalyzerNavbar";
import { FileDropzone } from "@/components/analyzer/FileDropzone";
import { JobDescriptionInput } from "@/components/analyzer/JobDescriptionInput";
import { SubmitButton } from "@/components/analyzer/SubmitButton";
import { useLoadingTips } from "@/hooks/useLoadingTips";
import { ResultsSection } from "@/components/analyzer/ResultsSection";

export function AnalyzePage() {
  const [resumeText, setResumeText] = useState("");
  const [jobDescription, setJobDescription] = useState("");
  const [isExtracting, setIsExtracting] = useState(false);
  const [isAnalyzing, setIsAnalyzing] = useState(false);
  const [extractError, setExtractError] = useState("");
  const { tip, start, clear } = useLoadingTips();
  const [results, setResults] = useState<AnalysisResponse | null>(null);

  const handleExtract = async (file: File) => {
    setIsExtracting(true);
    setExtractError("");
    try {
      const text = await extractText(file);
      setResumeText(text);
    } catch (e) {
      const msg =
        e instanceof Error
          ? e.message
          : "Could not read this file. Please try another PDF or DOCX.";
      setExtractError(msg);
    } finally {
      setIsExtracting(false);
    }
  };

  const handleRemove = () => {
    setResumeText("");
    setExtractError("");
  };

  const handleSubmit = async () => {
    setIsAnalyzing(true);
    start();
    try {
      const data = await analyzeResume(resumeText, jobDescription);
      setResults(data);
    } catch (e) {
      const msg =
        e instanceof Error
          ? e.message
          : "Unable to reach the analysis server. Please try again.";
      toast.error(msg);
    } finally {
      clear();
      setIsAnalyzing(false);
    }
  };

  const handleReset = () => {
    clear();
    setResumeText("");
    setJobDescription("");
    setIsExtracting(false);
    setIsAnalyzing(false);
    setExtractError("");
    setResults(null);
    window.scrollTo({ top: 0, behavior: "smooth" });
  };

  const canSubmit =
    resumeText.length > 0 && jobDescription.length > 0 && !isExtracting;

  return (
    <div className="min-h-screen flex flex-col bg-surface">
      <AnalyzerNavbar />
      <Toaster
        position="top-right"
        toastOptions={{
          style: {
            border: "2px solid black",
            borderRadius: "0",
            fontFamily: "Work Sans, system-ui, sans-serif",
          },
        }}
      />
      <main className="flex-1 w-full max-w-4xl mx-auto px-6 py-12">
        {results ? (
          <ResultsSection results={results} onReset={handleReset} />
        ) : (
          <div className="space-y-10">
            <div className="text-center">
              <h1 className="font-headline-md text-headline-md uppercase tracking-tighter text-black">
                ANALYZE YOUR RESUME
              </h1>
              <p className="font-body-lg text-body-lg text-on-surface-variant mt-2">
                Upload your resume, paste the job description, and get instant
                feedback.
              </p>
            </div>

            <FileDropzone
              onExtract={handleExtract}
              onRemove={handleRemove}
              extracting={isExtracting}
              error={extractError}
              hasFile={resumeText.length > 0}
            />

            <JobDescriptionInput
              value={jobDescription}
              onChange={setJobDescription}
            />

            <div className="flex justify-center">
              <SubmitButton
                disabled={!canSubmit}
                loading={isAnalyzing}
                tip={tip}
                onClick={handleSubmit}
              />
            </div>

          </div>
        )}
      </main>
    </div>
  );
}
