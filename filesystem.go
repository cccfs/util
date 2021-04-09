package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Exists checks whether the file or directory exists.
func Exists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CopyFile copies file from src to dest
func CopyFile(srcFile, destFile string) error {
	if output, err := exec.Command("cp", "-f", srcFile, destFile).CombinedOutput(); err != nil {
		stderr := strings.TrimSpace(string(output))
		return fmt.Errorf("unable to copy file from %s -> %s (%s): %v", srcFile, destFile, stderr, err)
	}
	return nil
}

