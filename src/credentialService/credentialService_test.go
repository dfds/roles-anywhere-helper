package credentialService

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCredentialsFilePath(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("No access to 'HOME' directory: ", err)
	}
	expected := home
	actual := GetDefaultCredentialsFilePath()
	assert.Equal(t, expected, actual)
}

func TestProcessCredentialProcessTemplate(t *testing.T) {
	expected := CredentialsFileTemplate{
		CredentialProcess: "aws_signing_helper credential-process --certificate cert --private-key key --trust-anchor-arn arn --profile-arn arn --role-arn arn",
		Region:            "eu-west-1",
	}
	actual := ProcessCredentialProcessTemplate("cert", "key", "arn", "arn", "arn", "eu-west-1")
	assert.Equal(t, expected, actual)
}

func TestCreateCredentialsFile(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir)
	file, err := CreateCredentialsFile(filePath)
	defer file.Close()

	assert.Error(t, err)
	assert.DirExists(t, filePath)
}
