//go:build windows

package executableFinder

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createFakeExecWin(dir, fileName string) error {
	var winExecExt = ".exe"
	fileName = fileName + winExecExt

	fakeExecFile, err := os.Create(filepath.Join(dir, fileName))
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

			prepareFakePathEnv(t, fakeExecFileNames, createFakeExecWin)

			assert.Equal(t, test.expected, executableExists(test.commandToCheck))
		})
	}
}

func TestCommandExists_CommandInPath(t *testing.T) {
	t.Run("Command exists in Path", func(t *testing.T) {

		var commandToCheck string = "test_command"
		fakeExecFileNames := []string{commandToCheck}
		prepareFakePathEnv(t, fakeExecFileNames, createFakeExecWin)

		err := CommandExists(commandToCheck)
		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}

func TestCommandExists_CommandNotInPath(t *testing.T) {
	t.Run("Command doesn't exist in Path", func(t *testing.T) {

		var commandToCheck string = "test_command"
		var expectedError error = fmt.Errorf("'%s' command not found in PATH", commandToCheck)
		var fakeExecFileNames []string
		prepareFakePathEnv(t, fakeExecFileNames, createFakeExecWin)

		err := CommandExists(commandToCheck)
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}

	})
}
