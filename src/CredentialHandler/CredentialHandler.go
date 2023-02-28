package CredentialHandler

import (
	"fmt"

	"go-aws-iam-roles-anywhere-credential-helper/ProfileHandler"

	"os"

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

func Configure(profileName string, certificateDirectory string, privateKeyDirectory string, trustAnchorArn string, profileArn string, roleArn string, region string) {
	file, err := createCredentialsFile(GetCredentialsFilePath()+"/.aws", "credentials")

	if err != nil {
		fmt.Println("Error writing credential file:", err)
	} else {
		fmt.Println("Credential file written successfully!")
	}

	profileName = ProfileHandler.SetProfileName(profileName)
	profileTemplate := ProcessCredentialProcessTemplate(certificateDirectory, privateKeyDirectory, trustAnchorArn, profileArn, roleArn, region)
	writeIniFile(profileTemplate, profileName)

	defer file.Close()
}

func ProcessCredentialProcessTemplate(certificateDirectory string, privateKeyDirectory string, trustAnchorArn string, profileArn string, roleArn string, region string) CredentialsFileTemplate {
	profileTemplate := CredentialsFileTemplate{
		CredentialProcess: fmt.Sprintf("aws_signing_helper credential-process --certificate %s --private-key %s --trust-anchor-arn %s --profile-arn %s --role-arn %s", certificateDirectory, privateKeyDirectory, trustAnchorArn, profileArn, roleArn),
		Region:            region,
	}
	return profileTemplate
}

func createCredentialsFile(filePath string, fileName string) (*os.File, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.MkdirAll(filePath, 0700)
	}

	_, e := os.Stat(filePath + "/" + fileName)

	if e == nil {
		return nil, os.ErrExist
	}
	return os.Create(filePath + "/" + fileName)
}

func writeIniFile(template CredentialsFileTemplate, profile string) {
	cfg, err := ini.Load(CredentialsFilePath)
	check(err)
	recreateSection(template, profile, cfg)
	cfg.SaveTo(CredentialsFilePath)
}

func recreateSection(template CredentialsFileTemplate, profile string, cfg *ini.File) {
	cfg.DeleteSection(profile)
	sec, err := cfg.NewSection(profile)
	check(err)
	err = sec.ReflectFrom(template)
}

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
