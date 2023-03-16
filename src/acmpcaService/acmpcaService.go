package acmpcaService

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/dfds/iam-anywhere-ninja/awsService"
	"github.com/dfds/iam-anywhere-ninja/certificateHandler"
	"github.com/dfds/iam-anywhere-ninja/fileNames"
	"time"
)

func ImportCertificate(profileName, acmpcaArn, commonName, organizationName, organizationalUnit, certificateDirectory string) string {

	ctx, cfg := awsService.ConfigureAws(profileName)
	println("Generating new certificate")

	privateKey := certificateHandler.GeneratePrivateKey()

	acmPCA := acmpca.NewFromConfig(cfg)

	certResp, err := acmPCA.IssueCertificate(ctx, &acmpca.IssueCertificateInput{
		CertificateAuthorityArn: aws.String(acmpcaArn),
		Csr:                     certificateHandler.CreateCsrPEM(commonName, organizationName, organizationalUnit, privateKey),
		SigningAlgorithm:        "SHA256WITHRSA",
		Validity: &types.Validity{
			Type:  "DAYS",
			Value: aws.Int64(6),
		},
	})
	if err != nil {
		fmt.Println("Failed to issue certificate:", err)
		panic(err)
	}

	waiter := acmpca.NewCertificateIssuedWaiter(acmPCA)

	certData, err := waiter.WaitForOutput(
		ctx,
		&acmpca.GetCertificateInput{
			CertificateArn:          certResp.CertificateArn,
			CertificateAuthorityArn: aws.String(acmpcaArn),
		},
		5*time.Second,
	)
	if err != nil {
		fmt.Println("Failed to get certificate data:", err)
		panic(err)
	}

	printf("Creating certificate Files.... in %s", certificateDirectory)

	certificateHandler.CreatePemFileFromString(*certData.Certificate, certificateDirectory, fileNames.Certificate)
	certificateHandler.CreatePemFileFromString(*certData.CertificateChain, certificateDirectory, fileNames.CertificateChain)
	certificateHandler.CreatePemFileFromPemBlock(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}, certificateDirectory, fileNames.PrivateKey)

	certArn := *certResp.CertificateArn
	println("---------- CertificateArn -----------")
	println(certArn)
	return certArn
}

func RevokeCertificate(profileName, certArn, pcaArn, revocationReason string) (string, error) {

	ctx, cfg := awsService.ConfigureAws(profileName)

	acmSvc := acm.NewFromConfig(cfg)

	certData, err := acmSvc.DescribeCertificate(context.TODO(), &acm.DescribeCertificateInput{CertificateArn: aws.String(certArn)})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	svc := acmpca.NewFromConfig(cfg)

	rci := &acmpca.RevokeCertificateInput{
		CertificateAuthorityArn: aws.String(pcaArn),
		CertificateSerial:       certData.Certificate.Serial,
		RevocationReason:        types.RevocationReason(revocationReason),
	}

	_, err = svc.RevokeCertificate(ctx, rci)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return "Successfully revoked certificate", nil
}
