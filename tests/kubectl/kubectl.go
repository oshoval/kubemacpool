package kubectl

import (
	"bytes"
	"os"
	"os/exec"
)

func Kubectl(command ...string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	var cmd *exec.Cmd

	kubectl, found := os.LookupEnv("KUBECTL")
	if found {
		cmd = exec.Command(kubectl, command...)
	} else {
		cmd = exec.Command("./cluster/kubectl.sh", command...)
		cmd.Dir = ".."
	}
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
