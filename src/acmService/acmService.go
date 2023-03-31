package acmService

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/dfds/roles-anywhere-helper/awsService"
	"github.com/dfds/roles-anywhere-helper/fileNames"
)

func ImportCertificate(creds awsService.AwsCredentialsObject, certificateDirectory, region string) string {

	ctx, cfg := awsService.ConfigureAws(creds, region)

	svc := acm.NewFromConfig(cfg)
	println("Importing Certificate")
	input := &acm.ImportCertificateInput{
		Certificate:      ReadFile(certificateDirectory, fileNames.Certificate),
		PrivateKey:       ReadFile(certificateDirectory, fileNames.PrivateKey),
		CertificateChain: ReadFile(certificateDirectory, fileNames.CertificateChain),
	}

	result, err := svc.ImportCertificate(ctx, input)
	if err != nil {
		println("Importing certificate error", err)
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
