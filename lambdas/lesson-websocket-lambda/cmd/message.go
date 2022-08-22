package cmd

import (
	"github.com/aws/aws-lambda-go/events"
)

// Default ..
func Default(request APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	//tableName := os.Getenv("CONNECTIONS_TABLE")
	//tableNameMessages := os.Getenv("MESSAGES_TABLE")
	//
	//websocketApi := os.Getenv("WEBSOCKET_API")
	//deploymentType := os.Getenv("DEPLOYMENT_TYPE")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
