package awsService

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func ConfigureAws(profileName string, region string) (context.Context, aws.Config) {

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
