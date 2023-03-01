package CredentialHandler

import (
	"fmt"

	"go-aws-iam-roles-anywhere-credential-helper/Flags"
	"go-aws-iam-roles-anywhere-credential-helper/ProfileHandler"

	"os"

	"github.com/spf13/cobra"
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

func Configure(cmd *cobra.Command, args []string) {
	profileName, _ := cmd.Flags().GetString(Flags.ProfileName)
	certificateDirectory, _ := cmd.Flags().GetString(Flags.CertificateDirectory)
	privateKeyDirectory, _ := cmd.Flags().GetString(Flags.PrivateKeyDirectory)
	trustAnchorArn, _ := cmd.Flags().GetString(Flags.TrustAnchor)
	profileArn, _ := cmd.Flags().GetString(Flags.ProfileArn)
	roleArn, _ := cmd.Flags().GetString(Flags.RoleArn)
	region, _ := cmd.Flags().GetString(Flags.Region)

	file, err := createCredentialsFile(GetCredentialsFilePath())
	defer file.Close()
	if err != nil {
		fmt.Println("Error writing credential file:", err)
	} else {
		fmt.Println("Credential file written successfully!")
	}

	profileName = ProfileHandler.SetProfileName(profileName)
	fmt.Printf("Profile Name set to %s", profileName)
	profileTemplate := ProcessCredentialProcessTemplate(certificateDirectory, privateKeyDirectory, trustAnchorArn, profileArn, roleArn, region)
	writeIniFile(&profileTemplate, profileName)

}

func ProcessCredentialProcessTemplate(certificateDirectory string, privateKeyDirectory string, trustAnchorArn string, profileArn string, roleArn string, region string) CredentialsFileTemplate {
	profileTemplate := CredentialsFileTemplate{
		CredentialProcess: fmt.Sprintf("aws_signing_helper credential-process --certificate %s --private-key %s --trust-anchor-arn %s --profile-arn %s --role-arn %s", certificateDirectory, privateKeyDirectory, trustAnchorArn, profileArn, roleArn),
		Region:            region,
	}
	fmt.Println(profileTemplate)

	return profileTemplate
}

func createCredentialsFile(filePath string) (*os.File, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.MkdirAll(filePath, 0755)
	}

	_, e := os.Stat(filePath)

	if e == nil {
		return nil, os.ErrExist
	}
	return os.Create(filePath)
}

func writeIniFile(template *CredentialsFileTemplate, profile string) {
	fmt.Println("writeIniFile")
	cfg, err := ini.Load(CredentialsFilePath)
	check(err)
	recreateSection(template, profile, cfg)
	cfg.SaveTo(CredentialsFilePath)
}

func recreateSection(template *CredentialsFileTemplate, profile string, cfg *ini.File) {
	fmt.Println("recreateSection")
	cfg.DeleteSection(profile)
	sec, err := cfg.NewSection(profile)
	check(err)
	err = sec.ReflectFrom(template)
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
