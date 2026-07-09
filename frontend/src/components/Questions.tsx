import { useState } from "react";
import { Plus, Minus } from 'lucide-react';

const questions = [
  {
    question: "Is my data secure?",
    answer:
      "Your resume and job description are sent securely to our analysis engine and never stored on our servers. We don't keep copies of your data after the analysis completes.",
  },
  {
    question: "Does it work for tech jobs?",
    answer:
      "Yes. We specialize in tech, finance, and marketing sectors where ATS systems are most aggressive.",
  },
  {
    question: "Can I use it more than once?",
    answer:
      "Absolutely. There's no limit, analyze as many resumes and job descriptions as you need, completely free.",
  },
];

export function Questions() {
  const [openIndex, setOpenIndex] = useState<number | null>(null);

  const toggle = (index: number) => {
    setOpenIndex(openIndex === index ? null : index);
  };

  return (
    <section
      id="faq"
      className="px-8 md:px-16 py-xl bg-surface border-t-4 border-black"
    >
      <h2 className="font-headline-lg text-headline-lg uppercase text-black mb-lg">
        Questions?
      </h2>
      <div className="max-w-4xl space-y-sm">
        {questions.map((q, index) => {
          const isOpen = openIndex === index;
          return (
            <div
              key={index}
              className={`border-4 border-black p-6 neobrutal-shadow ${
                isOpen ? "bg-primary-container" : "bg-white"
              }`}
            >
              <div
                className="flex justify-between items-center cursor-pointer"
                onClick={() => toggle(index)}
              >
                <h4 className="font-headline-md text-headline-md uppercase text-black">
                  {q.question}
                </h4>
                <span className="text-4xl text-black">
                  {isOpen ? <Minus size={36} /> : <Plus size={36} />}
                </span>
              </div>
              {isOpen && q.answer && (
                <div className="font-body-md mt-4 text-black border-t-2 border-black pt-4">
                  {q.answer}
                </div>
              )}
            </div>
          );
        })}
      </div>
    </section>
  );
}
