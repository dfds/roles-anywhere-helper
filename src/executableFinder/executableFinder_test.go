package executableFinder

import (
	"testing"
)

type createFakeExec func(dir string, fileName string) error

// Prepare fake PATH directory with executable files
func prepareFakePathEnv(t *testing.T, filesToCreate []string, fn createFakeExec) {
	tempPathDir := t.TempDir()
	for _, fileToCreate := range filesToCreate {
		if fileToCreate != "" {
			err := fn(tempPathDir, fileToCreate)
			if err != nil {
				t.Error(err)
			}
		}
	}

	t.Setenv("PATH", tempPathDir)
}
