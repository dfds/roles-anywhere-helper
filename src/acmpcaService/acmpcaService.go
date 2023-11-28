package acmpcaService

import (
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
	"github.com/dfds/roles-anywhere-helper/credentialService"
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

	// Load the default profile so we can save it back later
	defaultConfigSection := credentialService.LoadSection("default")

	certResp, err := acmPCA.IssueCertificate(ctx, &acmpca.IssueCertificateInput{
		CertificateAuthorityArn: aws.String(acmpcaArn),
		Csr:                     csrPem,
		SigningAlgorithm:        "SHA256WITHRSA",
		Validity: &types.Validity{
			Type:  "DAYS",
			Value: aws.Int64(expiryDays),
		},
	})

	if err != nil {
		fmt.Println("Failed to issue certificate:", err)
		return "", err
	}

	// Set the default profile back to the original
	var template credentialService.CredentialsFileTemplate
	template.CredentialProcess = fmt.Sprint(defaultConfigSection.Key("credential_process"))
	template.Region = fmt.Sprint(defaultConfigSection.Key("region"))
	credentialService.RecreateSection(&template, "default", credentialService.GetIniFile())
	credentialService.WriteIniFile(&template, "default")

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

func RevokeCertificate(acmCreds awsService.AwsCredentialsObject, pcaCreds awsService.AwsCredentialsObject, certArn, pcaArn, revocationReason, acmRegion, pcaRegion string) (string, error) {

	// Load the default profile so we can save it back later
	defaultConfigSection := credentialService.LoadSection("default")

	acmCtx, acmCfg := awsService.ConfigureAws(acmCreds, acmRegion)
	acmSvc := acm.NewFromConfig(acmCfg)

	pcaCtx, pcaCfg := awsService.ConfigureAws(pcaCreds, pcaRegion)
	pcaSvc := acmpca.NewFromConfig(pcaCfg)

	certData, err := acmSvc.DescribeCertificate(acmCtx, &acm.DescribeCertificateInput{CertificateArn: aws.String(certArn)})

	// Set the default profile back to the original
	var template credentialService.CredentialsFileTemplate
	template.CredentialProcess = fmt.Sprint(defaultConfigSection.Key("credential_process"))
	template.Region = fmt.Sprint(defaultConfigSection.Key("region"))
	credentialService.RecreateSection(&template, "default", credentialService.GetIniFile())
	credentialService.WriteIniFile(&template, "default")

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	rci := &acmpca.RevokeCertificateInput{
		CertificateAuthorityArn: aws.String(pcaArn),
		CertificateSerial:       certData.Certificate.Serial,
		RevocationReason:        types.RevocationReason(revocationReason),
	}

	_, err = pcaSvc.RevokeCertificate(pcaCtx, rci)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return "Successfully revoked certificate", nil
}
