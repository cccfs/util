package util

import (
	"bytes"
	"github.com/fatih/color"
	"os/exec"
	"strings"
)

func RunCommand(command string) string {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		//fmt.Errorf(fmt.Sprint(err) + ": " + stderr.String())
		color.HiRed("%s : %s", err, stderr.String())
		return stderr.String()
	}
	return strings.Trim(stdout.String(), "\n")
}
