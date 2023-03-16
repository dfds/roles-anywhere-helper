package acmService

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/dfds/iam-anywhere-ninja/awsService"
	"github.com/dfds/iam-anywhere-ninja/fileNames"
)

func ImportCertificate(profileName, certificateDirectory string) string {

	ctx, cfg := awsService.ConfigureAws(profileName)

	svc := acm.NewFromConfig(cfg)
	println("Importing Certificate")
	input := &acm.ImportCertificateInput{
		Certificate:      ReadFile(certificateDirectory, fileNames.Certificate),
		PrivateKey:       ReadFile(certificateDirectory, fileNames.PrivateKey),
		CertificateChain: ReadFile(certificateDirectory, fileNames.CertificateChain),
	}

	result, err := svc.ImportCertificate(ctx, input)
	if err != nil {
		println("Importing Certificate Error", err)
		panic(err)
	}

	certArn := *result.CertificateArn
	println("---------- CertificateArn -----------")
	println(certArn)

	return certArn
}

func ReadFile(directory string, fileName string) []byte {
	fileData, err := os.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}
	return fileData
}
