export function SuccessStories() {
  return (
    <section id="reviews" className="px-8 md:px-16 py-xl bg-[#E0BBE4]">
      <h2 className="font-headline-lg text-headline-lg uppercase text-black mb-lg">
        Success Stories
      </h2>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-md">
        {/* Mark R. */}
        <div className="bg-white border-4 border-black p-8 shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] space-y-sm">
          <div className="flex items-center gap-sm">
            <div className="w-16 h-16 border-4 border-black overflow-hidden bg-tertiary-container">
              <img
                alt="Portrait"
                className="w-full h-full object-cover"
                src="https://lh3.googleusercontent.com/aida-public/AB6AXuArM9zvrMHSEnUvsDsRtTQA9-U_1gvF2mjN_oKTm-enRHmq8DXNuOzpeQnHTmtgFRnJImRPicRmsrfPO7hVz0iPB2aM5bVl9Fn4tp4R_hNjUe_ONKYLkJcfRd8Y-2LLH_5GCOV5dnk2RPFogjrL30eMJ1u0S9SfqTMF4FhbnniHrp3wr2eRYAem3GAW0ExPxAIyEacG_UKQP3bxNPibRZMYrLKIGZBf8MvdUKR-l4GYZkTWwDjDL3StvtArVSYOWCEeSfymNrAzl52W"
              />
            </div>
            <div>
              <p className="font-label-bold uppercase text-black">
                MARK R. — Senior Developer
              </p>
              <div className="flex text-black">
                <span
                  className="material-symbols-outlined"
                  style={{ fontVariationSettings: "'FILL' 1" }}
                >
                  star
                </span>
                <span
                  className="material-symbols-outlined"
                  style={{ fontVariationSettings: "'FILL' 1" }}
                >
                  star
                </span>
                <span
                  className="material-symbols-outlined"
                  style={{ fontVariationSettings: "'FILL' 1" }}
                >
                  star
                </span>
                <span
                  className="material-symbols-outlined"
                  style={{ fontVariationSettings: "'FILL' 1" }}
                >
                  star
                </span>
                <span
                  className="material-symbols-outlined"
                  style={{ fontVariationSettings: "'FILL' 1" }}
                >
                  star
                </span>
              </div>
            </div>
          </div>
          <p className="font-body-lg italic text-black">
            &quot;I went from 0 callbacks to 5 interviews in one week. The ATS
            score prediction is scary accurate. Worth every penny.&quot;
          </p>
        </div>

        {/* Sarah L. */}
        <div className="bg-white border-4 border-black p-8 shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] space-y-sm">
          <div className="flex items-center gap-sm">
            <div className="w-16 h-16 border-4 border-black overflow-hidden bg-primary-container">
              <img
                alt="Portrait"
                className="w-full h-full object-cover"
                src="https://lh3.googleusercontent.com/aida-public/AB6AXuABEhfWIbXJqCVm4qWV2MGymoCNrgChJTVO22Xj2nwMjvRy-MuYi_6nRhVmhyKX5um-TZz03bZStqN9sRjMsiQKFjaCJDa5GRRsxJAQ1b8PDHw2ydTgy7U2rmzYbg-AGnoOk3QruB2hqBBkgUBnOgEaD3Yn1wkytPaltqXvWVvM4Ho7sNeW70vhpirv4lDnlLCrP-nhI10eav6EKOMiC2PmKj9HDvraTnqcQ-jDz5NgmoAkapw8PqfzCc0JQ43kgS02GBiBCvoMolSS"
              />
            </div>
            <div>
              <p className="font-label-bold uppercase text-black">
                SARAH L. — Marketing Lead
              </p>
              <div className="flex text-black">
                <span
                  className="material-symbols-outlined"
                  style={{ fontVariationSettings: "'FILL' 1" }}
                >
                  star
                </span>
                <span
                  className="material-symbols-outlined"
                  style={{ fontVariationSettings: "'FILL' 1" }}
                >
                  star
                </span>
                <span
                  className="material-symbols-outlined"
                  style={{ fontVariationSettings: "'FILL' 1" }}
                >
                  star
                </span>
                <span
                  className="material-symbols-outlined"
                  style={{ fontVariationSettings: "'FILL' 1" }}
                >
                  star
                </span>
                <span
                  className="material-symbols-outlined"
                  style={{ fontVariationSettings: "'FILL' 1" }}
                >
                  star
                </span>
              </div>
            </div>
          </div>
          <p className="font-body-lg italic text-black">
            &quot;The format optimization made my old resume look like a joke.
            The engine is fast, brutal, and effective.&quot;
          </p>
        </div>
      </div>
    </section>
  );
}
