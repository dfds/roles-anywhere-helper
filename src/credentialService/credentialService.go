package credentialService

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dfds/roles-anywhere-helper/executableFinder"
	"github.com/dfds/roles-anywhere-helper/fileHandler"
	"github.com/dfds/roles-anywhere-helper/profileHandler"

	"gopkg.in/ini.v1"
)

const awsSignHelpName = "aws_signing_helper"
const awsCredDir = "/.aws/credentials"

var CredentialsFilePath = GetCredentialsFilePath()

type CredentialsFileTemplate struct {
	CredentialProcess string `ini:"credential_process,omitempty"`
	Region            string `ini:"region,omitempty"`
}

func Configure(profileName string, certificatePath string, privateKeyPath string, trustAnchorArn string, profileArn string, roleArn string, region string) {
	fmt.Println("Configuring credential file")

	err := executableFinder.CommandExists(awsSignHelpName)
	Check(err)

	file, err := fileHandler.CreateFile(GetCredentialsFilePath())
	
	if err != nil {
		fmt.Println("Credential file already exists:", err)
	} else {
		fmt.Println("Credential file written successfully!")
	}
	defer file.Close()
	profileName = profileHandler.SetProfileName(profileName)
	profileTemplate := ProcessCredentialProcessTemplate(certificatePath, privateKeyPath, trustAnchorArn, profileArn, roleArn, region)
	WriteIniFile(&profileTemplate, profileName)
	fmt.Printf("Profile %s set", profileName)
}

func GetCredentialsFilePath() string {
	homeDir, err := os.UserHomeDir()
	Check(err)
	return filepath.Join(homeDir, awsCredDir)
}

func ProcessCredentialProcessTemplate(certificatePath string, privateKeyPath string, trustAnchorArn string, profileArn string, roleArn string, region string) CredentialsFileTemplate {
	profileTemplate := CredentialsFileTemplate{
		CredentialProcess: fmt.Sprintf(
			"%s credential-process --certificate %s --private-key %s --trust-anchor-arn %s --profile-arn %s --role-arn %s",
			awsSignHelpName,
			certificatePath,
			privateKeyPath,
			trustAnchorArn,
			profileArn,
			roleArn,
		),
		Region: region,
	}

	return profileTemplate
}

func CreateCredentialsFile(filePath string) (*os.File, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.MkdirAll(filePath, 0755)
	}

	_, e := os.Stat(filePath)

	if e == nil {
		return nil, os.ErrExist
	}
	return os.Create(filePath)
}

func WriteIniFile(template *CredentialsFileTemplate, profile string) {
	cfg, err := ini.Load(CredentialsFilePath)
	Check(err)
	RecreateSection(template, profile, cfg)
	cfg.SaveTo(CredentialsFilePath)
}

func RecreateSection(template *CredentialsFileTemplate, profile string, cfg *ini.File) {
	cfg.DeleteSection(profile)
	sec, err := cfg.NewSection(profile)
	Check(err)
	err = sec.ReflectFrom(template)
	Check(err)
}

func Check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
