package awsService

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	Profile         string
}

type AwsCredentialsObject struct {
	Value Credentials
}

func configureAwsByProfile(profileName string, region string) (context.Context, aws.Config) {

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(region),
		config.WithSharedConfigProfile(profileName))
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	return ctx, cfg
}

func configureAwsByCredentials(accessKey string, secret string, token string, region string) (context.Context, aws.Config) {

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKey,
			secret,
			token,
		)))
	if err != nil {
		log.Fatalf("failed to load configuration")
	}

	return ctx, cfg
}

func ConfigureAws(credentials AwsCredentialsObject, region string) (context.Context, aws.Config) {

	if credentials.Value.AccessKeyID == "" || credentials.Value.SecretAccessKey == "" {
		return configureAwsByProfile(credentials.Value.Profile, region)
	}

	return configureAwsByCredentials(credentials.Value.AccessKeyID, credentials.Value.SecretAccessKey, credentials.Value.SessionToken, region)
}

func NewAwsCredentialsObject(accessKeyID, secretAccessKey, sessionToken, profile string) AwsCredentialsObject {
	return AwsCredentialsObject{
		Value: Credentials{
			AccessKeyID:     accessKeyID,
			SecretAccessKey: secretAccessKey,
			SessionToken:    sessionToken,
			Profile:         profile,
		},
	}
}
