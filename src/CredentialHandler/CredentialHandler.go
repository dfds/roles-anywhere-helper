package CredentialHandler

import (
	"fmt"
	"path"

	"go-aws-iam-roles-anywhere-credential-helper/ProfileHandler"

	"os"

	"go.uber.org/zap"
	"gopkg.in/ini.v1"
)

var CredentialsFilePath = GetCredentialsFilePath()

type CredentialsFileTemplate struct {
	CredentialProcess string `ini:"credential_process,omitempty"`
	Region            string `ini:"region,omitempty"`
}

func GetCredentialsFilePath() string {
	homeDir, err := os.UserHomeDir()
	check(err)
	return homeDir + "/.aws/credentials"
}

func ProcessCredentialProcessTemplate(certificateDirectory string, privateKeyDirectory string, trustAnchorArn string, profileArn string, roleArn string, region string) CredentialsFileTemplate {
	profileTemplate := CredentialsFileTemplate{
		CredentialProcess: fmt.Sprintf("aws_signing_helper credential-process --certificate %s --private-key %s --trust-anchor-arn %s --profile-arn %s --role-arn %s", certificateDirectory, privateKeyDirectory, trustAnchorArn, profileArn, roleArn),
		Region:            region,
	}
	return profileTemplate
}

func Configure(profileName string) {
	linuxDir, err := os.UserHomeDir()

	file, err := Create(linuxDir+"/.aws", "credentials")

	if err != nil {
		fmt.Println("Error writing credential file:", err)
	} else {
		fmt.Println("Credential file written successfully!")
	}

	profileName = ProfileHandler.SetProfileName(profileName)
	// check for profile.

	// If exsits - replace

	// If not exisits - create

	defer file.Close()

}

func Create(filePath string, fileName string) (*os.File, error) {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.MkdirAll(filePath, 0700)
	}

	_, e := os.Stat(filePath + "/" + fileName)

	if e == nil {
		return nil, os.ErrExist
	}
	return os.Create(filePath + "/" + fileName)
}

func createCredentialsFile() {
	dir := path.Dir(CredentialsFilePath)
	err := os.MkdirAll(dir, 0755)
	check(err)
	f, err := os.OpenFile(CredentialsFilePath, os.O_CREATE, 0644)
	check(err)
	defer f.Close()
}

func writeIniFile(template *CredentialsFileTemplate, profile string) {
	cfg, err := ini.Load(CredentialsFilePath)
	check(err)

	recreateSection(template, profile, cfg)

	zap.S().Debugf("Saving ini file to %s", CredentialsFilePath)
	cfg.SaveTo(CredentialsFilePath)
}

func recreateSection(template *CredentialsFileTemplate, profile string, cfg *ini.File) {
	zap.S().Debugf("Deleting profile [%s] in credentials file", profile)
	cfg.DeleteSection(profile)
	sec, err := cfg.NewSection(profile)
	check(err)
	zap.S().Debugf("Reflecting profile [%s] in credentials file", profile)
	err = sec.ReflectFrom(template)
}

func WriteAWSCredentialsFile(template *CredentialsFileTemplate, profile string) {
	if !isFileOrFolderExisting(CredentialsFilePath) {
		createCredentialsFile()
	}
	writeIniFile(template, profile)
}

func Update() {
	fmt.Println("update called")
}

func Remove() {
	fmt.Println("remove")
}

func check(err error) {
	if err != nil {
		fmt.Println("Error writing credential file:", err)
	}
}
