package acmpcaService

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/dfds/roles-anywhere-helper/awsService"
	"github.com/dfds/roles-anywhere-helper/certificateHandler"
	"github.com/dfds/roles-anywhere-helper/fileNames"
)

func GenerateCertificate(creds awsService.AwsCredentialsObject, acmpcaArn, commonName, organizationName, organizationalUnit, country, locality, province, certificateDirectory, region string, expiryDays int64) (string, error) {

	ctx, cfg := awsService.ConfigureAws(creds, region)
	println("Generating new certificate")

	privateKey, err := certificateHandler.GeneratePrivateKey()
	if err != nil {
		return "", err
	}
	csrPem, err := certificateHandler.CreateCsrPEM(commonName, organizationName, organizationalUnit, country, locality, province, privateKey)
	if err != nil {
		return "", err
	}

	acmPCA := acmpca.NewFromConfig(cfg)

	// SOMEWHERE HERE THE DEFAULT AWS PROFILE IS OVERWRITTEN
	certResp, err := acmPCA.IssueCertificate(ctx, &acmpca.IssueCertificateInput{
		CertificateAuthorityArn: aws.String(acmpcaArn),
		Csr:                     csrPem,
		SigningAlgorithm:        "SHA256WITHRSA",
		Validity: &types.Validity{
			Type:  "DAYS",
			Value: aws.Int64(expiryDays),
		},
	})

	// ABOVE HERE
	if err != nil {
		fmt.Println("Failed to issue certificate:", err)
		return "", err
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
		return "", err
	}

	fmt.Printf("Creating certificate files.... in %s", certificateDirectory)

	certificateHandler.CreatePemFileFromString(*certData.Certificate, certificateDirectory, fileNames.Certificate)
	certificateHandler.CreatePemFileFromString(*certData.CertificateChain, certificateDirectory, fileNames.CertificateChain)
	certificateHandler.CreatePemFileFromPemBlock(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}, certificateDirectory, fileNames.PrivateKey)

	certArn := *certResp.CertificateArn
	println("---------- CertificateArn -----------")
	println(certArn)
	return certArn, nil
}

func RevokeCertificate(creds awsService.AwsCredentialsObject, certArn, pcaArn, revocationReason, region string) (string, error) {

	ctx, cfg := awsService.ConfigureAws(creds, region)

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
