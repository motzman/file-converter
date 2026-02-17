package converters

import (
	"fmt"
	"os/exec"
)

func checkTool(tool string) error {
	_, err := exec.LookPath(tool)
	if err != nil {
		return fmt.Errorf("required tool '%s' not found", tool)
	}
	return nil
}
