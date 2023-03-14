package acmService

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/dfds/iam-anywhere-ninja/fileNames"
)

func ImportCertificate(profileName string, certificateDirectory string) {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewSharedCredentials("", profileName),
	})
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	svc := acm.New(sess)

	// Load the certificate files into memory
	certData, err := ioutil.ReadFile(certificateDirectory + "/" + fileNames.Certificate)
	if err != nil {
		panic(err)
	}

	privateKeyData, err := ioutil.ReadFile(certificateDirectory + "/" + fileNames.PrivateKey)
	if err != nil {
		panic(err)
	}

	// Import the certificate into ACM
	input := &acm.ImportCertificateInput{
		Certificate: certData,
		PrivateKey:  privateKeyData,
	}
	result, err := svc.ImportCertificate(input)
	if err != nil {
		panic(err)
	}

	// Print the ARN of the imported certificate
	fmt.Println(*result.CertificateArn)
}
