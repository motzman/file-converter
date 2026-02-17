package converters

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func DOCXToPDF(inputPath, outputDir string) error {
	if err := checkTool("soffice"); err != nil {
		return err
	}
	inputPath, _ = filepath.Abs(inputPath)

	if outputDir == "" {
		outputDir = filepath.Dir(inputPath)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("could not create output: %w", err)
	}

	cmd := exec.Command("soffice",
		"--headless",
		"--norestore",
		"--convert-to", "pdf",
		"--outdir", outputDir,
		inputPath,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("libreoffice error: %w\n%s", err, string(out))
	}
	return nil
}

func PDFToDOCX(inputPath, outputDir string) error {
	if err := checkTool("soffice"); err != nil {
		return err
	}
	inputPath, _ = filepath.Abs(inputPath)

	if outputDir == "" {
		outputDir = filepath.Dir(inputPath)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("could not create output: %w", err)
	}

	cmd := exec.Command("soffice",
		"--headless",
		"--norestore",
		"--infilter=writer_pdf_import",
		"--outdir", outputDir,
		inputPath,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("libreoffice error: %w\n%s", err, string(out))
	}
	return nil
}

func ExpectedOutputPath(inputPath, newExt, outputDir string) string {
	base := filepath.Base(inputPath)
	nameWithoutExt := strings.TrimSuffix(base, filepath.Ext(base))
	if outputDir == "" {
		outputDir = filepath.Dir(inputPath)
	}
	return filepath.Join(outputDir, nameWithoutExt+"."+newExt)
}
