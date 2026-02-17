package main

import (
	"file-converter/converters"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	input := flag.String("input", "", "Input file path (req)")
	output := flag.String("output", "", "Output file path (opt)")
	to := flag.String("to", "", "Target format: pdf, docx, mp3 (req)")

	flag.Usage = func() {
		fmt.Println("Usage: file-converter -input <file> -to <format> [-output <path>]")
		fmt.Println("\nSupported conversions:")
		fmt.Println("  .docx  -> pdf")
		fmt.Println("  .pdf   -> docx")
		fmt.Println("  .mp4   -> mp3")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *input == "" || *to == "" {
		flag.Usage()
		os.Exit(1)
	}
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(*input), "."))
	target := strings.ToLower(*to)

	fmt.Printf("converting %s to %s \n", ext, target)

	var err error

	switch ext + "->" + target {
	case "docx->pdf":
		err = converters.DOCXToPDF(*input, *output)
		if err == nil {
			fmt.Println("Output: ", converters.ExpectedOutputPath(*input, "pdf", *output))
		}
	case "pdf->docx":
		err = converters.PDFToDOCX(*input, *output)
		if err == nil {
			fmt.Println("Output: ", converters.ExpectedOutputPath(*input, "docx", *output))
		}
	case "mp4->mp3":
		err = converters.MP4ToMP3(*input, *output)
		if err == nil {
			outPath := *output
			if outPath == "" {
				outPath = converters.ExpectedOutputPath(*input, "mp3", "")
			}
			fmt.Println("Output: ", outPath)
		}
	default:
		fmt.Printf("Not supported: .%s -> .%s\n", ext, target)
		flag.Usage()
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Println("Done")
}
