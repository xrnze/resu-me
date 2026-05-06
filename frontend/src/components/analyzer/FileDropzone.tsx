import { useRef, useState, type DragEvent, type ChangeEvent } from 'react';
import { Upload, FileText, X } from 'lucide-react';

interface FileDropzoneProps {
  onExtract: (file: File) => Promise<void>;
  onRemove: () => void;
  extracting: boolean;
  error: string;
  hasFile: boolean;
}

export function FileDropzone({ onExtract, onRemove, extracting, error, hasFile }: FileDropzoneProps) {
  const inputRef = useRef<HTMLInputElement>(null);
  const [isDragOver, setIsDragOver] = useState(false);
  const [fileName, setFileName] = useState('');
  const [fileSize, setFileSize] = useState('');

  const formatSize = (bytes: number) => {
    if (bytes < 1024) return `${bytes} B`;
    if (bytes < 1024 * 1024) return `${(bytes / 1024).toFixed(0)} KB`;
    return `${(bytes / (1024 * 1024)).toFixed(1)} MB`;
  };

  const handleFile = (file: File) => {
    setFileName(file.name);
    setFileSize(formatSize(file.size));
    onExtract(file);
  };

  const handleDrop = (e: DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setIsDragOver(false);
    const file = e.dataTransfer.files[0];
    if (file) handleFile(file);
  };

  const handleDragOver = (e: DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setIsDragOver(true);
  };

  const handleDragLeave = () => setIsDragOver(false);

  const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) handleFile(file);
  };

  const handleClick = () => inputRef.current?.click();

  if (hasFile) {
    return (
      <div className="card-neo flex items-center justify-between gap-4">
        <div className="flex items-center gap-4 min-w-0">
          <FileText size={24} className="shrink-0 text-black" />
          <div className="min-w-0">
            <p className="font-['Space_Grotesk'] font-bold text-black truncate">{fileName}</p>
            <p className="text-sm text-on-surface-variant">{fileSize}</p>
          </div>
        </div>
        <button
          onClick={onRemove}
          className="p-2 border-2 border-black hover:bg-black hover:text-white transition-colors"
          aria-label="Remove file"
        >
          <X size={16} />
        </button>
      </div>
    );
  }

  return (
    <div>
      <div
        onClick={handleClick}
        onDrop={handleDrop}
        onDragOver={handleDragOver}
        onDragLeave={handleDragLeave}
        className={`border-4 border-black p-10 text-center cursor-pointer transition-colors ${
          isDragOver ? 'bg-secondary-container' : 'bg-card'
        } ${extracting ? 'opacity-50 pointer-events-none' : ''}`}
        role="button"
        tabIndex={0}
        onKeyDown={(e) => { if (e.key === 'Enter' || e.key === ' ') handleClick(); }}
      >
        <input
          ref={inputRef}
          type="file"
          accept=".pdf,.docx"
          className="sr-only"
          onChange={handleInputChange}
          aria-hidden="true"
        />
        <Upload size={40} className="mx-auto mb-4 text-black" />
        <p className="font-['Space_Grotesk'] font-bold text-lg text-black uppercase">
          {extracting ? 'Reading file...' : 'Drop your resume here'}
        </p>
        <p className="text-on-surface-variant mt-1">
          or click to browse (PDF, DOCX — max 5MB)
        </p>
      </div>
      {error && (
        <p className="mt-2 font-['Space_Grotesk'] font-bold text-sm text-black bg-error px-4 py-2 border-2 border-black">
          {error}
        </p>
      )}
    </div>
  );
}
