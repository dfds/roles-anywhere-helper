package acmService

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/dfds/roles-anywhere-helper/awsService"
	"github.com/dfds/roles-anywhere-helper/credentialService"
	"github.com/dfds/roles-anywhere-helper/fileNames"
)

func ImportCertificate(creds awsService.AwsCredentialsObject, certificateDirectory, region string) (string, error) {

	// Load default config section before it gets overwritten
	defaultConfigSection := credentialService.LoadSection("default")

	ctx, cfg := awsService.ConfigureAws(creds, region)

	svc := acm.NewFromConfig(cfg)
	println("Importing Certificate")

	certificateData, err := readFile(certificateDirectory, fileNames.Certificate)
	if err != nil {
		return "", err
	}
	privateKeyData, err := readFile(certificateDirectory, fileNames.PrivateKey)
	if err != nil {
		return "", err
	}
	certificateChainData, err := readFile(certificateDirectory, fileNames.CertificateChain)
	if err != nil {
		return "", err
	}

	input := &acm.ImportCertificateInput{
		Certificate:      certificateData,
		PrivateKey:       privateKeyData,
		CertificateChain: certificateChainData,
	}

	result, err := svc.ImportCertificate(ctx, input)

	// Set the default profile back to the original
	var template credentialService.CredentialsFileTemplate
	template.CredentialProcess = fmt.Sprint(defaultConfigSection.Key("credential_process"))
	template.Region = fmt.Sprint(defaultConfigSection.Key("region"))
	credentialService.RecreateSection(&template, "default", credentialService.GetIniFile())
	credentialService.WriteIniFile(&template, "default")

	if err != nil {
		return "", fmt.Errorf("importing certificate error: %w", err)
	}

	certArn := *result.CertificateArn
	println("---------- CertificateArn -----------")
	println(certArn)

	return certArn, nil
}

func readFile(directory string, fileName string) ([]byte, error) {
	fileData, err := os.ReadFile(filepath.Join(directory, fileName))
	if err != nil {
		return nil, err
	}
	return fileData, nil
}

func DeleteCertificate(creds awsService.AwsCredentialsObject, certificateArn, region string) (bool, error) {

	defaultConfigSection := credentialService.LoadSection("default")

	ctx, cfg := awsService.ConfigureAws(creds, region)

	svc := acm.NewFromConfig(cfg)
	println("Deleting Certificate")

	input := &acm.DeleteCertificateInput{
		CertificateArn: &certificateArn,
	}

	_, err := svc.DeleteCertificate(ctx, input)

	// Set the default profile back to the original
	var template credentialService.CredentialsFileTemplate
	template.CredentialProcess = fmt.Sprint(defaultConfigSection.Key("credential_process"))
	template.Region = fmt.Sprint(defaultConfigSection.Key("region"))
	credentialService.RecreateSection(&template, "default", credentialService.GetIniFile())
	credentialService.WriteIniFile(&template, "default")

	if err != nil {
		return false, fmt.Errorf("Deleting certificate error: %w", err)
	}

	return true, nil

}
