package CredentialHandler

import (
	"fmt"

	"github.com/dfds/iam-anywhere-ninja/Flags"
	"github.com/dfds/iam-anywhere-ninja/ProfileHandler"

	"os"

	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

var CredentialsFilePath = GetCredentialsFilePath()

type CredentialsFileTemplate struct {
	CredentialProcess string `ini:"credential_process,omitempty"`
	Region            string `ini:"region,omitempty"`
}

func Configure(cmd *cobra.Command, args []string) {
	profileName, _ := cmd.Flags().GetString(Flags.ProfileName)
	certificateDirectory, _ := cmd.Flags().GetString(Flags.CertificateDirectory)
	privateKeyDirectory, _ := cmd.Flags().GetString(Flags.PrivateKeyDirectory)
	trustAnchorArn, _ := cmd.Flags().GetString(Flags.TrustAnchor)
	profileArn, _ := cmd.Flags().GetString(Flags.ProfileArn)
	roleArn, _ := cmd.Flags().GetString(Flags.RoleArn)
	region, _ := cmd.Flags().GetString(Flags.Region)

	file, err := CreateCredentialsFile(GetCredentialsFilePath())
	defer file.Close()

	if err != nil {
		fmt.Println("Credential file already exists:", err)
	} else {
		fmt.Println("Credential file written successfully!")
	}

	profileName = ProfileHandler.SetProfileName(profileName)
	fmt.Printf("Profile Name set to %s", profileName)
	profileTemplate := ProcessCredentialProcessTemplate(certificateDirectory, privateKeyDirectory, trustAnchorArn, profileArn, roleArn, region)
	WriteIniFile(&profileTemplate, profileName)
}

func GetCredentialsFilePath() string {
	homeDir, err := os.UserHomeDir()
	Check(err)
	return homeDir + "/.aws/credentials"
}

func ProcessCredentialProcessTemplate(certificateDirectory string, privateKeyDirectory string, trustAnchorArn string, profileArn string, roleArn string, region string) CredentialsFileTemplate {
	profileTemplate := CredentialsFileTemplate{
		CredentialProcess: fmt.Sprintf("aws_signing_helper credential-process --certificate %s --private-key %s --trust-anchor-arn %s --profile-arn %s --role-arn %s", certificateDirectory, privateKeyDirectory, trustAnchorArn, profileArn, roleArn),
		Region:            region,
	}
	fmt.Println(profileTemplate)

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
