package certificateService

package main

import (
    "crypto/rand"
    "crypto/rsa"
    "encoding/pem"
    "os"
)

func Generate(cmd *cobra.Command, args []string) {

	certificateDirectory, _ := cmd.Flags().GetString(Flags.CertificateDirectory)
	privateKeyDirectory, _ := cmd.Flags().GetString(Flags.PrivateKeyDirectory)

    // Generate a 2048-bit RSA key pair
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        panic(err)
    }

    // Encode the private key to PEM format
    privateKeyPem := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
    }
    privatePemFile, err := os.Create(privateKeyDirectory)
    if err != nil {
        panic(err)
    }
    pem.Encode(privatePemFile, privateKeyPem)
    privatePemFile.Close()

    // Encode the public key to PEM format
    publicKey := privateKey.PublicKey
    publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
    if err != nil {
        panic(err)
    }
    publicKeyPem := &pem.Block{
        Type:  "RSA PUBLIC KEY",
        Bytes: publicKeyBytes,
    }
    publicPemFile, err := os.Create(certificateDirectory)
    if err != nil {
        panic(err)
    }
    pem.Encode(publicPemFile, publicKeyPem)
    publicPemFile.Close()
}0