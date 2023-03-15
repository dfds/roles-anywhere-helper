package acmpcaService

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/acmpca"
	"github.com/dfds/iam-anywhere-ninja/awsService"
	"github.com/dfds/iam-anywhere-ninja/certificateHandler"
	"github.com/dfds/iam-anywhere-ninja/fileNames"
)

func ImportCertificate(profileName string, acmpcaArn string, commonName string, organizationName string, organizationalUnit string, certificateDirectory string) string {

	sess := awsService.SetAwsSession(profileName)

	privateKey := certificateHandler.GeneratePrivateKey()

	acmPCA := acmpca.New(sess)

	certResp, err := acmPCA.IssueCertificate(&acmpca.IssueCertificateInput{
		CertificateAuthorityArn: aws.String(acmpcaArn),
		Csr:                     certificateHandler.CreateCsrPEM(commonName, organizationName, organizationalUnit, privateKey),
		SigningAlgorithm:        aws.String(acmpca.SigningAlgorithmSha256withrsa),
		Validity: &acmpca.Validity{
			Type:  aws.String(acmpca.ValidityPeriodTypeDays),
			Value: aws.Int64(6),
		},
	})
	if err != nil {
		fmt.Println("Failed to issue certificate:", err)
		panic(err)
	}

	certData, err := acmPCA.GetCertificate(&acmpca.GetCertificateInput{
		CertificateArn:          certResp.CertificateArn,
		CertificateAuthorityArn: aws.String(acmpcaArn),
	})
	if err != nil {
		fmt.Println("Failed to get certificate data:", err)
		panic(err)
	}

	certificateHandler.CreatePemFileFromString(aws.StringValue(certData.Certificate), certificateDirectory, fileNames.Certificate)
	certificateHandler.CreatePemFileFromString(aws.StringValue(certData.CertificateChain), certificateDirectory, fileNames.CertificateChain)
	certificateHandler.CreatePemFileFromPemBlock(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}, certificateDirectory, fileNames.PrivateKey)

	certArn := aws.StringValue(certResp.CertificateArn)
	println("---------- CertificateArn -----------")
	println(certArn)
	return certArn
}
