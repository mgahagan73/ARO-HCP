package cmdline

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// From rosa e2e repo
func RunCMD(cmd string) (stdout string, stderr string, err error) {
	var stdoutput bytes.Buffer
	var stderroutput bytes.Buffer
	CMD := exec.Command("bash", "-c", cmd)
	CMD.Stderr = &stderroutput
	CMD.Stdout = &stdoutput
	err = CMD.Run()
	stdout = strings.TrimPrefix(stdoutput.String(), "\n")
	stderr = strings.TrimPrefix(stderroutput.String(), "\n")
	if err != nil {
		err = fmt.Errorf("%s:%s", err.Error(), stderr)
	}
	return
}
