package kubectl

import (
	"bytes"
	"os/exec"
)

func Kubectl(command ...string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("./cluster/kubectl.sh", command...)
	cmd.Dir = ".."
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

