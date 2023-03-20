package certificateHandler

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"os"
)

func CreatePemFileFromPemBlock(pemData *pem.Block, directory string, fileName string) {

	fileOut, err := os.Create(directory + fileName)

	if err != nil {
		panic(err)
	}

	defer fileOut.Close()

	pem.Encode(fileOut, pemData)

	fmt.Printf("%s created", fileName)
}

func CreatePemFileFromString(pemData string, directory string, fileName string) {

	fileOut, err := os.Create(directory + fileName)

	if err != nil {
		panic(err)
	}

	defer fileOut.Close()

	fileOut.WriteString(pemData)

	fmt.Printf("%s created", fileName)
}

func CreateCsrPEM(commonName string, organizationName string, organizationalUnit string, country string, locality string, province string, privateKey *rsa.PrivateKey) []byte {
	csrTemplate := GenerateCsrTemplate(commonName, organizationName, organizationalUnit, country, locality, province)

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)
	if err != nil {
		fmt.Println("Failed to create CSR:", err)
		panic(err)
	}

	csrPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})
	return csrPEM
}

func GeneratePrivateKey() *rsa.PrivateKey {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("acme: Failed to generate private key:", err)
		panic(err)
	}
	return key
}

func GenerateCsrTemplate(commonName string, organizationName string, organizationalUnit string, country string, locality string, province string) x509.CertificateRequest {
	csrTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:         commonName,
			Organization:       []string{organizationName},
			OrganizationalUnit: []string{organizationalUnit},
			Country:            []string{country},
			Locality:           []string{locality},
			Province:           []string{province},
		},
	}
	return csrTemplate
}
