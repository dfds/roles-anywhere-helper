package acmService

import (
	"os"

	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/dfds/iam-anywhere-ninja/awsService"
	"github.com/dfds/iam-anywhere-ninja/fileNames"
)

func ImportCertificate(profileName string, certificateDirectory string) string {

	sess := awsService.SetAwsSession(profileName)

	svc := acm.New(sess)

	input := &acm.ImportCertificateInput{
		Certificate:      ReadFile(certificateDirectory, fileNames.Certificate),
		PrivateKey:       ReadFile(certificateDirectory, fileNames.PrivateKey),
		CertificateChain: ReadFile(certificateDirectory, fileNames.CertificateChain),
	}

	result, err := svc.ImportCertificate(input)
	if err != nil {
		panic(err)
	}

	return *result.CertificateArn
}

func ReadFile(directory string, fileName string) []byte {
	fileData, err := os.ReadFile(directory + fileName)
	if err != nil {
		panic(err)
	}
	return fileData
}
