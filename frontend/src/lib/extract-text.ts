const MAX_SIZE = 5 * 1024 * 1024;

const VALID_TYPES = new Set([
  'application/pdf',
  'application/vnd.openxmlformats-officedocument.wordprocessingml.document',
]);

let workerConfigured = false;

export async function extractText(file: File): Promise<string> {
  if (!VALID_TYPES.has(file.type)) {
    throw new Error('Please upload a PDF or DOCX file.');
  }

  if (file.size > MAX_SIZE) {
    throw new Error('File size must be under 5MB.');
  }

  const buffer = await file.arrayBuffer();

  if (file.type === 'application/pdf') {
    const pdfjsLib = await import('pdfjs-dist');

    if (!workerConfigured) {
      pdfjsLib.GlobalWorkerOptions.workerSrc = '/pdf.worker.min.js';
      workerConfigured = true;
    }

    const pdf = await pdfjsLib.getDocument({ data: new Uint8Array(buffer) }).promise;
    const pages: string[] = [];

    for (let i = 1; i <= pdf.numPages; i++) {
      const page = await pdf.getPage(i);
      const content = await page.getTextContent();
      const text = content.items
        .filter((item) => 'str' in item)
        .map((item) => (item as { str: string }).str)
        .join(' ');
      pages.push(text);
    }

    return pages.join('\n');
  }

  const mammoth = await import('mammoth');
  const result = await mammoth.extractRawText({ arrayBuffer: buffer });
  return result.value;
}
