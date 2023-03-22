package executableFinder

import (
	"fmt"
	"os/exec"
)


func executableExists(name string) bool {
	_, err := exec.LookPath(name)

	return err == nil
}

func CommandExists(cmd string) error {

	if !executableExists(cmd) {
		return fmt.Errorf("'%s' command not found in PATH", cmd)
	}

	return nil
}
