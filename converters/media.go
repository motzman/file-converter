package converters

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func MP4ToMP3(inputPath, outputPath string) error {
	if err := checkTool("ffmpeg"); err != nil {
		return err
	}

	inputPath, _ = filepath.Abs(inputPath)

	if outputPath == "" {
		base := strings.TrimSuffix(inputPath, filepath.Ext(inputPath))
		outputPath = base + ".mp3"
	}

	cmd := exec.Command("ffmpeg",
		"-i", inputPath, //? input file
		"-q:a", "0", //? best audio qualitity
		"-map", "a", //? extract only audio
		"-y", //? overwrite output if exists
		outputPath,
	)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("ffmpeg error : %w\n%s", err, string(out))
	}
	return nil
}
