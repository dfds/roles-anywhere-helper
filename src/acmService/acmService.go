package acmService

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/dfds/roles-anywhere-helper/awsService"
	"github.com/dfds/roles-anywhere-helper/fileNames"
)

func ImportCertificate(profileName, certificateDirectory, region string) (string, error) {

	ctx, cfg := awsService.ConfigureAws(profileName, region)

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
