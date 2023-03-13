package certificateService

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"

	"github.com/dfds/iam-anywhere-ninja/flags"
	"github.com/spf13/cobra"
)

func Generate(cmd *cobra.Command, args []string) {

	certificateDirectory, _ := cmd.Flags().GetString(flags.CertificateDirectory)
	privateKeyDirectory, _ := cmd.Flags().GetString(flags.PrivateKeyDirectory)

	// Generate a 2048-bit RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Create a new self-signed certificate template
	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "localhost"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(6 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{"appsec.dfds.cloud"},
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
	}

	// Create a new self-signed certificate using the private key and template
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		panic(err)
	}

	// Write the private key and self-signed certificate to disk
	certOut, err := os.Create(certificateDirectory + "certificate.pem")
	if err != nil {
		panic(err)
	}
	defer certOut.Close()

	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	keyOut, err := os.Create(privateKeyDirectory + "privatekey.pem")
	if err != nil {
		panic(err)
	}
	defer keyOut.Close()

	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
}
