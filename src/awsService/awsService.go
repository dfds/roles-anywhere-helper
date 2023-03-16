package awsService

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"
)

func ConfigureAws(profileName string) (context.Context, aws.Config) {

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("eu-central-1"),
		config.WithSharedConfigProfile(profileName))
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	return ctx, cfg
}
