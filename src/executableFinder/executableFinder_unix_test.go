//go:build unix

package executableFinder

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createFakeExecutableUnix(filePath string) error {
	fakeExecFile, err := os.Create(filepath.Join(filePath))
	if err != nil {
		return err
	}
	fakeExecFile.Close()
	
	err = os.Chmod(fakeExecFile.Name(), 0111)
	if err != nil {
		return err
	}

	return nil
}

// Prepare fake PATH directory with executable files
func prepareFakePathEnvUnix(t *testing.T, filesToCreate []string) {
	tempPathDir := t.TempDir()
	for _, fileToCreate := range filesToCreate {
		if fileToCreate != "" {
			err := createFakeExecutableUnix(filepath.Join(tempPathDir, fileToCreate))
			if err != nil {
				t.Error(err)
			}
		}
	}

	t.Setenv("PATH", tempPathDir)
}

func TestExecutableExists(t *testing.T) {
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

			prepareFakePathEnvUnix(t, fakeExecFileNames)

			assert.Equal(t, test.expected, executableExists(test.commandToCheck))
		})
	}
}

func TestCommandExists_CommandInPath(t *testing.T) {
	t.Run("Command exists in Path", func(t *testing.T) {

		var commandToCheck string = "test_command"
		fakeExecFileNames := []string{commandToCheck}
		prepareFakePathEnvUnix(t, fakeExecFileNames)

		assert.NoError(t, CommandExists(commandToCheck))
	})
}

func TestCommandExists_CommandNotInPath(t *testing.T) {
	t.Run("Command doesn't exist in Path", func(t *testing.T) {

		var commandToCheck string = "test_command"
		var expectedError error = fmt.Errorf("'%s' command not found in PATH", commandToCheck)
		var fakeExecFileNames []string
		prepareFakePathEnvUnix(t, fakeExecFileNames)

		err := CommandExists(commandToCheck)
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}

	})
}
