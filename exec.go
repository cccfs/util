package util

import (
	"bytes"
	"fmt"
	"strings"
	"os/exec"
)

func RunCommand(command string) string {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Errorf(fmt.Sprint(err) + ": " + stderr.String())
	}
	return strings.Trim(stdout.String(), "\n")
}
