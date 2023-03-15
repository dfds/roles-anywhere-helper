package acmpcaService

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/dfds/iam-anywhere-ninja/fileNames"
	"log"
	"os"
	"time"
)

func ImportCertificate(profileName, acmpcaArn, commonName, organizationName, organizationalUnit, certificateDirectory string) string {

	// Load config
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("eu-central-1"),
		config.WithSharedConfigProfile(profileName))
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	csrTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:         commonName,
			Organization:       []string{organizationName},
			OrganizationalUnit: []string{organizationalUnit},
		},
	}

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("acme: Failed to generate private key:", err)
		panic(err)
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, key)
	if err != nil {
		fmt.Println("Failed to create CSR:", err)
		panic(err)
	}

	csrPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})
	acmPCA := acmpca.NewFromConfig(cfg)

	certResp, err := acmPCA.IssueCertificate(ctx, &acmpca.IssueCertificateInput{
		CertificateAuthorityArn: aws.String(acmpcaArn),
		Csr:                     csrPEM,
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

	certPEM := *certData.Certificate

	certificateOut, err := os.Create(certificateDirectory + fileNames.Certificate)

	if err != nil {
		panic(err)
	}

	defer certificateOut.Close()

	certificateOut.WriteString(certPEM)

	certChainPEM := *certData.CertificateChain
	certChainOut, err := os.Create(certificateDirectory + fileNames.CertificateChain)

	if err != nil {
		panic(err)
	}
	defer certChainOut.Close()

	certChainOut.WriteString(certChainPEM)

	privateKeyOut, err := os.Create(certificateDirectory + fileNames.PrivateKey)

	if err != nil {
		panic(err)
	}
	defer privateKeyOut.Close()

	pem.Encode(privateKeyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})

	certArn := certResp.CertificateArn
	println("---------- CertificateArn -----------")
	println(certArn)
	return *certArn
}

func RevokeCertificate(profileName, certArn, pcaArn, revocationReason string) (string, error) {
	// Load config
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("eu-central-1"),
		config.WithSharedConfigProfile(profileName))
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	// Retrieve certificate serial
	acmSvc := acm.NewFromConfig(cfg)

	dco, err := acmSvc.DescribeCertificate(context.TODO(), &acm.DescribeCertificateInput{CertificateArn: aws.String(certArn)})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Create an ACM PCA client
	svc := acmpca.NewFromConfig(cfg)

	//Revoke certificate
	rci := &acmpca.RevokeCertificateInput{
		CertificateAuthorityArn: aws.String(pcaArn),
		CertificateSerial:       dco.Certificate.Serial,
		RevocationReason:        types.RevocationReason(revocationReason),
	}

	_, err = svc.RevokeCertificate(ctx, rci)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return "Successfully revoked certificate", nil
}
