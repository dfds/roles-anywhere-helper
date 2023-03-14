package acmpcaService

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acmpca"
	"github.com/dfds/iam-anywhere-ninja/fileNames"
)

func ImportCertificate(profileName string, acmpcaArn string, commonName string, organizationName []string, organizationalUnit []string, certificateDirectory string) string {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewSharedCredentials("", profileName),
	})
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	csrTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:         commonName,
			Organization:       organizationName,
			OrganizationalUnit: organizationalUnit,
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

	acmPCA := acmpca.New(sess)
	certResp, err := acmPCA.IssueCertificate(&acmpca.IssueCertificateInput{
		CertificateAuthorityArn: aws.String(acmpcaArn),
		Csr:                     csrPEM,
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

	certPEM := aws.StringValue(certData.Certificate)

	certificateOut, err := os.Create(certificateDirectory + fileNames.Certificate)

	if err != nil {
		panic(err)
	}

	defer certificateOut.Close()

	certificateOut.WriteString(certPEM)

	certChainPEM := aws.StringValue(certData.CertificateChain)
	certChainOut, err := os.Create(certificateDirectory + fileNames.CertificateChain)

	if err != nil {
		panic(err)
	}
	defer certChainOut.Close()

	certificateOut.WriteString(certChainPEM)

	privateKeyOut, err := os.Create(certificateDirectory + fileNames.PrivateKey)

	if err != nil {
		panic(err)
	}
	defer privateKeyOut.Close()

	pem.Encode(privateKeyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})

	println("---------- CertificateArn -----------")
	println(certResp.CertificateArn)
	return *certResp.CertificateArn
}
