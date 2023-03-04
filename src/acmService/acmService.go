package acmService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
	"io/ioutil"
)

func ImportCertificate(cmd *cobra.Command, args []string) (CertificateArn string) {
	profileName, _ := cmd.Flags().GetString(Flags.ProfileName)
	certificateDirectory, _ := cmd.Flags().GetString(Flags.CertificateDirectory)
	privateKeyDirectory, _ := cmd.Flags().GetString(Flags.PrivateKeyDirectory)

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           profileName,
	})
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	svc := acm.New(sess)

	// Load the certificate files into memory
	certData, err := ioutil.ReadFile(certificateDirectory)
	if err != nil {
		panic(err)
	}

	privateKeyData, err := ioutil.ReadFile(privateKeyDirectory)
	if err != nil {
		panic(err)
	}

	// Convert the certificate and private key data to strings
	certString := string(certData)
	privateKeyString := string(privateKeyData)

	// Import the certificate into ACM
	input := &acm.ImportCertificateInput{
		Certificate: &certString,
		PrivateKey:  &privateKeyString,
	}
	result, err := svc.ImportCertificate(input)
	if err != nil {
		panic(err)
	}

	// Print the ARN of the imported certificate
	return fmt.Println(*result.CertificateArn)
}
