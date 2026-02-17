# file-converter

A simple CLI tool to convert files written in Go.

## Supported conversions

- `.docx` → `.pdf`
- `.pdf` → `.docx`
- `.mp4` → `.mp3`

## Prerequisites

Install these tools first:

- [LibreOffice](https://www.libreoffice.org/download/download-libreoffice/) (for document conversion)
- [ffmpeg](https://ffmpeg.org/download.html) (for media conversion)

Make sure both are in your PATH:
\```
soffice --version
ffmpeg -version
\```

## Installation

### Option A — Download the binary

Go to [Releases](../../releases) and download the latest `.exe` for Windows.

### Option B — Build from source

\```bash
git clone https://github.com/motzman/file-converter.git
cd file-converter
go build -o file-converter.exe .
\```

## Usage

\```powershell

# DOCX to PDF

./file-converter.exe -input report.docx -to pdf

# PDF to DOCX

./file-converter.exe -input scan.pdf -to docx

# MP4 to MP3

./file-converter.exe -input video.mp4 -to mp3

# Custom output location

./file-converter.exe -input report.docx -to pdf -output "C:\Users\Moritz\Desktop"
\```
