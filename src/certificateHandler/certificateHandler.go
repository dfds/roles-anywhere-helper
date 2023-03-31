package certificateHandler

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"path/filepath"

	"github.com/dfds/roles-anywhere-helper/fileHandler"
)

func CreatePemFileFromPemBlock(pemData *pem.Block, directory string, fileName string) error {

	fileOut, err := fileHandler.CreateFile(filepath.Join(directory, fileName))

	if err != nil {
		return err
	}

	defer fileOut.Close()

	pem.Encode(fileOut, pemData)

	fmt.Printf("%s created", fileName)
	return nil
}

func CreatePemFileFromString(pemData string, directory string, fileName string) error {

	fileOut, err := fileHandler.CreateFile(filepath.Join(directory, fileName))

	if err != nil {
		return err
	}

	defer fileOut.Close()

	fileOut.WriteString(pemData)

	fmt.Printf("%s created", fileName)
	return nil
}

func CreateCsrPEM(commonName string, organizationName string, organizationalUnit string, country string, locality string, province string, privateKey *rsa.PrivateKey) ([]byte, error) {
	csrTemplate := generateCsrTemplate(commonName, organizationName, organizationalUnit, country, locality, province)

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)
	if err != nil {
		fmt.Println("Failed to create CSR:", err)
		return nil, err
	}

	csrPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})
	return csrPEM, nil
}

func GeneratePrivateKey() (*rsa.PrivateKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("acme: Failed to generate private key:", err)
		return nil, err
	}
	return key, nil
}

func generateCsrTemplate(commonName string, organizationName string, organizationalUnit string, country string, locality string, province string) x509.CertificateRequest {
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
