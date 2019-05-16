package util

import (
	"os/exec"
)

func CopyDir(src string, dst string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	cmd := exec.Command("cp", "-r", src, dst)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
