package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"strings"
)

const (
	DefaultRegion = "eu-central-1" // Used when region not provided
)

// AWS client

type Client struct {
	DynamoCl *dynamodb.Client
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
	}
}

// GetSession creates a new aws session and returns it
func GetSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})
	if err != nil {
		log.Fatalln("Unable to create AWS session", err.Error())
	}

	return sess
}
