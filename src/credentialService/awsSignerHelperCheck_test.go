package credentialService

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func createFakeExecutable(filePath string) error {
	fakeExecFile, err := os.Create(filepath.Join(filePath))
	if err != nil {
		return err
	}

	err = os.Chmod(fakeExecFile.Name(), 0111)
	if err != nil {
		return err
	}

	return nil
}

// Prepare fake PATH directory with executable files
func prepareFakePathEnvironment(t *testing.T, filesToCreate []string) {
	tempPathDir := t.TempDir()
	for _, fileToCreate := range filesToCreate {
		if fileToCreate != "" {
			err := createFakeExecutable(filepath.Join(tempPathDir, fileToCreate))
			if err != nil {
				t.Error(err)
			}
		}
	}

	t.Setenv("PATH", tempPathDir)
}

func TestCommandExists(t *testing.T) {
	type commandExistsTestCase struct {
		name             string
		commandToCheck   string
		fakeExecFileName string
		expected         bool
	}
	var commandExistsTestCases = []commandExistsTestCase{
		{"Executable exists", "mytool", "mytool", true},
		{"Executable doesn't exist", "mytool", "", false},
	}
	for _, test := range commandExistsTestCases {
		t.Run(test.name, func(t *testing.T) {

			var fakeExecFileNames []string
			if test.fakeExecFileName != "" {
				fakeExecFileNames = append(fakeExecFileNames, test.fakeExecFileName)
			}

			prepareFakePathEnvironment(t, fakeExecFileNames)

			assert.Equal(t, test.expected, commandExists(test.commandToCheck))
		})
	}
}

func TestAwsSignHelpExists_SigningHelperInPath(t *testing.T) {
	t.Run("AWS Signing helper exists in Path", func(t *testing.T) {

		fakeExecFileNames := []string{"aws_signing_helper"}
		prepareFakePathEnvironment(t, fakeExecFileNames)

		assert.NoError(t, awsSignHelpExists())
	})
}

func TestAwsSignHelpExists_SigningHelperNotInPath(t *testing.T) {
	t.Run("AWS Signing helper doesn't exist in Path", func(t *testing.T) {

		var fakeExecFileNames []string
		prepareFakePathEnvironment(t, fakeExecFileNames)

		assert.Error(t, awsSignHelpExists())
	})
}
