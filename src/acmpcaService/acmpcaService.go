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
	"github.com/dfds/iam-anywhere-ninja/flags"
	"github.com/spf13/cobra"
)

func ImportCertificate(cmd *cobra.Command, args []string) {
	profileName, _ := cmd.Flags().GetString(flags.ProfileName)
	acmpcaArn, _ := cmd.Flags().GetString(flags.AcmpcaArn)
	commonName, _ := cmd.Flags().GetString(flags.CommonName)
	organizationName, _ := cmd.Flags().GetStringArray(flags.OrganizationName)
	organizationalUnit, _ := cmd.Flags().GetStringArray(flags.OrganizationalUnit)

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
		return
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, key)
	if err != nil {
		fmt.Println("Failed to create CSR:", err)
		return
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
		return
	}

	certData, err := acmPCA.GetCertificate(&acmpca.GetCertificateInput{
		CertificateArn:          certResp.CertificateArn,
		CertificateAuthorityArn: aws.String(acmpcaArn),
	})
	if err != nil {
		fmt.Println("Failed to get certificate data:", err)
		return
	}

	println("### Certificate ##")
	certPEM := aws.StringValue(certData.Certificate)

	println(certPEM)

	certificateOut, err := os.Create("certificate.pem")

	if err != nil {
		panic(err)
	}
	defer certificateOut.Close()

	certificateOut.WriteString(certPEM)

	println("### Certificate Chain ##")
	certChainPEM := aws.StringValue(certData.CertificateChain)

	println(certChainPEM)

	privateKeyOut, err := os.Create("privatekey.pem")

	if err != nil {
		panic(err)
	}
	defer privateKeyOut.Close()

	pem.Encode(privateKeyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})

}
