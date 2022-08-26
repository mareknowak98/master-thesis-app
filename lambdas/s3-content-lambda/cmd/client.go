package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"strings"
)

const (
	DefaultRegion = "eu-central-1" // Used when region not provided
)

// AWS client

type Client struct {
	DynamoCl *dynamodb.Client
	S3Cl     *s3.Client
}

func NewClient(region string) *Client {

	// If no region provided use default one
	if strings.TrimSpace(region) == "" {
		region = DefaultRegion
	}

	// Load default config and assign region
	cfg, _ := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))

	// Return new client
	return &Client{
		DynamoCl: dynamodb.NewFromConfig(cfg),
		S3Cl:     s3.NewFromConfig(cfg),
	}
}
