package utils

import (
	"bytes"
	"os/exec"
	"syscall"
)

// RunCommand runs a shell command
func RunCommand(command string, args []string) (outStr string, errStr string, exitCode int) {
	var outBuffer, errBuffer bytes.Buffer
	defaultFailedCode := 11
	cmd := exec.Command(command, args...)
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer
	err := cmd.Run()
	outStr, errStr = outBuffer.String(), errBuffer.String()
	if err != nil {
		exitError, ok := err.(*exec.ExitError)
		if ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			exitCode = defaultFailedCode
			if errStr == "" {
				errStr = err.Error()
			}
		}
	} else {
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}
	return
}
