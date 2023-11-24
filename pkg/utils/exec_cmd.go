package utils

import (
	"os/exec"

	"changeme/pkg/log"
)

func ExecCmdRetOut(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("exec cmd [%s %v] failed: [%v]", name, args, err)
		return "", err
	}
	log.Debugf("exec cmd [%s %v] output: [%s]", name, args, string(out))
	return string(out), nil
}
