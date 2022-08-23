package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
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
	DynamoCl       *dynamodb.Client
	WebsocketApiCL *apigatewaymanagementapi.Client
}

func NewClient(region string) *Client {
	//websocketApi := os.Getenv("WEBSOCKET_API")
	//deploymentType := os.Getenv("DEPLOYMENT_TYPE")

	// If no region provided use default one
	if strings.TrimSpace(region) == "" {
		region = DefaultRegion
	}

	// Load default config and assign region
	cfg, _ := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	//cfg2, _ := config.LoadDefaultConfig(context.Background(), config.WithRegion(region), config.With)
	//fmt.Println(cfg2)

	// Return new client
	return &Client{
		DynamoCl: dynamodb.NewFromConfig(cfg),
		//WebsocketApiCL: apigatewaymanagementapi.New(apigatewaymanagementapi.Options{
		//	Region:           region,
		//	EndpointResolver: apigatewaymanagementapi.EndpointResolverFromURL(fmt.Sprintf("%s.execute-api.eu-central-1.amazonaws.com/%s", websocketApi, deploymentType)),
		//}),
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
