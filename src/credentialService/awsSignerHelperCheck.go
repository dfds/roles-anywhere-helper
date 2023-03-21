package credentialService

import (
	"errors"
	"os/exec"
)

const awsSignHelpName = "aws_signing_helper"

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)

	return err == nil
}

func awsSignHelpExists() error {

	if !commandExists(awsSignHelpName) {
		return errors.New("AWS Signing Helper is not found in PATH")
	}

	return nil
}
